package gssapi

import (
	"time"

	"github.com/go-logr/logr"
	"github.com/jcmturner/gokrb5/v8/types"
)

// Option is the signature for all constructor options.
type Option[T Initiator | Acceptor] func(*T) error

// WithConfig permits passing krb5.conf contents directly to an Initiator.
func WithConfig[T Initiator](config string) Option[T] {
	return func(a *T) error {
		if x, ok := any(a).(*Initiator); ok {
			x.config = config
		}

		return nil
	}
}

// WithLogger configures a logr.Logger in either an Initiator or Acceptor.
func WithLogger[T Initiator | Acceptor](logger logr.Logger) Option[T] {
	return func(a *T) error {
		switch x := any(a).(type) {
		case *Initiator:
			x.logger = logger.WithName("initiator")
			x.context.logger = x.logger.WithName("ctx")
		case *Acceptor:
			x.logger = logger.WithName("acceptor")
			x.context.logger = x.logger.WithName("ctx")
		}

		return nil
	}
}

// WithDomain sets the Kerberos domain in the Initiator.
func WithDomain[T Initiator](domain string) Option[T] {
	return func(a *T) error {
		if x, ok := any(a).(*Initiator); ok {
			x.domain = domain
		}

		return nil
	}
}

// WithRealm is an alias for WithDomain.
func WithRealm[T Initiator](realm string) Option[T] {
	return WithDomain[T](realm)
}

// WithUsername sets the username in the Initiator.
func WithUsername[T Initiator](username string) Option[T] {
	return func(a *T) error {
		if x, ok := any(a).(*Initiator); ok {
			x.username = username
		}

		return nil
	}
}

// WithPassword sets the password in the Initiator.
func WithPassword[T Initiator](password string) Option[T] {
	return func(a *T) error {
		if x, ok := any(a).(*Initiator); ok {
			x.password = password
			x.keytab = nil
		}

		return nil
	}
}

// WithKeytab sets the keytab path in either an Initiator or Acceptor.
func WithKeytab[T Initiator | Acceptor](keytab string) Option[T] {
	return func(a *T) error {
		switch x := any(a).(type) {
		case *Initiator:
			x.keytab = &keytab
			x.password = ""
		case *Acceptor:
			x.keytab = keytab
		}

		return nil
	}
}

// WithServicePrincipal sets the principal that is looked up in the keytab.
func WithServicePrincipal[T Acceptor](principal *types.PrincipalName) Option[T] {
	return func(a *T) error {
		if x, ok := any(a).(*Acceptor); ok {
			x.principal = principal
		}

		return nil
	}
}

// WithClockSkew sets the permitted amount of clock skew allowed between the
// Initiator and Acceptor.
func WithClockSkew[T Acceptor](clockSkew time.Duration) Option[T] {
	return func(a *T) error {
		if x, ok := any(a).(*Acceptor); ok {
			x.clockSkew = clockSkew
		}

		return nil
	}
}
