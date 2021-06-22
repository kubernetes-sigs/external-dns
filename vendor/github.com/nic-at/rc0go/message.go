// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"errors"
)

type MessageService service

type Message struct {
	ID 		int   	`json:"id"`
	Domain 	string `json:"domain"`
	Date 	string `json:"date"`
	Type 	string `json:"type"`
	Comment string `json:"comment"`
}

// Retrieves the oldest unacknowledged message from the message queue
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-message-queue-poll-message-get
func (s *MessageService) GetLatest()(*Message, error) {

	resp, err := s.client.NewRequest().
		Get(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0Messages,
		)

	if err != nil {
		return nil, err
	}

	var message *Message

	err = json.Unmarshal(resp.Body(), &message)

	if err != nil {
		return nil, err
	}

	return message, nil

}

// Acknowlegdes (and deletes) the message with the given id
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-message-queue-ack-message-delete
func (s *MessageService) AckAndDelete(id int) (*StatusResponse, error){

	resp, err := s.client.NewRequest().
		Delete(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0AckMessage,
		)

	if resp.StatusCode() == 404 {
		status, err := s.client.ResponseToRC0StatusResponse(resp)

		if err != nil {
			return nil, err
		}

		return nil, errors.New(status.Message)
	}

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)

}