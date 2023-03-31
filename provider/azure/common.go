package azure

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
	"github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
	"github.com/Azure/go-autorest/autorest/to"
)

// Helper function (shared with test code)
func parseMxTarget[T dns.MxRecord | privatedns.MxRecord](mxTarget string) (T, error) {
	targetParts := strings.SplitN(mxTarget, " ", 2)

	if len(targetParts) != 2 {
		return T{}, fmt.Errorf("mx target needs to be of form '10 example.com'")
	}

	preferenceRaw, exchange := targetParts[0], targetParts[1]
	preference, err := strconv.ParseInt(preferenceRaw, 10, 32)

	if err != nil {
		return T{}, fmt.Errorf("invalid preference specified")
	}
	return T{
		Preference: to.Int32Ptr(int32(preference)),
		Exchange:   to.StringPtr(exchange),
	}, nil
}
