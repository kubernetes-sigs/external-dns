package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/clustering/v1/policies"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListPolicies(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandlePolicyList(t)

	listOpts := policies.ListOpts{
		Limit: 1,
	}

	count := 0
	err := policies.List(fake.ServiceClient(), listOpts).EachPage(func(page pagination.Page) (bool, error) {
		actual, err := policies.ExtractPolicies(page)
		if err != nil {
			t.Errorf("Failed to extract policies: %v", err)
			return false, err
		}

		th.AssertDeepEquals(t, ExpectedPolicies[count], actual)
		count++

		return true, nil
	})

	th.AssertNoErr(t, err)

	if count != 2 {
		t.Errorf("Expected 2 pages, got %d", count)
	}
}
