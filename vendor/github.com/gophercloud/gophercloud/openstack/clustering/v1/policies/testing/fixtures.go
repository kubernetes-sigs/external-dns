package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/openstack/clustering/v1/policies"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

const PolicyListBody1 = `
{
  "policies": [
    {
      "created_at": "2018-04-02T21:43:30.000000",
      "data": {},
      "domain": null,
      "id": "PolicyListBodyID1",
      "name": "delpol",
      "project": "018cd0909fb44cd5bc9b7a3cd664920e",
      "spec": {
        "description": "A policy for choosing victim node(s) from a cluster for deletion.",
        "properties": {
          "criteria": "OLDEST_FIRST",
          "destroy_after_deletion": true,
          "grace_period": 60,
          "reduce_desired_capacity": false
        },
        "type": "senlin.policy.deletion",
        "version": 1
      },
      "type": "senlin.policy.deletion-1.0",
      "updated_at": null,
      "user": "fe43e41739154b72818565e0d2580819"
    }
  ]
}
`

const PolicyListBody2 = `
{
  "policies": [
    {
      "created_at": "2018-04-02T22:29:36.000000",
      "data": {},
      "domain": null,
      "id": "PolicyListBodyID2",
      "name": "delpol2",
      "project": "018cd0909fb44cd5bc9b7a3cd664920e",
      "spec": {
        "description": "A policy for choosing victim node(s) from a cluster for deletion.",
        "properties": {
          "criteria": "OLDEST_FIRST",
          "destroy_after_deletion": true,
          "grace_period": 60,
          "reduce_desired_capacity": false
        },
        "type": "senlin.policy.deletion",
        "version": 1
      },
      "type": "senlin.policy.deletion-1.0",
      "updated_at": null,
      "user": "fe43e41739154b72818565e0d2580819"
    }
  ]
}
`

var (
	ExpectedPolicyCreatedAt1, _ = time.Parse(time.RFC3339, "2018-04-02T21:43:30.000000Z")
	ExpectedPolicyCreatedAt2, _ = time.Parse(time.RFC3339, "2018-04-02T22:29:36.000000Z")
	ZeroTime, _                 = time.Parse(time.RFC3339, "1-01-01T00:00:00.000000Z")

	ExpectedPolicies = [][]policies.Policy{
		{
			{
				CreatedAt: ExpectedPolicyCreatedAt1,
				Data:      map[string]interface{}{},
				Domain:    "",
				ID:        "PolicyListBodyID1",
				Name:      "delpol",
				Project:   "018cd0909fb44cd5bc9b7a3cd664920e",

				Spec: map[string]interface{}{
					"description": "A policy for choosing victim node(s) from a cluster for deletion.",
					"properties": map[string]interface{}{
						"criteria":                "OLDEST_FIRST",
						"destroy_after_deletion":  true,
						"grace_period":            float64(60),
						"reduce_desired_capacity": false,
					},
					"type":    "senlin.policy.deletion",
					"version": float64(1),
				},
				Type:      "senlin.policy.deletion-1.0",
				User:      "fe43e41739154b72818565e0d2580819",
				UpdatedAt: ZeroTime,
			},
		},
		{
			{
				CreatedAt: ExpectedPolicyCreatedAt2,
				Data:      map[string]interface{}{},
				Domain:    "",
				ID:        "PolicyListBodyID2",
				Name:      "delpol2",
				Project:   "018cd0909fb44cd5bc9b7a3cd664920e",

				Spec: map[string]interface{}{
					"description": "A policy for choosing victim node(s) from a cluster for deletion.",
					"properties": map[string]interface{}{
						"criteria":                "OLDEST_FIRST",
						"destroy_after_deletion":  true,
						"grace_period":            float64(60),
						"reduce_desired_capacity": false,
					},
					"type":    "senlin.policy.deletion",
					"version": float64(1),
				},
				Type:      "senlin.policy.deletion-1.0",
				User:      "fe43e41739154b72818565e0d2580819",
				UpdatedAt: ZeroTime,
			},
		},
	}
)

func HandlePolicyList(t *testing.T) {
	th.Mux.HandleFunc("/v1/policies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, PolicyListBody1)
		case "PolicyListBodyID1":
			fmt.Fprintf(w, PolicyListBody2)
		case "PolicyListBodyID2":
			fmt.Fprintf(w, `{"policies":[]}`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}
