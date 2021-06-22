# Infoblox Go Client

An Infoblox NIOS WAPI client library in Golang.
The library enables us to do a CRUD oprations on NIOS Objects.

This library is compatible with Go 1.2+

- [Prerequisites](#Prerequisites)
- [Installation](#Installation)
- [Usage](#Usage)

## Build Status

| Master                                                                                                                                          | Develop                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [![Build Status](https://travis-ci.org/infobloxopen/infoblox-go-client.svg?branch=master)](https://travis-ci.org/infobloxopen/infoblox-go-client) | [![Build Status](https://travis-ci.org/infobloxopen/infoblox-go-client.svg?branch=develop)](https://travis-ci.org/infobloxopen/infoblox-go-client) |


## Prerequisites
   * Infoblox GRID with 2.5 or above WAPI support
   * Go 1.2 or above

## Installation
   go get github.com/infobloxopen/infoblox-go-client

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
   * CreateZoneAuth
   * CreateCNAMERecord
   * CreateDefaultNetviews
   * CreateEADefinition
   * CreateHostRecord
   * CreateNetwork
   * CreateNetworkContainer
   * CreateNetworkView
   * CreatePTRRecord
   * DeleteARecord
   * DeleteZoneAuth
   * DeleteCNAMERecord
   * DeleteFixedAddress
   * DeleteHostRecord
   * DeleteNetwork
   * DeleteNetworkView
   * DeletePTRRecord
   * GetAllMembers
   * GetARecordByRef
   * GetCapacityReport
   * GetCNAMERecordByRef
   * GetEADefinition
   * GetFixedAddress
   * GetFixedAddressByRef
   * GetHostRecord
   * GetHostRecordByRef
   * GetIpAddressFromHostRecord
   * GetNetwork
   * GetNetworkContainer
   * GetNetworkView
   * GetPTRRecordByRef
   * GetUpgradeStatus (2.7 or above)
   * GetZoneAuthByRef
   * ReleaseIP
   * UpdateFixedAddress
   * UpdateHostRecord
   * UpdateNetworkViewEA
   * UpdateARecord
   * UpdateCNAMERecord

The newly developed features will be available under `develop` branch. After validation they would be merged to `master`.
