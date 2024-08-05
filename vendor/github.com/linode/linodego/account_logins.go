package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

type Login struct {
	ID         int        `json:"id"`
	Datetime   *time.Time `json:"datetime"`
	IP         string     `json:"ip"`
	Restricted bool       `json:"restricted"`
	Username   string     `json:"username"`
	Status     string     `json:"status"`
}

type LoginsPagedResponse struct {
	*PageOptions
	Data []Login `json:"data"`
}

func (LoginsPagedResponse) endpoint(_ ...any) string {
	return "account/logins"
}

func (resp *LoginsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(LoginsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*LoginsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

func (c *Client) ListLogins(ctx context.Context, opts *ListOptions) ([]Login, error) {
	response := LoginsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *Login) UnmarshalJSON(b []byte) error {
	type Mask Login

	l := struct {
		*Mask
		Datetime *parseabletime.ParseableTime `json:"datetime"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &l); err != nil {
		return err
	}

	i.Datetime = (*time.Time)(l.Datetime)

	return nil
}

func (c *Client) GetLogin(ctx context.Context, loginID int) (*Login, error) {
	req := c.R(ctx).SetResult(&Login{})
	e := fmt.Sprintf("account/logins/%d", loginID)
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*Login), nil
}
