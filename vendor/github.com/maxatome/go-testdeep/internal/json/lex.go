// Copyright (c) 2020, 2021, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package json

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/maxatome/go-testdeep/internal/util"
)

type Position struct {
	bpos int
	Pos  int
	Line int
	Col  int
}

func (p Position) incHoriz(bytes int, runes ...int) Position {
	p.bpos += bytes

	r := bytes
	if len(runes) > 0 {
		r = runes[0]
	}
	p.Pos += r
	p.Col += r
	return p
}

func (p Position) String() string {
	return fmt.Sprintf("at line %d:%d (pos %d)", p.Line, p.Col, p.Pos)
}

type json struct {
	buf          []byte
	pos          Position
	lastTokenPos Position
	stackPos     []Position
	curSize      int
	curRune      rune
	value        interface{}
	errs         []*Error
	opts         ParseOpts
}

type ParseOpts struct {
	Placeholders       []interface{}
	PlaceholdersByName map[string]interface{}
	OpShortcutFn       func(string, Position) (interface{}, bool)
	OpFn               func(Operator, Position) (interface{}, error)
}

func Parse(buf []byte, opts ...ParseOpts) (interface{}, error) {
	yyErrorVerbose = true

	j := json{
		buf: buf,
		pos: Position{Line: 1},
	}
	if len(opts) > 0 {
		j.opts = opts[0]
	}
	yyParse(&j)

	if len(j.errs) > 0 {
		if len(j.errs) == 1 {
			return nil, j.errs[0]
		}

		errStr := bytes.NewBufferString(j.errs[0].Error())
		for _, err := range j.errs[1:] {
			errStr.WriteByte('\n')
			errStr.WriteString(err.Error())
		}
		return nil, errors.New(errStr.String())
	}

	return j.value, nil
}

// Lex implements yyLexer interface.
func (j *json) Lex(lval *yySymType) int {
	return j.nextToken(lval)
}

// Error implements yyLexer interface.
func (j *json) Error(s string) {
	if len(j.errs) == 0 || !j.errs[len(j.errs)-1].fatal {
		const syntaxErrorUnexpected = "syntax error: unexpected "
		if s == syntaxErrorUnexpected+"$unk" {
			switch {
			case unicode.IsPrint(j.curRune):
				s = syntaxErrorUnexpected + "'" + string(j.curRune) + "'"
			case j.curRune <= 0xffff:
				s = fmt.Sprintf(syntaxErrorUnexpected+`'\u%04x'`, j.curRune)
			default:
				s = fmt.Sprintf(syntaxErrorUnexpected+`'\U%08x'`, j.curRune)
			}
		} else if strings.HasPrefix(s, syntaxErrorUnexpected+"$end") {
			s = strings.Replace(s, "$end", "EOF", 1)
		}
		j.fatal(s, j.lastTokenPos)
	}
}

func (j *json) pushPos(pos Position) {
	j.stackPos = append(j.stackPos, pos)
}

func (j *json) popPos() Position {
	last := len(j.stackPos) - 1
	pos := j.stackPos[last]
	j.stackPos = j.stackPos[:last]
	return pos
}

func (j *json) moveHoriz(bytes int, runes ...int) {
	j.pos = j.pos.incHoriz(bytes, runes...)
	j.curSize = 0
}

func (j *json) getOperatorShortcut(operator string, opPos Position) (interface{}, bool) {
	if j.opts.OpShortcutFn == nil {
		return nil, false
	}
	return j.opts.OpShortcutFn(operator, opPos)
}

func (j *json) getOperator(operator Operator, opPos Position) (interface{}, error) {
	if j.opts.OpFn == nil {
		return nil, fmt.Errorf("unknown operator %q", operator.Name)
	}
	return j.opts.OpFn(operator, opPos)
}

func (j *json) nextToken(lval *yySymType) int {
	if !j.skipWs() {
		return 0
	}

	j.lastTokenPos = j.pos

	r, _ := j.getRune()

	switch r {
	case '"':
		firstPos := j.pos.incHoriz(1)
		s, ok := j.parseString()
		if !ok {
			return 0
		}

		// Check for placeholder ($1 or $name) or operator shortcut ($^Nil)
		if len(s) <= 1 || !strings.HasPrefix(s, "$") {
			lval.string = s
			return STRING
		}
		// Double $$ at start of strings escape a $
		if strings.HasPrefix(s[1:], "$") {
			lval.string = s[1:]
			return STRING
		}

		token, value := j.parseDollarToken(s[1:], firstPos)
		if token != 0 {
			lval.value = value
		}
		return token

	case 'n': // null
		if j.remain() >= 4 && bytes.Equal(j.buf[j.pos.bpos+1:j.pos.bpos+4], []byte(`ull`)) {
			j.skip(3)
			lval.value = nil
			return NULL
		}

	case 't': // true
		if j.remain() >= 4 && bytes.Equal(j.buf[j.pos.bpos+1:j.pos.bpos+4], []byte(`rue`)) {
			j.skip(3)
			lval.value = true
			return TRUE
		}

	case 'f': // false
		if j.remain() >= 5 && bytes.Equal(j.buf[j.pos.bpos+1:j.pos.bpos+5], []byte(`alse`)) { //nolint: misspell
			j.skip(4)
			lval.value = false
			return FALSE
		}

	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		n, ok := j.parseNumber()
		if !ok {
			return 0
		}
		lval.value = n
		return NUMBER

	case '$':
		var dollarToken string
		end := bytes.IndexAny(j.buf[j.pos.bpos+1:], " \t\r\n,}])")
		if end >= 0 {
			dollarToken = string(j.buf[j.pos.bpos+1 : j.pos.bpos+1+end])
		} else {
			dollarToken = string(j.buf[j.pos.bpos+1:])
		}

		if dollarToken == "" {
			return '$'
		}

		token, value := j.parseDollarToken(dollarToken, j.pos)
		if token != 0 {
			lval.value = value
		}
		j.moveHoriz(1+len(dollarToken), 1+utf8.RuneCountInString(dollarToken))
		return token

	default:
		if r >= 'A' && r <= 'Z' {
			operator, ok := j.parseOperator()
			if !ok {
				return 0
			}
			j.pushPos(j.lastTokenPos)
			lval.string = operator
			return OPERATOR
		}
	}

	return int(r)
}

func hex(b []byte) (rune, bool) {
	var r rune
	for i := 0; i < 4; i++ {
		r <<= 4
		switch {
		case b[i] >= '0' && b[i] <= '9':
			r += rune(b[i]) - '0'
		case b[i] >= 'a' && b[i] <= 'f':
			r += rune(b[i]) - 'a' + 10
		case b[i] >= 'A' && b[i] <= 'F':
			r += rune(b[i]) - 'A' + 10
		default:
			return 0, false
		}
	}
	return r, true
}

func (j *json) parseString() (string, bool) {
	// j.buf[j.pos.bpos] == '"' → caller responsibility

	var b *bytes.Buffer

	from := j.pos.bpos + 1
	savePos := j.pos

	appendBuffer := func(r rune) {
		if b == nil {
			b = bytes.NewBuffer(j.buf[from : j.pos.bpos-1])
		}
		b.WriteRune(r)
	}

str:
	for {
		r, ok := j.getRune()
		if !ok {
			break
		}

		switch r {
		case '"':
			if b == nil {
				return string(j.buf[from:j.pos.bpos]), true
			}
			return b.String(), true

		case '\\':
			r, ok := j.getRune()
			if !ok {
				break str
			}

			switch r {
			case '"', '\\', '/':
				appendBuffer(r)
			case 'b':
				appendBuffer('\b')
			case 'f':
				appendBuffer('\f')
			case 'n':
				appendBuffer('\n')
			case 'r':
				appendBuffer('\r')
			case 't':
				appendBuffer('\t')
			case 'u':
				if j.remain() >= 5 {
					r, ok = hex(j.buf[j.pos.bpos+1 : j.pos.bpos+5])
					if ok {
						appendBuffer(r)
						j.pos = j.pos.incHoriz(4)
						break
					}
				}
				fallthrough
			default:
				j.fatal("invalid escape sequence")
				return "", false
			}

		default:
			if r < ' ' || r > utf8.MaxRune {
				j.fatal("invalid character in string")
				return "", false
			}
			if b != nil {
				b.WriteRune(r)
			}
		}
	}

	j.fatal("unterminated string", savePos)
	return "", false
}

func (j *json) parseNumber() (float64, bool) {
	// j.buf[j.pos.bpos] == '[0-9]' → caller responsibility

	i := j.pos.bpos + 1
	l := len(j.buf)
num:
	for ; i < l; i++ {
		switch j.buf[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E', '+', '-':
		default:
			break num
		}
	}

	f, err := strconv.ParseFloat(string(j.buf[j.pos.bpos:i]), 64)
	if err != nil {
		j.fatal("invalid number")
		return 0, false
	}

	j.curSize = 0
	j.pos = j.pos.incHoriz(i - j.pos.bpos)
	return f, true
}

// parseDollarToken parses a $123 or $tag or $^Shortcut token.
// dollarToken is never empty, does not contain '$' and dollarPos
// is the '$' position.
func (j *json) parseDollarToken(dollarToken string, dollarPos Position) (int, interface{}) {
	firstRune, _ := utf8.DecodeRuneInString(dollarToken)

	// Test for $123
	if firstRune >= '0' && firstRune <= '9' {
		np, err := strconv.ParseUint(dollarToken, 10, 64)
		if err != nil {
			j.error("invalid numeric placeholder", dollarPos)
			return PLACEHOLDER, nil // continue parsing
		}
		if np == 0 {
			j.error(
				fmt.Sprintf(`invalid numeric placeholder "$%s", it should start at "$1"`, dollarToken),
				dollarPos)
			return PLACEHOLDER, nil // continue parsing
		}
		if numParams := len(j.opts.Placeholders); np > uint64(numParams) {
			switch numParams {
			case 0:
				j.error(
					fmt.Sprintf(`numeric placeholder "$%s", but no params given`, dollarToken),
					dollarPos)
			case 1:
				j.error(
					fmt.Sprintf(`numeric placeholder "$%s", but only one param given`, dollarToken),
					dollarPos)
			default:
				j.error(
					fmt.Sprintf(`numeric placeholder "$%s", but only %d params given`,
						dollarToken, numParams),
					dollarPos)
			}
			return PLACEHOLDER, nil // continue parsing
		}
		return PLACEHOLDER, j.opts.Placeholders[np-1]
	}

	// Test for operator shortcut
	if firstRune == '^' {
		op, ok := j.getOperatorShortcut(dollarToken[1:], dollarPos)
		if !ok {
			j.error(
				fmt.Sprintf(`bad operator shortcut "$%s"`, dollarToken),
				dollarPos)
			// continue parsing
		}
		return OPERATOR_SHORTCUT, op
	}

	// Test for $tag
	err := util.CheckTag(dollarToken)
	if err != nil {
		j.error(
			fmt.Sprintf(`bad placeholder "$%s"`, dollarToken),
			dollarPos)
		return PLACEHOLDER, nil // continue parsing
	}
	op, ok := j.opts.PlaceholdersByName[dollarToken]
	if !ok {
		j.error(
			fmt.Sprintf(`unknown placeholder "$%s"`, dollarToken),
			dollarPos)
		// continue parsing
	}
	return PLACEHOLDER, op
}

func (j *json) parseOperator() (string, bool) {
	// j.buf[j.pos.bpos] == '[A-Z]' → caller responsibility

	i := j.pos.bpos + 1
	l := len(j.buf)
operator:
	for ; i < l; i++ {
		switch r := j.buf[i]; r {
		case ' ', '\t', '\r', '\n', ',', '}', ']', '(':
			break operator

		default:
			if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') {
				j.fatal(fmt.Sprintf(`invalid operator name %q`, string(j.buf[j.pos.bpos:i+1])))
				return "", false
			}
		}
	}

	s := string(j.buf[j.pos.bpos:i])
	j.moveHoriz(i - j.pos.bpos)
	return s, true
}

func (j *json) skipWs() bool {
ws:
	for {
		r, ok := j.getRune()
		if !ok {
			return false
		}

		switch r {
		case ' ', '\n', '\r', '\t':

		case '/':
			if j.remain() < 2 {
				break ws
			}

			switch j.buf[j.pos.bpos+1] {
			case '/': // comment till eol
				j.curSize = 0
				if end := indexAfterEol(j.buf[j.pos.bpos+2:]); end >= 0 {
					lineLen := 2 + utf8.RuneCount(j.buf[j.pos.bpos+2:j.pos.bpos+2+end])
					j.pos.Pos += lineLen
					j.pos.Line++
					j.pos.Col = 0
					j.pos.bpos += 2 + end
					continue ws
				}
				lineLen := 2 + utf8.RuneCount(j.buf[j.pos.bpos+2:])
				j.pos.Pos += lineLen
				j.pos.Col += lineLen
				j.pos.bpos = len(j.buf) // till eof
				return false

			case '*': // multi-lines comment
				j.curSize = 0
				if end := bytes.Index(j.buf[j.pos.bpos+2:], []byte("*/")); end >= 0 {
					comment := j.buf[j.pos.bpos+2 : j.pos.bpos+2+end]
					commentLen := 4 + utf8.RuneCount(comment)
					// Count \r\n as only one rune
					if crnl := bytes.Count(comment, []byte("\r\n")); crnl > 0 {
						commentLen -= crnl
					}
					j.pos.Pos += commentLen
					j.pos.bpos += 4 + end

					nLines := countEol(comment)
					if nLines > 0 {
						j.pos.Line += nLines
						j.pos.Col = len(comment) - bytes.LastIndexAny(comment, "\r\n") + 1
					} else {
						j.pos.Col += commentLen
					}
					continue ws
				}
				j.fatal("multi-lines comment not terminated")
				return false

			default:
				break ws
			}

		default:
			break ws
		}
	}

	j.curSize = 0
	return true
}

// indexAfterEol returns the index of the byte just after the first
// instance of an end-of-line ('\n' alone, '\r' alone or "\r\n") in
// buf, or -1 if no end-of-line is found.
func indexAfterEol(buf []byte) int {
	// new line for:
	// - \n alone
	// - \r\n
	// - \r alone
	for i, b := range buf {
		switch b {
		case '\n':
			return i + 1
		case '\r':
			if i+1 == len(buf) || buf[i+1] != '\n' {
				return i + 1
			}
			return i + 2
		}
	}
	return -1
}

// countEol returns the number of end-of-line ('\n' alone, '\r' alone
// or "\r\n") occurrences in buf.
func countEol(buf []byte) int {
	// new line for:
	// - \n alone
	// - \r\n
	// - \r alone
	num := 0
	for {
		eol := indexAfterEol(buf)
		if eol < 0 {
			return num
		}
		buf = buf[eol:]
		num++
	}
}

func (j *json) getRune() (rune, bool) {
	if j.curSize > 0 {
		// new line for:
		// - \n alone
		// - \r\n (+ consider it as one rune)
		// - \r alone
		switch j.buf[j.pos.bpos] {
		case '\n':
			if j.pos.bpos > 0 && j.buf[j.pos.bpos-1] == '\r' {
				// \r\n → already handled
				break
			}
			fallthrough
		case '\r':
			j.pos.Line++
			j.pos.Col = 0
			j.pos.Pos++
		default:
			j.pos.Col++
			j.pos.Pos++
		}
		j.pos.bpos += j.curSize
		j.curSize = 0
	}

	if j.remain() == 0 {
		return 0, false
	}

	r, size := utf8.DecodeRune(j.buf[j.pos.bpos:])
	j.curSize = size
	j.curRune = r
	return r, true
}

func (j *json) skip(n int) {
	j.pos.Pos += n
	j.pos.Col += n
	j.pos.bpos += n
}

func (j *json) remain() int {
	return len(j.buf) - j.pos.bpos
}

func (j *json) appendError(mesg string, fatal bool, pos ...Position) {
	err := Error{
		mesg:  mesg,
		Pos:   j.pos,
		fatal: fatal,
	}
	if len(pos) > 0 {
		err.Pos = pos[0]
	}
	j.errs = append(j.errs, &err)
}

func (j *json) error(mesg string, pos ...Position) {
	j.appendError(mesg, false, pos...)
}

func (j *json) fatal(mesg string, pos ...Position) {
	j.appendError(mesg, true, pos...)
}

type Error struct {
	mesg  string
	Pos   Position
	fatal bool
}

func (e *Error) Error() string {
	return e.mesg + " " + e.Pos.String()
}
