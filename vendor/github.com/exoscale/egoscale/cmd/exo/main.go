package main

import (
	"log"

	"github.com/exoscale/egoscale/cmd/exo/cmd"
)

func main() {
	log.SetFlags(0)
	cmd.Execute()
}
