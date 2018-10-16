package main

import (
	"log"

	"github.com/exoscale/egoscale/cmd/exo/cmd"
)

func main() {
	if err := cmd.RootCmd.GenBashCompletionFile("bash_completion"); err != nil {
		log.Fatal(err)
	}
}
