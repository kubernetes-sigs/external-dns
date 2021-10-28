package tsig

import "github.com/miekg/dns"

type multiProvider struct {
	providers []dns.TsigProvider
}

func (mp *multiProvider) Generate(msg []byte, t *dns.TSIG) (b []byte, err error) {
	for _, p := range mp.providers {
		b, err = p.Generate(msg, t)
		switch err {
		case dns.ErrKeyAlg:
			break
		default:
			return
		}
	}
	return nil, dns.ErrKeyAlg
}

func (mp *multiProvider) Verify(msg []byte, t *dns.TSIG) (err error) {
	for _, p := range mp.providers {
		err = p.Verify(msg, t)
		switch err {
		case dns.ErrKeyAlg:
			break
		default:
			return
		}
	}
	return dns.ErrKeyAlg
}

// MultiProvider creates a dns.TsigProvider that chains the provided input
// providers. This allows multiple TSIG algorithms.
//
// Each provider is called in turn and if it returns dns.ErrKeyAlg the next
// provider in the list is tried. On success or any other error, the result is
// returned; it does not continue down the list.
func MultiProvider(providers ...dns.TsigProvider) dns.TsigProvider {
	allProviders := make([]dns.TsigProvider, 0, len(providers))
	for _, p := range providers {
		if mp, ok := p.(*multiProvider); ok {
			allProviders = append(allProviders, mp.providers...)
		} else {
			allProviders = append(allProviders, p)
		}
	}
	return &multiProvider{allProviders}
}
