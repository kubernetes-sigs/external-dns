package credentials

import "os"

type EnvProvider struct {
	retrieved bool
}

func NewEnvCredentials() *Credentials {
	return NewCredentials(&EnvProvider{})
}

// Retrieve retrieves the keys from the environment.
func (e *EnvProvider) Retrieve() (Value, error) {
	e.retrieved = false

	v := Value{
		APIKey:    os.Getenv("EXOSCALE_API_KEY"),
		APISecret: os.Getenv("EXOSCALE_API_SECRET"),
	}

	if !v.IsSet() {
		return Value{}, ErrMissingIncomplete
	}

	e.retrieved = true

	return v, nil
}

// IsExpired returns if the credentials have been retrieved.
func (e *EnvProvider) IsExpired() bool {
	return !e.retrieved
}
