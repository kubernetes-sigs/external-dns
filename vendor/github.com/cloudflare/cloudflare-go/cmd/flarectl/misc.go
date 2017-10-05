package main

import (
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/codegangsta/cli"
)

func ips(*cli.Context) {
	ips, _ := cloudflare.IPs()
	fmt.Println("IPv4 ranges:")
	for _, r := range ips.IPv4CIDRs {
		fmt.Println(" ", r)
	}
	fmt.Println()
	fmt.Println("IPv6 ranges:")
	for _, r := range ips.IPv6CIDRs {
		fmt.Println(" ", r)
	}
}

func userInfo(*cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	user, err := api.UserDetails()
	if err != nil {
		fmt.Println(err)
		return
	}
	var output []table
	output = append(output, table{
		"ID":       user.ID,
		"Email":    user.Email,
		"Username": user.Username,
		"Name":     user.FirstName + " " + user.LastName,
		"2FA":      fmt.Sprintf("%t", user.TwoFA),
	})
	makeTable(output, "ID", "Email", "Username", "Name", "2FA")
}

func userUpdate(*cli.Context) {
}

func pageRules(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone"); err != nil {
		return
	}
	zone := c.String("zone")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	rules, err := api.ListPageRules(zoneID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%3s %-32s %-8s %s\n", "Pri", "ID", "Status", "URL")
	for _, r := range rules {
		var settings []string
		fmt.Printf("%3d %s %-8s %s\n", r.Priority, r.ID, r.Status, r.Targets[0].Constraint.Value)
		for _, a := range r.Actions {
			var s string
			switch v := a.Value.(type) {
			case int:
				s = fmt.Sprintf("%s: %d", cloudflare.PageRuleActions[a.ID], v)
			case float64:
				s = fmt.Sprintf("%s: %.f", cloudflare.PageRuleActions[a.ID], v)
			case map[string]interface{}:
				s = fmt.Sprintf("%s: %.f - %s", cloudflare.PageRuleActions[a.ID], v["status_code"], v["url"])
			case nil:
				s = fmt.Sprintf("%s", cloudflare.PageRuleActions[a.ID])
			default:
				vs := fmt.Sprintf("%s", v)
				s = fmt.Sprintf("%s: %s", cloudflare.PageRuleActions[a.ID], strings.Title(strings.Replace(vs, "_", " ", -1)))
			}
			settings = append(settings, s)
		}
		fmt.Println("   ", strings.Join(settings, ", "))
	}
}

func railgun(*cli.Context) {
}
