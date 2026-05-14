package civogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// DiskImage represents a serialized structure
type DiskImage struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	Version             string    `json:"version"`
	State               string    `json:"state"`
	InitialUser         string    `json:"initial_user,omitempty"`
	Distribution        string    `json:"distribution"`
	OS                  string    `json:"os,omitempty"`
	Description         string    `json:"description"`
	Label               string    `json:"label"`
	DiskImageURL        string    `json:"disk_image_url,omitempty"`
	DiskImageSizeBytes  int64     `json:"disk_image_size_bytes,omitempty"`
	LogoURL             string    `json:"logo_url,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	CreatedBy           string    `json:"created_by,omitempty"` // User information (because multiple users can operate under the same account)
	DistributionDefault bool      `json:"distribution_default"`
}

// CreateDiskImageParams represents the parameters for creating a new disk image
type CreateDiskImageParams struct {
	Name           string `json:"name"`
	Distribution   string `json:"distribution"`
	Version        string `json:"version"`
	Source         string `json:"source"`
	OS             string `json:"os,omitempty"`
	InitialUser    string `json:"initial_user,omitempty"`
	Region         string `json:"region,omitempty"`
	ImageSHA256    string `json:"image_sha256"`
	ImageMD5       string `json:"image_md5"`
	LogoBase64     string `json:"logo_base64,omitempty"`
	ImageSizeBytes int64  `json:"image_size_bytes"` // Size of the image in bytes
}

// CreateDiskImageResponse represents the response from creating a new disk image
type CreateDiskImageResponse struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	Distribution        string    `json:"distribution"`
	Version             string    `json:"version"`
	OS                  string    `json:"os"`
	Region              string    `json:"region"`
	Status              string    `json:"status"`
	InitialUser         string    `json:"initial_user,omitempty"`
	DiskImageURL        string    `json:"disk_image_url"`
	DiskImageSizeBytes  int64     `json:"disk_image_size_bytes,omitempty"`
	LogoURL             string    `json:"logo_url"`
	ImageSize           int64     `json:"image_size"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	CreatedBy           string    `json:"created_by,omitempty"`
	DistributionDefault bool      `json:"distribution_default,omitempty"`
}

// ListDiskImages return all disk image in system
// includeCustom when true will also return custom images (default: false)
func (c *Client) ListDiskImages(includeCustom ...bool) ([]DiskImage, error) {
	includeCustomFlag := false
	if len(includeCustom) > 0 {
		includeCustomFlag = includeCustom[0]
	}

	url := "/v2/disk_images"
	if includeCustomFlag {
		url += "?type=custom"
	}

	resp, err := c.SendGetRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	diskImages := make([]DiskImage, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&diskImages); err != nil {
		return nil, err
	}

	filteredDiskImages := make([]DiskImage, 0)
	for _, diskImage := range diskImages {
		if !strings.Contains(diskImage.Name, "k3s") && !strings.Contains(diskImage.Name, "talos") {
			filteredDiskImages = append(filteredDiskImages, diskImage)
		}
	}

	return filteredDiskImages, nil
}

// GetDiskImage get one disk image using the id
func (c *Client) GetDiskImage(id string) (*DiskImage, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/disk_images/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	diskImage := &DiskImage{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&diskImage); err != nil {
		return nil, err
	}

	return diskImage, nil
}

// FindDiskImage finds a disk image by either part of the ID or part of the name
func (c *Client) FindDiskImage(search string) (*DiskImage, error) {
	templateList, err := c.ListDiskImages()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := DiskImage{}

	for _, value := range templateList {
		if value.Name == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.Name, search) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// GetDiskImageByName finds the DiskImage for an account with the specified code
func (c *Client) GetDiskImageByName(name string) (*DiskImage, error) {
	resp, err := c.ListDiskImages()
	if err != nil {
		return nil, decodeError(err)
	}

	for _, diskimage := range resp {
		if diskimage.Name == name {
			return &diskimage, nil
		}
	}

	return nil, errors.New("diskimage not found")
}

// CreateDiskImage creates a new disk image entry and returns a pre-signed URL for uploading
func (c *Client) CreateDiskImage(params *CreateDiskImageParams) (*CreateDiskImageResponse, error) {
	url := "/v2/disk_images"
	resp, err := c.SendPostRequest(url, params)

	if err != nil {
		return nil, decodeError(err)
	}

	diskImage := &CreateDiskImageResponse{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&diskImage); err != nil {
		return nil, err
	}

	return diskImage, nil
}

// DeleteDiskImage deletes a disk image by its ID
func (c *Client) DeleteDiskImage(id string) error {
	_, err := c.SendDeleteRequest(fmt.Sprintf("/v2/disk_images/%s", id))
	if err != nil {
		return decodeError(err)
	}

	return nil
}
