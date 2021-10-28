// +build windows

package gss

import (
	"encoding/hex"
	"errors"
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

// WithConfig sets the Kerberos configuration used
func WithConfig(config string) func(*Client) error {
	return func(c *Client) error {
		return errNotSupported
	}
}

// New performs any library initialization necessary.
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

// Generate generates the TSIG MAC based on the established context.
// It is called with the bytes of the DNS message, and the partial TSIG
// record containing the algorithm and name which is the negotiated TKEY
// for this context.
// It returns the bytes for the TSIG MAC and any error that occurred.
func (c *Client) Generate(msg []byte, t *dns.TSIG) ([]byte, error) {

	if dns.CanonicalName(t.Algorithm) != tsig.GSS {
		return nil, dns.ErrKeyAlg
	}

	c.m.RLock()
	defer c.m.RUnlock()

	ctx, ok := c.ctx[t.Hdr.Name]
	if !ok {
		return nil, dns.ErrSecret
	}

	token, err := ctx.MakeSignature(msg, 0, 0)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Verify verifies the TSIG MAC based on the established context.
// It is called with the bytes of the DNS message, and the TSIG record
// containing the algorithm, MAC, and name which is the negotiated TKEY
// for this context.
// It returns any error that occurred.
func (c *Client) Verify(stripped []byte, t *dns.TSIG) error {

	if dns.CanonicalName(t.Algorithm) != tsig.GSS {
		return dns.ErrKeyAlg
	}

	c.m.RLock()
	defer c.m.RUnlock()

	ctx, ok := c.ctx[t.Hdr.Name]
	if !ok {
		return dns.ErrSecret
	}

	token, err := hex.DecodeString(t.MAC)
	if err != nil {
		return err
	}

	if _, err = ctx.VerifySignature(stripped, token, 0); err != nil {
		return err
	}

	return nil
}

func (c *Client) negotiateContext(host string, creds *sspi.Credentials) (string, time.Time, error) {

	hostname, _, err := net.SplitHostPort(host)
	if err != nil {
		return "", time.Time{}, err
	}

	keyname := generateTKEYName(hostname)

	ctx, output, err := negotiate.NewClientContext(creds, generateSPN(hostname))
	if err != nil {
		return "", time.Time{}, err
	}

	var completed bool
	var tkey *dns.TKEY

	for ok := false; !ok; ok = completed {

		var errs error

		// We don't care about non-TKEY answers, no additional RR's to send, and no signing
		if tkey, _, err = util.ExchangeTKEY(c.client, host, keyname, tsig.GSS, util.TkeyModeGSS, 3600, output, nil, "", ""); err != nil {
			errs = multierror.Append(errs, err)
			errs = multierror.Append(errs, ctx.Release())
			return "", time.Time{}, errs
		}

		if tkey.Header().Name != keyname {
			errs = multierror.Append(errs, errors.New("TKEY name does not match"))
			errs = multierror.Append(errs, ctx.Release())
			return "", time.Time{}, errs
		}

		input, err := hex.DecodeString(tkey.Key)
		if err != nil {
			errs = multierror.Append(errs, err)
			errs = multierror.Append(errs, ctx.Release())
			return "", time.Time{}, errs
		}

		if completed, output, err = ctx.Update(input); err != nil {
			errs = multierror.Append(errs, err)
			errs = multierror.Append(errs, ctx.Release())
			return "", time.Time{}, errs
		}
	}

	expiry := time.Unix(int64(tkey.Expiration), 0)

	c.m.Lock()
	defer c.m.Unlock()

	c.ctx[keyname] = ctx

	return keyname, expiry, nil
}

// NegotiateContext exchanges RFC 2930 TKEY records with the indicated DNS
// server to establish a security context using the current user.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
func (c *Client) NegotiateContext(host string) (string, time.Time, error) {

	creds, err := negotiate.AcquireCurrentUserCredentials()
	if err != nil {
		return "", time.Time{}, err
	}
	defer creds.Release()

	return c.negotiateContext(host, creds)
}

// NegotiateContextWithCredentials exchanges RFC 2930 TKEY records with the
// indicated DNS server to establish a security context using the provided
// credentials.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
func (c *Client) NegotiateContextWithCredentials(host, domain, username, password string) (string, time.Time, error) {

	creds, err := negotiate.AcquireUserCredentials(domain, username, password)
	if err != nil {
		return "", time.Time{}, err
	}
	defer creds.Release()

	return c.negotiateContext(host, creds)
}

// NegotiateContextWithKeytab exchanges RFC 2930 TKEY records with the
// indicated DNS server to establish a security context using the provided
// keytab.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
func (c *Client) NegotiateContextWithKeytab(host, domain, username, path string) (string, time.Time, error) {

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
		return errors.New("No such context")
	}

	if err := ctx.Release(); err != nil {
		return err
	}

	delete(c.ctx, keyname)

	return nil
}
