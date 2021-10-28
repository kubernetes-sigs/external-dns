// +build !windows,apcera

package gss

import (
	"encoding/hex"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/bodgit/tsig"
	"github.com/bodgit/tsig/internal/util"
	"github.com/go-logr/logr"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/miekg/dns"
	"github.com/openshift/gssapi"
)

// Client maps the TKEY name to the context that negotiated it as
// well as any other internal state.
type Client struct {
	m      sync.RWMutex
	lib    *gssapi.Lib
	client *dns.Client
	ctx    map[string]*gssapi.CtxId
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

	lib, err := gssapi.Load(nil)
	if err != nil {
		return nil, err
	}

	c := &Client{
		lib:    lib,
		client: client,
		ctx:    make(map[string]*gssapi.CtxId),
		logger: logr.Discard(),
	}

	if err := c.setOption(options...); err != nil {
		return nil, multierror.Append(err, c.lib.Unload())
	}

	return c, nil
}

// Close deletes any active contexts and unloads any underlying libraries as
// necessary.
// It returns any error that occurred.
func (c *Client) Close() error {

	return multierror.Append(c.close(), c.lib.Unload())
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

	message, err := c.lib.MakeBufferBytes(msg)
	if err != nil {
		return nil, err
	}
	defer message.Release()

	token, err := ctx.GetMIC(gssapi.GSS_C_QOP_DEFAULT, message)
	if err != nil {
		return nil, err
	}
	defer token.Release()

	return token.Bytes(), nil
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

	// Turn the TSIG-stripped message bytes into a *gssapi.Buffer
	message, err := c.lib.MakeBufferBytes(stripped)
	if err != nil {
		return err
	}
	defer message.Release()

	mac, err := hex.DecodeString(t.MAC)
	if err != nil {
		return err
	}

	// Turn the TSIG MAC bytes into a *gssapi.Buffer
	token, err := c.lib.MakeBufferBytes(mac)
	if err != nil {
		return err
	}
	defer token.Release()

	// This is the actual verification bit
	if _, err = ctx.VerifyMIC(message, token); err != nil {
		return err
	}

	return nil
}

// NegotiateContext exchanges RFC 2930 TKEY records with the indicated DNS
// server to establish a security context using the current user.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
func (c *Client) NegotiateContext(host string) (string, time.Time, error) {

	hostname, _, err := net.SplitHostPort(host)
	if err != nil {
		return "", time.Time{}, err
	}

	keyname := generateTKEYName(hostname)

	buffer, err := c.lib.MakeBufferString(generateSPN(hostname))
	if err != nil {
		return "", time.Time{}, err
	}
	defer buffer.Release()

	service, err := buffer.Name(c.lib.GSS_KRB5_NT_PRINCIPAL_NAME)
	if err != nil {
		return "", time.Time{}, err
	}
	defer service.Release()

	var input *gssapi.Buffer
	var ctx *gssapi.CtxId
	var tkey *dns.TKEY

	for ok := true; ok; ok = c.lib.LastStatus.Major.ContinueNeeded() {
		nctx, _, output, _, _, err := c.lib.InitSecContext(
			c.lib.GSS_C_NO_CREDENTIAL,
			ctx, // nil initially
			service,
			c.lib.GSS_C_NO_OID,
			gssapi.GSS_C_MUTUAL_FLAG|gssapi.GSS_C_REPLAY_FLAG|gssapi.GSS_C_INTEG_FLAG,
			0,
			c.lib.GSS_C_NO_CHANNEL_BINDINGS,
			input)
		defer output.Release()
		ctx = nctx
		if err != nil {
			if !c.lib.LastStatus.Major.ContinueNeeded() {
				return "", time.Time{}, err
			}
		} else {
			// There is no further token to send
			break
		}

		var errs error

		// We don't care about non-TKEY answers, no additional RR's to send, and no signing
		if tkey, _, err = util.ExchangeTKEY(c.client, host, keyname, tsig.GSS, util.TkeyModeGSS, 3600, output.Bytes(), nil, "", ""); err != nil {
			errs = multierror.Append(errs, err)
			errs = multierror.Append(errs, ctx.DeleteSecContext())
			return "", time.Time{}, errs
		}

		if tkey.Header().Name != keyname {
			errs = multierror.Append(errs, errors.New("TKEY name does not match"))
			errs = multierror.Append(errs, ctx.DeleteSecContext())
			return "", time.Time{}, errs
		}

		key, err := hex.DecodeString(tkey.Key)
		if err != nil {
			errs = multierror.Append(errs, err)
			errs = multierror.Append(errs, ctx.DeleteSecContext())
			return "", time.Time{}, errs
		}

		if input, err = c.lib.MakeBufferBytes(key); err != nil {
			errs = multierror.Append(errs, err)
			errs = multierror.Append(errs, ctx.DeleteSecContext())
			return "", time.Time{}, errs
		}
		defer input.Release()
	}

	expiry := time.Unix(int64(tkey.Expiration), 0)

	c.m.Lock()
	defer c.m.Unlock()

	c.ctx[keyname] = ctx

	return keyname, expiry, nil
}

// NegotiateContextWithCredentials exchanges RFC 2930 TKEY records with the
// indicated DNS server to establish a security context using the provided
// credentials.
// It returns the negotiated TKEY name, expiration time, and any error that
// occurred.
func (c *Client) NegotiateContextWithCredentials(host, domain, username, password string) (string, time.Time, error) {

	return "", time.Time{}, errNotSupported
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

	if err := ctx.DeleteSecContext(); err != nil {
		return err
	}

	delete(c.ctx, keyname)

	return nil
}
