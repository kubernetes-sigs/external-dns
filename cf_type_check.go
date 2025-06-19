package main

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go/v4"
)

func main() {
	var r cloudflare.DNSRecord
	fmt.Printf("Cloudflare DNSRecord type: %+v\n", r)
}
