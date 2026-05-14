// Copyright (c) 2020-2022, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package json

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/maxatome/go-testdeep/internal/util"
)

const delimiters = " \t\r\n,}]()"

func init() {
	yyErrorVerbose = true
}

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
	value        any
	errs         []*Error
	opts         ParseOpts
}

type ParseOpts struct {
	Placeholders       []any
	PlaceholdersByName map[string]any
	OpFn               func(Operator, Position) (any, error)
}

func Parse(buf []byte, opts ...ParseOpts) (any, error) {
	j := json{
		buf: buf,
		pos: Position{Line: 1},
	}
	if len(opts) > 0 {
		j.opts = opts[0]
	}

	if !j.parse() {
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

// parse returns true if no errors occurred during parsing.
func (j *json) parse() bool {
	yyParse(j)
	return len(j.errs) == 0
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

func (j *json) newOperator(name string, params []any) any {
	if name == "" {
		return nil // an operator error is in progress
	}
	opPos := j.popPos()
	op, err := j.getOperator(Operator{Name: name, Params: params}, opPos)
	if err != nil {
		j.fatal(err.Error(), opPos)
		return nil
	}
	return op
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

func (j *json) getOperator(operator Operator, opPos Position) (any, error) {
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
		return j.analyzeStringContent(s, firstPos, lval)

	case 'r': // raw string, aka r!str! or r<str> (ws possible bw r & start delim)
		if !j.skipWs() {
			j.fatal("cannot find r start delimiter")
			return 0
		}

		firstPos := j.pos.incHoriz(1)
		s, ok := j.parseRawString()
		if !ok {
			return 0
		}
		return j.analyzeStringContent(s, firstPos, lval)

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

	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'+', '.': // '+' & '.' are not normally accepted by JSON spec
		n, ok := j.parseNumber()
		if !ok {
			return 0
		}
		lval.value = n
		return NUMBER

	case '$':
		var dollarToken string
		end := bytes.IndexAny(j.buf[j.pos.bpos+1:], delimiters)
		if end >= 0 {
			dollarToken = string(j.buf[j.pos.bpos+1 : j.pos.bpos+1+end])
		} else {
			dollarToken = string(j.buf[j.pos.bpos+1:])
		}

		if dollarToken == "" {
			return '$'
		}

		token, value := j.parseDollarToken(dollarToken, j.pos, false)
		if token == OPERATOR {
			lval.string = value.(string)
			return OPERATOR
		}
		lval.value = value
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

	var b *strings.Builder

	from := j.pos.bpos + 1
	savePos := j.pos

	appendBuffer := func(r rune) {
		if b == nil {
			b = &strings.Builder{}
			b.Write(j.buf[from : j.pos.bpos-1])
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

		default: //nolint: gocritic
			if r < ' ' || r > utf8.MaxRune {
				j.fatal("invalid character in string")
				return "", false
			}
			fallthrough

		case '\n', '\r', '\t': // not normally accepted by JSON spec
			if b != nil {
				b.WriteRune(r)
			}
		}
	}

	j.fatal("unterminated string", savePos)
	return "", false
}

func (j *json) parseRawString() (string, bool) {
	// j.buf[j.pos.bpos] == first non-ws rune after 'r' → caller responsibility

	savePos := j.pos
	startDelim, _ := j.getRune() // cannot fail, caller called j.skipWs()

	var endDelim rune
	switch startDelim {
	case '(':
		endDelim = ')'
	case '{':
		endDelim = '}'
	case '[':
		endDelim = ']'
	case '<':
		endDelim = '>'
	default:
		if startDelim == '_' ||
			(!unicode.IsPunct(startDelim) && !unicode.IsSymbol(startDelim)) {
			j.fatal(fmt.Sprintf("invalid r delimiter %q, should be either a punctuation or a symbol rune, excluding '_'",
				startDelim))
			return "", false
		}
		endDelim = startDelim
	}

	from := j.pos.bpos + j.curSize

	for innerDelim := 0; ; {
		r, ok := j.getRune()
		if !ok {
			break
		}

		switch r {
		case startDelim:
			if startDelim == endDelim {
				return string(j.buf[from:j.pos.bpos]), true
			}
			innerDelim++

		case endDelim:
			if innerDelim == 0 {
				return string(j.buf[from:j.pos.bpos]), true
			}
			innerDelim--

		case '\n', '\r', '\t': // accept these raw bytes
		default:
			if r < ' ' || r > utf8.MaxRune {
				j.fatal("invalid character in raw string")
				return "", false
			}
		}
	}

	j.fatal("unterminated raw string", savePos)
	return "", false
}

// analyzeStringContent checks whether s contains $ prefix or not. If
// yes, it tries to parse it.
func (j *json) analyzeStringContent(s string, strPos Position, lval *yySymType) int {
	if len(s) <= 1 || !strings.HasPrefix(s, "$") {
		lval.string = s
		return STRING
	}
	// Double $$ at start of strings escape a $
	if strings.HasPrefix(s[1:], "$") {
		lval.string = s[1:]
		return STRING
	}

	// Check for placeholder ($1 or $name) or operator call as $^Empty
	// or $^Re(q<\d+>)
	token, value := j.parseDollarToken(s[1:], strPos, true)
	// in string, j.parseDollarToken can never return an OPERATOR
	// token. In case an operator is embedded in string, a SUB_PARSER is
	// returned instead.
	lval.value = value
	return token
}

const (
	numInt = 1 << iota
	numFloat
	numGoExt
)

var numBytes = [...]uint8{
	'+': numInt, '-': numInt,
	'0': numInt,
	'1': numInt,
	'2': numInt,
	'3': numInt,
	'4': numInt,
	'5': numInt,
	'6': numInt,
	'7': numInt,
	'8': numInt,
	'9': numInt,
	'_': numGoExt,
	// bases 2, 8, 16
	'b': numInt, 'B': numInt, 'o': numInt, 'O': numInt, 'x': numInt, 'X': numInt,
	'a': numInt, 'A': numInt,
	'c': numInt, 'C': numInt,
	'd': numInt, 'D': numInt,
	'e': numInt | numFloat, 'E': numInt | numFloat,
	'f': numInt, 'F': numInt,
	// floats
	'.': numFloat, 'p': numFloat, 'P': numFloat,
}

func (j *json) parseNumber() (float64, bool) {
	// j.buf[j.pos.bpos] == '[-+0-9.]' → caller responsibility

	numKind := numBytes[j.buf[j.pos.bpos]]
	i := j.pos.bpos + 1
	for l := len(j.buf); i < l; i++ {
		b := int(j.buf[i])
		if b >= len(numBytes) || numBytes[b] == 0 {
			break
		}
		numKind |= numBytes[b]
	}

	s := string(j.buf[j.pos.bpos:i])

	var (
		f   float64
		err error
	)
	// Differentiate float/int parsing to accept old octal notation:
	// 0600 → 384 as int64, but 600 as float64
	if (numKind & numFloat) != 0 {
		// strconv.ParseFloat does not handle "_"
		var bf *big.Float
		bf, _, err = new(big.Float).Parse(s, 0)
		if err == nil {
			f, _ = bf.Float64()
		}
	} else { // numInt and/or numGoExt
		var i64 int64
		i64, err = strconv.ParseInt(s, 0, 64)
		if err == nil {
			f = float64(i64)
		}
	}

	if err != nil {
		j.fatal("invalid number")
		return 0, false
	}

	j.curSize = 0
	j.pos = j.pos.incHoriz(i - j.pos.bpos)
	return f, true
}

// parseDollarToken parses a $123 or $tag or $^Operator or
// $^Operator(PARAMS…) token. dollarToken is never empty, does not
// contain '$' and dollarPos is the '$' position.
func (j *json) parseDollarToken(dollarToken string, dollarPos Position, inString bool) (int, any) {
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

	// Test for operator call $^Operator or $^Operator(…)
	if firstRune == '^' {
		nextRune, _ := utf8.DecodeRuneInString(dollarToken[1:])
		if nextRune < 'A' || nextRune > 'Z' {
			j.error(`$^ must be followed by an operator name`, dollarPos)
			if inString {
				return SUB_PARSER, nil // continue parsing
			}
			return OPERATOR, "" // continue parsing
		}

		if inString {
			jr := json{
				buf: []byte(dollarToken[1:]),
				pos: Position{
					Pos:  dollarPos.Pos + 2,
					Line: dollarPos.Line,
					Col:  dollarPos.Col + 2,
				},
				opts: j.opts,
			}
			if !jr.parse() {
				j.errs = append(j.errs, jr.errs...)
				return SUB_PARSER, nil // continue parsing
			}
			return SUB_PARSER, jr.value
		}

		j.moveHoriz(2)
		j.lastTokenPos = j.pos
		operator, ok := j.parseOperator()
		if !ok {
			return OPERATOR, ""
		}
		j.pushPos(j.lastTokenPos)
		return OPERATOR, operator
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
	for ; i < l; i++ {
		if bytes.ContainsAny(j.buf[i:i+1], delimiters) {
			break
		}
		if r := j.buf[i]; (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') {
			j.fatal(fmt.Sprintf(`invalid operator name %q`, string(j.buf[j.pos.bpos:i+1])))
			j.moveHoriz(i - j.pos.bpos)
			return "", false
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
