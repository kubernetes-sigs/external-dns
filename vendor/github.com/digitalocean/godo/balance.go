package godo

import (
	"context"
	"net/http"
	"time"
)

// BalanceService is an interface for interfacing with the Balance
// endpoints of the DigitalOcean API
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// See: https://docs.digitalocean.com/reference/api/api-reference/#operation/get_customer_balance
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2/#balance
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// See: https://developers.digitalocean.com/documentation/v2/#balance
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#operation/get_customer_balance
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2/#balance
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// See: https://developers.digitalocean.com/documentation/v2/#balance
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#operation/get_customer_balance
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2/#balance
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// See: https://developers.digitalocean.com/documentation/v2/#balance
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#operation/get_customer_balance
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2/#balance
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// See: https://developers.digitalocean.com/documentation/v2/#balance
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#operation/balance_get
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
type BalanceService interface {
	Get(context.Context) (*Balance, *Response, error)
}

// BalanceServiceOp handles communication with the Balance related methods of
// the DigitalOcean API.
type BalanceServiceOp struct {
	client *Client
}

var _ BalanceService = &BalanceServiceOp{}

// Balance represents a DigitalOcean Balance
type Balance struct {
	MonthToDateBalance string    `json:"month_to_date_balance"`
	AccountBalance     string    `json:"account_balance"`
	MonthToDateUsage   string    `json:"month_to_date_usage"`
	GeneratedAt        time.Time `json:"generated_at"`
}

func (r Balance) String() string {
	return Stringify(r)
}

// Get DigitalOcean balance info
func (s *BalanceServiceOp) Get(ctx context.Context) (*Balance, *Response, error) {
	path := "v2/customers/my/balance"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Balance)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
