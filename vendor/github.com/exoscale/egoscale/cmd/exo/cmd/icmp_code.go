package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/camelcase"
)

type icmpCode uint16

// nolint
const (
	//destinationUnreachable
	netUnreachable                                                 icmpCode = 0x0300
	hostUnreachable                                                icmpCode = 0x0301
	protocolUnreachable                                            icmpCode = 0x0302
	portUnreachable                                                icmpCode = 0x0303
	fragmentationNeededAndDoNotFragmentWasSet                      icmpCode = 0x0304
	sourceRouteFailed                                              icmpCode = 0x0305
	destinationNetworkUnknown                                      icmpCode = 0x0306
	destinationHostUnknown                                         icmpCode = 0x0307
	sourceHostIsolated                                             icmpCode = 0x0308
	communicationWithDestinationNetworkIsAdminstrativelyProhibited icmpCode = 0x0309
	communicationWithDestinationHostIsAdminstrativelyProhibited    icmpCode = 0x030A
	destinationNetworkUnreachableForTypeOfService                  icmpCode = 0x030B
	destinationHostUnreachableForTypeOfService                     icmpCode = 0x030C
	communicationAdministrativelyProhibited                        icmpCode = 0x030D
	hostPrecedenceViolation                                        icmpCode = 0x030E
	precedenceCutoffInEffect                                       icmpCode = 0x030F
	//redirect
	redirectDatagramForTheNetwork                 icmpCode = 0x0500
	redirectDatagramForTheHost                    icmpCode = 0x0501
	redirectDatagramForTheTypeOfServiceAndNetwork icmpCode = 0x0502
	redirectSDatagramForTheTypeOfSerivceAndHost   icmpCode = 0x0503
	//alternateHostAddress
	alternateAddressForHost icmpCode = 0x0600
	//timeExceeded
	timeToLiveExceededInTransit    icmpCode = 0x0B00
	fragmentReassemblyTimeExceeded icmpCode = 0x0B01
	//parameterProblem
	pointerIndicatesTheError icmpCode = 0x0C00
	missingARequiredOption   icmpCode = 0x0C01
	badLength                icmpCode = 0x0C02
)

func (i icmpCode) StringFormatted() string {
	if i.String() == fmt.Sprintf("icmpCode(%d)", i) {
		return ""
	}
	splitted := camelcase.Split(i.String())
	res := ""
	for i, v := range splitted {
		if i == 0 {
			res += v
			continue
		}
		res += " " + v
	}
	return strings.Title(res)
}

func (i icmpCode) icmpType() icmpType {
	return icmpType(i >> 8)
}
