package egoscale

import (
	"net/url"
	"testing"
)

func TestSecurityGroup(t *testing.T) {
	instance := &SecurityGroup{
		Account: "Hello",
		Name:    "world",
	}
	if instance.ResourceType() != "SecurityGroup" {
		t.Error("ResourceType doesn't match")
	}

	usg := instance.UserSecurityGroup()
	if usg.Account != instance.Account && usg.Group != instance.Name {
		t.Error("got bad user security group")
	}
}

func TestAuthorizeSecurityGroupEgress(t *testing.T) {
	req := &AuthorizeSecurityGroupEgress{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*SecurityGroup)
}

func TestAuthorizeSecurityGroupIngress(t *testing.T) {
	req := &AuthorizeSecurityGroupIngress{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*SecurityGroup)
}

func TestCreateSecurityGroup(t *testing.T) {
	req := &CreateSecurityGroup{}
	_ = req.response().(*SecurityGroup)
}

func TestDeleteSecurityGroup(t *testing.T) {
	req := &DeleteSecurityGroup{}
	_ = req.response().(*booleanResponse)
}

func TestListSecurityGroupsApiName(t *testing.T) {
	req := &ListSecurityGroups{}
	_ = req.response().(*ListSecurityGroupsResponse)
}

func TestRevokeSecurityGroupEgress(t *testing.T) {
	req := &RevokeSecurityGroupEgress{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestRevokeSecurityGroupIngress(t *testing.T) {
	req := &RevokeSecurityGroupIngress{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestAuthorizeSecurityGroupEgressOnBeforeSendICMP(t *testing.T) {
	req := &AuthorizeSecurityGroupEgress{
		Protocol: "icmp",
	}
	params := url.Values{}

	if err := req.onBeforeSend(params); err != nil {
		t.Error(err)
	}

	if _, ok := params["icmpcode"]; !ok {
		t.Errorf("icmpcode should have been set")
	}
	if _, ok := params["icmptype"]; !ok {
		t.Errorf("icmptype should have been set")
	}
}

func TestAuthorizeSecurityGroupEgressOnBeforeSendICMPv6(t *testing.T) {
	req := &AuthorizeSecurityGroupEgress{
		Protocol: "ICMPv6",
	}
	params := url.Values{}

	if err := req.onBeforeSend(params); err != nil {
		t.Error(err)
	}

	if _, ok := params["icmpcode"]; !ok {
		t.Errorf("icmpcode should have been set")
	}
	if _, ok := params["icmptype"]; !ok {
		t.Errorf("icmptype should have been set")
	}
}

func TestAuthorizeSecurityGroupEgressOnBeforeSendTCP(t *testing.T) {
	req := &AuthorizeSecurityGroupEgress{
		Protocol:  "TCP",
		StartPort: 0,
		EndPort:   1024,
	}
	params := url.Values{}

	if err := req.onBeforeSend(params); err != nil {
		t.Error(err)
	}

	if _, ok := params["startport"]; !ok {
		t.Errorf("startport should have been set")
	}
}

func TestGetSecurityGroup(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listsecuritygroupsresponse": {
	"count": 1,
	"securitygroup": [
		{
			"account": "yoan.blanc@exoscale.ch",
			"description": "dummy (for test)",
			"domain": "yoan.blanc@exoscale.ch",
			"domainid": "2da0d0d3-e7b2-42ef-805d-eb2ea90ae7ef",
			"egressrule": [],
			"id": "4bfe1073-a6d4-48bd-8f24-2ab586674092",
			"ingressrule": [
				{
					"cidr": "0.0.0.0/0",
					"description": "SSH",
					"endport": 22,
					"protocol": "tcp",
					"ruleid": "fc03b5b1-1d15-4933-99c3-afa0b8f2ab25",
					"startport": 22,
					"tags": []
				}
			],
			"name": "ssh",
			"tags": []
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	sg := &SecurityGroup{
		ID: MustParseUUID("4bfe1073-a6d4-48bd-8f24-2ab586674092"),
	}
	if err := cs.Get(sg); err != nil {
		t.Error(err)
	}

	if sg.Account != "yoan.blanc@exoscale.ch" {
		t.Errorf("Account doesn't match, got %v", sg.Account)
	}
}

func TestListSecurityGroups(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
		{"listsecuritygroupsresponse":{
			"count": 2,
			"securitygroup": [
			  {
				"account": "exoscale-1",
				"description": "test",
				"domain": "exoscale-1",
				"domainid": "5b2f621e-3eb6-4a14-a315-d4d7d62f28ff",
				"egressrule": [],
				"id": "55c3b385-0a9b-4970-a5d9-ad1e7f13157d",
				"ingressrule": [
				  {
					"cidr": "0.0.0.0/0",
					"icmpcode": 0,
					"icmptype": 8,
					"protocol": "icmp",
					"ruleid": "1d64f828-9267-4aeb-9cf9-703c3ed99627",
					"tags": []
				  },
				  {
					"cidr": "0.0.0.0/0",
					"endport": 22,
					"protocol": "tcp",
					"ruleid": "1a84c747-1ad5-48ea-a75d-521638c403ea",
					"startport": 22,
					"tags": []
				  },
				  {
					"cidr": "0.0.0.0/0",
					"endport": 3389,
					"protocol": "tcp",
					"ruleid": "f2ab2e27-65a1-40b8-b8c2-9252dc75b5b3",
					"startport": 3389,
					"tags": []
				  }
				],
				"name": "hello",
				"tags": []
			  },
			  {
				"account": "exoscale-1",
				"description": "Default Security Group",
				"domain": "exoscale-1",
				"domainid": "5b2f621e-3eb6-4a14-a315-d4d7d62f28ff",
				"egressrule": [],
				"id": "b1b05d21-11de-4c38-804e-c9bdacdaaa70",
				"egressrule": [
				  {
					"cidr": "0.0.0.0/0",
					"endport": 25,
					"protocol": "tcp",
					"ruleid": "4aa47b2c-9d1f-4856-9893-286fb8befa76",
					"startport": 25,
					"tags": []
				  }
				],
				"name": "default",
				"tags": []
			  }
			]
		  }}`})

	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	sgs, err := cs.List(&SecurityGroup{})
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(sgs) != 2 {
		t.Errorf("Expected two sg, got %d", len(sgs))
	}

	sg := sgs[0].(*SecurityGroup)

	if !sg.ID.Equal(*MustParseUUID("55c3b385-0a9b-4970-a5d9-ad1e7f13157d")) {
		t.Errorf("Wrong security group")
	}

	ig, _ := sg.RuleByID(*MustParseUUID("f2ab2e27-65a1-40b8-b8c2-9252dc75b5b3"))
	if ig == nil {
		t.Errorf("ingress rule not found")
	}

	sg = sgs[1].(*SecurityGroup)

	_, eg := sg.RuleByID(*MustParseUUID("4aa47b2c-9d1f-4856-9893-286fb8befa76"))
	if eg == nil {
		t.Errorf("egress rule not found")
	}

}
