package rest

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	commaRegexp      = regexp.MustCompile(`,\s{0,}`)
	valueCommaRegexp = regexp.MustCompile(`([^"]),`)
	equalRegexp      = regexp.MustCompile(` *= *`)
	keyRegexp        = regexp.MustCompile(`[a-z*]+`)
	linkRegexp       = regexp.MustCompile(`\A<(.+)>;(.+)\z`)
	semiRegexp       = regexp.MustCompile(`; +`)
	valRegexp        = regexp.MustCompile(`"+([^"]+)"+`)
)

// Links represents a Link Header, keyed by the Rel attribute
type Links map[string]*Link

// Link has a URI and its relation (next/prev/last/etc)
type Link struct {
	URI   string
	Rel   string
	Extra map[string]string
}

// Next gets the URI for "next", if present
func (l Links) Next() string {
	for k, v := range l {
		if k == "next" {
			return v.URI
		}
	}
	return ""
}

// ParseLink parses a Link header value into a Links, mainly cribbed from:
// https://github.com/peterhellberg/link/blob/master/link.go
// The forceHTTPS parameter will rewrite any HTTP URLs it finds to HTTPS.
func ParseLink(s string, forceHTTPS bool) Links {
	if s == "" {
		return nil
	}

	links := Links{}

	s = valueCommaRegexp.ReplaceAllString(s, "$1")

	for _, l := range commaRegexp.Split(s, -1) {
		linkMatches := linkRegexp.FindAllStringSubmatch(l, -1)

		if len(linkMatches) == 0 {
			return nil
		}

		pieces := linkMatches[0]

		// Make sure we have a reasonable URL
		uri := ""
		if url, err := url.ParseRequestURI(pieces[1]); err == nil {

			// See PLAT-188
			if forceHTTPS && url.Scheme == "http" {
				url.Scheme = "https"
			}

			uri = url.String()
		}

		link := &Link{URI: uri, Extra: map[string]string{}}

		for _, extra := range semiRegexp.Split(pieces[2], -1) {
			vals := equalRegexp.Split(extra, -1)

			key := keyRegexp.FindString(vals[0])
			val := valRegexp.FindStringSubmatch(vals[1])[1]

			if key == "rel" {
				vals := strings.Split(val, " ")
				rels := []string{vals[0]}

				if len(vals) > 1 {
					for _, v := range vals[1:] {
						if !strings.HasPrefix(v, "http") {
							rels = append(rels, v)
						}
					}
				}

				rel := strings.Join(rels, " ")

				link.Rel = rel
				links[rel] = link
			} else {
				link.Extra[key] = val
			}
		}
	}

	return links
}
