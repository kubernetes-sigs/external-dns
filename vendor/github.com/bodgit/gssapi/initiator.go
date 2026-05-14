package gssapi

import (
	"errors"
	"fmt"
	"math"
	"math/bits"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/crypto"
	ianaflags "github.com/jcmturner/gokrb5/v8/iana/flags"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/spnego"
	"github.com/jcmturner/gokrb5/v8/types"
)

// Initiator represents the client side of the GSSAPI protocol.
type Initiator struct {
	context

	config   string
	domain   string
	username string
	password string
	keytab   *string

	client *client.Client

	logger logr.Logger
}

func (ctx *Initiator) loadConfig() (*config.Config, error) {
	if ctx.config != "" {
		return config.NewFromString(ctx.config)
	}

	return loadConfig(ctx.logger)
}

func (ctx *Initiator) usePassword() bool {
	return ctx.domain != "" && ctx.username != "" && ctx.password != ""
}

func (ctx *Initiator) useKeytab() bool {
	return ctx.domain != "" && ctx.username != "" && ctx.keytab != nil
}

func (ctx *Initiator) newClient() (*client.Client, error) {
	cfg, err := ctx.loadConfig()
	if err != nil {
		return nil, err
	}

	settings := []func(*client.Settings){
		client.DisablePAFXFAST(true),
	}

	switch {
	case ctx.usePassword():
		return client.NewWithPassword(ctx.username, ctx.domain, ctx.password, cfg, settings...), nil
	case ctx.useKeytab():
		var kt *keytab.Keytab

		if *ctx.keytab != "" {
			kt, err = keytab.Load(*ctx.keytab)
		} else {
			kt, err = loadClientKeytab(ctx.logger)
		}

		if err != nil {
			return nil, err
		}

		return client.NewWithKeytab(ctx.username, ctx.domain, kt, cfg, settings...), nil
	}

	ctx.logger.Info("using default session")

	cache, err := loadCCache(ctx.logger)
	if err != nil {
		return nil, err
	}

	return client.NewFromCCache(cache, cfg, settings...)
}

// NewInitiator returns a new Initiator.
func NewInitiator(options ...Option[Initiator]) (*Initiator, error) {
	ctx := &Initiator{
		context: context{
			sequenceMask: math.MaxUint32,
			logger:       logr.Discard(),
		},
		logger: logr.Discard(),
	}

	var err error

	for _, option := range options {
		if err = option(ctx); err != nil {
			return nil, err
		}
	}

	if ctx.client, err = ctx.newClient(); err != nil {
		return nil, err
	}

	if err = ctx.client.AffirmLogin(); err != nil {
		return nil, err
	}

	return ctx, nil
}

// Close releases any resources held by the Initiator.
func (ctx *Initiator) Close() error {
	ctx.client.Destroy()

	return nil
}

// Initiate creates a new context targeting the service with the desired flags
// along with the initial input token, which will initially be nil. The output
// token is returned and whether another round is required.
//
//nolint:cyclop,funlen
func (ctx *Initiator) Initiate(service string, flags int, input []byte) ([]byte, bool, error) {
	if ctx.established {
		return nil, false, nil
	}

	var err error

	//nolint:nestif
	if len(input) == 0 {
		ctx.context.flags = flags & supportedFlags

		// BUG(bodgit): see https://github.com/jcmturner/gokrb5/issues/529
		ctx.context.expiry = time.Now().Add(ctx.client.Config.LibDefaults.TicketLifetime)

		var ticket messages.Ticket

		if ticket, ctx.context.key, err = ctx.client.GetServiceTicket(strings.ReplaceAll(service, "@", "/")); err != nil {
			return nil, false, err
		}

		ctx.context.peerName = fmt.Sprintf("%s@%s", ticket.SName.PrincipalNameString(), ticket.Realm)

		f := make([]int, 0, bits.OnesCount(uint(ctx.context.flags)))

		for i := 0; i < bits.Len(supportedFlags); i++ {
			if ctx.context.flags&(1<<i) != 0 {
				f = append(f, 1<<i)
			}
		}

		apreq, err := spnego.NewKRB5TokenAPREQ(ctx.client, ticket, ctx.context.key, f, nil)
		if err != nil {
			return nil, false, err
		}

		if ctx.context.doMutual() {
			types.SetFlag(&apreq.APReq.APOptions, ianaflags.APOptionMutualRequired)
		}

		if err = apreq.APReq.DecryptAuthenticator(ctx.context.key); err != nil {
			return nil, false, err
		}

		ctx.context.sequenceNumber = uint64(apreq.APReq.Authenticator.SeqNumber)

		ctx.context.ctime = apreq.APReq.Authenticator.CTime
		ctx.context.cusec = apreq.APReq.Authenticator.Cusec

		output, err := apreq.Marshal()
		if err != nil {
			return nil, false, err
		}

		if !ctx.context.doMutual() {
			ctx.context.established = true
			ctx.context.baseSequenceNumber = ctx.context.sequenceNumber
		}

		return output, true, nil
	}

	if !ctx.context.doMutual() {
		return nil, false, errors.New("not mutual")
	}

	var aprep spnego.KRB5Token
	if err = aprep.Unmarshal(input); err != nil {
		return nil, false, err
	}

	if aprep.IsKRBError() {
		return nil, false, errors.New("received Kerberos error")
	}

	if !aprep.IsAPRep() {
		return nil, false, errors.New("didn't receive an AP-REP")
	}

	b, err := crypto.DecryptEncPart(aprep.APRep.EncPart, ctx.context.key, keyusage.AP_REP_ENCPART)
	if err != nil {
		return nil, false, krberror.Errorf(err, krberror.DecryptingError, "error decrypting AP-REP enc-part")
	}

	var payload messages.EncAPRepPart
	if err = payload.Unmarshal(b); err != nil {
		return nil, false, krberror.Errorf(err, krberror.EncodingError, "error unmarshalling decrypted AP-REP enc-part")
	}

	ctx.context.baseSequenceNumber = uint64(payload.SequenceNumber)

	if payload.Subkey.KeyType != 0 {
		ctx.context.peerSubkey = payload.Subkey
	}

	// Use Round()) to strip off any monotonic clock reading
	if !ctx.context.ctime.Round(0).Equal(payload.CTime.UTC()) || ctx.context.cusec != payload.Cusec {
		return nil, false, errors.New("mutual failed")
	}

	ctx.context.established = true

	return nil, false, nil
}
