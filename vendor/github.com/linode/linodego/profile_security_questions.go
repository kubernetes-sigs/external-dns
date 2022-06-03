package linodego

import (
	"context"
	"encoding/json"
)

type SecurityQuestion struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Response string `json:"response"`
}

type SecurityQuestionsListResponse struct {
	SecurityQuestions []SecurityQuestion `json:"security_questions"`
}

type SecurityQuestionsAnswerQuestion struct {
	QuestionID int    `json:"question_id"`
	Response   string `json:"response"`
}

type SecurityQuestionsAnswerOptions struct {
	SecurityQuestions []SecurityQuestionsAnswerQuestion `json:"security_questions"`
}

// SecurityQuestionsList returns a collection of security questions and their responses, if any, for your User Profile.
func (c *Client) SecurityQuestionsList(ctx context.Context) (*SecurityQuestionsListResponse, error) {
	e, err := c.ProfileSecurityQuestions.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&SecurityQuestionsListResponse{})

	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*SecurityQuestionsListResponse), nil
}

// SecurityQuestionsAnswer adds security question responses for your User.
func (c *Client) SecurityQuestionsAnswer(ctx context.Context, opts SecurityQuestionsAnswerOptions) error {
	var body string
	e, err := c.ProfileSecurityQuestions.Endpoint()
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
