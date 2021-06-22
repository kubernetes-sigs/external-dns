// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package ctxerr

import (
	"strconv"
	"strings"

	"github.com/maxatome/go-testdeep/internal/util"
)

// Path defines a structure depth path, typically used to mark a
// position during a deep traversal in case of error.
type Path []pathLevel

type pathLevelKind uint8

type pathLevel struct {
	Content  string
	Pointers int
	Kind     pathLevelKind
}

const (
	levelStruct pathLevelKind = iota
	levelArray
	levelMap
	levelFunc
	levelCustom
)

// NewPath returns a new Path initialized with "root" root node.
func NewPath(root string) Path {
	return Path{
		{
			Kind:    levelCustom,
			Content: root,
		},
	}
}

// Len returns the number of levels, excluding pointers ones.
func (p Path) Len() int {
	return len(p)
}

// Equal returns true if "p" and "o" are equal, false otherwise.
func (p Path) Equal(o Path) bool {
	if len(p) != len(o) {
		return false
	}
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] != o[i] {
			return false
		}
	}
	return true
}

func (p Path) addLevel(level pathLevel) Path {
	new := make(Path, len(p), len(p)+1)
	copy(new, p)
	return append(new, level)
}

// Copy returns a new Path, exact but independent copy of "p".
func (p Path) Copy() Path {
	if p == nil {
		return nil
	}

	new := make(Path, len(p))
	copy(new, p)
	return new
}

// AddField adds a level corresponding to a struct field.
func (p Path) AddField(field string) Path {
	if p == nil {
		return nil
	}

	new := p.addLevel(pathLevel{
		Kind:    levelStruct,
		Content: field,
	})

	if len(new) > 1 && new[len(new)-2].Pointers > 0 {
		new[len(new)-2].Pointers--
	}

	return new
}

// AddArrayIndex adds a level corresponding to an array index.
func (p Path) AddArrayIndex(index int) Path {
	if p == nil {
		return nil
	}

	return p.addLevel(pathLevel{
		Kind:    levelArray,
		Content: strconv.Itoa(index),
	})
}

// AddMapKey adds a level corresponding to a map key.
func (p Path) AddMapKey(key interface{}) Path {
	if p == nil {
		return nil
	}

	return p.addLevel(pathLevel{
		Kind:    levelMap,
		Content: util.ToString(key),
	})
}

// AddPtr adds "num" pointers levels.
func (p Path) AddPtr(num int) Path {
	if p == nil {
		return nil
	}

	new := p.Copy()
	// Do not check len(new) > 0, as it should
	new[len(new)-1].Pointers += num
	return new
}

// AddFunctionCall adds a level corresponding to a function call.
func (p Path) AddFunctionCall(fn string) Path {
	if p == nil {
		return nil
	}

	return p.addLevel(pathLevel{
		Kind:    levelFunc,
		Content: fn,
	})
}

// AddCustomLevel adds a custom level.
func (p Path) AddCustomLevel(custom string) Path {
	if p == nil {
		return nil
	}

	return p.addLevel(pathLevel{
		Kind:    levelCustom,
		Content: custom,
	})
}

func (p Path) String() string {
	if len(p) == 0 {
		return ""
	}

	var str string

	for i, level := range p {
		var ptrs string
		if level.Pointers > 0 {
			ptrs = strings.Repeat("*", level.Pointers)
		}

		if level.Kind == levelFunc {
			str = ptrs + level.Content + "(" + str + ")"
		} else {
			if i > 0 && p[i-1].Pointers > 0 {
				// Last level contains pointer(s), protect them
				str = ptrs + "(" + str + ")"
			} else {
				str = ptrs + str
			}

			switch level.Kind {
			case levelStruct:
				str += "." + level.Content
			case levelArray, levelMap:
				str += "[" + level.Content + "]"
			default:
				str += level.Content
			}
		}
	}

	return str
}

/*
func setPtrs(buf []byte, num int) {
	for i := 0; i < num; i++ {
		buf[i] = '*'
	}
}

func (p Path) String() string {
	if len(p) == 0 {
		return ""
	}

	size := 0
	for i, level := range p {
		size += level.Pointers + len(level.Content)

		if level.Kind == levelFunc || (i > 0 && p[i-1].Pointers > 0) {
			size += 2 // () ⇒ content(x) || (x)content
		}

		switch level.Kind {
		case levelStruct:
			size++ // "."
		case levelArray, levelMap:
			size += 2 // []
		}
	}

	buf := make([]byte, size)
	curLen := 0

	for i, level := range p {
		if level.Kind == levelFunc {
			// **content(prev)
			levelLen := level.Pointers + len(level.Content) + 1
			copy(buf[levelLen:], buf[:curLen])
			setPtrs(buf, level.Pointers)
			copy(buf[level.Pointers:], []byte(level.Content))
			buf[levelLen-1] = '('
			curLen += levelLen
			buf[curLen] = ')'
			curLen++
		} else {
			if i > 0 && p[i-1].Pointers > 0 {
				// **(prev)content
				copy(buf[level.Pointers+1:], buf[:curLen])
				setPtrs(buf, level.Pointers)
				buf[level.Pointers] = '('
				curLen += level.Pointers + 1
				buf[curLen] = ')'
				curLen++
			} else {
				// **prevcontent
				if level.Pointers > 0 {
					copy(buf[level.Pointers:], buf[:curLen])
					setPtrs(buf, level.Pointers)
					curLen += level.Pointers
				}
			}
			switch level.Kind {
			case levelStruct:
				buf[curLen] = '.'
				curLen++
				copy(buf[curLen:], []byte(level.Content))
				curLen += len(level.Content)
			case levelArray, levelMap:
				buf[curLen] = '['
				curLen++
				copy(buf[curLen:], []byte(level.Content))
				curLen += len(level.Content)
				buf[curLen] = ']'
				curLen++
			default:
				copy(buf[curLen:], []byte(level.Content))
				curLen += len(level.Content)
			}
		}
	}

	return string(buf)
}
*/
