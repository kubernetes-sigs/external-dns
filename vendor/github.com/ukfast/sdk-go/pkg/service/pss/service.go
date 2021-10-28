package pss

import (
	"io"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// PSSService is an interface for managing PSS
type PSSService interface {
	CreateRequest(req CreateRequestRequest) (int, error)
	GetRequests(parameters connection.APIRequestParameters) ([]Request, error)
	GetRequestsPaginated(parameters connection.APIRequestParameters) (*PaginatedRequest, error)
	GetRequest(requestID int) (Request, error)
	PatchRequest(requestID int, req PatchRequestRequest) error

	GetRequestFeedback(requestID int) (Feedback, error)
	CreateRequestFeedback(requestID int, req CreateFeedbackRequest) (int, error)

	CreateRequestReply(requestID int, req CreateReplyRequest) (string, error)
	GetRequestReplies(solutionID int, parameters connection.APIRequestParameters) ([]Reply, error)
	GetRequestRepliesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedReply, error)
	GetRequestConversation(requestID int, parameters connection.APIRequestParameters) ([]Reply, error)
	GetRequestConversationPaginated(requestID int, parameters connection.APIRequestParameters) (*PaginatedReply, error)

	GetReply(replyID string) (Reply, error)

	DownloadReplyAttachmentStream(replyID string, attachmentName string) (contentStream io.ReadCloser, err error)
	UploadReplyAttachmentStream(replyID string, attachmentName string, fileStream io.Reader) (err error)
	DeleteReplyAttachment(replyID string, attachmentName string) error
}

// Service implements PSSService for managing
// PSS via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of PSSService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
