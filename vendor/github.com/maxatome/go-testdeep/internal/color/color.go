// Copyright (c) 2019, 2020 Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package color

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
)

const (
	// EnvColor is the name of the environment variable allowing to
	// enable/disable coloring feature.
	EnvColor = "TESTDEEP_COLOR"
	// EnvColorTestName is the name of the environment variable
	// containing the color of test names in error reports.
	EnvColorTestName = "TESTDEEP_COLOR_TEST_NAME"
	// EnvColorTitle is the name of the environment variable
	// containing the color of failure reason in error reports.
	EnvColorTitle = "TESTDEEP_COLOR_TITLE"
	// EnvColorOK is the name of the environment variable
	// containing the color of "expected" in error reports.
	EnvColorOK = "TESTDEEP_COLOR_OK"
	// EnvColorBad is the name of the environment variable
	// containing the color of "got" in error reports.
	EnvColorBad = "TESTDEEP_COLOR_BAD"
)

var (
	// TestNameOn contains the ANSI color escape sequence to turn test
	// name color on.
	TestNameOn string
	// TestNameOff contains the ANSI color escape sequence to turn test
	// name color off.
	TestNameOff string
	// TitleOn contains the ANSI color escape sequence to turn title color on.
	TitleOn string
	// TitleOff contains the ANSI color escape sequence to turn title color off.
	TitleOff string
	// OKOn contains the ANSI color escape sequence to turn "expected" color on.
	OKOn string
	// OKOnBold contains the ANSI color escape sequence to turn
	// "expected" color and bold on.
	OKOnBold string
	// OKOff contains the ANSI color escape sequence to turn "expected" color off.
	OKOff string
	// BadOn contains the ANSI color escape sequence to turn "got" color on.
	BadOn string
	// BadOnBold contains the ANSI color escape sequence to turn "got"
	// color and bold on.
	BadOnBold string
	// BadOff contains the ANSI color escape sequence to turn "got" color off.
	BadOff string
)

var initOnce sync.Once

// Init initializes all the colors from the environment. It can be
// called several times concurrently, but only the first call is
// effective.
func Init() {
	initOnce.Do(func() {
		_, TestNameOn, TestNameOff = FromEnv(EnvColorTestName, "yellow")
		_, TitleOn, TitleOff = FromEnv(EnvColorTitle, "cyan")
		OKOn, OKOnBold, OKOff = FromEnv(EnvColorOK, "green")
		BadOn, BadOnBold, BadOff = FromEnv(EnvColorBad, "red")
	})
}

// SaveState saves the "TESTDEEP_COLOR" environment variable
// value, sets it to "on" (if true passed as on) or "false" (if on not
// passed or set to false), resets the colors initialization and
// returns a function to be called in a defer statement. Only intended
// to be used in tests like:
//
//   defer color.SaveState()()
//
// It is not thread-safe.
func SaveState(on ...bool) func() {
	colorState, set := os.LookupEnv(EnvColor)
	if len(on) == 0 || !on[0] {
		os.Setenv(EnvColor, "off") // nolint: errcheck
	} else {
		os.Setenv(EnvColor, "on") // nolint: errcheck
	}
	initOnce = sync.Once{}
	return func() {
		if set {
			os.Setenv(EnvColor, colorState) // nolint: errcheck
		} else {
			os.Unsetenv(EnvColor) // nolint: errcheck
		}
		initOnce = sync.Once{}
	}
}

var colors = map[string]byte{
	"black":   '0',
	"red":     '1',
	"green":   '2',
	"yellow":  '3',
	"blue":    '4',
	"magenta": '5',
	"cyan":    '6',
	"white":   '7',
	"gray":    '7',
}

// FromEnv returns the light, bold and end ANSI sequences for the
// color contained in the environment variable env. defaultColor is
// used if the environment variable does exist or is empty.
//
// If coloring is disabled, returns "", "", "".
func FromEnv(env, defaultColor string) (string, string, string) {
	var color string
	switch os.Getenv(EnvColor) {
	case "on", "":
		if curColor := os.Getenv(env); curColor != "" {
			color = curColor
		} else {
			color = defaultColor
		}
	default: // "off" or any other value
		color = ""
	}

	if color == "" {
		return "", "", ""
	}

	names := strings.SplitN(color, ":", 2)

	light := [...]byte{
		//   0    1    2    4    4    5    6
		'\x1b', '[', '0', ';', '3', 'y', 'm', // foreground
		//   7    8    9   10   11
		'\x1b', '[', '4', 'z', 'm', // background
	}
	bold := [...]byte{
		//   0    1    2    4    4    5    6
		'\x1b', '[', '1', ';', '3', 'y', 'm', // foreground
		//   7    8    9   10   11
		'\x1b', '[', '4', 'z', 'm', // background
	}

	var start, end int

	// Foreground
	if names[0] != "" {
		c := colors[names[0]]
		if c == 0 {
			c = colors[defaultColor]
		}

		light[5] = c
		bold[5] = c

		end = 7
	} else {
		start = 7
	}

	// Background
	if len(names) > 1 && names[1] != "" {
		c := colors[names[1]]
		if c != 0 {
			light[10] = c
			bold[10] = c

			end = 12
		}
	}

	return string(light[start:end]), string(bold[start:end]), "\x1b[0m"
}

// AppendTestNameOn enables test name color in b.
func AppendTestNameOn(b *bytes.Buffer) {
	Init()
	b.WriteString(TestNameOn)
}

// AppendTestNameOff disables test name color in b.
func AppendTestNameOff(b *bytes.Buffer) {
	Init()
	b.WriteString(TestNameOff)
}

// Bad returns a string surrounded by BAD color. If len(args) is > 0,
// s and args are given to fmt.Sprintf.
//
// Typically used in panic() when the user made a mistake.
func Bad(s string, args ...interface{}) string {
	Init()
	if len(args) == 0 {
		return BadOnBold + s + BadOff
	}
	return fmt.Sprintf(BadOnBold+s+BadOff, args...)
}

// BadUsage returns a string surrounded by BAD color to notice the
// user he passes a bad parameter to a function. Typically used in a
// panic().
func BadUsage(usage string, param interface{}, pos int, kind bool) string {
	Init()

	var b bytes.Buffer
	fmt.Fprintf(&b, "%susage: %s, but received ", BadOnBold, usage)

	if param == nil {
		b.WriteString("nil")
	} else {
		t := reflect.TypeOf(param)
		if kind && t.String() != t.Kind().String() {
			fmt.Fprintf(&b, "%s (%s)", t, t.Kind())
		} else {
			b.WriteString(t.String())
		}
	}

	b.WriteString(" as ")
	switch pos {
	case 1:
		b.WriteString("1st")
	case 2:
		b.WriteString("2nd")
	case 3:
		b.WriteString("3rd")
	default:
		fmt.Fprintf(&b, "%dth", pos)
	}
	b.WriteString(" parameter")
	b.WriteString(BadOff)
	return b.String()
}

// TooManyParams returns a string surrounded by BAD color to notice
// the user he called a variadic function with too many
// parameters. Typically used in a panic().
func TooManyParams(usage string) string {
	Init()
	return BadOnBold + "usage: " + usage + ", too many parameters" + BadOff
}

// UnBad returns "s" with bad color prefix & suffix removed.
func UnBad(s string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, BadOnBold), BadOff)
}
