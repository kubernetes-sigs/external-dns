package durationstring

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Duration struct {
	Years        int
	Months       int
	Days         int
	Hours        int
	Minutes      int
	Seconds      int
	Milliseconds int
	Microseconds int
	Nanoseconds  int
}

func NewDuration(years, months, days, hours, minutes, seconds, milliseconds, microseconds, nanoseconds int) *Duration {
	return &Duration{
		Years:        years,
		Months:       months,
		Days:         days,
		Hours:        hours,
		Minutes:      minutes,
		Seconds:      seconds,
		Milliseconds: milliseconds,
		Microseconds: microseconds,
		Nanoseconds:  nanoseconds,
	}
}

// Parse takes string s and returns the duration in years, months, days, hours, minutes, seconds, nanoseconds, with
// a non-nil error if there was an error with parsing. Expects s in format ([0-9]+(y|mo|d|h|m|s|ns))+, and accepts any order e.g.
// 5y4m, 5s4d
func Parse(s string) (d *Duration, err error) {
	var digitBuf bytes.Buffer
	var unitBuf bytes.Buffer

	flushBuffers := func() {
		digitBuf = bytes.Buffer{}
		unitBuf = bytes.Buffer{}
	}

	d = &Duration{}

	flush := func() error {
		digit := digitBuf.String()
		unit := unitBuf.String()
		flushBuffers()

		if len(digit) < 1 {
			return fmt.Errorf("Digit not supplied for unit '%s'", unit)
		}
		if len(unit) < 1 {
			return fmt.Errorf("Unit not supplied for digit '%s'", digit)
		}

		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			return fmt.Errorf("Failed to parse digit '%s' as int: %s", digit, err.Error())
		}

		switch strings.ToUpper(unit) {
		case "Y":
			d.Years = digitInt
		case "MO":
			d.Months = digitInt
		case "D":
			d.Days = digitInt
		case "H":
			d.Hours = digitInt
		case "M":
			d.Minutes = digitInt
		case "S":
			d.Seconds = digitInt
		case "MS":
			d.Milliseconds = digitInt
		case "US":
			fallthrough
		case "µS":
			d.Microseconds = digitInt
		case "NS":
			d.Nanoseconds = digitInt
		default:
			return fmt.Errorf("invalid unit '%s'", unit)
		}

		return nil
	}

	isUnit := false
	flushBuffers()
	for i, char := range s {
		if unicode.IsSpace(char) {
			continue
		}
		if unicode.IsDigit(char) {
			digitBuf.WriteRune(char)
			isUnit = false
		} else {
			unitBuf.WriteRune(char)
			isUnit = true
		}

		// if we're at the last rune in iteration, or looking at a unit and next rune is either a digit or whitespace, flush
		if len(s)-1 == i || (isUnit && (unicode.IsDigit(rune(s[i+1])) || unicode.IsSpace(rune(s[i+1])))) {
			err := flush()
			if err != nil {
				return d, err
			}
		}
	}

	return d, err
}

func (d *Duration) String() string {
	buf := bytes.Buffer{}

	add := func(digit int, unit string) {
		if digit < 1 {
			return
		}

		buf.WriteString(fmt.Sprintf("%d%s", digit, unit))
	}

	add(d.Years, "y")
	add(d.Months, "mo")
	add(d.Days, "d")
	add(d.Hours, "h")
	add(d.Minutes, "m")
	add(d.Seconds, "s")
	add(d.Milliseconds, "ms")
	add(d.Microseconds, "µS")
	add(d.Nanoseconds, "ns")

	return buf.String()
}
