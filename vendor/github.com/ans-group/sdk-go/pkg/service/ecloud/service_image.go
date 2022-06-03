package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetImages retrieves a list of images
func (s *Service) GetImages(parameters connection.APIRequestParameters) ([]Image, error) {
	return connection.InvokeRequestAll(s.GetImagesPaginated, parameters)
}

// GetImagesPaginated retrieves a paginated list of images
func (s *Service) GetImagesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Image], error) {
	body, err := s.getImagesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetImagesPaginated), err
}

func (s *Service) getImagesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Image], error) {
	body := &connection.APIResponseBodyData[[]Image]{}

	response, err := s.connection.Get("/ecloud/v2/images", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetImage retrieves a single Image by ID
func (s *Service) GetImage(imageID string) (Image, error) {
	body, err := s.getImageResponseBody(imageID)

	return body.Data, err
}

func (s *Service) getImageResponseBody(imageID string) (*connection.APIResponseBodyData[Image], error) {
	body := &connection.APIResponseBodyData[Image]{}

	if imageID == "" {
		return body, fmt.Errorf("invalid image id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/images/%s", imageID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ImageNotFoundError{ID: imageID}
		}

		return nil
	})
}

// UpdateImage removes a single Image by ID
func (s *Service) UpdateImage(imageID string, req UpdateImageRequest) (TaskReference, error) {
	body, err := s.updateImageResponseBody(imageID, req)

	return body.Data, err
}

func (s *Service) updateImageResponseBody(imageID string, req UpdateImageRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	if imageID == "" {
		return &connection.APIResponseBodyData[TaskReference]{}, fmt.Errorf("invalid image id")
	}

	return connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/images/%s", imageID), &req, connection.NotFoundResponseHandler(&ImageNotFoundError{ID: imageID}))
}

// DeleteImage removes a single Image by ID
func (s *Service) DeleteImage(imageID string) (string, error) {
	body, err := s.deleteImageResponseBody(imageID)

	return body.Data.TaskID, err
}

func (s *Service) deleteImageResponseBody(imageID string) (*connection.APIResponseBodyData[TaskReference], error) {
	if imageID == "" {
		return &connection.APIResponseBodyData[TaskReference]{}, fmt.Errorf("invalid image id")
	}

	return connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/images/%s", imageID), nil, connection.NotFoundResponseHandler(&ImageNotFoundError{ID: imageID}))
}

// GetImageParameters retrieves a list of parameters
func (s *Service) GetImageParameters(imageID string, parameters connection.APIRequestParameters) ([]ImageParameter, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ImageParameter], error) {
		return s.GetImageParametersPaginated(imageID, p)
	}, parameters)
}

// GetImageParametersPaginated retrieves a paginated list of domains
func (s *Service) GetImageParametersPaginated(imageID string, parameters connection.APIRequestParameters) (*connection.Paginated[ImageParameter], error) {
	body, err := s.getImageParametersPaginatedResponseBody(imageID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ImageParameter], error) {
		return s.GetImageParametersPaginated(imageID, p)
	}), err
}

func (s *Service) getImageParametersPaginatedResponseBody(imageID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ImageParameter], error) {
	body := &connection.APIResponseBodyData[[]ImageParameter]{}

	if imageID == "" {
		return body, fmt.Errorf("invalid image id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/images/%s/parameters", imageID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ImageNotFoundError{ID: imageID}
		}

		return nil
	})
}

// GetImageMetadata retrieves a list of metadata
func (s *Service) GetImageMetadata(imageID string, parameters connection.APIRequestParameters) ([]ImageMetadata, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ImageMetadata], error) {
		return s.GetImageMetadataPaginated(imageID, p)
	}, parameters)
}

// GetImageMetadataPaginated retrieves a paginated list of domains
func (s *Service) GetImageMetadataPaginated(imageID string, parameters connection.APIRequestParameters) (*connection.Paginated[ImageMetadata], error) {
	body, err := s.getImageMetadataPaginatedResponseBody(imageID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ImageMetadata], error) {
		return s.GetImageMetadataPaginated(imageID, p)
	}), err
}

func (s *Service) getImageMetadataPaginatedResponseBody(imageID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ImageMetadata], error) {
	body := &connection.APIResponseBodyData[[]ImageMetadata]{}

	if imageID == "" {
		return body, fmt.Errorf("invalid image id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/images/%s/metadata", imageID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ImageNotFoundError{ID: imageID}
		}

		return nil
	})
}
