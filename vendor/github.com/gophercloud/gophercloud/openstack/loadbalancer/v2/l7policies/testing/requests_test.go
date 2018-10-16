package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/l7policies"
	fake "github.com/gophercloud/gophercloud/openstack/networking/v2/common"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestCreateL7Policy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleL7PolicyCreationSuccessfully(t, SingleL7PolicyBody)

	actual, err := l7policies.Create(fake.ServiceClient(), l7policies.CreateOpts{
		Name:        "redirect-example.com",
		ListenerID:  "023f2e34-7806-443b-bfae-16c324569a3d",
		Action:      l7policies.ActionRedirectToURL,
		RedirectURL: "http://www.example.com",
	}).Extract()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, L7PolicyToURL, *actual)
}

func TestRequiredL7PolicyCreateOpts(t *testing.T) {
	// no param specified.
	res := l7policies.Create(fake.ServiceClient(), l7policies.CreateOpts{})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}

	// Action is invalid.
	res = l7policies.Create(fake.ServiceClient(), l7policies.CreateOpts{
		ListenerID: "023f2e34-7806-443b-bfae-16c324569a3d",
		Action:     l7policies.Action("invalid"),
	})
	if res.Err == nil {
		t.Fatalf("Expected error, but got none")
	}
}

func TestListL7Policies(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleL7PolicyListSuccessfully(t)

	pages := 0
	err := l7policies.List(fake.ServiceClient(), l7policies.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := l7policies.ExtractL7Policies(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 l7policies, got %d", len(actual))
		}
		th.CheckDeepEquals(t, L7PolicyToURL, actual[0])
		th.CheckDeepEquals(t, L7PolicyToPool, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllL7Policies(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleL7PolicyListSuccessfully(t)

	allPages, err := l7policies.List(fake.ServiceClient(), l7policies.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := l7policies.ExtractL7Policies(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, L7PolicyToURL, actual[0])
	th.CheckDeepEquals(t, L7PolicyToPool, actual[1])
}
