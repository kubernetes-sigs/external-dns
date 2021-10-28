# Infoblox Go Client

An Infoblox NIOS WAPI client library in Golang.
The library enables us to do a CRUD oprations on NIOS Objects.

This library is compatible with Go 1.2+

- [Prerequisites](#Prerequisites)
- [Installation](#Installation)
- [Usage](#Usage)

## Build Status
[![Build Status](https://travis-ci.org/infobloxopen/infoblox-go-client.svg?branch=master)](https://travis-ci.org/infobloxopen/infoblox-go-client) 


## Prerequisites
   * Infoblox GRID with 2.5 or above WAPI support
   * Go 1.2 or above

## Installation
   To get the latest released version [v2.1.0](https://github.com/infobloxopen/infoblox-go-client/releases/tag/v2.1.0) of Go Client use below command.
   
   `go get github.com/infobloxopen/infoblox-go-client/v2`

   To get the previous major version [v1.1.1](https://github.com/infobloxopen/infoblox-go-client/releases/tag/v1.1.1) use below command.
   
   `go get github.com/infobloxopen/infoblox-go-client`

   Note: Go client version v2.0.0 and above have breaking changes and are not backward compatible.

## Usage

   The following is a very simple example for the client usage:

       package main
       import (
   	    "fmt"
   	    ibclient "github.com/infobloxopen/infoblox-go-client"
       )

       func main() {
   	    hostConfig := ibclient.HostConfig{
   		    Host:     "<NIOS grid IP>",
   		    Version:  "<WAPI version>",
   		    Port:     "PORT",
   		    Username: "username",
   		    Password: "password",
   	    }
   	    transportConfig := ibclient.NewTransportConfig("false", 20, 10)
   	    requestBuilder := &ibclient.WapiRequestBuilder{}
   	    requestor := &ibclient.WapiHttpRequestor{}
   	    conn, err := ibclient.NewConnector(hostConfig, transportConfig, requestBuilder, requestor)
   	    if err != nil {
   		    fmt.Println(err)
   	    }
   	    defer conn.Logout()
   	    objMgr := ibclient.NewObjectManager(conn, "myclient", "")
   	    //Fetches grid information
   	    fmt.Println(objMgr.GetLicense())
       }

## Supported NIOS operations

   * AllocateIP
   * AllocateNetwork
   * CreateARecord
   * CreateAAAARecord
   * CreateZoneAuth
   * CreateCNAMERecord
   * CreateDefaultNetviews
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
   * GetEADefinition
   * GetFixedAddress
   * GetFixedAddressByRef
   * GetHostRecord
   * GetHostRecordByRef
   * SearchHostRecordByAltId
   * GetIpAddressFromHostRecord
   * GetNetwork
   * GetNetworkByRef
   * GetNetworkContainer
   * GetNetworkContainerByRef
   * GetNetworkView
   * GetNetworkViewByRef
   * GetPTRRecordByRef
   * GetPTRRecord
   * GetZoneAuthByRef
   * GetZoneDelegated
   * GetUpgradeStatus (2.7 or above)
   * GetAllMembers
   * GetGridInfo
   * GetGridLicense
   * ReleaseIP
   * UpdateAAAARecord
   * UpdateCNAMERecord
   * UpdateFixedAddress
   * UpdateHostRecord
   * UpdateNetwork
   * UpdateNetworkContainer
   * UpdateNetworkView
   * UpdatePTRRecord
   * UpdateARecord
   * UpdateZoneDelegated


