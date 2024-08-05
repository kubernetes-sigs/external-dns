package dns

// TSIGKey wraps an NS1 /tsig resource
type TSIGKey struct {
	Name      string `json:"name,omitempty"`
	Algorithm string `json:"algorithm,omitempty"`
	Secret    string `json:"secret,omitempty"`
}

// NewTsigKey takes a name, algorithm and secret and creates a new TSIG key.
func NewTsigKey(name string, algorithm string, secret string) *TSIGKey {
	tsigKey := TSIGKey{
		Name:      name,
		Algorithm: algorithm,
		Secret:    secret,
	}
	return &tsigKey
}
