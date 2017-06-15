package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizations_ListOrganizations(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/organizations", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
"success": true,
"errors": [],
"messages": [],
"result": [
    {
      "id": "01a7362d577a6c3019a474fd6f485823",
      "name": "Cloudflare, Inc.",
      "status": "member",
      "permissions": [
        "#zones:read"
      ],
      "roles": [
        "All Privileges - Super Administrator"
      ]
    }
  ],
"result_info": {
  "page": 1,
  "per_page": 20,
  "count": 1,
  "total_count": 2000
  }
}`)
	})

	user, paginator, err := client.ListOrganizations()

	want := []Organization{{
		ID:          "01a7362d577a6c3019a474fd6f485823",
		Name:        "Cloudflare, Inc.",
		Status:      "member",
		Permissions: []string{"#zones:read"},
		Roles:       []string{"All Privileges - Super Administrator"},
	}}

	if assert.NoError(t, err) {
		assert.Equal(t, user, want)
	}

	want_pagination := ResultInfo{
		Page:    1,
		PerPage: 20,
		Count:   1,
		Total:   2000,
	}
	assert.Equal(t, paginator, want_pagination)
}
