package credentials

// A StaticProvider is a set of credentials which are set programmatically,
// and will never expire.
type StaticProvider struct {
	creds     Value
	retrieved bool
}

func NewStaticCredentials(apiKey, apiSecret string) *Credentials {
	return NewCredentials(
		&StaticProvider{creds: Value{APIKey: apiKey, APISecret: apiSecret}},
	)
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (s *StaticProvider) Retrieve() (Value, error) {
	if !s.creds.IsSet() {
		return Value{}, ErrMissingIncomplete
	}

	s.retrieved = true

	return s.creds, nil
}

// IsExpired returns if the credentials are expired.
//
// For StaticProvider, the credentials never expired.
func (s *StaticProvider) IsExpired() bool {
	return !s.retrieved
}
