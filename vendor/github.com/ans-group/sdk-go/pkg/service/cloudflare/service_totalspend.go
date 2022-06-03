package cloudflare

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTotalSpendMonthToDate retrieves a total spend for current month
func (s *Service) GetTotalSpendMonthToDate() (TotalSpend, error) {
	body, err := s.getTotalSpendMonthToDateResponseBody()

	return body.Data, err
}

func (s *Service) getTotalSpendMonthToDateResponseBody() (*connection.APIResponseBodyData[TotalSpend], error) {
	body := &connection.APIResponseBodyData[TotalSpend]{}
	response, err := s.connection.Get("/cloudflare/v1/total-spend/month-to-date", connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
