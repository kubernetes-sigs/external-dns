package civogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/mod/semver"
)

// DiskImage represents a DiskImage for launching instances from
type DiskImage struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Version      string `json:"version,omitempty"`
	State        string `json:"state,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	Description  string `json:"description,omitempty"`
	Label        string `json:"label,omitempty"`
}

// ListDiskImages return all disk image in system
func (c *Client) ListDiskImages() ([]DiskImage, error) {
	resp, err := c.SendGetRequest("/v2/disk_images")
	if err != nil {
		return nil, decodeError(err)
	}

	diskImages := make([]DiskImage, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&diskImages); err != nil {
		return nil, err
	}

	filteredDiskImages := make([]DiskImage, 0)
	for _, diskImage := range diskImages {
		if !strings.Contains(diskImage.Name, "k3s") {
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

// GetMostRecentDistro finds the highest version of a specified distro
func (c *Client) GetMostRecentDistro(name string) (*DiskImage, error) {
	resp, err := c.ListDiskImages()
	if err != nil {
		return nil, decodeError(err)
	}

	var highestVersionDistro *DiskImage

	for _, diskimage := range resp {
		if strings.Contains(diskimage.Name, name) {
			if highestVersionDistro == nil {
				highestVersionDistro = &diskimage
			} else {
				if semver.Compare(highestVersionDistro.Version, diskimage.Version) < 0 {
					highestVersionDistro = &diskimage
				}
			}
		}
	}
	if highestVersionDistro == nil {
		return nil, fmt.Errorf("%s image not found", name)
	}

	return highestVersionDistro, nil
}
