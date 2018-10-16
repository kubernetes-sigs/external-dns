package cmd

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Templates details",
}

func getTemplateByName(zoneID *egoscale.UUID, name string) (*egoscale.Template, error) {

	// Find by name, then by ID
	template := &egoscale.Template{
		IsFeatured: true,
		ZoneID:     zoneID,
	}

	id, errUUID := egoscale.ParseUUID(name)
	if errUUID != nil {
		template.Name = name
	} else {
		template.ID = id
	}

	if err := cs.GetWithContext(gContext, template); err == nil {
		return template, err
	}

	// attempts a fuzzy search
	sortedTemplates, err := findTemplates(zoneID, name)
	if err != nil {
		return nil, err
	}

	if len(sortedTemplates) > 1 {
		return nil, fmt.Errorf("more than one templates found")
	}
	if len(sortedTemplates) == 0 {
		return nil, fmt.Errorf("template %q not found", name)
	}

	return &sortedTemplates[0], nil
}

func findTemplates(zoneID *egoscale.UUID, filters ...string) ([]egoscale.Template, error) {
	allOS := make(map[string]*egoscale.Template)

	reLinux := regexp.MustCompile(`^Linux (?P<name>.+?) (?P<version>[0-9]+(\.[0-9]+)?)`)
	reVersion := regexp.MustCompile(`(?P<version>[0-9]+(\.[0-9]+)?)`)

	req := &egoscale.ListTemplates{
		TemplateFilter: "featured",
		ZoneID:         zoneID,
		Keyword:        strings.Join(filters, " "),
	}

	var err error
	cs.PaginateWithContext(gContext, req, func(i interface{}, e error) bool {
		if e != nil {
			err = e
			return false
		}
		template := i.(*egoscale.Template)
		size := template.Size >> 30 // Size in GiB

		if strings.HasPrefix(template.Name, "Linux") {
			m := reSubMatchMap(reLinux, template.DisplayText)
			if len(m) > 0 {
				if size > 10 {
					// Skipping big, legacy images
					return true
				}

				version, errParse := strconv.ParseFloat(m["version"], 64)
				if errParse != nil {
					log.Printf("Malformed Linux version. got %q in %q", m["version"], template.Name)
					return true
				}
				res := fmt.Sprintf("%.5f", 10000-version)

				// fix Container Linux sorting
				name := strings.Replace(m["name"], "stable ", "", 1)
				key := fmt.Sprintf("Linux %s %s", name, res)
				allOS[key] = template
				return true
			}
			// skip
			log.Printf("Malformed Linux. %q", template.DisplayText)
			return true
		}

		if strings.HasPrefix(template.Name, "Windows Server") || strings.HasPrefix(template.Name, "OpenBSD") {
			m := reSubMatchMap(reVersion, template.DisplayText)
			if len(m) > 0 {
				version, errParse := strconv.ParseFloat(m["version"], 64)
				if errParse != nil {
					log.Printf("Malformed Windows/OpenBSD version. %q", template.Name)
					return true
				}
				key := fmt.Sprintf("%s %.5f %5d", template.Name[:7], 10000-version, size)
				allOS[key] = template
				return true
			}

			log.Printf("Malformed Windows/OpenBSD. %q", template.DisplayText)
			return true
		}

		// In doubt, use it directly
		allOS[template.Name] = template
		return true
	})
	if err != nil {
		return nil, err
	}

	var keys []string
	for k := range allOS {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	reDate := regexp.MustCompile(`.* \((?P<date>.*)\)$`)

	templates := make([]egoscale.Template, len(keys))
	for i, k := range keys {
		t := allOS[k]
		m := reSubMatchMap(reDate, t.DisplayText)
		if m["date"] != "" {
			t.Created = m["date"]
		} else if len(t.Created) > 10 {
			t.Created = t.Created[0:10]
		}
		templates[i] = *t
	}
	return templates, nil
}

func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && len(match) > 0 {
			subMatchMap[name] = match[i]
		}
	}
	return subMatchMap
}

func init() {
	RootCmd.AddCommand(templateCmd)
}
