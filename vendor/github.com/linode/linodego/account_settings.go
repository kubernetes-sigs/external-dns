package linodego

import (
	"context"
	"encoding/json"
)

// AccountSettings are the account wide flags or plans that effect new resources
type AccountSettings struct {
	// The default backups enrollment status for all new Linodes for all users on the account.  When enabled, backups are mandatory per instance.
	BackupsEnabled bool `json:"backups_enabled"`

	// Wether or not Linode Managed service is enabled for the account.
	Managed bool `json:"managed"`

	// Wether or not the Network Helper is enabled for all new Linode Instance Configs on the account.
	NetworkHelper bool `json:"network_helper"`

	// A plan name like "longview-3"..."longview-100", or a nil value for to cancel any existing subscription plan.
	LongviewSubscription *string `json:"longview_subscription"`

	// A string like "disabled", "suspended", or "active" describing the status of this account’s Object Storage service enrollment.
	ObjectStorage *string `json:"object_storage"`
}

// AccountSettingsUpdateOptions are the updateable account wide flags or plans that effect new resources.
type AccountSettingsUpdateOptions struct {
	// The default backups enrollment status for all new Linodes for all users on the account.  When enabled, backups are mandatory per instance.
	BackupsEnabled *bool `json:"backups_enabled,omitempty"`

	// A plan name like "longview-3"..."longview-100", or a nil value for to cancel any existing subscription plan.
	// Deprecated: Use PUT /longview/plan instead to update the LongviewSubscription
	LongviewSubscription *string `json:"longview_subscription,omitempty"`

	// The default network helper setting for all new Linodes and Linode Configs for all users on the account.
	NetworkHelper *bool `json:"network_helper,omitempty"`
}

// GetAccountSettings gets the account wide flags or plans that effect new resources
func (c *Client) GetAccountSettings(ctx context.Context) (*AccountSettings, error) {
<<<<<<< HEAD
	e, err := c.AccountSettings.Endpoint()
	if err != nil {
		return nil, err
	}

	r, err := coupleAPIErrors(c.R(ctx).SetResult(&AccountSettings{}).Get(e))
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	if err != nil {
		return nil, err
	}

	return r.Result().(*AccountSettings), nil
}

// UpdateAccountSettings updates the settings associated with the account
func (c *Client) UpdateAccountSettings(ctx context.Context, settings AccountSettingsUpdateOptions) (*AccountSettings, error) {
	var body string

	e, err := c.AccountSettings.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&AccountSettings{})

	if bodyData, err := json.Marshal(settings); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of 5ce8c7613 (update vendored files)

=======
>>>>>>> 5ce8c7613 (update vendored files)
	if err != nil {
		return nil, err
	}

	return r.Result().(*AccountSettings), nil
}

// UpdateAccountSettings updates the settings associated with the account
func (c *Client) UpdateAccountSettings(ctx context.Context, settings AccountSettingsUpdateOptions) (*AccountSettings, error) {
	var body string

	e, err := c.AccountSettings.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&AccountSettings{})

	if bodyData, err := json.Marshal(settings); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
<<<<<<< HEAD

>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)

=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of 6b7ce455e (update vendored files)

=======
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		return nil, err
	}

	return r.Result().(*AccountSettings), nil
}

// UpdateAccountSettings updates the settings associated with the account
func (c *Client) UpdateAccountSettings(ctx context.Context, settings AccountSettingsUpdateOptions) (*AccountSettings, error) {
	var body string

	e, err := c.AccountSettings.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&AccountSettings{})

	if bodyData, err := json.Marshal(settings); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
<<<<<<< HEAD

>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of 4d7e5ad26 (update vendored files)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
	if err != nil {
		return nil, err
	}

	return r.Result().(*AccountSettings), nil
}

// UpdateAccountSettings updates the settings associated with the account
func (c *Client) UpdateAccountSettings(ctx context.Context, settings AccountSettingsUpdateOptions) (*AccountSettings, error) {
	var body string

	e, err := c.AccountSettings.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&AccountSettings{})

	if bodyData, err := json.Marshal(settings); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
<<<<<<< HEAD

>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	e, err := c.AccountSettings.Endpoint()
	if err != nil {
		return nil, err
	}

	r, err := coupleAPIErrors(c.R(ctx).SetResult(&AccountSettings{}).Get(e))

=======
	req := c.R(ctx).SetResult(&AccountSettings{})
	e := "account/settings"
	r, err := coupleAPIErrors(req.Get(e))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return nil, err
	}

	return r.Result().(*AccountSettings), nil
}

// UpdateAccountSettings updates the settings associated with the account
func (c *Client) UpdateAccountSettings(ctx context.Context, opts AccountSettingsUpdateOptions) (*AccountSettings, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	req := c.R(ctx).SetResult(&AccountSettings{})

	if bodyData, err := json.Marshal(settings); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))

>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	req := c.R(ctx).SetResult(&AccountSettings{})

	if bodyData, err := json.Marshal(settings); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))

=======
	req := c.R(ctx).SetResult(&AccountSettings{}).SetBody(string(body))
	e := "account/settings"
	r, err := coupleAPIErrors(req.Put(e))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return nil, err
	}

	return r.Result().(*AccountSettings), nil
}
