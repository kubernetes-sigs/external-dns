package cloudflare

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/google/go-querystring/query"
)

// buildURI assembles the base path and queries.
func buildURI(path string, options interface{}) string {
	v, _ := query.Values(options)
	return (&url.URL{Path: path, RawQuery: v.Encode()}).String()
}

// loadFixture takes a series of path components and returns the JSON fixture at
// that locationassociated.
func loadFixture(parts ...string) string {
	paths := []string{"testdata", "fixtures"}
	paths = append(paths, parts...)
	b, err := os.ReadFile(filepath.Join(paths...) + ".json")
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
