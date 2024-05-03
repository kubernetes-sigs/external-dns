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
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

type Zone struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	TargetCname string    `json:"target_cname"`
	CustomCname string    `json:"custom_cname"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	OriginalNS  []string  `json:"ns_keys"`
	NameServers []string  `json:"current_ns"`
	Plan        int8      `json:"plan_level"`
}

type DnsType string

const (
	AType     DnsType = "A"
	AAAAType  DnsType = "AAAA"
	NSType    DnsType = "NS"
	TXTType   DnsType = "TXT"
	CNAMEType DnsType = "CNAME"
	ANAMEType DnsType = "ANAME"
	MXType    DnsType = "MX"
	SRVType   DnsType = "SRV"
	SPFType   DnsType = "SPF"
	DKIMType  DnsType = "DKIM"
	PTRType   DnsType = "PTR"
	TLSAType  DnsType = "TLSA"
	CAAType   DnsType = "CAA"
)

type RecordMarshaler interface {
	MarshalRec() ([]byte, error)
}

type RecordUnmarshaler interface {
	UnmarshalRec(data []byte) error
}

type DnsRecord struct {
	CreatedAt   time.Time   `json:"created_at"`
	UpdateAt    time.Time   `json:"updated_at"`
	ID          string      `json:"id"`
	Zone        string      `json:"-"`
	Name        string      `json:"name"`
	Type        DnsType     `json:"type"`
	Value       interface{} `json:"value"`
	Contents    []string    `json:"-"`
	TTL         int64       `json:"ttl"`
	Cloud       bool        `json:"cloud"`
	IsProtected bool        `json:"is_protected"`
}

type aliasDnsRecord DnsRecord

type parseDnsRecord struct {
	*aliasDnsRecord
	Value *json.RawMessage `json:"value"`
	Type  string           `json:"type"`
}

func (d *DnsRecord) UnmarshalJSON(data []byte) error {
	ns := &parseDnsRecord{
		aliasDnsRecord: (*aliasDnsRecord)(d),
	}
	if err := json.Unmarshal(data, &ns); err != nil {
		return err
	}

	d.Type = DnsType(strings.ToUpper(ns.Type))

	if ns.Value == nil || (ns.Value != nil && len(*ns.Value) == 0) {
		return nil
	}

	switch d.Type {
	case AType:
		var value []ARecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		d.Contents = make([]string, 0, len(value))
		for _, v := range value {
			content, err := v.MarshalRec()
			if err != nil {
				return err
			}
			d.Contents = append(d.Contents, string(content))
		}
	case AAAAType:
		var value []AAAARecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		d.Contents = make([]string, 0, len(value))
		for _, v := range value {
			content, err := v.MarshalRec()
			if err != nil {
				return err
			}
			d.Contents = append(d.Contents, string(content))
		}
	case NSType:
		var value NSRecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case TXTType:
		var value TXTRecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case CNAMEType:
		var value CNAMERecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case ANAMEType:
		var value ANAMERecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case MXType:
		var value MXRecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case SRVType:
		var value SRVRecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case SPFType:
		var value SPFRecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case DKIMType:
		var value DKIMRecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case PTRType:
		var value PTRRecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case TLSAType:
		var value TLSARecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	case CAAType:
		var value CAARecord
		if err := d.valueParse(ns, &value); err != nil {
			return err
		}

		content, err := value.MarshalRec()
		if err != nil {
			return err
		}
		d.Contents = []string{string(content)}
	default:
		return NewParseRecordError()
	}

	return nil
}

func (d DnsRecord) MarshalJSON() ([]byte, error) {
	if d.Value == nil && len(d.Contents) > 0 {
		d.Value = nil

		switch d.Type {
		case AType:
			var data []interface{}
			for _, v := range d.Contents {
				var val ARecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = append(data, val)
			}
			d.Value = data
		case AAAAType:
			var data []interface{}
			for _, v := range d.Contents {
				var val AAAARecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = append(data, val)
			}
			d.Value = data
		case NSType:
			var data interface{}
			for _, v := range d.Contents {
				var val NSRecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case TXTType:
			var data interface{}
			for _, v := range d.Contents {
				var val TXTRecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case CNAMEType:
			var data interface{}
			for _, v := range d.Contents {
				var val CNAMERecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case ANAMEType:
			var data interface{}
			for _, v := range d.Contents {
				var val ANAMERecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case MXType:
			var data interface{}
			for _, v := range d.Contents {
				var val MXRecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case SRVType:
			var data interface{}
			for _, v := range d.Contents {
				var val SRVRecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case SPFType:
			var data interface{}
			for _, v := range d.Contents {
				var val SPFRecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case DKIMType:
			var data interface{}
			for _, v := range d.Contents {
				var val DKIMRecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case PTRType:
			var data interface{}
			for _, v := range d.Contents {
				var val PTRRecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case TLSAType:
			var data interface{}
			for _, v := range d.Contents {
				var val TLSARecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		case CAAType:
			var data interface{}
			for _, v := range d.Contents {
				var val CAARecord
				if err := val.UnmarshalRec([]byte(v)); err != nil {
					return nil, err
				}
				data = val
			}
			d.Value = data
		}
	}

	return json.Marshal(&struct{ *aliasDnsRecord }{aliasDnsRecord: (*aliasDnsRecord)(&d)})
}

func (d *DnsRecord) valueParse(ns *parseDnsRecord, valuePtr interface{}) error {
	vpt := reflect.TypeOf(valuePtr)
	if vpt.Kind() != reflect.Pointer {
		return NewNonPointerError()
	}
	if err := json.Unmarshal(*ns.Value, valuePtr); err != nil {
		return err
	}

	d.Value = reflect.ValueOf(valuePtr).Elem().Interface()

	return nil
}

type DnsChangeAction string

const (
	CreateDns DnsChangeAction = "CREATE"
	UpdateDns DnsChangeAction = "UPDATE"
	DeleteDns DnsChangeAction = "DELETE"
)

type DnsChange struct {
	Action DnsChangeAction
	Record DnsRecord
}
