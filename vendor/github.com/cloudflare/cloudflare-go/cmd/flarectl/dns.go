package main

import (
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/codegangsta/cli"
)

func dnsCreate(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone", "name", "type", "content"); err != nil {
		return
	}
	zone := c.String("zone")
	name := c.String("name")
	rtype := c.String("type")
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	record := cloudflare.DNSRecord{
		Name:    name,
		Type:    strings.ToUpper(rtype),
		Content: content,
		TTL:     ttl,
		Proxied: proxy,
	}
	// TODO: Print the result.
	_, err = api.CreateDNSRecord(zoneID, record)
	if err != nil {
		fmt.Println("Error creating DNS record:", err)
	}
}

func dnsCreateOrUpdate(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone", "name", "type", "content"); err != nil {
		return
	}
	zone := c.String("zone")
	name := c.String("name")
	rtype := strings.ToUpper(c.String("type"))
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Look for an existing record
	rr := cloudflare.DNSRecord{
		Name: name + "." + zone,
	}
	records, err := api.DNSRecords(zoneID, rr)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(records) > 0 {
		// Record exists - find the ID and update it.
		// This is imprecise without knowing the original content; if a label
		// has multiple RRs we'll just update the first one.
		for _, r := range records {
			if r.Type == rtype {
				rr.ID = r.ID
				rr.Type = r.Type
				rr.Content = content
				rr.TTL = ttl
				rr.Proxied = proxy
				err := api.UpdateDNSRecord(zoneID, r.ID, rr)
				if err != nil {
					fmt.Println("Error updating DNS record:", err)
				}
			}
		}
	} else {
		// Record doesn't exist - create it
		rr.Type = rtype
		rr.Content = content
		rr.TTL = ttl
		rr.Proxied = proxy
		// TODO: Print the response.
		_, err := api.CreateDNSRecord(zoneID, rr)
		if err != nil {
			fmt.Println("Error creating DNS record:", err)
		}
	}
}

func dnsUpdate(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone", "id"); err != nil {
		return
	}
	zone := c.String("zone")
	recordID := c.String("id")
	name := c.String("name")
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	record := cloudflare.DNSRecord{
		ID:      recordID,
		Name:    name,
		Content: content,
		TTL:     ttl,
		Proxied: proxy,
	}
	err = api.UpdateDNSRecord(zoneID, recordID, record)
	if err != nil {
		fmt.Println("Error updating DNS record:", err)
	}
}

func dnsDelete(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone", "id"); err != nil {
		return
	}
	zone := c.String("zone")
	recordID := c.String("id")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = api.DeleteDNSRecord(zoneID, recordID)
	if err != nil {
		fmt.Println("Error deleting DNS record:", err)
	}
}
