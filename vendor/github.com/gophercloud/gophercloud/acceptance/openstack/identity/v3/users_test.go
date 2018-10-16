// +build acceptance

package v3

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/groups"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
)

func TestUsersList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	var iTrue bool = true
	listOpts := users.ListOpts{
		Enabled: &iTrue,
	}

	allPages, err := users.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to list users: %v", err)
	}

	allUsers, err := users.ExtractUsers(allPages)
	if err != nil {
		t.Fatalf("Unable to extract users: %v", err)
	}

	for _, user := range allUsers {
		tools.PrintResource(t, user)
		tools.PrintResource(t, user.Extra)
	}
}

func TestUsersGet(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	allPages, err := users.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list users: %v", err)
	}

	allUsers, err := users.ExtractUsers(allPages)
	if err != nil {
		t.Fatalf("Unable to extract users: %v", err)
	}

	user := allUsers[0]
	p, err := users.Get(client, user.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get user: %v", err)
	}

	tools.PrintResource(t, p)
}

func TestUserCRUD(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	project, err := CreateProject(t, client, nil)
	if err != nil {
		t.Fatalf("Unable to create project: %v", err)
	}
	defer DeleteProject(t, client, project.ID)

	tools.PrintResource(t, project)

	createOpts := users.CreateOpts{
		DefaultProjectID: project.ID,
		Password:         "foobar",
		DomainID:         "default",
		Options: map[users.Option]interface{}{
			users.IgnorePasswordExpiry: true,
			users.MultiFactorAuthRules: []interface{}{
				[]string{"password", "totp"},
				[]string{"password", "custom-auth-method"},
			},
		},
		Extra: map[string]interface{}{
			"email": "jsmith@example.com",
		},
	}

	user, err := CreateUser(t, client, &createOpts)
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}
	defer DeleteUser(t, client, user.ID)

	tools.PrintResource(t, user)
	tools.PrintResource(t, user.Extra)

	iFalse := false
	updateOpts := users.UpdateOpts{
		Enabled: &iFalse,
		Options: map[users.Option]interface{}{
			users.MultiFactorAuthRules: nil,
		},
		Extra: map[string]interface{}{
			"disabled_reason": "DDOS",
		},
	}

	newUser, err := users.Update(client, user.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update user: %v", err)
	}

	tools.PrintResource(t, newUser)
	tools.PrintResource(t, newUser.Extra)
}

func TestUserChangePassword(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	createOpts := users.CreateOpts{
		Password: "secretsecret",
		DomainID: "default",
		Options: map[users.Option]interface{}{
			users.IgnorePasswordExpiry: true,
			users.MultiFactorAuthRules: []interface{}{
				[]string{"password", "totp"},
				[]string{"password", "custom-auth-method"},
			},
		},
		Extra: map[string]interface{}{
			"email": "jsmith@example.com",
		},
	}

	user, err := CreateUser(t, client, &createOpts)
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}
	defer DeleteUser(t, client, user.ID)

	tools.PrintResource(t, user)
	tools.PrintResource(t, user.Extra)

	changePasswordOpts := users.ChangePasswordOpts{
		OriginalPassword: "secretsecret",
		Password:         "new_secretsecret",
	}
	err = users.ChangePassword(client, user.ID, changePasswordOpts).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to change password for user: %v", err)
	}
}

func TestUsersListGroups(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}
	allUserPages, err := users.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list users: %v", err)
	}

	allUsers, err := users.ExtractUsers(allUserPages)
	if err != nil {
		t.Fatalf("Unable to extract users: %v", err)
	}

	user := allUsers[0]

	allGroupPages, err := users.ListGroups(client, user.ID).AllPages()
	if err != nil {
		t.Fatalf("Unable to list groups: %v", err)
	}

	allGroups, err := groups.ExtractGroups(allGroupPages)
	if err != nil {
		t.Fatalf("Unable to extract groups: %v", err)
	}

	for _, group := range allGroups {
		tools.PrintResource(t, group)
		tools.PrintResource(t, group.Extra)
	}
}

func TestUsersAddToGroup(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	createOpts := users.CreateOpts{
		Password: "foobar",
		DomainID: "default",
		Options: map[users.Option]interface{}{
			users.IgnorePasswordExpiry: true,
			users.MultiFactorAuthRules: []interface{}{
				[]string{"password", "totp"},
				[]string{"password", "custom-auth-method"},
			},
		},
		Extra: map[string]interface{}{
			"email": "jsmith@example.com",
		},
	}

	user, err := CreateUser(t, client, &createOpts)
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}
	defer DeleteUser(t, client, user.ID)

	tools.PrintResource(t, user)
	tools.PrintResource(t, user.Extra)

	createGroupOpts := groups.CreateOpts{
		Name:     "testgroup",
		DomainID: "default",
		Extra: map[string]interface{}{
			"email": "testgroup@example.com",
		},
	}

	// Create Group in the default domain
	group, err := CreateGroup(t, client, &createGroupOpts)
	if err != nil {
		t.Fatalf("Unable to create group: %v", err)
	}
	defer DeleteGroup(t, client, group.ID)

	tools.PrintResource(t, group)
	tools.PrintResource(t, group.Extra)

	err = users.AddToGroup(client, group.ID, user.ID).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to add user to group: %v", err)
	}
}

func TestUsersRemoveFromGroup(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	createOpts := users.CreateOpts{
		Password: "foobar",
		DomainID: "default",
		Options: map[users.Option]interface{}{
			users.IgnorePasswordExpiry: true,
			users.MultiFactorAuthRules: []interface{}{
				[]string{"password", "totp"},
				[]string{"password", "custom-auth-method"},
			},
		},
		Extra: map[string]interface{}{
			"email": "jsmith@example.com",
		},
	}

	user, err := CreateUser(t, client, &createOpts)
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}
	defer DeleteUser(t, client, user.ID)

	tools.PrintResource(t, user)
	tools.PrintResource(t, user.Extra)

	createGroupOpts := groups.CreateOpts{
		Name:     "testgroup",
		DomainID: "default",
		Extra: map[string]interface{}{
			"email": "testgroup@example.com",
		},
	}

	// Create Group in the default domain
	group, err := CreateGroup(t, client, &createGroupOpts)
	if err != nil {
		t.Fatalf("Unable to create group: %v", err)
	}
	defer DeleteGroup(t, client, group.ID)

	tools.PrintResource(t, group)
	tools.PrintResource(t, group.Extra)

	err = users.AddToGroup(client, group.ID, user.ID).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to add user to group: %v", err)
	}

	err = users.RemoveFromGroup(client, group.ID, user.ID).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to remove user from group: %v", err)
	}
}

func TestUsersListProjects(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}
	allUserPages, err := users.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list users: %v", err)
	}

	allUsers, err := users.ExtractUsers(allUserPages)
	if err != nil {
		t.Fatalf("Unable to extract users: %v", err)
	}

	user := allUsers[0]

	allProjectPages, err := users.ListProjects(client, user.ID).AllPages()
	if err != nil {
		t.Fatalf("Unable to list projects: %v", err)
	}

	allProjects, err := projects.ExtractProjects(allProjectPages)
	if err != nil {
		t.Fatalf("Unable to extract projects: %v", err)
	}

	for _, project := range allProjects {
		tools.PrintResource(t, project)
	}
}

func TestUsersListInGroup(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}
	allGroupPages, err := groups.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list groups: %v", err)
	}

	allGroups, err := groups.ExtractGroups(allGroupPages)
	if err != nil {
		t.Fatalf("Unable to extract groups: %v", err)
	}

	group := allGroups[0]

	allUserPages, err := users.ListInGroup(client, group.ID, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list users: %v", err)
	}

	allUsers, err := users.ExtractUsers(allUserPages)
	if err != nil {
		t.Fatalf("Unable to extract users: %v", err)
	}

	for _, user := range allUsers {
		tools.PrintResource(t, user)
		tools.PrintResource(t, user.Extra)
	}
}
