package common

type Logger interface {
	Printf(format string, args ...interface{})
}

func (c *Client) WithLogger(logger Logger) *Client {
	c.logger = logger
	return c
}
