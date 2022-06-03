package connection

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

// Date represents date string from API
type Date string

// Time returns Time struct for DateTime
func (c Date) Time() time.Time {
	t, _ := time.Parse("2006-01-02", c.String())

	return t
}

func (c Date) String() string {
	return string(c)
}

// DateTime represents datetime string from API
type DateTime string

// Time returns Time struct for DateTime
func (c DateTime) Time() time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05-0700", c.String())

	return t
}

func (c DateTime) String() string {
	return string(c)
}

// IPAddress represents ip address string from API
type IPAddress string

func (i IPAddress) IP() net.IP {
	return net.ParseIP(i.String())
}

func (i IPAddress) String() string {
	return string(i)
}

type ErrInvalidEnumValue struct {
	Message string
}

func NewErrInvalidEnumValue(msg string) *ErrInvalidEnumValue {
	return &ErrInvalidEnumValue{Message: msg}
}

func (e *ErrInvalidEnumValue) Error() string {
	return e.Message
}

type Enum interface {
	String() string
}

type EnumSlice []Enum

// ParseEnum parses string s against array of enums, returning parsed enum and nil error, or nil with error
func ParseEnum(s string, enums EnumSlice) (Enum, error) {
	if len(enums) < 1 {
		return nil, errors.New("Must provide at least one enum")
	}

	for _, e := range enums {
		if strings.ToUpper(s) == strings.ToUpper(e.String()) {
			return e, nil
		}
	}

	return nil, NewErrInvalidEnumValue(fmt.Sprintf("Invalid %T. Valid values: %s", enums[0], enums.String()))
}

// StringSlice returns a slice of strings containing the string values of enums for EnumSlice
func (enums EnumSlice) StringSlice() []string {
	var values []string
	for _, enum := range enums {
		values = append(values, enum.String())
	}
	return values
}

// String returns string containing a comma separated list of enum string values
func (enums EnumSlice) String() string {
	return strings.Join(enums.StringSlice(), ", ")
}
