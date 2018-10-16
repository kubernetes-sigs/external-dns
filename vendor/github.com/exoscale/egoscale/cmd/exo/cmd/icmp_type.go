package cmd

import (
	"strings"

	"github.com/fatih/camelcase"
)

//go:generate stringer -type=icmpType

type icmpType uint8

// nolint
const (
	echoReply                 icmpType = 0
	destinationUnreachable    icmpType = 3
	sourceQuench              icmpType = 4
	redirect                  icmpType = 5
	alternateHostAddress      icmpType = 6
	echo                      icmpType = 8
	routerAdvertisement       icmpType = 9
	routerSelection           icmpType = 10
	timeExceeded              icmpType = 11
	parameterProblem          icmpType = 12
	timestamp                 icmpType = 13
	timestampReply            icmpType = 14
	informationRequest        icmpType = 15
	informationReply          icmpType = 16
	addressMaskRequest        icmpType = 17
	addressMaskReply          icmpType = 18
	traceroute                icmpType = 30
	datagramConversionError   icmpType = 31
	mobileHostRedirect        icmpType = 32
	ipv6WhereAreYou           icmpType = 33
	ipv6IAmHere               icmpType = 34
	mobileRegistrationRequest icmpType = 35
	mobileRegistrationReply   icmpType = 36
	domainNameRequest         icmpType = 37
	domainNameReply           icmpType = 38
	skip                      icmpType = 39
	photuris                  icmpType = 40
	echoRequest               icmpType = 128
)

func (i icmpType) StringFormatted() string {
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
