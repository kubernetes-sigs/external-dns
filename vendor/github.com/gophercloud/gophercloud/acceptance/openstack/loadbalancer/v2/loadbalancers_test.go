// +build acceptance networking loadbalancer loadbalancers

package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	networking "github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/listeners"
	"github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/loadbalancers"
	"github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/monitors"
	"github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/pools"
)

func TestLoadbalancersList(t *testing.T) {
	client, err := clients.NewLoadBalancerV2Client()
	if err != nil {
		t.Fatalf("Unable to create a loadbalancer client: %v", err)
	}

	allPages, err := loadbalancers.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list loadbalancers: %v", err)
	}

	allLoadbalancers, err := loadbalancers.ExtractLoadBalancers(allPages)
	if err != nil {
		t.Fatalf("Unable to extract loadbalancers: %v", err)
	}

	for _, lb := range allLoadbalancers {
		tools.PrintResource(t, lb)
	}
}

func TestLoadbalancersCRUD(t *testing.T) {
	netClient, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a networking client: %v", err)
	}

	lbClient, err := clients.NewLoadBalancerV2Client()
	if err != nil {
		t.Fatalf("Unable to create a loadbalancer client: %v", err)
	}

	network, err := networking.CreateNetwork(t, netClient)
	if err != nil {
		t.Fatalf("Unable to create network: %v", err)
	}
	defer networking.DeleteNetwork(t, netClient, network.ID)

	subnet, err := networking.CreateSubnet(t, netClient, network.ID)
	if err != nil {
		t.Fatalf("Unable to create subnet: %v", err)
	}
	defer networking.DeleteSubnet(t, netClient, subnet.ID)

	lb, err := CreateLoadBalancer(t, lbClient, subnet.ID)
	if err != nil {
		t.Fatalf("Unable to create loadbalancer: %v", err)
	}
	defer DeleteLoadBalancer(t, lbClient, lb.ID)

	newLB, err := loadbalancers.Get(lbClient, lb.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get loadbalancer: %v", err)
	}

	tools.PrintResource(t, newLB)

	// Because of the time it takes to create a loadbalancer,
	// this test will include some other resources.

	// Listener
	listener, err := CreateListener(t, lbClient, lb)
	if err != nil {
		t.Fatalf("Unable to create listener: %v", err)
	}
	defer DeleteListener(t, lbClient, lb.ID, listener.ID)

	updateListenerOpts := listeners.UpdateOpts{
		Description: "Some listener description",
	}
	_, err = listeners.Update(lbClient, listener.ID, updateListenerOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update listener")
	}

	if err := WaitForLoadBalancerState(lbClient, lb.ID, "ACTIVE", loadbalancerActiveTimeoutSeconds); err != nil {
		t.Fatalf("Timed out waiting for loadbalancer to become active")
	}

	newListener, err := listeners.Get(lbClient, listener.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get listener")
	}

	tools.PrintResource(t, newListener)

	// L7 policy
	_, err = CreateL7Policy(t, lbClient, listener, lb)
	if err != nil {
		t.Fatalf("Unable to create l7 policy: %v", err)
	}

	// Pool
	pool, err := CreatePool(t, lbClient, lb)
	if err != nil {
		t.Fatalf("Unable to create pool: %v", err)
	}
	defer DeletePool(t, lbClient, lb.ID, pool.ID)

	updatePoolOpts := pools.UpdateOpts{
		Description: "Some pool description",
	}
	_, err = pools.Update(lbClient, pool.ID, updatePoolOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update pool")
	}

	if err := WaitForLoadBalancerState(lbClient, lb.ID, "ACTIVE", loadbalancerActiveTimeoutSeconds); err != nil {
		t.Fatalf("Timed out waiting for loadbalancer to become active")
	}

	newPool, err := pools.Get(lbClient, pool.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get pool")
	}

	tools.PrintResource(t, newPool)

	// Member
	member, err := CreateMember(t, lbClient, lb, newPool, subnet.ID, subnet.CIDR)
	if err != nil {
		t.Fatalf("Unable to create member: %v", err)
	}
	defer DeleteMember(t, lbClient, lb.ID, pool.ID, member.ID)

	newWeight := tools.RandomInt(11, 100)
	updateMemberOpts := pools.UpdateMemberOpts{
		Weight: newWeight,
	}
	_, err = pools.UpdateMember(lbClient, pool.ID, member.ID, updateMemberOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update pool")
	}

	if err := WaitForLoadBalancerState(lbClient, lb.ID, "ACTIVE", loadbalancerActiveTimeoutSeconds); err != nil {
		t.Fatalf("Timed out waiting for loadbalancer to become active")
	}

	newMember, err := pools.GetMember(lbClient, pool.ID, member.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get member")
	}

	tools.PrintResource(t, newMember)

	// Monitor
	monitor, err := CreateMonitor(t, lbClient, lb, newPool)
	if err != nil {
		t.Fatalf("Unable to create monitor: %v", err)
	}
	defer DeleteMonitor(t, lbClient, lb.ID, monitor.ID)

	newDelay := tools.RandomInt(20, 30)
	updateMonitorOpts := monitors.UpdateOpts{
		Delay: newDelay,
	}
	_, err = monitors.Update(lbClient, monitor.ID, updateMonitorOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update monitor")
	}

	if err := WaitForLoadBalancerState(lbClient, lb.ID, "ACTIVE", loadbalancerActiveTimeoutSeconds); err != nil {
		t.Fatalf("Timed out waiting for loadbalancer to become active")
	}

	newMonitor, err := monitors.Get(lbClient, monitor.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get monitor")
	}

	tools.PrintResource(t, newMonitor)

}

func TestLoadbalancersCascadeCRUD(t *testing.T) {
	netClient, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a networking client: %v", err)
	}

	lbClient, err := clients.NewLoadBalancerV2Client()
	if err != nil {
		t.Fatalf("Unable to create a loadbalancer client: %v", err)
	}

	network, err := networking.CreateNetwork(t, netClient)
	if err != nil {
		t.Fatalf("Unable to create network: %v", err)
	}
	defer networking.DeleteNetwork(t, netClient, network.ID)

	subnet, err := networking.CreateSubnet(t, netClient, network.ID)
	if err != nil {
		t.Fatalf("Unable to create subnet: %v", err)
	}
	defer networking.DeleteSubnet(t, netClient, subnet.ID)

	lb, err := CreateLoadBalancer(t, lbClient, subnet.ID)
	if err != nil {
		t.Fatalf("Unable to create loadbalancer: %v", err)
	}
	defer CascadeDeleteLoadBalancer(t, lbClient, lb.ID)

	newLB, err := loadbalancers.Get(lbClient, lb.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get loadbalancer: %v", err)
	}

	tools.PrintResource(t, newLB)

	// Because of the time it takes to create a loadbalancer,
	// this test will include some other resources.

	// Listener
	listener, err := CreateListener(t, lbClient, lb)
	if err != nil {
		t.Fatalf("Unable to create listener: %v", err)
	}

	updateListenerOpts := listeners.UpdateOpts{
		Description: "Some listener description",
	}
	_, err = listeners.Update(lbClient, listener.ID, updateListenerOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update listener")
	}

	if err := WaitForLoadBalancerState(lbClient, lb.ID, "ACTIVE", loadbalancerActiveTimeoutSeconds); err != nil {
		t.Fatalf("Timed out waiting for loadbalancer to become active")
	}

	newListener, err := listeners.Get(lbClient, listener.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get listener")
	}

	tools.PrintResource(t, newListener)

	// Pool
	pool, err := CreatePool(t, lbClient, lb)
	if err != nil {
		t.Fatalf("Unable to create pool: %v", err)
	}

	updatePoolOpts := pools.UpdateOpts{
		Description: "Some pool description",
	}
	_, err = pools.Update(lbClient, pool.ID, updatePoolOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update pool")
	}

	if err := WaitForLoadBalancerState(lbClient, lb.ID, "ACTIVE", loadbalancerActiveTimeoutSeconds); err != nil {
		t.Fatalf("Timed out waiting for loadbalancer to become active")
	}

	newPool, err := pools.Get(lbClient, pool.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get pool")
	}

	tools.PrintResource(t, newPool)

	// Member
	member, err := CreateMember(t, lbClient, lb, newPool, subnet.ID, subnet.CIDR)
	if err != nil {
		t.Fatalf("Unable to create member: %v", err)
	}

	newWeight := tools.RandomInt(11, 100)
	updateMemberOpts := pools.UpdateMemberOpts{
		Weight: newWeight,
	}
	_, err = pools.UpdateMember(lbClient, pool.ID, member.ID, updateMemberOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update pool")
	}

	if err := WaitForLoadBalancerState(lbClient, lb.ID, "ACTIVE", loadbalancerActiveTimeoutSeconds); err != nil {
		t.Fatalf("Timed out waiting for loadbalancer to become active")
	}

	newMember, err := pools.GetMember(lbClient, pool.ID, member.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get member")
	}

	tools.PrintResource(t, newMember)

	// Monitor
	monitor, err := CreateMonitor(t, lbClient, lb, newPool)
	if err != nil {
		t.Fatalf("Unable to create monitor: %v", err)
	}

	newDelay := tools.RandomInt(20, 30)
	updateMonitorOpts := monitors.UpdateOpts{
		Delay: newDelay,
	}
	_, err = monitors.Update(lbClient, monitor.ID, updateMonitorOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update monitor")
	}

	if err := WaitForLoadBalancerState(lbClient, lb.ID, "ACTIVE", loadbalancerActiveTimeoutSeconds); err != nil {
		t.Fatalf("Timed out waiting for loadbalancer to become active")
	}

	newMonitor, err := monitors.Get(lbClient, monitor.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get monitor")
	}

	tools.PrintResource(t, newMonitor)

}
