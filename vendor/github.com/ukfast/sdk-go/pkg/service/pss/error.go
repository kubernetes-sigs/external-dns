package pss

import "fmt"

// RequestNotFoundError indicates a request was not found
type RequestNotFoundError struct {
	ID int
}

func (e *RequestNotFoundError) Error() string {
	return fmt.Sprintf("Request not found with id [%d]", e.ID)
}

// ReplyNotFoundError indicates a reply was not found
type ReplyNotFoundError struct {
	ID string
}

func (e *ReplyNotFoundError) Error() string {
	return fmt.Sprintf("Reply not found with id [%s]", e.ID)
}

// AttachmentNotFoundError indicates a attachment was not found
type AttachmentNotFoundError struct {
	Name string
}

func (e *AttachmentNotFoundError) Error() string {
	return fmt.Sprintf("Attachment not found with name [%s]", e.Name)
}

// RequestFeedbackNotFoundError indicates feedback for a request was not found
type RequestFeedbackNotFoundError struct {
	RequestID int
}

func (e *RequestFeedbackNotFoundError) Error() string {
	return fmt.Sprintf("Feedback not found for request [%d]", e.RequestID)
}
