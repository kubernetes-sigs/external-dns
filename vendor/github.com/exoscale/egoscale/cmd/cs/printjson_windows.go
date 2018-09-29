// +build windows

package main

import (
	"fmt"
	"os"
)

func printJSON(out, theme string) {
	fmt.Fprintln(os.Stdout, out)
}
