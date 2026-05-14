package gssapi

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/go-logr/logr"
	"github.com/jcmturner/gokrb5/v8/gssapi"
	"github.com/jcmturner/gokrb5/v8/iana/errorcode"
	ianaflags "github.com/jcmturner/gokrb5/v8/iana/flags"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/spnego"
	"github.com/jcmturner/gokrb5/v8/types"
)

// Acceptor represents the server side of the GSSAPI protocol.
type Acceptor struct {
	context

	keytab    string
	principal *types.PrincipalName
	clockSkew time.Duration

	logger logr.Logger
}

// NewAcceptor returns a new Acceptor.
func NewAcceptor(options ...Option[Acceptor]) (*Acceptor, error) {
	ctx := &Acceptor{
		context: context{
			acceptor:     true,
			sequenceMask: math.MaxUint32,
			logger:       logr.Discard(),
		},
		clockSkew: 10 * time.Second,
		logger:    logr.Discard(),
	}

	var err error

	for _, option := range options {
		if err = option(ctx); err != nil {
			return nil, err
		}
	}

	return ctx, nil
}

// Close releases any resources held by the Acceptor.
func (ctx *Acceptor) Close() error {
	return nil
}

func verifyAPReq(apreq *messages.APReq, kt *keytab.Keytab, skew time.Duration, sname *types.PrincipalName) error {
	err := apreq.Ticket.DecryptEncPart(kt, sname)

	if _, ok := err.(messages.KRBError); ok { //nolint:errorlint
		return err
	} else if err != nil {
		return messages.NewKRBError(apreq.Ticket.SName, apreq.Ticket.Realm,
			errorcode.KRB_AP_ERR_BAD_INTEGRITY, "could not decrypt ticket")
	}

	if _, err := apreq.Ticket.Valid(skew); err != nil {
		return err
	}

	if err := apreq.DecryptAuthenticator(apreq.Ticket.DecryptedEncPart.Key); err != nil {
		return messages.NewKRBError(apreq.Ticket.SName, apreq.Ticket.Realm,
			errorcode.KRB_AP_ERR_BAD_INTEGRITY, "could not decrypt authenticator")
	}

	if !apreq.Authenticator.CName.Equal(apreq.Ticket.DecryptedEncPart.CName) {
		return messages.NewKRBError(apreq.Ticket.SName, apreq.Ticket.Realm,
			errorcode.KRB_AP_ERR_BADMATCH, "CName in Authenticator does not match that in service ticket")
	}

	if time.Now().UTC().Sub(
		apreq.Authenticator.CTime.Add(
			time.Duration(apreq.Authenticator.Cusec)*time.Microsecond)).Abs() > skew {
		return messages.NewKRBError(apreq.Ticket.SName, apreq.Ticket.Realm,
			errorcode.KRB_AP_ERR_SKEW, fmt.Sprintf("clock skew with client too large, greater than %v seconds", skew))
	}

	return nil
}

func getAPRepMessage(tkt messages.Ticket, key types.EncryptionKey, ctime time.Time, cusec int) (*apRep, uint64, error) {
	seq, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	if err != nil {
		return nil, 0, err
	}

	seqNum := seq.Int64() & 0x3fffffff

	encPart := encAPRepPart{
		CTime:          ctime,
		Cusec:          cusec,
		SequenceNumber: seqNum,
	}

	aprep, err := newAPRep(tkt, key, encPart)
	if err != nil {
		return nil, 0, fmt.Errorf("gssapi: %w", err)
	}

	return &aprep, uint64(seqNum), nil
}

// Accept responds to the token from the Initiator, returning a token to be
// sent back to the Initiator and whether another round is required.
//
//nolint:cyclop,funlen
func (ctx *Acceptor) Accept(input []byte) ([]byte, bool, error) {
	if ctx.context.established {
		return nil, false, nil
	}

	var apreq spnego.KRB5Token
	if err := apreq.Unmarshal(input); err != nil {
		return nil, false, err
	}

	if apreq.IsKRBError() {
		return nil, false, errors.New("received kerberos error")
	}

	// FIXME check invalid token ID

	if !apreq.IsAPReq() {
		return nil, false, errors.New("didn't receive an AP-REQ")
	}

	kt, err := loadKeytab(ctx.logger)
	if err != nil {
		return nil, false, err
	}

	var output []byte

	// if _, err := apreq.APReq.Verify(kt, ctx.clockSkew, FIXME, nil); err != nil {
	if err = verifyAPReq(&apreq.APReq, kt, ctx.clockSkew, ctx.principal); err != nil {
		var krbError messages.KRBError

		if errors.As(err, &krbError) {
			tb, _ := hex.DecodeString(spnego.TOK_ID_KRB_ERROR)

			m := krb5Token{
				oid:      gssapi.OIDKRB5.OID(),
				tokID:    tb,
				krbError: &krbError,
			}

			if output, err = m.marshal(); err == nil {
				return output, true, nil
			}
		}

		return nil, false, err
	}

	ctx.context.baseSequenceNumber = uint64(apreq.APReq.Authenticator.SeqNumber)

	ctx.context.ctime = apreq.APReq.Authenticator.CTime
	ctx.context.cusec = apreq.APReq.Authenticator.Cusec

	ctx.context.key = apreq.APReq.Ticket.DecryptedEncPart.Key

	if apreq.APReq.Authenticator.SubKey.KeyType != 0 {
		ctx.context.peerSubkey = apreq.APReq.Authenticator.SubKey
	}

	ctx.context.flags = int(supportedFlags & binary.LittleEndian.Uint32(apreq.APReq.Authenticator.Cksum.Checksum[20:24]))

	ctx.context.expiry = apreq.APReq.Ticket.DecryptedEncPart.EndTime

	ctx.context.peerName = fmt.Sprintf("%s@%s", apreq.APReq.Ticket.DecryptedEncPart.CName.PrincipalNameString(),
		apreq.APReq.Ticket.DecryptedEncPart.CRealm)

	if types.IsFlagSet(&apreq.APReq.APOptions, ianaflags.APOptionMutualRequired) {
		var aprep *apRep

		aprep, ctx.context.sequenceNumber, err = getAPRepMessage(apreq.APReq.Ticket, ctx.context.key,
			ctx.context.ctime, ctx.context.cusec)
		if err != nil {
			return nil, false, err
		}

		tb, _ := hex.DecodeString(spnego.TOK_ID_KRB_AP_REP)

		m := krb5Token{
			oid:   gssapi.OIDKRB5.OID(),
			tokID: tb,
			apRep: aprep,
		}

		output, err = m.marshal()
		if err != nil {
			return nil, false, err
		}
	} else {
		ctx.context.sequenceNumber = ctx.context.baseSequenceNumber
	}

	ctx.context.established = true

	return output, false, nil
}
