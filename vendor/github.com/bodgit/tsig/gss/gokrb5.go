//go:build !windows && !apcera
// +build !windows,!apcera

package gss

import (
	"encoding/hex"
	"net"
	"sync"
	"time"

	wrapper "github.com/bodgit/gssapi"
	"github.com/bodgit/tsig"
	"github.com/bodgit/tsig/internal/util"
	"github.com/go-logr/logr"
	"github.com/jcmturner/gokrb5/v8/gssapi"
	"github.com/miekg/dns"
)

// Client maps the TKEY name to the context that negotiated it as
// well as any other internal state.
type Client struct {
	m      sync.RWMutex
	client *dns.Client
	config string
	ctx    map[string]*wrapper.Initiator
	logger logr.Logger
}

// WithConfig sets the Kerberos configuration used.
func WithConfig(config string) func(*Client) error {
	return func(c *Client) error {
		c.config = config

		return nil
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
		ctx:    make(map[string]*wrapper.Initiator),
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

func (c *Client) generate(ctx *wrapper.Initiator, msg []byte) ([]byte, error) {
	return ctx.MakeSignature(msg)
}

func (c *Client) verify(ctx *wrapper.Initiator, stripped, mac []byte) error {
	return ctx.VerifySignature(stripped, mac)
}

func (c *Client) negotiateContext(host string, options []wrapper.Option[wrapper.Initiator]) (string, time.Time, error) {
	options = append(options, wrapper.WithConfig(c.config), wrapper.WithLogger[wrapper.Initiator](c.logger))

	ctx, err := wrapper.NewInitiator(options...)
	if err != nil {
		return "", time.Time{}, err
	}

	hostname, _, err := net.SplitHostPort(host)
	if err != nil {
		return "", time.Time{}, err
	}

	keyname, err := generateTKEYName(hostname)
	if err != nil {
		return "", time.Time{}, err
	}

	spn := generateSPN(hostname)

	flags := gssapi.ContextFlagMutual | gssapi.ContextFlagReplay | gssapi.ContextFlagInteg

	output, cont, err := ctx.Initiate(spn, flags, nil)
	if err != nil {
		return "", time.Time{}, err
	}

	var tkey *dns.TKEY

	for cont {
		// We don't care about non-TKEY answers, no additional RR's to send, and no signing
		tkey, _, err = util.ExchangeTKEY(c.client, host, keyname, tsig.GSS, util.TkeyModeGSS, 3600, output, nil, "", "")
		if err != nil {
			return "", time.Time{}, err
		}

		if tkey.Header().Name != keyname {
			return "", time.Time{}, errDoesNotMatch
		}

		var input []byte

		if input, err = hex.DecodeString(tkey.Key); err != nil {
			return "", time.Time{}, err
		}

		output, cont, err = ctx.Initiate(spn, flags, input)
		if err != nil {
			return "", time.Time{}, err
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
func (c *Client) NegotiateContext(host string) (string, time.Time, error) {
	return c.negotiateContext(host, nil)
}

// NegotiateContextWithCredentials exchanges RFC 2930 TKEY records with the
// indicated DNS server to establish a security context using the provided
// credentials.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
func (c *Client) NegotiateContextWithCredentials(host, domain, username, password string) (string, time.Time, error) {
	options := []wrapper.Option[wrapper.Initiator]{
		wrapper.WithDomain(domain),
		wrapper.WithUsername(username),
		wrapper.WithPassword(password),
	}

	return c.negotiateContext(host, options)
}

// NegotiateContextWithKeytab exchanges RFC 2930 TKEY records with the
// indicated DNS server to establish a security context using the provided
// keytab.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
func (c *Client) NegotiateContextWithKeytab(host, domain, username, path string) (string, time.Time, error) {
	options := []wrapper.Option[wrapper.Initiator]{
		wrapper.WithDomain(domain),
		wrapper.WithUsername(username),
		wrapper.WithKeytab[wrapper.Initiator](path),
	}

	return c.negotiateContext(host, options)
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

	if err := ctx.Close(); err != nil {
		return err
	}

	delete(c.ctx, keyname)

	return nil
}
