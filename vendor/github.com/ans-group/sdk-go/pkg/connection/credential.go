package connection

// AuthHeaders is a string map of authorization headers
type AuthHeaders map[string]string

type Credentials interface {
	GetAuthHeaders() AuthHeaders
}

type APIKeyCredentials struct {
	APIKey string
}

// GetAuthHeaders returns the Authorization header for API key
func (c *APIKeyCredentials) GetAuthHeaders() AuthHeaders {
	h := make(map[string]string)
	h["Authorization"] = c.APIKey

	return h
}
