package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

// The details and enrollment information of a Beta program that an account is enrolled in.
type AccountBetaProgram struct {
	Label       string     `json:"label"`
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Started     *time.Time `json:"-"`
	Ended       *time.Time `json:"-"`

	// Date the account was enrolled in the beta program
	Enrolled *time.Time `json:"-"`
}

// AccountBetaProgramCreateOpts fields are those accepted by JoinBetaProgram
type AccountBetaProgramCreateOpts struct {
	ID string `json:"id"`
}

// AccountBetasPagedResponse represents a paginated Account Beta Programs API response
type AccountBetasPagedResponse struct {
	*PageOptions
	Data []AccountBetaProgram `json:"data"`
}

// endpoint gets the endpoint URL for AccountBetaProgram
func (AccountBetasPagedResponse) endpoint(_ ...any) string {
	return "/account/betas"
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (cBeta *AccountBetaProgram) UnmarshalJSON(b []byte) error {
	type Mask AccountBetaProgram

	p := struct {
		*Mask
		Started  *parseabletime.ParseableTime `json:"started"`
		Ended    *parseabletime.ParseableTime `json:"ended"`
		Enrolled *parseabletime.ParseableTime `json:"enrolled"`
	}{
		Mask: (*Mask)(cBeta),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	cBeta.Started = (*time.Time)(p.Started)
	cBeta.Ended = (*time.Time)(p.Ended)
	cBeta.Enrolled = (*time.Time)(p.Enrolled)

	return nil
}

func (resp *AccountBetasPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(AccountBetasPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*AccountBetasPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListAccountBetaPrograms lists all beta programs an account is enrolled in.
func (c *Client) ListAccountBetaPrograms(ctx context.Context, opts *ListOptions) ([]AccountBetaProgram, error) {
	response := AccountBetasPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetAccountBetaProgram gets the details of a beta program an account is enrolled in.
func (c *Client) GetAccountBetaProgram(ctx context.Context, betaID string) (*AccountBetaProgram, error) {
	req := c.R(ctx).SetResult(&AccountBetaProgram{})
	betaID = url.PathEscape(betaID)
	b := fmt.Sprintf("/account/betas/%s", betaID)
	r, err := coupleAPIErrors(req.Get(b))
	if err != nil {
		return nil, err
	}

	return r.Result().(*AccountBetaProgram), nil
}

// JoinBetaProgram enrolls an account into a beta program.
func (c *Client) JoinBetaProgram(ctx context.Context, opts AccountBetaProgramCreateOpts) (*AccountBetaProgram, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := "account/betas"
	req := c.R(ctx).SetResult(&AccountBetaProgram{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*AccountBetaProgram), nil
}
