package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetImages retrieves a list of images
func (s *Service) GetImages(parameters connection.APIRequestParameters) ([]Image, error) {
	var images []Image

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetImagesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, image := range response.(*PaginatedImage).Items {
			images = append(images, image)
		}
	}

	return images, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetImagesPaginated retrieves a paginated list of images
func (s *Service) GetImagesPaginated(parameters connection.APIRequestParameters) (*PaginatedImage, error) {
	body, err := s.getImagesPaginatedResponseBody(parameters)

	return NewPaginatedImage(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetImagesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getImagesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetImageSliceResponseBody, error) {
	body := &GetImageSliceResponseBody{}

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

func (s *Service) getImageResponseBody(imageID string) (*GetImageResponseBody, error) {
	body := &GetImageResponseBody{}

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

// GetImageParameters retrieves a list of parameters
func (s *Service) GetImageParameters(imageID string, parameters connection.APIRequestParameters) ([]ImageParameter, error) {
	var appParameters []ImageParameter

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetImageParametersPaginated(imageID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, parameter := range response.(*PaginatedImageParameter).Items {
			appParameters = append(appParameters, parameter)
		}
	}

	return appParameters, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetImageParametersPaginated retrieves a paginated list of domains
func (s *Service) GetImageParametersPaginated(imageID string, parameters connection.APIRequestParameters) (*PaginatedImageParameter, error) {
	body, err := s.getImageParametersPaginatedResponseBody(imageID, parameters)

	return NewPaginatedImageParameter(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetImageParametersPaginated(imageID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getImageParametersPaginatedResponseBody(imageID string, parameters connection.APIRequestParameters) (*GetImageParameterSliceResponseBody, error) {
	body := &GetImageParameterSliceResponseBody{}

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
	var appMetadata []ImageMetadata

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetImageMetadataPaginated(imageID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, parameter := range response.(*PaginatedImageMetadata).Items {
			appMetadata = append(appMetadata, parameter)
		}
	}

	return appMetadata, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetImageMetadataPaginated retrieves a paginated list of domains
func (s *Service) GetImageMetadataPaginated(imageID string, parameters connection.APIRequestParameters) (*PaginatedImageMetadata, error) {
	body, err := s.getImageMetadataPaginatedResponseBody(imageID, parameters)

	return NewPaginatedImageMetadata(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetImageMetadataPaginated(imageID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getImageMetadataPaginatedResponseBody(imageID string, parameters connection.APIRequestParameters) (*GetImageMetadataSliceResponseBody, error) {
	body := &GetImageMetadataSliceResponseBody{}

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
