package dns

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/ns1/ns1-go.v2/rest/model/data"
)

// AnswerFeed wraps the values of an Answer's "feed" attribute
type AnswerFeed struct {
	FeedID   string `json:"feed"`
	SourceID string `json:"source"`
}

// Answer wraps the values of a Record's "filters" attribute
type Answer struct {
	ID string `json:"id,omitempty"`

	Meta  *data.Meta   `json:"meta,omitempty"`
	Feeds []AnswerFeed `json:"feeds,omitempty"`

	// Answer response data. eg:
	// Av4: ["1.1.1.1"]
	// Av6: ["2001:db8:85a3::8a2e:370:7334"]
	// MX:  [10, "2.2.2.2"]
	Rdata []string `json:"answer"`

	// Region(grouping) that answer belongs to.
	RegionName string `json:"region,omitempty"`
}

// Alias is used as an alias for an answer so that the custom marshaler isn't used.
type Alias struct {
	Rdata []interface{} `json:"answer"`
	*AliasAnswer
}

// AliasAnswer is a duplicate of Answer.
type AliasAnswer Answer

// UnmarshalJSON parses responses to Answer and attempts to convert Rdata
// elements to string.
func (a *Answer) UnmarshalJSON(data []byte) error {
	aux := &Alias{
		AliasAnswer: (*AliasAnswer)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var rdata []string
	for _, record := range aux.Rdata {
		switch v := record.(type) {
		case string:
			rdata = append(rdata, v)
		case float64:
			rdata = append(rdata, strconv.Itoa(int(v)))
		default:
			return fmt.Errorf(
				"Could not unmarshal Rdata value %[1]v (type %[1]T) as type string", v,
			)
		}
	}
	a.Rdata = rdata

	return nil
}

func (a Answer) String() string {
	return strings.Trim(fmt.Sprint(a.Rdata), "[]")
}

// SetRegion associates a region with this answer.
func (a *Answer) SetRegion(name string) {
	a.RegionName = name
}

// NewAnswer creates a generic Answer with given rdata.
func NewAnswer(rdata []string) *Answer {
	return &Answer{
		Meta:  &data.Meta{},
		Rdata: rdata,
	}
}

// NewAv4Answer creates an Answer for A record.
func NewAv4Answer(host string) *Answer {
	return &Answer{
		Meta:  &data.Meta{},
		Rdata: []string{host},
	}
}

// NewAv6Answer creates an Answer for AAAA record.
func NewAv6Answer(host string) *Answer {
	return &Answer{
		Meta:  &data.Meta{},
		Rdata: []string{host},
	}
}

// NewALIASAnswer creates an Answer for ALIAS record.
func NewALIASAnswer(host string) *Answer {
	return &Answer{
		Meta:  &data.Meta{},
		Rdata: []string{host},
	}
}

// NewCNAMEAnswer creates an Answer for CNAME record.
func NewCNAMEAnswer(name string) *Answer {
	return &Answer{
		Meta:  &data.Meta{},
		Rdata: []string{name},
	}
}

// NewTXTAnswer creates an Answer for TXT record.
func NewTXTAnswer(text string) *Answer {
	return &Answer{
		Meta:  &data.Meta{},
		Rdata: []string{text},
	}
}

// NewMXAnswer creates an Answer for MX record.
func NewMXAnswer(pri int, host string) *Answer {
	return &Answer{
		Meta:  &data.Meta{},
		Rdata: []string{strconv.Itoa(pri), host},
	}
}

// NewSRVAnswer creates an Answer for SRV record.
func NewSRVAnswer(priority, weight, port int, target string) *Answer {
	return &Answer{
		Meta: &data.Meta{},
		Rdata: []string{
			strconv.Itoa(priority),
			strconv.Itoa(weight),
			strconv.Itoa(port),
			target,
		},
	}
}

// NewDSAnswer creates an Answer for DS record.
func NewDSAnswer(key string, algorithm string, t string, digest string) *Answer {
	return &Answer{
		Meta: &data.Meta{},
		Rdata: []string{
			key,
			algorithm,
			t,
			digest,
		},
	}
}

// NewCAAAnswer creates an Answer for a CAA record.
func NewCAAAnswer(flag int, tag, value string) *Answer {
	return &Answer{
		Meta:  &data.Meta{},
		Rdata: []string{strconv.Itoa(flag), tag, value},
	}
}

// NewURLFWDAnswer creates an Answer for URLFWD record.
func NewURLFWDAnswer(from, to string, redirectType, pathForwardingMode, queryForwarding int) *Answer {
	return &Answer{
		Meta: &data.Meta{},
		Rdata: []string{
			from,
			to,
			strconv.Itoa(redirectType),
			strconv.Itoa(pathForwardingMode),
			strconv.Itoa(queryForwarding),
		},
	}
}

// return Answer with Rdata as list of interface, with elements of the correct
// type for API.
func prepareURLFWDAnswer(a *Answer) (interface{}, error) {
	if len(a.Rdata) < 5 {
		return nil, errors.New("invalid number of arguments for Rdata")
	}

	redirectType, err := strconv.Atoi(a.Rdata[2])
	if err != nil {
		return nil, err
	}
	pathForwardingMode, err := strconv.Atoi(a.Rdata[3])
	if err != nil {
		return nil, err
	}
	queryForwarding, err := strconv.Atoi(a.Rdata[4])
	if err != nil {
		return nil, err
	}

	prepared := &Alias{
		Rdata: []interface{}{
			a.Rdata[0],
			a.Rdata[1],
			redirectType,
			pathForwardingMode,
			queryForwarding,
		},
		AliasAnswer: (*AliasAnswer)(a),
	}

	return prepared, nil
}
