# Infoblox Go Client

An Infoblox NIOS WAPI client library in Golang.
The library enables us to do CRUD operations on NIOS Objects.

This library is compatible with Go 1.25.8 or later.

- [Prerequisites](#Prerequisites)
- [Installation](#Installation)
- [Usage](#Usage)

## Build Status
[![Build Status](https://travis-ci.org/infobloxopen/infoblox-go-client.svg?branch=master)](https://travis-ci.org/infobloxopen/infoblox-go-client) 


## Prerequisites
   * Infoblox GRID with 2.9 or above WAPI support
   * Go 1.25.8 or above
   
    * Note: DNS resource record types HTTPS and SVCB Records (Compatible with NIOS 9.0.7 using WAPI v2.13.7).

## Installation
   To get the latest released version [v2.12.0](https://github.com/infobloxopen/infoblox-go-client/releases/tag/v2.12.0) of Go Client use below command.
   
   `go get github.com/infobloxopen/infoblox-go-client/v2`

   To get the previous major version [v1.1.1](https://github.com/infobloxopen/infoblox-go-client/releases/tag/v1.1.1) use below command.
   
   `go get github.com/infobloxopen/infoblox-go-client`

   Note: Go client version v2.0.0 and above have breaking changes and are not backward compatible.

## Usage

   The following is a very simple example for the client usage:

       package main

       import (
         "fmt"
         ibclient "github.com/infobloxopen/infoblox-go-client/v2"
       )

       func main() {
         hostConfig := ibclient.HostConfig{
            Scheme:  "https",
         	Host:    "<NIOS grid IP>",
            Version: "<WAPI version>",
            Port:    "PORT",
         }
         authConfig := ibclient.AuthConfig{
            Username: "username",
            Password: "password",
         }
         transportConfig := ibclient.NewTransportConfig("false", 20, 10)
         requestBuilder := &ibclient.WapiRequestBuilder{}
         requestor := &ibclient.WapiHttpRequestor{}
         conn, err := ibclient.NewConnector(hostConfig, authConfig, transportConfig, requestBuilder, requestor)
         if err != nil {
         	fmt.Println(err)
         }
         defer conn.Logout()
         objMgr := ibclient.NewObjectManager(conn, "myclient", "")
         //Fetches grid information
         fmt.Println(objMgr.GetGridLicense())
       } 


## Supported NIOS operations

   * AllocateIP
   * AllocateNextAvailableIp
   * AllocateNetwork
   * AllocateNetworkByEA
   * AllocateNetworkContainerByEA
   * AllocateNetworkContainer
   * CreateARecord
   * CreateAAAARecord
   * CreateZoneAuth
   * CreateCNAMERecord
   * CreateDefaultNetviews
   * CreateZoneForward
   * CreateEADefinition
   * CreateHostRecord
   * CreateNetwork
   * CreateNetworkContainer
   * CreateNetworkView
   * CreatePTRRecord
   * CreateTXTRecord
   * CreateZoneDelegated
   * DeleteARecord
   * DeleteAAAARecord
   * DeleteZoneAuth
   * DeleteZoneForward
   * DeleteCNAMERecord
   * DeleteFixedAddress
   * DeleteHostRecord
   * DeleteNetwork
   * DeleteNetworkView
   * DeletePTRRecord
   * DeleteTXTRecord
   * DeleteZoneDelegated
   * GetAllMembers
   * GetARecordByRef
   * GetARecord
   * GetAAAARecordByRef
   * GetAAAARecord
   * GetCapacityReport
   * GetCNAMERecordByRef
   * GetCNAMERecord
   * GetDhcpMember
   * GetDnsMember
   * GetEADefinition
   * GetFixedAddress
   * GetFixedAddressByRef
   * GetHostRecord
   * GetHostRecordByRef
   * SearchHostRecordByAltId
   * SearchObjectByAltId
   * GetIpAddressFromHostRecord
   * GetNetwork
   * GetNetworkByRef
   * GetNetworkContainer
   * GetNetworkContainerByRef
   * GetNetworkView
   * GetNetworkViewByRef
   * GetPTRRecordByRef
   * GetPTRRecord
   * GetTXTRecord
   * GetTXTRecordByRef
   * GetZoneAuthByRef
   * GetZoneDelegated
   * GetUpgradeStatus (2.7 or above)
   * GetAllMembers
   * GetZoneForwardByRef
   * GetZoneForwardFilters
   * GetGridInfo
   * GetGridLicense
   * ReleaseIP
   * UpdateAAAARecord
   * UpdateCNAMERecord
   * UpdateDhcpStatus
   * UpdateDnsStatus
   * UpdateFixedAddress
   * UpdateHostRecord
   * UpdateNetwork
   * UpdateNetworkContainer
   * UpdateNetworkView
   * UpdatePTRRecord
   * UpdateTXTRecord
   * UpdateARecord
   * UpdateZoneDelegated
   * UpdateZoneForward
   * CreateDtcLbdn
   * CreateDtcPool
   * CreateDtcServer
   * DeleteDtcLbdn
   * DeleteDtcPool
   * DeleteDtcServer
   * GetAllDtcPool
   * GetDtcPool
   * GetDtcPoolByRef
   * GetAllDtcServer
   * GetDtcServer
   * GetDtcServerByRef
   * GetAllDtcLbdn
   * GetDtcLbdn
   * GetDtcLbdnByRef
   * UpdateDtcPool
   * UpdateDtcServer
   * UpdateDtcLbdn
   * CreateAliasRecord
   * CreateNSRecord
   * CreateIpv4SharedNetwork
   * CreateNetworkRange
   * CreateRangeTemplate
   * DeleteAliasRecord
   * DeleteNSRecord
   * DeleteIpv4SharedNetwork
   * DeleteNetworkRange
   * DeleteRangeTemplate
   * GetAllAliasRecord
   * GetAllRecordNS
   * GetAllIpv4SharedNetwork
   * GetAllFixedAddress
   * GetNetworkRange
   * GetAllRangeTemplate
   * GetAliasRecordByRef
   * GetNSRecordByRef
   * GetIpv4SharedNetworkByRef
   * GetNetworkRangeByRef
   * GetRangeTemplateByRef
   * UpdateAliasRecord
   * UpdateNSRecord
   * UpdateIpv4SharedNetwork
   * UpdateNetworkRange
   * UpdateRangeTemplate
