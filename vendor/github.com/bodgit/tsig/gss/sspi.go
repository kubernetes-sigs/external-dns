//go:build windows
// +build windows

package gss

import (
	"encoding/hex"
	"net"
	"sync"
	"time"

	"github.com/alexbrainman/sspi"
	"github.com/alexbrainman/sspi/negotiate"
	"github.com/bodgit/tsig"
	"github.com/bodgit/tsig/internal/util"
	"github.com/go-logr/logr"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/miekg/dns"
)

// Client maps the TKEY name to the context that negotiated it as
// well as any other internal state.
type Client struct {
	m      sync.RWMutex
	client *dns.Client
	ctx    map[string]*negotiate.ClientContext
	logger logr.Logger
}

// WithConfig sets the Kerberos configuration used.
func WithConfig(_ string) func(*Client) error {
	return func(c *Client) error {
		return errNotSupported
	}
}

// NewClient performs any library initialization necessary.
// It returns a context handle for any further functions along with any error
// that occurred.
func NewClient(dnsClient *dns.Client, options ...func(*Client) error) (*Client, error) {
	client, err := util.CopyDNSClient(dnsClient)
	if err != nil {
		return nil, err
	}

	client.TsigProvider = new(gssNoVerify)

	c := &Client{
		client: client,
		ctx:    make(map[string]*negotiate.ClientContext),
		logger: logr.Discard(),
	}

	if err := c.setOption(options...); err != nil {
		return nil, err
	}

	return c, nil
}

// Close deletes any active contexts and unloads any underlying libraries as
// necessary.
// It returns any error that occurred.
func (c *Client) Close() error {
	return c.close()
}

func (c *Client) generate(ctx *negotiate.ClientContext, msg []byte) ([]byte, error) {
	return ctx.MakeSignature(msg, 0, 0)
}

func (c *Client) verify(ctx *negotiate.ClientContext, stripped, mac []byte) error {
	_, err := ctx.VerifySignature(stripped, mac, 0)

	return err
}

func (c *Client) negotiateContext(host string, creds *sspi.Credentials) (string, time.Time, error) {
	hostname, _, err := net.SplitHostPort(host)
	if err != nil {
		return "", time.Time{}, err
	}

	keyname, err := generateTKEYName(hostname)
	if err != nil {
		return "", time.Time{}, err
	}

	ctx, output, err := negotiate.NewClientContext(creds, generateSPN(hostname))
	if err != nil {
		return "", time.Time{}, err
	}

	var (
		completed bool
		tkey      *dns.TKEY
	)

	for ok := false; !ok; ok = completed {
		//nolint:lll
		if tkey, _, err = util.ExchangeTKEY(c.client, host, keyname, tsig.GSS, util.TkeyModeGSS, 3600, output, nil, "", ""); err != nil {
			return "", time.Time{}, multierror.Append(err, ctx.Release())
		}

		if tkey.Header().Name != keyname {
			return "", time.Time{}, multierror.Append(errDoesNotMatch, ctx.Release())
		}

		input, err := hex.DecodeString(tkey.Key)
		if err != nil {
			return "", time.Time{}, multierror.Append(err, ctx.Release())
		}

		if completed, output, err = ctx.Update(input); err != nil {
			return "", time.Time{}, multierror.Append(err, ctx.Release())
		}
	}

	c.m.Lock()
	defer c.m.Unlock()

	c.ctx[keyname] = ctx

	return keyname, ctx.Expiry(), nil
}

// NegotiateContext exchanges RFC 2930 TKEY records with the indicated DNS
// server to establish a security context using the current user.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
func (c *Client) NegotiateContext(host string) (keyname string, expiry time.Time, err error) {
	creds, err := negotiate.AcquireCurrentUserCredentials()
	if err != nil {
		return "", time.Time{}, err
	}

	defer func() {
		err = multierror.Append(err, creds.Release()).ErrorOrNil()
	}()

	return c.negotiateContext(host, creds)
}

// NegotiateContextWithCredentials exchanges RFC 2930 TKEY records with the
// indicated DNS server to establish a security context using the provided
// credentials.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
//
//nolint:lll
func (c *Client) NegotiateContextWithCredentials(host, domain, username, password string) (keyname string, expiry time.Time, err error) {
	creds, err := negotiate.AcquireUserCredentials(domain, username, password)
	if err != nil {
		return "", time.Time{}, err
	}

	defer func() {
		err = multierror.Append(err, creds.Release()).ErrorOrNil()
	}()

	return c.negotiateContext(host, creds)
}

// NegotiateContextWithKeytab exchanges RFC 2930 TKEY records with the
// indicated DNS server to establish a security context using the provided
// keytab.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
func (c *Client) NegotiateContextWithKeytab(_, _, _, _ string) (string, time.Time, error) {
	return "", time.Time{}, errNotSupported
}

// DeleteContext deletes the active security context associated with the given
// TKEY name.
// It returns any error that occurred.
func (c *Client) DeleteContext(keyname string) error {
	c.m.Lock()
	defer c.m.Unlock()

	ctx, ok := c.ctx[keyname]
	if !ok {
		return errNoSuchContext
	}

	if err := ctx.Release(); err != nil {
		return err
	}

	delete(c.ctx, keyname)

	return nil
}
