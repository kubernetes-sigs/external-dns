// +build linux darwin openbsd freebsd

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"syscall"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"golang.org/x/crypto/ssh/terminal"
)

const nocolorsTheme = "nocolors"

func highlight(w io.Writer, source, style string) error {
	l := lexers.Get("json")
	if l == nil {
		l = lexers.Fallback
	}

	l = chroma.Coalesce(l)
	f := formatters.Get("terminal16m")
	if f == nil {
		f = formatters.Fallback
	}

	s := styles.Get(style)
	if s == nil {
		s = styles.Fallback
	}

	sb := chroma.NewStyleBuilder("my" + style)
	for _, tt := range s.Types() {
		se := s.Get(tt)
		se.Background = 0
		if style == nocolorsTheme {
			se.Colour = 0
		}
		sb.AddEntry(tt, se)
	}

	s, err := sb.Build()
	if err != nil {
		return err
	}

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}

	return f.Format(w, s, it)
}

func printJSON(out, theme string) {
	if terminal.IsTerminal(syscall.Stdout) {
		if err := highlight(os.Stdout, out, theme); err != nil {
			log.Fatal(err)
		}
		if _, err := fmt.Fprintln(os.Stdout, ""); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := fmt.Fprintln(os.Stdout, out); err != nil {
			log.Fatal(err)
		}
	}
}
