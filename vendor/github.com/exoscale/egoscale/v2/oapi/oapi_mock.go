package oapi

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type oapiClientMock struct {
	ClientWithResponsesInterface
	mock.Mock
}

func (m *oapiClientMock) GetOperationWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...RequestEditorFn,
) (*GetOperationResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*GetOperationResponse), args.Error(1)
}
