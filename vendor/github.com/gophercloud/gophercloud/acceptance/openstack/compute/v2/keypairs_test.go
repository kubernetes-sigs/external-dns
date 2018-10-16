// +build acceptance compute keypairs

package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	th "github.com/gophercloud/gophercloud/testhelper"
)

const keyName = "gophercloud_test_key_pair"

func TestKeypairsCreateDelete(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	keyPair, err := CreateKeyPair(t, client)
	th.AssertNoErr(t, err)
	defer DeleteKeyPair(t, client, keyPair)

	tools.PrintResource(t, keyPair)

	allPages, err := keypairs.List(client).AllPages()
	th.AssertNoErr(t, err)

	allKeys, err := keypairs.ExtractKeyPairs(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, kp := range allKeys {
		tools.PrintResource(t, kp)

		if kp.Name == keyPair.Name {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestKeypairsImportPublicKey(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	publicKey, err := createKey()
	th.AssertNoErr(t, err)

	keyPair, err := ImportPublicKey(t, client, publicKey)
	th.AssertNoErr(t, err)
	defer DeleteKeyPair(t, client, keyPair)

	tools.PrintResource(t, keyPair)
}

func TestKeypairsServerCreateWithKey(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	publicKey, err := createKey()
	th.AssertNoErr(t, err)

	keyPair, err := ImportPublicKey(t, client, publicKey)
	th.AssertNoErr(t, err)
	defer DeleteKeyPair(t, client, keyPair)

	server, err := CreateServerWithPublicKey(t, client, keyPair.Name)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	server, err = servers.Get(client, server.ID).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, server.KeyName, keyPair.Name)
}
