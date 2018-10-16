/*
Package policies provides information and interaction with the policies through
the OpenStack Clustering service.

Example to List Policies

	listOpts := policies.ListOpts{
		Limit: 2,
	}

	allPages, err := policies.List(clusteringClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allPolicies, err := policies.ExtractPolicies(allPages)
	if err != nil {
		panic(err)
	}

	for _, policy := range allPolicies {
		fmt.Printf("%+v\n", policy)
	}

*/
package policies
