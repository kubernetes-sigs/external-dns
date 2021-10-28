package pss

import (
	"fmt"
	"io"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetReply retrieves a single reply by id
func (s *Service) GetReply(replyID string) (Reply, error) {
	body, err := s.getReplyResponseBody(replyID)

	return body.Data, err
}

func (s *Service) getReplyResponseBody(replyID string) (*GetReplyResponseBody, error) {
	body := &GetReplyResponseBody{}

	if replyID == "" {
		return body, fmt.Errorf("invalid reply id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/pss/v1/replies/%s", replyID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ReplyNotFoundError{ID: replyID}
		}

		return nil
	})
}

// DownloadReplyAttachmentStream downloads the provided attachment, returning
// a stream of the file contents and an error
func (s *Service) DownloadReplyAttachmentStream(replyID string, attachmentName string) (contentStream io.ReadCloser, err error) {
	response, err := s.downloadReplyAttachmentResponse(replyID, attachmentName)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (s *Service) downloadReplyAttachmentResponse(replyID string, attachmentName string) (*connection.APIResponse, error) {
	body := &connection.APIResponseBody{}
	response := &connection.APIResponse{}

	if replyID == "" {
		return response, fmt.Errorf("invalid reply id")
	}
	if attachmentName == "" {
		return response, fmt.Errorf("invalid attachment name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/pss/v1/replies/%s/attachments/%s", replyID, attachmentName), connection.APIRequestParameters{})
	if err != nil {
		return response, err
	}

	if response.StatusCode == 404 {
		return response, &AttachmentNotFoundError{Name: attachmentName}
	}

	return response, response.ValidateStatusCode([]int{}, body)
}

// UploadReplyAttachmentStream uploads the provided attachment
func (s *Service) UploadReplyAttachmentStream(replyID string, attachmentName string, stream io.Reader) (err error) {
	_, err = s.uploadReplyAttachmentStreamResponseBody(replyID, attachmentName, stream)

	return err
}

func (s *Service) uploadReplyAttachmentStreamResponseBody(replyID string, attachmentName string, stream io.Reader) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if replyID == "" {
		return body, fmt.Errorf("invalid reply id")
	}
	if attachmentName == "" {
		return body, fmt.Errorf("invalid attachment name")
	}
	if stream == nil {
		return body, fmt.Errorf("invalid stream")
	}

	response, err := s.connection.Post(fmt.Sprintf("/pss/v1/replies/%s/attachments/%s", replyID, attachmentName), stream)
	if err != nil {
		return body, err
	}

	if response.StatusCode == 404 {
		return body, &ReplyNotFoundError{ID: replyID}
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteReplyAttachment removes a reply attachment
func (s *Service) DeleteReplyAttachment(replyID string, attachmentName string) error {
	_, err := s.deleteReplyAttachmentResponseBody(replyID, attachmentName)

	return err
}

func (s *Service) deleteReplyAttachmentResponseBody(replyID string, attachmentName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if replyID == "" {
		return body, fmt.Errorf("invalid reply id")
	}
	if attachmentName == "" {
		return body, fmt.Errorf("invalid attachment name")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/pss/v1/replies/%s/attachments/%s", replyID, attachmentName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AttachmentNotFoundError{Name: attachmentName}
		}

		return nil
	})
}
