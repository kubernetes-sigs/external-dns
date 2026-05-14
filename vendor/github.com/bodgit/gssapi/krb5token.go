package gssapi

import (
	"encoding/hex"
	"fmt"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/spnego"
)

// This is a 1:1 copy of the type from github.com/jcmturner/gokrb5/v8 with a
// marshal method that supports all token IDs instead of just the AP-REQ.
// If/when upstream fixes this omission it can be removed.

type krb5Token struct {
	oid   asn1.ObjectIdentifier
	tokID []byte
	// apReq    *messages.APReq
	apRep    *apRep
	krbError *messages.KRBError
}

func (m *krb5Token) marshal() ([]byte, error) {
	b, _ := asn1.Marshal(m.oid)
	b = append(b, m.tokID...)

	var (
		tb  []byte
		err error
	)

	switch hex.EncodeToString(m.tokID) {
	/*
		case spnego.TOK_ID_KRB_AP_REQ:
			tb, err = m.apReq.Marshal()
			if err != nil {
				return []byte{}, fmt.Errorf("error marshalling AP_REQ for MechToken: %w", err)
			}
	*/
	case spnego.TOK_ID_KRB_AP_REP:
		tb, err = m.apRep.marshal()
		if err != nil {
			return []byte{}, fmt.Errorf("error marshalling AP_REP for MechToken: %w", err)
		}
	case spnego.TOK_ID_KRB_ERROR:
		tb, err = m.krbError.Marshal()
		if err != nil {
			return []byte{}, fmt.Errorf("error marshalling KRB_ERROR for MechToken: %w", err)
		}
	}

	if err != nil {
		return nil, err
	}

	b = append(b, tb...)

	return asn1tools.AddASNAppTag(b, 0), nil
}
