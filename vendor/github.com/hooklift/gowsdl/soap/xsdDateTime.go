// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package soap

import (
	"encoding/xml"
	"strings"
	"time"
)

const (
	dateLayout = "2006-01-02Z07:00"
	timeLayout = "15:04:05.999999999Z07:00"
)

//
// DateTime struct
//

// XSDDateTime is a type for representing xsd:datetime in Golang
type XSDDateTime struct {
	innerTime time.Time
	hasTz     bool
}

// StripTz removes TZ information from the datetime
func (xdt *XSDDateTime) StripTz() {
	xdt.hasTz = false
}

// ToGoTime converts the time to time.Time by checking if a TZ is specified.
// If there is a TZ, that TZ is used, otherwise local TZ is used
func (xdt *XSDDateTime) ToGoTime() time.Time {
	if xdt.hasTz {
		return xdt.innerTime
	}
	return time.Date(xdt.innerTime.Year(), xdt.innerTime.Month(), xdt.innerTime.Day(),
		xdt.innerTime.Hour(), xdt.innerTime.Minute(), xdt.innerTime.Second(),
		xdt.innerTime.Nanosecond(), time.Local)
}

// MarshalXML implementation on DateTime to skip "zero" time values. It also checks if nanoseconds and TZ exist.
func (xdt XSDDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !xdt.innerTime.IsZero() {
		dateTimeLayout := time.RFC3339Nano
		if xdt.innerTime.Nanosecond() == 0 {
			dateTimeLayout = time.RFC3339
		}
		dtString := xdt.innerTime.Format(dateTimeLayout)
		if !xdt.hasTz {
			// split off time portion
			dateAndTime := strings.SplitN(dtString, "T", 2)
			toks := strings.SplitN(dateAndTime[1], "Z", 2)
			toks = strings.SplitN(toks[0], "+", 2)
			toks = strings.SplitN(toks[0], "-", 2)
			dtString = dateAndTime[0] + "T" + toks[0]
		}
		e.EncodeElement(dtString, start)
	}
	return nil
}

// UnmarshalXML implementation on DateTimeg to use dateTimeLayout
func (xdt *XSDDateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	xdt.innerTime, xdt.hasTz, err = unmarshalTime(d, start, time.RFC3339Nano)
	return err
}

// CreateXsdDateTime creates an object represent xsd:datetime object in Golang
func CreateXsdDateTime(dt time.Time, hasTz bool) XSDDateTime {
	return XSDDateTime{
		innerTime: dt,
		hasTz:     hasTz,
	}
}

func unmarshalTime(d *xml.Decoder, start xml.StartElement, format string) (time.Time, bool, error) {
	var t time.Time
	var content string
	err := d.DecodeElement(&content, &start)
	if err != nil {
		return t, true, err
	}
	if content == "" {
		return t, true, nil
	}
	hasTz := false
	if strings.Contains(content, "T") { // check if we have a time portion
		// split into date and time portion
		dateAndTime := strings.SplitN(content, "T", 2)
		if len(dateAndTime) > 1 {
			if strings.Contains(dateAndTime[1], "Z") ||
				strings.Contains(dateAndTime[1], "+") ||
				strings.Contains(dateAndTime[1], "-") {
				hasTz = true
			}
		}
		if !hasTz {
			content += "Z"
		}
		if content == "0001-01-01T00:00:00Z" {
			return t, true, nil
		}
	} else {
		// we don't see to have a time portion, check timezone
		if strings.Contains(content, "Z") ||
			strings.Contains(content, ":") {
			hasTz = true
		}
		if !hasTz {
			content += "Z"
		}
	}
	t, err = time.Parse(format, content)
	return t, hasTz, nil
}

// XSDDate is a type for representing xsd:date in Golang
type XSDDate struct {
	innerDate time.Time
	hasTz     bool
}

// StripTz removes the TZ information from the date
func (xd *XSDDate) StripTz() {
	xd.hasTz = false
}

// ToGoTime converts the date to Golang time.Time by checking if a TZ is specified.
// If there is a TZ, that TZ is used, otherwise local TZ is used
func (xd *XSDDate) ToGoTime() time.Time {
	if xd.hasTz {
		return xd.innerDate
	}
	return time.Date(xd.innerDate.Year(), xd.innerDate.Month(), xd.innerDate.Day(),
		0, 0, 0, 0, time.Local)
}

// MarshalXML implementation on DateTimeg to skip "zero" time values
func (xd XSDDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !xd.innerDate.IsZero() {
		dateString := xd.innerDate.Format(dateLayout) // serialize with TZ
		if !xd.hasTz {
			if strings.Contains(dateString, "Z") {
				// UTC Tz
				toks := strings.SplitN(dateString, "Z", 2)
				dateString = toks[0]
			} else {
				// [+-]00:00 Tz, remove last 6 chars
				if len(dateString) > 5 { // this should always be true
					start := len(dateString) - 6 // locate at "-"
					dateString = dateString[0:start]
				}
			}
		}
		e.EncodeElement(dateString, start)
	}
	return nil
}

// UnmarshalXML implementation on DateTimeg to use dateTimeLayout
func (xd *XSDDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	xd.innerDate, xd.hasTz, err = unmarshalTime(d, start, dateLayout)
	return err
}

// CreateXsdDate creates an object represent xsd:datetime object in Golang
func CreateXsdDate(date time.Time, hasTz bool) XSDDate {
	return XSDDate{
		innerDate: date,
		hasTz:     hasTz,
	}
}

// XSDTime is a type for representing xsd:time
type XSDTime struct {
	innerTime time.Time
	hasTz     bool
}

// MarshalXML implementation on DateTimeg to skip "zero" time values
func (xt XSDTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !xt.innerTime.IsZero() {
		dateTimeLayout := time.RFC3339Nano
		if xt.innerTime.Nanosecond() == 0 {
			dateTimeLayout = time.RFC3339
		}
		// split off date portion
		dateAndTime := strings.SplitN(xt.innerTime.Format(dateTimeLayout), "T", 2)
		timeString := dateAndTime[1]
		if !xt.hasTz {
			toks := strings.SplitN(timeString, "Z", 2)
			toks = strings.SplitN(toks[0], "+", 2)
			toks = strings.SplitN(toks[0], "-", 2)
			timeString = toks[0]
		}
		e.EncodeElement(timeString, start)
	}
	return nil
}

// UnmarshalXML implementation on DateTimeg to use dateTimeLayout
func (xt *XSDTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var t time.Time
	var err error
	var content string
	err = d.DecodeElement(&content, &start)
	if err != nil {
		return err
	}
	if content == "" {
		xt.innerTime = t
		return nil
	}
	xt.hasTz = false
	if strings.Contains(content, "Z") ||
		strings.Contains(content, "+") ||
		strings.Contains(content, "-") {
		xt.hasTz = true
	}
	if !xt.hasTz {
		content += "Z"
	}
	xt.innerTime, err = time.Parse(timeLayout, content)
	return err
}

// Hour returns hour of the xsd:time
func (xt XSDTime) Hour() int {
	return xt.innerTime.Hour()
}

// Minute returns minutes of the xsd:time
func (xt XSDTime) Minute() int {
	return xt.innerTime.Minute()
}

// Second returns seconds of the xsd:time
func (xt XSDTime) Second() int {
	return xt.innerTime.Second()
}

// Nanosecond returns nanosecond of the xsd:time
func (xt XSDTime) Nanosecond() int {
	return xt.innerTime.Nanosecond()
}

// Location returns the TZ information of the xsd:time
func (xt XSDTime) Location() *time.Location {
	if xt.hasTz {
		return xt.innerTime.Location()
	}
	return nil
}

// CreateXsdTime creates an object representing xsd:time in Golang
func CreateXsdTime(hour int, min int, sec int, nsec int, loc *time.Location) XSDTime {
	realLoc := loc
	if loc == nil {
		realLoc = time.Local
	}
	return XSDTime{
		innerTime: time.Date(1951, 10, 22, hour, min, sec, nsec, realLoc),
		hasTz:     loc != nil,
	}
}
