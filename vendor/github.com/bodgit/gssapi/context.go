package gssapi

import (
	"errors"
	"time"

	"github.com/go-logr/logr"
	"github.com/jcmturner/gokrb5/v8/gssapi"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/types"
)

var (
	errDuplicateToken = errors.New("duplicate per-message token detected")
	errOldToken       = errors.New("timed-out per-message token detected")
	errUnseqToken     = errors.New("reordered (early) per-message token detected")
	errGapToken       = errors.New("skipped predecessor token(s) detected")
)

type context struct {
	acceptor    bool
	established bool

	key        types.EncryptionKey
	subkey     types.EncryptionKey
	peerSubkey types.EncryptionKey
	flags      int
	ctime      time.Time
	cusec      int
	expiry     time.Time

	peerName string

	sequenceNumber uint64

	baseSequenceNumber uint64
	nextSequenceNumber uint64
	receiveMask        uint64
	sequenceMask       uint64

	logger logr.Logger
}

func (ctx *context) hasSubkey() bool {
	return ctx.subkey.KeyType != 0
}

func (ctx *context) hasPeerSubkey() bool {
	return ctx.peerSubkey.KeyType != 0
}

func (ctx *context) doMutual() bool {
	return ctx.flags&gssapi.ContextFlagMutual != 0
}

func (ctx *context) doReplay() bool {
	return ctx.flags&gssapi.ContextFlagReplay != 0
}

func (ctx *context) doSequence() bool {
	return ctx.flags&gssapi.ContextFlagSequence != 0
}

//nolint:cyclop
func (ctx *context) checkSequenceNumber(sequenceNumber uint64) error {
	if !ctx.doReplay() && !ctx.doSequence() {
		return nil
	}

	relativeSequenceNumber := (sequenceNumber - ctx.baseSequenceNumber) & ctx.sequenceMask

	if relativeSequenceNumber >= ctx.nextSequenceNumber {
		offset := relativeSequenceNumber - ctx.nextSequenceNumber
		ctx.receiveMask = ctx.receiveMask<<(offset+1) | 1
		ctx.nextSequenceNumber = (relativeSequenceNumber + 1) & ctx.sequenceMask

		if offset > 0 && ctx.doSequence() {
			return errGapToken
		}

		return nil
	}

	offset := ctx.nextSequenceNumber - relativeSequenceNumber

	if offset > 64 {
		if ctx.doSequence() {
			return errUnseqToken
		}

		return errOldToken
	}

	bit := uint64(1) << (offset - 1)
	if ctx.doReplay() && ctx.receiveMask&bit != 0 {
		return errDuplicateToken
	}

	ctx.receiveMask |= bit

	if ctx.doSequence() {
		return errUnseqToken
	}

	return nil
}

// PeerName returns the peer Kerberos principal.
func (ctx *context) PeerName() string {
	return ctx.peerName
}

// Established returns the context state.
func (ctx *context) Established() bool {
	return ctx.established
}

// Expiry returns the ticket expiry for the context.
func (ctx *context) Expiry() time.Time {
	return ctx.expiry
}

// MakeSignature creates a MIC token against the provided input.
func (ctx *context) MakeSignature(message []byte) ([]byte, error) {
	var (
		flags byte
		usage uint32 = keyusage.GSSAPI_INITIATOR_SIGN
	)

	if ctx.acceptor {
		flags |= gssapi.MICTokenFlagSentByAcceptor
		usage = keyusage.GSSAPI_ACCEPTOR_SIGN
	}

	key := ctx.key
	if ctx.hasSubkey() {
		key = ctx.subkey

		if ctx.acceptor {
			flags |= gssapi.MICTokenFlagAcceptorSubkey
		}
	}

	token := gssapi.MICToken{
		Flags:     flags,
		SndSeqNum: ctx.sequenceNumber,
		Payload:   message,
	}

	if err := token.SetChecksum(key, usage); err != nil {
		return nil, err
	}

	signature, err := token.Marshal()
	if err != nil {
		return nil, err
	}

	ctx.sequenceNumber++

	return signature, nil
}

// VerifySignature verifies the MIC token against the provided input.
func (ctx *context) VerifySignature(message, signature []byte) error {
	var (
		token gssapi.MICToken
		err   error
	)

	if err = token.Unmarshal(signature, !ctx.acceptor); err != nil {
		return err
	}

	token.Payload = message

	if err = ctx.checkSequenceNumber(token.SndSeqNum); err != nil {
		return err
	}

	var usage uint32 = keyusage.GSSAPI_ACCEPTOR_SIGN
	if ctx.acceptor {
		usage = keyusage.GSSAPI_INITIATOR_SIGN
	}

	key := ctx.key
	if ctx.hasPeerSubkey() {
		key = ctx.peerSubkey
	}

	if _, err = token.Verify(key, usage); err != nil {
		return err
	}

	return nil
}
