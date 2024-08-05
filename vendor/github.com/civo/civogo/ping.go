package civogo

// Ping checks if Civo API is reachable and responding. Returns no error if API is reachable and running.
func (c *Client) Ping() error {
	url := "/v2/ping"

	_, err := c.SendGetRequest(url)
	if err != nil {
		return decodeError(err)
	}

	return nil
}
