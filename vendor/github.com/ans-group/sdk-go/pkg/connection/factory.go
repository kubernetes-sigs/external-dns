package connection

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"github.com/ans-group/sdk-go/pkg/config"
)

type ConnectionFactory interface {
	NewConnection() (Connection, error)
}

type DefaultConnectionFactoryOption func(f *DefaultConnectionFactory)

type DefaultConnectionFactory struct {
	apiUserAgent string
}

func WithDefaultConnectionUserAgent(userAgent string) DefaultConnectionFactoryOption {
	return func(p *DefaultConnectionFactory) {
		p.apiUserAgent = userAgent
	}
}

func NewDefaultConnectionFactory(opts ...DefaultConnectionFactoryOption) *DefaultConnectionFactory {
	f := &DefaultConnectionFactory{}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *DefaultConnectionFactory) NewConnection() (Connection, error) {
	apiKey := config.GetString("api_key")
	if len(apiKey) < 1 {
		return nil, errors.New("Missing api_key")
	}

	conn := NewAPIConnection(&APIKeyCredentials{APIKey: apiKey})
	conn.UserAgent = f.apiUserAgent
	apiURI := config.GetString("api_uri")
	if apiURI != "" {
		conn.APIURI = apiURI
	}
	apiTimeoutSeconds := config.GetInt("api_timeout_seconds")
	if apiTimeoutSeconds > 0 {
		conn.HTTPClient.Timeout = (time.Duration(apiTimeoutSeconds) * time.Second)
	}
	if config.GetBool("api_insecure") {
		conn.HTTPClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
	apiHeaders := config.GetStringMapString("api_headers")
	if apiHeaders != nil {
		conn.Headers = http.Header{}
		for headerKey, headerValue := range apiHeaders {
			conn.Headers.Add(headerKey, headerValue)
		}
	}

	return conn, nil
}
