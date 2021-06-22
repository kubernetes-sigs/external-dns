/*
Package openstack contains resources for the individual OpenStack projects
supported in Gophercloud. It also includes functions to authenticate to an
OpenStack cloud and for provisioning various service-level clients.

Example of Creating a Service Client

	ao, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(ao)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	client, err := openstack.NewNetworkV2(client, gophercloud.EndpointOpts{
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	client, err := openstack.NewNetworkV2(client, gophercloud.EndpointOpts{
=======
	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	client, err := openstack.NewNetworkV2(client, gophercloud.EndpointOpts{
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	client, err := openstack.NewNetworkV2(client, gophercloud.EndpointOpts{
=======
	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	client, err := openstack.NewNetworkV2(client, gophercloud.EndpointOpts{
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
		Region: os.Getenv("OS_REGION_NAME"),
	})
*/
package openstack
