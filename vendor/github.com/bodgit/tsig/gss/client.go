package gss

import (
	"encoding/hex"

	"github.com/bodgit/tsig"
	"github.com/go-logr/logr"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/miekg/dns"
)

var _ dns.TsigProvider = new(Client)

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

	return c.generate(ctx, msg)
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

	mac, err := hex.DecodeString(t.MAC)
	if err != nil {
		return err
	}

	return c.verify(ctx, stripped, mac)
}

func (c *Client) close() error {
	c.m.RLock()

	keys := make([]string, 0, len(c.ctx))
	for k := range c.ctx {
		keys = append(keys, k)
	}

	c.m.RUnlock()

	var errs error
	for _, k := range keys {
		errs = multierror.Append(errs, c.DeleteContext(k))
	}

	return errs
}

func (c *Client) setOption(options ...func(*Client) error) error {
	for _, option := range options {
		if err := option(c); err != nil {
			return err
		}
	}

	return nil
}

// SetConfig sets the Kerberos configuration used by c.
func (c *Client) SetConfig(config string) error {
	return c.setOption(WithConfig(config))
}

// WithLogger sets the logger used.
func WithLogger(logger logr.Logger) func(*Client) error {
	return func(c *Client) error {
		c.logger = logger.WithName("client")

		return nil
	}
}

// SetLogger sets the logger used by c.
func (c *Client) SetLogger(logger logr.Logger) error {
	return c.setOption(WithLogger(logger))
}
