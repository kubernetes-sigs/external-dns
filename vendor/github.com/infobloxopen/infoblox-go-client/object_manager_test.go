package ibclient

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
)

type fakeConnector struct {
	createObjectObj interface{}

	getObjectObj interface{}
	getObjectRef string

	deleteObjectRef string

	updateObjectObj interface{}
	updateObjectRef string

	resultObject interface{}

	fakeRefReturn string
}

func (c *fakeConnector) CreateObject(obj IBObject) (string, error) {
	Expect(obj).To(Equal(c.createObjectObj))

	return c.fakeRefReturn, nil
}

func (c *fakeConnector) GetObject(obj IBObject, ref string, res interface{}) (err error) {
	Expect(obj).To(Equal(c.getObjectObj))
	Expect(ref).To(Equal(c.getObjectRef))

	if ref == "" {
		switch obj.(type) {
		case *NetworkView:
			*res.(*[]NetworkView) = c.resultObject.([]NetworkView)
		case *NetworkContainer:
			*res.(*[]NetworkContainer) = c.resultObject.([]NetworkContainer)
		case *Network:
			*res.(*[]Network) = c.resultObject.([]Network)
		case *FixedAddress:
			*res.(*[]FixedAddress) = c.resultObject.([]FixedAddress)
		case *EADefinition:
			*res.(*[]EADefinition) = c.resultObject.([]EADefinition)
		}
	} else {
		switch obj.(type) {
		case *NetworkView:
			*res.(*NetworkView) = c.resultObject.(NetworkView)
		}
	}

	err = nil
	return
}

func (c *fakeConnector) DeleteObject(ref string) (string, error) {
	Expect(ref).To(Equal(c.deleteObjectRef))

	return c.fakeRefReturn, nil
}

func (c *fakeConnector) UpdateObject(obj IBObject, ref string) (string, error) {
	Expect(obj).To(Equal(c.updateObjectObj))
	Expect(ref).To(Equal(c.updateObjectRef))

	return c.fakeRefReturn, nil
}

var _ = Describe("Object Manager", func() {

	Describe("Create Network View", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "Default View"
		fakeRefReturn := "networkview/ZG5zLm5ldHdvcmtfdmlldyQyMw:global_view/false"
		nvFakeConnector := &fakeConnector{
			createObjectObj: NewNetworkView(NetworkView{Name: netviewName}),
			resultObject:    NewNetworkView(NetworkView{Name: netviewName, Ref: fakeRefReturn}),
			fakeRefReturn:   fakeRefReturn,
		}

		objMgr := NewObjectManager(nvFakeConnector, cmpType, tenantID)
		nvFakeConnector.createObjectObj.(*NetworkView).Ea = objMgr.getBasicEA(false)
		nvFakeConnector.resultObject.(*NetworkView).Ea = objMgr.getBasicEA(false)

		var actualNetworkView *NetworkView
		var err error
		It("should pass expected NetworkView Object to CreateObject", func() {
			actualNetworkView, err = objMgr.CreateNetworkView(netviewName)
		})
		It("should return expected NetworkView Object", func() {
			Expect(actualNetworkView).To(Equal(nvFakeConnector.resultObject))
			Expect(err).To(BeNil())
		})
	})

	Describe("Update Network View", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "Global View"
		fakeRefReturn := "networkview/ZG5zLm5ldHdvcmtfdmlldyQyMw:global_view/false"

		returnGetObject := NetworkView{Name: netviewName, Ref: fakeRefReturn, Ea: EA{"network-name": "net1", "Lock": "Removed"}}
		returnUpdateObject := NetworkView{Name: netviewName, Ref: fakeRefReturn, Ea: EA{"network-name": "net2", "New": "Added"}}
		getObjectObj := &NetworkView{}
		getObjectObj.returnFields = []string{"extattrs"}
		nvFakeConnector := &fakeConnector{
			getObjectObj:    getObjectObj,
			getObjectRef:    fakeRefReturn,
			fakeRefReturn:   fakeRefReturn,
			resultObject:    returnGetObject,
			updateObjectObj: &returnUpdateObject,
			updateObjectRef: fakeRefReturn,
		}

		objMgr := NewObjectManager(nvFakeConnector, cmpType, tenantID)

		var err error
		It("should pass expected updated object to UpdateObject", func() {
			addEA := EA{"network-name": "net2", "New": "Added"}
			delEA := EA{"Lock": "Removed"}
			err = objMgr.UpdateNetworkViewEA(fakeRefReturn, addEA, delEA)
		})
		It("should updated the GetObject with new EA and with no error", func() {
			Expect(returnGetObject).To(Equal(returnUpdateObject))
			Expect(err).To(BeNil())
		})
	})

	Describe("Create Network Container", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "Default View"
		cidr := "43.0.11.0/24"
		fakeRefReturn := "networkcontainer/ZG5zLm5ldHdvcmtfdmlldyQyMw:global_view/false"
		ncFakeConnector := &fakeConnector{
			createObjectObj: NewNetworkContainer(NetworkContainer{NetviewName: netviewName, Cidr: cidr}),
			resultObject:    NewNetworkContainer(NetworkContainer{NetviewName: netviewName, Cidr: cidr, Ref: fakeRefReturn}),
			fakeRefReturn:   fakeRefReturn,
		}

		objMgr := NewObjectManager(ncFakeConnector, cmpType, tenantID)
		ncFakeConnector.createObjectObj.(*NetworkContainer).Ea = objMgr.getBasicEA(true)
		ncFakeConnector.resultObject.(*NetworkContainer).Ea = objMgr.getBasicEA(true)

		var actualNetworkContainer *NetworkContainer
		var err error
		It("should pass expected NetworkContainer Object to CreateObject", func() {
			actualNetworkContainer, err = objMgr.CreateNetworkContainer(netviewName, cidr)
		})
		It("should return expected NetworkContainer Object", func() {
			Expect(actualNetworkContainer).To(Equal(ncFakeConnector.resultObject))
			Expect(err).To(BeNil())
		})
	})

	Describe("Create Network", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "default_view"
		cidr := "43.0.11.0/24"
		networkName := "private-net"
		fakeRefReturn := "network/ZG5zLm5ldHdvcmskODkuMC4wLjAvMjQvMjU:43.0.11.0/24/default_view"
		nwFakeConnector := &fakeConnector{
			createObjectObj: NewNetwork(Network{NetviewName: netviewName, Cidr: cidr}),
			resultObject:    NewNetwork(Network{NetviewName: netviewName, Cidr: cidr, Ref: fakeRefReturn}),
			fakeRefReturn:   fakeRefReturn,
		}

		objMgr := NewObjectManager(nwFakeConnector, cmpType, tenantID)

		nwFakeConnector.createObjectObj.(*Network).Ea = objMgr.getBasicEA(true)
		nwFakeConnector.createObjectObj.(*Network).Ea["Network Name"] = networkName

		nwFakeConnector.resultObject.(*Network).Ea = objMgr.getBasicEA(true)
		nwFakeConnector.resultObject.(*Network).Ea["Network Name"] = networkName

		var actualNetwork *Network
		var err error
		It("should pass expected Network Object to CreateObject", func() {
			actualNetwork, err = objMgr.CreateNetwork(netviewName, cidr, networkName)
		})
		It("should return expected Network Object", func() {
			Expect(actualNetwork).To(Equal(nwFakeConnector.resultObject))
			Expect(err).To(BeNil())
		})
	})

	Describe("Allocate Network", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "default_view"
		cidr := "142.0.22.0/24"
		prefixLen := uint(24)
		networkName := "private-net"
		fakeRefReturn := fmt.Sprintf("network/ZG5zLm5ldHdvcmskODkuMC4wLjAvMjQvMjU:%s/%s", cidr, netviewName)
		anFakeConnector := &fakeConnector{
			createObjectObj: NewNetwork(Network{
				NetviewName: netviewName,
				Cidr:        fmt.Sprintf("func:nextavailablenetwork:%s,%s,%d", cidr, netviewName, prefixLen),
			}),
			resultObject:  BuildNetworkFromRef(fakeRefReturn),
			fakeRefReturn: fakeRefReturn,
		}

		objMgr := NewObjectManager(anFakeConnector, cmpType, tenantID)

		anFakeConnector.createObjectObj.(*Network).Ea = objMgr.getBasicEA(true)
		anFakeConnector.createObjectObj.(*Network).Ea["Network Name"] = networkName

		var actualNetwork *Network
		var err error
		It("should pass expected Network Object to CreateObject", func() {
			actualNetwork, err = objMgr.AllocateNetwork(netviewName, cidr, prefixLen, networkName)
		})
		It("should return expected Network Object", func() {
			Expect(actualNetwork).To(Equal(anFakeConnector.resultObject))
			Expect(err).To(BeNil())
		})
	})

	Describe("Allocate Specific IP", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "private"
		cidr := "53.0.0.0/24"
		ipAddr := "53.0.0.21"
		macAddr := "01:23:45:67:80:ab"
		vmID := "93f9249abc039284"
		fakeRefReturn := fmt.Sprintf("fixedaddress/ZG5zLmJpbmRfY25h:%s/private", ipAddr)

		asiFakeConnector := &fakeConnector{
			createObjectObj: NewFixedAddress(FixedAddress{
				NetviewName: netviewName,
				Cidr:        cidr,
				IPAddress:   ipAddr,
				Mac:         macAddr,
			}),
			resultObject: NewFixedAddress(FixedAddress{
				NetviewName: netviewName,
				Cidr:        cidr,
				IPAddress:   GetIPAddressFromRef(fakeRefReturn),
				Mac:         macAddr,
				Ref:         fakeRefReturn,
			}),
			fakeRefReturn: fakeRefReturn,
		}

		objMgr := NewObjectManager(asiFakeConnector, cmpType, tenantID)

		asiFakeConnector.createObjectObj.(*FixedAddress).Ea = objMgr.getBasicEA(true)
		asiFakeConnector.createObjectObj.(*FixedAddress).Ea["VM ID"] = vmID

		asiFakeConnector.resultObject.(*FixedAddress).Ea = objMgr.getBasicEA(true)
		asiFakeConnector.resultObject.(*FixedAddress).Ea["VM ID"] = vmID

		var actualIP *FixedAddress
		var err error
		It("should pass expected Fixed Address Object to CreateObject", func() {
			actualIP, err = objMgr.AllocateIP(netviewName, cidr, ipAddr, macAddr, vmID)
		})
		It("should return expected Fixed Address Object", func() {
			Expect(actualIP).To(Equal(asiFakeConnector.resultObject))
			Expect(err).To(BeNil())
		})
	})

	Describe("Allocate Next Available IP", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "private"
		cidr := "53.0.0.0/24"
		ipAddr := fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netviewName)
		macAddr := "01:23:45:67:80:ab"
		vmID := "93f9249abc039284"
		resultIP := "53.0.0.32"
		fakeRefReturn := fmt.Sprintf("fixedaddress/ZG5zLmJpbmRfY25h:%s/private", resultIP)

		aniFakeConnector := &fakeConnector{
			createObjectObj: NewFixedAddress(FixedAddress{
				NetviewName: netviewName,
				Cidr:        cidr,
				IPAddress:   ipAddr,
				Mac:         macAddr,
			}),
			resultObject: NewFixedAddress(FixedAddress{
				NetviewName: netviewName,
				Cidr:        cidr,
				IPAddress:   resultIP,
				Mac:         macAddr,
				Ref:         fakeRefReturn,
			}),
			fakeRefReturn: fakeRefReturn,
		}

		objMgr := NewObjectManager(aniFakeConnector, cmpType, tenantID)

		aniFakeConnector.createObjectObj.(*FixedAddress).Ea = objMgr.getBasicEA(true)
		aniFakeConnector.createObjectObj.(*FixedAddress).Ea["VM ID"] = vmID

		aniFakeConnector.resultObject.(*FixedAddress).Ea = objMgr.getBasicEA(true)
		aniFakeConnector.resultObject.(*FixedAddress).Ea["VM ID"] = vmID

		var actualIP *FixedAddress
		var err error
		It("should pass expected Fixed Address Object to CreateObject", func() {
			actualIP, err = objMgr.AllocateIP(netviewName, cidr, "", macAddr, vmID)
		})
		It("should return expected Fixed Address Object", func() {
			Expect(actualIP).To(Equal(aniFakeConnector.resultObject))
			Expect(err).To(BeNil())
		})
	})

	Describe("Create EA Definition", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		comment := "Test Extensible Attribute"
		flags := "CGV"
		listValues := []EADefListValue{"True", "False"}
		name := "TestEA"
		eaType := "string"
		allowedTypes := []string{"arecord", "aaarecord", "ptrrecord"}
		ead := EADefinition{
			Name:               name,
			Comment:            comment,
			Flags:              flags,
			ListValues:         listValues,
			Type:               eaType,
			AllowedObjectTypes: allowedTypes}
		fakeRefReturn := "extensibleattributedef/ZG5zLm5ldHdvcmtfdmlldyQyMw:TestEA"
		eadFakeConnector := &fakeConnector{
			createObjectObj: NewEADefinition(ead),
			resultObject:    NewEADefinition(ead),
			fakeRefReturn:   fakeRefReturn,
		}
		eadFakeConnector.resultObject.(*EADefinition).Ref = fakeRefReturn

		objMgr := NewObjectManager(eadFakeConnector, cmpType, tenantID)

		var actualEADef *EADefinition
		var err error
		It("should pass expected EA Definintion Object to CreateObject", func() {
			actualEADef, err = objMgr.CreateEADefinition(ead)
		})
		It("should return expected EA Definition Object", func() {
			Expect(actualEADef).To(Equal(eadFakeConnector.resultObject))
			Expect(err).To(BeNil())
		})
	})

	Describe("Get Network View", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "Default View"
		fakeRefReturn := "networkview/ZG5zLm5ldHdvcmtfdmlldyQyMw:global_view/false"
		nvFakeConnector := &fakeConnector{
			getObjectObj: NewNetworkView(NetworkView{Name: netviewName}),
			getObjectRef: "",
			resultObject: []NetworkView{*NewNetworkView(NetworkView{Name: netviewName, Ref: fakeRefReturn})},
		}

		objMgr := NewObjectManager(nvFakeConnector, cmpType, tenantID)

		var actualNetworkView *NetworkView
		var err error
		It("should pass expected NetworkView Object to GetObject", func() {
			actualNetworkView, err = objMgr.GetNetworkView(netviewName)
		})
		It("should return expected NetworkView Object", func() {
			Expect(*actualNetworkView).To(Equal(nvFakeConnector.resultObject.([]NetworkView)[0]))
			Expect(err).To(BeNil())
		})
	})

	Describe("Get Network Container", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "Default View"
		cidr := "43.0.11.0/24"
		fakeRefReturn := "networkcontainer/ZG5zLm5ldHdvcmtfdmlldyQyMw:global_view/false"
		ncFakeConnector := &fakeConnector{
			getObjectObj: NewNetworkContainer(NetworkContainer{NetviewName: netviewName, Cidr: cidr}),
			getObjectRef: "",
			resultObject: []NetworkContainer{*NewNetworkContainer(NetworkContainer{NetviewName: netviewName, Cidr: cidr, Ref: fakeRefReturn})},
		}

		objMgr := NewObjectManager(ncFakeConnector, cmpType, tenantID)

		var actualNetworkContainer *NetworkContainer
		var err error
		It("should pass expected NetworkContainer Object to GetObject", func() {
			actualNetworkContainer, err = objMgr.GetNetworkContainer(netviewName, cidr)
		})
		It("should return expected NetworkContainer Object", func() {
			Expect(*actualNetworkContainer).To(Equal(ncFakeConnector.resultObject.([]NetworkContainer)[0]))
			Expect(err).To(BeNil())
		})
	})

	Describe("Get Network", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "default_view"
		cidr := "28.0.42.0/24"
		networkName := "private-net"
		ea := EA{"Network Name": networkName}
		fakeRefReturn := fmt.Sprintf("network/ZG5zLm5ldHdvcmskODkuMC4wLjAvMjQvMjU:%s/%s", cidr, netviewName)
		nwFakeConnector := &fakeConnector{
			getObjectObj: NewNetwork(Network{NetviewName: netviewName, Cidr: cidr}),
			getObjectRef: "",
			resultObject: []Network{*NewNetwork(Network{NetviewName: netviewName, Cidr: cidr, Ref: fakeRefReturn})},
		}

		nwFakeConnector.getObjectObj.(*Network).eaSearch = EASearch(ea)
		nwFakeConnector.resultObject.([]Network)[0].eaSearch = EASearch(ea)

		objMgr := NewObjectManager(nwFakeConnector, cmpType, tenantID)

		var actualNetwork *Network
		var err error
		It("should pass expected Network Object to GetObject", func() {
			actualNetwork, err = objMgr.GetNetwork(netviewName, cidr, ea)
		})
		It("should return expected Network Object", func() {
			Expect(*actualNetwork).To(Equal(nwFakeConnector.resultObject.([]Network)[0]))
			Expect(err).To(BeNil())
		})
	})

	Describe("Get Fixed Address", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "private"
		cidr := "53.0.0.0/24"
		ipAddr := "53.0.0.21"
		macAddr := "01:23:45:67:80:ab"
		fakeRefReturn := fmt.Sprintf("fixedaddress/ZG5zLmJpbmRfY25h:%s/private", ipAddr)

		fipFakeConnector := &fakeConnector{
			getObjectObj: NewFixedAddress(FixedAddress{
				NetviewName: netviewName,
				Cidr:        cidr,
				IPAddress:   ipAddr,
				Mac:         macAddr,
			}),
			getObjectRef: "",
			resultObject: []FixedAddress{*NewFixedAddress(FixedAddress{
				NetviewName: netviewName,
				Cidr:        cidr,
				IPAddress:   GetIPAddressFromRef(fakeRefReturn),
				Mac:         macAddr,
				Ref:         fakeRefReturn,
			})},
			fakeRefReturn: fakeRefReturn,
		}

		objMgr := NewObjectManager(fipFakeConnector, cmpType, tenantID)

		var actualIP *FixedAddress
		var err error
		It("should pass expected Fixed Address Object to GetObject", func() {
			actualIP, err = objMgr.GetFixedAddress(netviewName, cidr, ipAddr, macAddr)
		})
		It("should return expected Fixed Address Object", func() {
			Expect(*actualIP).To(Equal(fipFakeConnector.resultObject.([]FixedAddress)[0]))
			Expect(err).To(BeNil())
		})
	})

	Describe("Get EA Definition", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		comment := "Test Extensible Attribute"
		flags := "CGV"
		listValues := []EADefListValue{"True", "False"}
		name := "TestEA"
		eaType := "string"
		allowedTypes := []string{"arecord", "aaarecord", "ptrrecord"}
		ead := EADefinition{
			Name: name,
		}
		fakeRefReturn := "extensibleattributedef/ZG5zLm5ldHdvcmtfdmlldyQyMw:TestEA"
		eadRes := EADefinition{
			Name:               name,
			Comment:            comment,
			Flags:              flags,
			ListValues:         listValues,
			Type:               eaType,
			AllowedObjectTypes: allowedTypes,
			Ref:                fakeRefReturn,
		}

		eadFakeConnector := &fakeConnector{
			getObjectObj:  NewEADefinition(ead),
			getObjectRef:  "",
			resultObject:  []EADefinition{*NewEADefinition(eadRes)},
			fakeRefReturn: fakeRefReturn,
		}

		objMgr := NewObjectManager(eadFakeConnector, cmpType, tenantID)

		var actualEADef *EADefinition
		var err error
		It("should pass expected EA Definintion Object to GetObject", func() {
			actualEADef, err = objMgr.GetEADefinition(name)
		})
		It("should return expected EA Definition Object", func() {
			Expect(*actualEADef).To(Equal(eadFakeConnector.resultObject.([]EADefinition)[0]))
			Expect(err).To(BeNil())
		})
	})

	Describe("Delete Network", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "default_view"
		cidr := "28.0.42.0/24"
		deleteRef := fmt.Sprintf("network/ZG5zLm5ldHdvcmskODkuMC4wLjAvMjQvMjU:%s/%s", cidr, netviewName)
		fakeRefReturn := deleteRef
		nwFakeConnector := &fakeConnector{
			deleteObjectRef: deleteRef,
			fakeRefReturn:   fakeRefReturn,
		}

		objMgr := NewObjectManager(nwFakeConnector, cmpType, tenantID)

		var actualRef string
		var err error
		It("should pass expected Network Ref to DeleteObject", func() {
			actualRef, err = objMgr.DeleteNetwork(deleteRef, netviewName)
		})
		It("should return expected Network Ref", func() {
			Expect(actualRef).To(Equal(fakeRefReturn))
			Expect(err).To(BeNil())
		})
	})

	Describe("Delete Fixed Address", func() {
		cmpType := "Docker"
		tenantID := "01234567890abcdef01234567890abcdef"
		netviewName := "private"
		cidr := "83.0.101.0/24"
		ipAddr := "83.0.101.68"
		macAddr := "01:23:45:67:80:ab"
		fakeRefReturn := fmt.Sprintf("fixedaddress/ZG5zLmJpbmRfY25h:%s/private", ipAddr)

		fipFakeConnector := &fakeConnector{
			getObjectObj: NewFixedAddress(FixedAddress{
				NetviewName: netviewName,
				Cidr:        cidr,
				IPAddress:   ipAddr,
				Mac:         macAddr,
			}),
			getObjectRef: "",
			resultObject: []FixedAddress{*NewFixedAddress(FixedAddress{
				NetviewName: netviewName,
				Cidr:        cidr,
				IPAddress:   GetIPAddressFromRef(fakeRefReturn),
				Mac:         macAddr,
				Ref:         fakeRefReturn,
			})},
			deleteObjectRef: fakeRefReturn,
			fakeRefReturn:   fakeRefReturn,
		}

		objMgr := NewObjectManager(fipFakeConnector, cmpType, tenantID)

		var actualRef string
		var err error
		It("should pass expected Fixed Address Object to GetObject and DeleteObject", func() {
			actualRef, err = objMgr.ReleaseIP(netviewName, cidr, ipAddr, macAddr)
		})
		It("should return expected Fixed Address Ref", func() {
			Expect(actualRef).To(Equal(fakeRefReturn))
			Expect(err).To(BeNil())
		})
	})

	Describe("BuildNetworkViewFromRef", func() {
		netviewName := "default_view"
		netviewRef := fmt.Sprintf("networkview/ZG5zLm5ldHdvcmtfdmlldyQyMw:%s/false", netviewName)

		expectedNetworkView := NetworkView{Ref: netviewRef, Name: netviewName}
		It("should return expected Network View Object", func() {
			Expect(*BuildNetworkViewFromRef(netviewRef)).To(Equal(expectedNetworkView))
		})
		It("should failed if bad Network View Ref is provided", func() {
			Expect(BuildNetworkViewFromRef("bad")).To(BeNil())
		})
	})

	Describe("BuildNetworkFromRef", func() {
		netviewName := "test_view"
		cidr := "23.11.0.0/24"
		networkRef := fmt.Sprintf("network/ZG5zLm5ldHdvcmtfdmlldyQyMw:%s/%s", cidr, netviewName)

		expectedNetwork := Network{Ref: networkRef, NetviewName: netviewName, Cidr: cidr}
		It("should return expected Network Object", func() {
			Expect(*BuildNetworkFromRef(networkRef)).To(Equal(expectedNetwork))
		})
		It("should failed if bad Network Ref is provided", func() {
			Expect(BuildNetworkFromRef("network/ZG5zLm5ldHdvcmtfdmlldyQyMw")).To(BeNil())
		})
	})

})
