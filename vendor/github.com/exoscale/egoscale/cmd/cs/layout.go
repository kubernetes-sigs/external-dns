package main

import (
	"github.com/exoscale/egoscale"
)

type cmd struct {
	command egoscale.Command
	hidden  bool
}

var methods = map[string][]cmd{
	"network": {
		{&egoscale.CreateNetwork{}, false},
		{&egoscale.DeleteNetwork{}, false},
		{&egoscale.ListNetworkOfferings{}, false},
		{&egoscale.ListNetworks{}, false},
		{&egoscale.RestartNetwork{}, true},
		{&egoscale.UpdateNetwork{}, false},
	},
	"virtual machine": {
		{&egoscale.AddNicToVirtualMachine{}, false},
		{&egoscale.ChangeServiceForVirtualMachine{}, false},
		{&egoscale.DeployVirtualMachine{}, false},
		{&egoscale.DestroyVirtualMachine{}, false},
		{&egoscale.ExpungeVirtualMachine{}, false},
		{&egoscale.GetVMPassword{}, false},
		{&egoscale.GetVirtualMachineUserData{}, false},
		{&egoscale.ListVirtualMachines{}, false},
		{&egoscale.MigrateVirtualMachine{}, true},
		{&egoscale.RebootVirtualMachine{}, false},
		{&egoscale.RecoverVirtualMachine{}, false},
		{&egoscale.RemoveNicFromVirtualMachine{}, false},
		{&egoscale.ResetPasswordForVirtualMachine{}, false},
		{&egoscale.RestoreVirtualMachine{}, false},
		{&egoscale.ScaleVirtualMachine{}, false},
		{&egoscale.StartVirtualMachine{}, false},
		{&egoscale.StopVirtualMachine{}, false},
		{&egoscale.UpdateDefaultNicForVirtualMachine{}, true},
		{&egoscale.UpdateVirtualMachine{}, false},
	},
	"volume": {
		{&egoscale.ListVolumes{}, false},
		{&egoscale.ResizeVolume{}, false},
	},
	"template": {
		{&egoscale.CopyTemplate{}, true},
		{&egoscale.CreateTemplate{}, true},
		{&egoscale.ListTemplates{}, false},
		{&egoscale.PrepareTemplate{}, true},
		{&egoscale.RegisterTemplate{}, true},
		{&egoscale.ListOSCategories{}, true},
	},
	"account": {
		{&egoscale.EnableAccount{}, true},
		{&egoscale.DisableAccount{}, true},
		{&egoscale.ListAccounts{}, false},
	},
	"zone": {
		{&egoscale.ListZones{}, false},
	},
	"snapshot": {
		{&egoscale.CreateSnapshot{}, false},
		{&egoscale.DeleteSnapshot{}, false},
		{&egoscale.ListSnapshots{}, false},
		{&egoscale.RevertSnapshot{}, false},
	},
	"user": {
		{&egoscale.CreateUser{}, true},
		{&egoscale.DeleteUser{}, true},
		//{&egoscale.DisableUser{}, true},
		//{&egoscale.GetUser{}, true},
		{&egoscale.UpdateUser{}, true},
		{&egoscale.ListUsers{}, false},
		{&egoscale.RegisterUserKeys{}, false},
	},
	"security group": {
		{&egoscale.AuthorizeSecurityGroupEgress{}, false},
		{&egoscale.AuthorizeSecurityGroupIngress{}, false},
		{&egoscale.CreateSecurityGroup{}, false},
		{&egoscale.DeleteSecurityGroup{}, false},
		{&egoscale.ListSecurityGroups{}, false},
		{&egoscale.RevokeSecurityGroupEgress{}, false},
		{&egoscale.RevokeSecurityGroupIngress{}, false},
	},
	"ssh": {
		{&egoscale.RegisterSSHKeyPair{}, false},
		{&egoscale.ListSSHKeyPairs{}, false},
		{&egoscale.CreateSSHKeyPair{}, false},
		{&egoscale.DeleteSSHKeyPair{}, false},
		{&egoscale.ResetSSHKeyForVirtualMachine{}, false},
	},
	"affinity group": {
		{&egoscale.CreateAffinityGroup{}, false},
		{&egoscale.DeleteAffinityGroup{}, false},
		{&egoscale.ListAffinityGroups{}, false},
		{&egoscale.UpdateVMAffinityGroup{}, false},
	},
	"vm group": {
		{&egoscale.CreateInstanceGroup{}, false},
		{&egoscale.ListInstanceGroups{}, false},
	},
	"tags": {
		{&egoscale.CreateTags{}, false},
		{&egoscale.DeleteTags{}, false},
		{&egoscale.ListTags{}, false},
	},
	"nic": {
		{&egoscale.ActivateIP6{}, false},
		{&egoscale.AddIPToNic{}, false},
		{&egoscale.ListNics{}, false},
		{&egoscale.RemoveIPFromNic{}, false},
	},
	"address": {
		{&egoscale.AssociateIPAddress{}, false},
		{&egoscale.DisassociateIPAddress{}, false},
		{&egoscale.ListPublicIPAddresses{}, false},
		{&egoscale.UpdateIPAddress{}, false},
	},
	"async job": {
		{&egoscale.QueryAsyncJobResult{}, false},
		{&egoscale.ListAsyncJobs{}, false},
	},
	"apis": {
		{&egoscale.ListAPIs{}, false},
	},
	"event": {
		{&egoscale.ListEventTypes{}, false},
		{&egoscale.ListEvents{}, false},
	},
	"offerings": {
		{&egoscale.ListResourceDetails{}, false},
		{&egoscale.ListResourceLimits{}, false},
		{&egoscale.ListServiceOfferings{}, false},
	},
	"host": {
		{&egoscale.ListHosts{}, true},
		{&egoscale.UpdateHost{}, true},
	},
	"reversedns": {
		{&egoscale.DeleteReverseDNSFromPublicIPAddress{}, false},
		{&egoscale.QueryReverseDNSForPublicIPAddress{}, false},
		{&egoscale.UpdateReverseDNSForPublicIPAddress{}, false},
		{&egoscale.DeleteReverseDNSFromVirtualMachine{}, false},
		{&egoscale.QueryReverseDNSForVirtualMachine{}, false},
		{&egoscale.UpdateReverseDNSForVirtualMachine{}, false},
	},
}
