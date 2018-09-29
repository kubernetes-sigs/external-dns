/*
Package l7policies provides information and interaction with L7Policies and
Rules of the LBaaS v2 extension for the OpenStack Networking service.
Example to Create a L7Policy
	createOpts := l7policies.CreateOpts{
		Name:        "redirect-example.com",
		ListenerID:  "023f2e34-7806-443b-bfae-16c324569a3d",
		Action:      l7policies.ActionRedirectToURL,
		RedirectURL: "http://www.example.com",
	}
	l7policy, err := l7policies.Create(networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to List L7Policies
	listOpts := l7policies.ListOpts{
		ListenerID: "c79a4468-d788-410c-bf79-9a8ef6354852",
	}
	allPages, err := l7policies.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}
	allL7Policies, err := l7policies.ExtractL7Policies(allPages)
	if err != nil {
		panic(err)
	}
	for _, l7policy := range allL7Policies {
		fmt.Printf("%+v\n", l7policy)
	}
*/
package l7policies
