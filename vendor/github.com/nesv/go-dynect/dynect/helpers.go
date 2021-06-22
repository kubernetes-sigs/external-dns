package dynect

import "fmt"

func GetAllDSFServicesDetailed(c *Client) (error, []DSFService) {
	var dsfsResponse AllDSFDetailedResponse
	requestData := struct {
		Detail string `json:"detail"`
	}{Detail: "Y"}

	if err := c.Do("GET", "DSF", requestData, &dsfsResponse); err != nil {
		return err, nil
	}

	return nil, dsfsResponse.Data
}

func GetDSFServiceDetailed(c *Client, id string) (error, DSFService) {
	var dsfsResponse DSFResponse
	requestData := struct {
		Detail string `json:"detail"`
	}{Detail: "Y"}

	loc := fmt.Sprintf("DSF/%s", id)

	if err := c.Do("GET", loc, requestData, &dsfsResponse); err != nil {
		return err, DSFService{}
	}
	return nil, dsfsResponse.Data
}
