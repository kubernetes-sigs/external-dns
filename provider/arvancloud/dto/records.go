/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dto

import (
	"bytes"
	"strconv"
	"strings"
)

type Uint16 uint16

func (u *Uint16) UnmarshalJSON(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		return nil
	}
	l := len(text)
	if l > 0 && text[0] == '"' && text[l-1] == '"' {
		text = bytes.Trim(text, `"`)
	}

	byeToInt, err := strconv.Atoi(string(text))
	if err != nil {
		return err
	}
	*u = Uint16(byeToInt)

	return nil
}

var _ RecordMarshaler = (*ARecord)(nil)
var _ RecordUnmarshaler = (*ARecord)(nil)

type ARecord struct {
	IP      string `json:"ip"`
	Country string `json:"country,omitempty"`
	Port    Uint16 `json:"port,omitempty"`
	Weight  Uint16 `json:"weight,omitempty"`
}

func (r *ARecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	enc.WriteString(r.IP)

	return enc.Bytes(), nil
}

func (r *ARecord) UnmarshalRec(data []byte) error {
	r.IP = string(data)

	return nil
}

var _ RecordMarshaler = (*AAAARecord)(nil)
var _ RecordUnmarshaler = (*AAAARecord)(nil)

type AAAARecord ARecord

func (r *AAAARecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	enc.WriteString(r.IP)

	return enc.Bytes(), nil
}

func (r *AAAARecord) UnmarshalRec(data []byte) error {
	r.IP = string(data)

	return nil
}

var _ RecordMarshaler = (*NSRecord)(nil)
var _ RecordUnmarshaler = (*NSRecord)(nil)

type NSRecord struct {
	Host string `json:"host"`
}

func (r *NSRecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	enc.WriteString(r.Host)

	return enc.Bytes(), nil
}

func (r *NSRecord) UnmarshalRec(data []byte) error {
	r.Host = string(data)

	return nil
}

var _ RecordMarshaler = (*TXTRecord)(nil)
var _ RecordUnmarshaler = (*TXTRecord)(nil)

type TXTRecord struct {
	Text string `json:"text"`
}

func (r *TXTRecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	enc.WriteString(r.Text)

	return enc.Bytes(), nil
}

func (r *TXTRecord) UnmarshalRec(data []byte) error {
	l := len(data)
	var in []byte

	if l > 0 && data[0] == '"' && data[l-1] == '"' {
		in = make([]byte, l-2, cap(data)-2)
		copy(in, data[1:l-1])
	} else {
		in = data
	}

	r.Text = string(in)

	return nil
}

var _ RecordMarshaler = (*CNAMERecord)(nil)
var _ RecordUnmarshaler = (*CNAMERecord)(nil)

type CNAMERecord struct {
	Host       string `json:"host"`
	HostHeader string `json:"host_header"`
	Port       Uint16 `json:"port,omitempty"`
}

func (r *CNAMERecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	enc.WriteString(r.Host)

	return enc.Bytes(), nil
}

func (r *CNAMERecord) UnmarshalRec(data []byte) error {
	r.Host = string(data)

	return nil
}

var _ RecordMarshaler = (*ANAMERecord)(nil)
var _ RecordUnmarshaler = (*ANAMERecord)(nil)

type ANAMERecord struct {
	Location   string `json:"location"`
	HostHeader string `json:"host_header"`
	Port       Uint16 `json:"port,omitempty"`
}

func (r *ANAMERecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	enc.WriteString(r.Location)

	return enc.Bytes(), nil
}

func (r *ANAMERecord) UnmarshalRec(data []byte) error {
	r.Location = string(data)

	return nil
}

var _ RecordMarshaler = (*MXRecord)(nil)
var _ RecordUnmarshaler = (*MXRecord)(nil)

type MXRecord struct {
	Host     string `json:"host"`
	Priority Uint16 `json:"priority"`
}

func (r *MXRecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	wl := [...]string{strconv.Itoa(int(r.Priority)), " ", r.Host}
	for _, w := range wl {
		enc.WriteString(w)
	}

	return enc.Bytes(), nil
}

func (r *MXRecord) UnmarshalRec(data []byte) error {
	record := string(data)
	values := strings.Split(record, " ")
	if len(values) != 2 {
		return NewParseMXRecordError()
	}

	priority, err := strconv.ParseUint(values[0], 10, 16)
	if err != nil {
		return err
	}

	r.Priority = Uint16(priority)
	r.Host = values[1]

	return nil
}

var _ RecordMarshaler = (*SRVRecord)(nil)
var _ RecordUnmarshaler = (*SRVRecord)(nil)

type SRVRecord struct {
	Target   string `json:"target"`
	Port     Uint16 `json:"port"`
	Weight   Uint16 `json:"weight,omitempty"`
	Priority Uint16 `json:"priority,omitempty"`
}

func (r *SRVRecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	wl := [...]string{strconv.Itoa(int(r.Priority)), " ", strconv.Itoa(int(r.Weight)), " ", strconv.Itoa(int(r.Port)), " ", r.Target}
	for _, w := range wl {
		enc.WriteString(w)
	}

	return enc.Bytes(), nil
}

func (r *SRVRecord) UnmarshalRec(data []byte) error {
	record := string(data)
	values := strings.Split(record, " ")
	if len(values) != 4 {
		return NewParseSRVRecordError()
	}

	priority, err := strconv.ParseUint(values[0], 10, 16)
	if err != nil {
		return err
	}
	weight, err := strconv.ParseUint(values[1], 10, 16)
	if err != nil {
		return err
	}
	port, err := strconv.ParseUint(values[2], 10, 16)
	if err != nil {
		return err
	}

	r.Priority = Uint16(priority)
	r.Weight = Uint16(weight)
	r.Port = Uint16(port)
	r.Target = values[3]

	return nil
}

var _ RecordMarshaler = (*SPFRecord)(nil)
var _ RecordUnmarshaler = (*SPFRecord)(nil)

type SPFRecord TXTRecord

func (r *SPFRecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	enc.WriteString(r.Text)

	return enc.Bytes(), nil
}

func (r *SPFRecord) UnmarshalRec(data []byte) error {
	r.Text = string(data)

	return nil
}

var _ RecordMarshaler = (*DKIMRecord)(nil)
var _ RecordUnmarshaler = (*DKIMRecord)(nil)

type DKIMRecord TXTRecord

func (r *DKIMRecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	enc.WriteString(r.Text)

	return enc.Bytes(), nil
}

func (r *DKIMRecord) UnmarshalRec(data []byte) error {
	r.Text = string(data)

	return nil
}

var _ RecordMarshaler = (*PTRRecord)(nil)
var _ RecordUnmarshaler = (*PTRRecord)(nil)

type PTRRecord struct {
	Domain string `json:"domain"`
}

func (r *PTRRecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	enc.WriteString(r.Domain)

	return enc.Bytes(), nil
}

func (r *PTRRecord) UnmarshalRec(data []byte) error {
	r.Domain = string(data)

	return nil
}

var _ RecordMarshaler = (*TLSARecord)(nil)
var _ RecordUnmarshaler = (*TLSARecord)(nil)

type TLSARecord struct {
	Usage        string `json:"usage"`
	Selector     string `json:"selector"`
	MatchingType string `json:"matching_type"`
	Certificate  string `json:"certificate,omitempty"`
}

func (r *TLSARecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	wl := [...]string{r.Usage, " ", r.Selector, " ", r.MatchingType, " ", r.Certificate}
	for _, w := range wl {
		enc.WriteString(w)
	}

	return enc.Bytes(), nil
}

func (r *TLSARecord) UnmarshalRec(data []byte) error {
	record := string(data)
	values := strings.Split(record, " ")
	if len(values) == 3 {
		r.Usage = values[0]
		r.Selector = values[1]
		r.MatchingType = values[2]

		return nil
	}
	if len(values) == 4 {
		r.Usage = values[0]
		r.Selector = values[1]
		r.MatchingType = values[2]
		r.Certificate = values[3]

		return nil
	}

	return NewParseTLSARecordError()
}

var _ RecordMarshaler = (*CAARecord)(nil)
var _ RecordUnmarshaler = (*CAARecord)(nil)

type CAARecord struct {
	Value string `json:"value"`
	Tag   string `json:"tag"`
}

func (r *CAARecord) MarshalRec() ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	var enc bytes.Buffer
	wl := [...]string{r.Tag, " ", r.Value}
	for _, w := range wl {
		enc.WriteString(w)
	}

	return enc.Bytes(), nil
}

func (r *CAARecord) UnmarshalRec(data []byte) error {
	record := string(data)
	values := strings.Split(record, " ")
	if len(values) != 2 {
		return NewCAARecordError()
	}

	r.Tag = values[0]
	r.Value = values[1]

	return nil
}
