// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"strings"
)

type RRSetService service

type RRSetServiceInterface interface {
	List(zone string, options *ListOptions) ([]*RRType, *Page, error)
	Create(zone string, rrsetCreate []*RRSetChange) (*StatusResponse, error)
	Edit(zone string, rrsetEdit []*RRSetChange) (*StatusResponse, error)
	Delete(zone string, rrsetDelete []*RRSetChange) (*StatusResponse, error)
	SubmitChangeSet(zone string, changeSet []*RRSetChange) (*StatusResponse, error)
	EncryptTXT(key []byte, rrType *RRSetChange)
	DecryptTXT(key []byte, rrType *RRType)
}

type RRType struct {
	Name    string    `json:"name, omitempty"`
	Type    string    `json:"type, omitempty"`
	TTL     int       `json:"ttl, omitempty"`
	Records []*Record `json:"records, omitempty"`
}

type Record struct {
	Content  string `json:"content, omitempty"`
	Disabled bool   `json:"disabled, omitempty"`
}

type RRSetChange struct {
	Name 		string    `json:"name, omitempty"`
	Type 		string    `json:"type, omitempty"`
	ChangeType  	string    `json:"changetype, omitempty"`
	TTL     	int       `json:"ttl,omitempty"`
	Records 	[]*Record `json:"records, omitempty"`
}

const (
	ChangeTypeADD 	 = "add"
	ChangeTypeUPDATE = "update"
	ChangeTypeDELETE = "delete"
)

// List all RRSets
//
// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-rrsets-get
func (s *RRSetService) List(zone string, options *ListOptions) ([]*RRType, *Page, error) {

	resp, err := s.client.NewRequest().
		SetQueryParam("page_size",	options.PageSizeAsString()).
		SetQueryParam("page", 		options.PageNumberAsString()).
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		Get(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0ZoneRRSets,
		)

	if err != nil {
		return nil, nil, err
	}

	var page *Page

	err = json.Unmarshal(resp.Body(), &page)

	if err != nil {
		return nil, nil, err
	}

	var rrset []*RRType

	err = mapstructure.WeakDecode(page.Data, &rrset)
	if err != nil {
		return nil, nil, err
	}

	return rrset, page, nil
}

func (s *RRSetService) Create(zone string, rrsetCreate []*RRSetChange) (*StatusResponse, error) {

	return s.SubmitChangeSet(zone, rrsetCreate)
}


func (s *RRSetService) Edit(zone string, rrsetEdit []*RRSetChange) (*StatusResponse, error) {

	return s.SubmitChangeSet(zone, rrsetEdit)
}

func (s *RRSetService) Delete(zone string, rrsetDelete []*RRSetChange) (*StatusResponse, error) {

		return s.SubmitChangeSet(zone, rrsetDelete)
}

func (s *RRSetService) SubmitChangeSet(zone string, changeSet []*RRSetChange) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		SetBody(changeSet).
		Patch(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0ZoneRRSets,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)
}

func (s *RRSetService) EncryptTXT(key []byte, rrType *RRSetChange) {

	for _, c := range rrType.Records {
		c.Content = "\"ENC:"+encrypt(key, c.Content) +"\""
	}

}

func (s *RRSetService) DecryptTXT(key []byte, rrType *RRType) {

	for _, c := range rrType.Records {
		if strings.HasPrefix(c.Content, "\"ENC:") {
			c.Content = decrypt(key, strings.TrimSuffix(strings.TrimPrefix(c.Content, "\"ENC:"), "\""))
		}
	}

}

// https://gist.github.com/manishtpatel/8222606
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// https://gist.github.com/manishtpatel/8222606
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

