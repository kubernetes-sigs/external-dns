//go:build !windows && apcera
// +build !windows,apcera

package gss

import (
	"encoding/hex"
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

func (c *Client) generate(ctx *gssapi.CtxId, msg []byte) ([]byte, error) {
	message, err := c.lib.MakeBufferBytes(msg)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = multierror.Append(err, message.Release()).ErrorOrNil()
	}()

	token, err := ctx.GetMIC(gssapi.GSS_C_QOP_DEFAULT, message)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = multierror.Append(err, token.Release()).ErrorOrNil()
	}()

	return token.Bytes(), nil
}

func (c *Client) verify(ctx *gssapi.CtxId, stripped, mac []byte) error {
	// Turn the TSIG-stripped message bytes into a *gssapi.Buffer
	message, err := c.lib.MakeBufferBytes(stripped)
	if err != nil {
		return err
	}

	defer func() {
		err = multierror.Append(err, message.Release()).ErrorOrNil()
	}()

	// Turn the TSIG MAC bytes into a *gssapi.Buffer
	token, err := c.lib.MakeBufferBytes(mac)
	if err != nil {
		return err
	}

	defer func() {
		err = multierror.Append(err, token.Release()).ErrorOrNil()
	}()

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
//
//nolint:cyclop,funlen
func (c *Client) NegotiateContext(host string) (keyname string, expiry time.Time, err error) {
	hostname, _, err := net.SplitHostPort(host)
	if err != nil {
		return "", time.Time{}, err
	}

	keyname, err = generateTKEYName(hostname)
	if err != nil {
		return "", time.Time{}, err
	}

	buffer, err := c.lib.MakeBufferString(generateSPN(hostname))
	if err != nil {
		return "", time.Time{}, err
	}

	defer func() {
		err = multierror.Append(err, buffer.Release()).ErrorOrNil()
	}()

	service, err := buffer.Name(c.lib.GSS_KRB5_NT_PRINCIPAL_NAME)
	if err != nil {
		return "", time.Time{}, err
	}

	defer func() {
		err = multierror.Append(err, service.Release()).ErrorOrNil()
	}()

	var (
		input *gssapi.Buffer
		ctx   *gssapi.CtxId
	)

	for ok := true; ok; ok = c.lib.LastStatus.Major.ContinueNeeded() {
		nctx, _, output, _, duration, err := c.lib.InitSecContext(
			c.lib.GSS_C_NO_CREDENTIAL,
			ctx, // nil initially
			service,
			c.lib.GSS_C_NO_OID,
			gssapi.GSS_C_MUTUAL_FLAG|gssapi.GSS_C_REPLAY_FLAG|gssapi.GSS_C_INTEG_FLAG,
			0,
			c.lib.GSS_C_NO_CHANNEL_BINDINGS,
			input)

		ctx, expiry = nctx, time.Now().UTC().Add(duration)

		defer func() {
			err = multierror.Append(err, output.Release()).ErrorOrNil()
		}()

		if err != nil {
			if !c.lib.LastStatus.Major.ContinueNeeded() {
				return "", time.Time{}, err
			}
		} else {
			// There is no further token to send
			break
		}

		//nolint:lll
		tkey, _, err := util.ExchangeTKEY(c.client, host, keyname, tsig.GSS, util.TkeyModeGSS, 3600, output.Bytes(), nil, "", "")
		if err != nil {
			return "", time.Time{}, multierror.Append(err, ctx.DeleteSecContext())
		}

		if tkey.Header().Name != keyname {
			return "", time.Time{}, multierror.Append(errDoesNotMatch, ctx.DeleteSecContext())
		}

		key, err := hex.DecodeString(tkey.Key)
		if err != nil {
			return "", time.Time{}, multierror.Append(err, ctx.DeleteSecContext())
		}

		if input, err = c.lib.MakeBufferBytes(key); err != nil {
			return "", time.Time{}, multierror.Append(err, ctx.DeleteSecContext())
		}

		defer func() {
			err = multierror.Append(err, input.Release()).ErrorOrNil()
		}()
	}

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
func (c *Client) NegotiateContextWithCredentials(_, _, _, _ string) (string, time.Time, error) {
	return "", time.Time{}, errNotSupported
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

	if err := ctx.DeleteSecContext(); err != nil {
		return err
	}

	delete(c.ctx, keyname)

	return nil
}
