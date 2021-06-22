package dnsv2

import (
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

type AuthorityResponse struct {
	Contracts []struct {
		ContractID  string   `json:"contractId"`
		Authorities []string `json:"authorities"`
	} `json:"contracts"`
}

func NewAuthorityResponse(contract string) *AuthorityResponse {
	authorities := &AuthorityResponse{}
	return authorities
}

func GetAuthorities(contractId string) (*AuthorityResponse, error) {
	authorities := NewAuthorityResponse(contractId)

	req, err := client.NewRequest(
		Config,
		"GET",
		"/config-dns/v2/data/authorities?contractIds="+contractId,
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpResponse(res, true)

	if client.IsError(res) && res.StatusCode != 404 {
		return nil, client.NewAPIError(res)
	} else if res.StatusCode == 404 {
		return nil, &ZoneError{zoneName: contractId}
	} else {

		err = client.BodyJSON(res, authorities)
		if err != nil {
			return nil, err
		}
		return authorities, nil
	}
}

func GetNameServerRecordList(contractId string) ([]string, error) {

	NSrecords, err := GetAuthorities(contractId)

	if err != nil {
		return nil, err
	}

	var arrLength int
	for _, c := range NSrecords.Contracts {
		arrLength = len(c.Authorities)
	}

	ns := make([]string, 0, arrLength)

	for _, r := range NSrecords.Contracts {
		for _, n := range r.Authorities {
			ns = append(ns, n)
		}
	}
	return ns, nil
}
