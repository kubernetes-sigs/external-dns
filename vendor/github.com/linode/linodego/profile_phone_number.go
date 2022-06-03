package linodego

import (
	"context"
	"encoding/json"
	"fmt"
)

// SendPhoneNumberVerificationCodeOptions fields are those accepted by SendPhoneNumberVerificationCode
type SendPhoneNumberVerificationCodeOptions struct {
	ISOCode     string `json:"iso_code"`
	PhoneNumber string `json:"phone_number"`
}

// VerifyPhoneNumberOptions fields are those accepted by VerifyPhoneNumber
type VerifyPhoneNumberOptions struct {
	OTPCode string `json:"otp_code"`
}

// SendPhoneNumberVerificationCode sends a one-time verification code via SMS message to the submitted phone number.
func (c *Client) SendPhoneNumberVerificationCode(ctx context.Context, opts SendPhoneNumberVerificationCodeOptions) error {
	var body string
	e, err := c.ProfilePhoneNumber.Endpoint()
	if err != nil {
		return err
	}

	req := c.R(ctx)

	if bodyData, err := json.Marshal(opts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	if _, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e)); err != nil {
		return err
	}
	return nil
}

// DeletePhoneNumber deletes the verified phone number for the User making this request.
func (c *Client) DeletePhoneNumber(ctx context.Context) error {
	e, err := c.ProfilePhoneNumber.Endpoint()
	if err != nil {
		return err
	}

	req := c.R(ctx)

	if _, err := coupleAPIErrors(req.
		Delete(e)); err != nil {
		return err
	}
	return nil
}

// VerifyPhoneNumber verifies a phone number by confirming the one-time code received via SMS message after accessing the Phone Verification Code Send command.
func (c *Client) VerifyPhoneNumber(ctx context.Context, opts VerifyPhoneNumberOptions) error {
	var body string
	e, err := c.ProfilePhoneNumber.Endpoint()
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/verify", e)

	req := c.R(ctx)

	if bodyData, err := json.Marshal(opts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	if _, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e)); err != nil {
		return err
	}
	return nil
}
