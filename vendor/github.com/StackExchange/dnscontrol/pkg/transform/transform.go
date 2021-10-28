package transform

import (
	"net"
	"strings"

	"github.com/pkg/errors"
)

// IpConversion describes an IP conversion.
type IpConversion struct {
	Low, High net.IP
	NewBases  []net.IP
	NewIPs    []net.IP
}

func ipToUint(i net.IP) (uint32, error) {
	parts := i.To4()
	if parts == nil || len(parts) != 4 {
		return 0, errors.Errorf("%s is not an ipv4 address", parts.String())
	}
	r := uint32(parts[0])<<24 | uint32(parts[1])<<16 | uint32(parts[2])<<8 | uint32(parts[3])
	return r, nil
}

// UintToIP convert a 32-bit into into a net.IP.
func UintToIP(u uint32) net.IP {
	return net.IPv4(
		byte((u>>24)&255),
		byte((u>>16)&255),
		byte((u>>8)&255),
		byte((u)&255))
}

// DecodeTransformTable turns a string-encoded table into a list of conversions.
func DecodeTransformTable(transforms string) ([]IpConversion, error) {
	result := []IpConversion{}
	rows := strings.Split(transforms, ";")
	for ri, row := range rows {
		items := strings.Split(row, "~")
		if len(items) != 4 {
			return nil, errors.Errorf("transform_table rows should have 4 elements. (%v) found in row (%v) of %#v", len(items), ri, transforms)
		}
		for i, item := range items {
			items[i] = strings.TrimSpace(item)
		}

		con := IpConversion{
			Low:  net.ParseIP(items[0]),
			High: net.ParseIP(items[1]),
		}
		parseList := func(s string) ([]net.IP, error) {
			ips := []net.IP{}
			for _, ip := range strings.Split(s, ",") {

				if ip == "" {
					continue
				}
				addr := net.ParseIP(ip)
				if addr == nil {
					return nil, errors.Errorf("%s is not a valid ip address", ip)
				}
				ips = append(ips, addr)
			}
			return ips, nil
		}
		var err error
		if con.NewBases, err = parseList(items[2]); err != nil {
			return nil, err
		}
		if con.NewIPs, err = parseList(items[3]); err != nil {
			return nil, err
		}

		low, _ := ipToUint(con.Low)
		high, _ := ipToUint(con.High)
		if low > high {
			return nil, errors.Errorf("transform_table Low should be less than High. row (%v) %v>%v (%v)", ri, con.Low, con.High, transforms)
		}
		if len(con.NewBases) > 0 && len(con.NewIPs) > 0 {
			return nil, errors.Errorf("transform_table_rows should only specify one of NewBases or NewIPs, Not both")
		}
		result = append(result, con)
	}

	return result, nil
}

// TransformIP transforms a single ip address. If the transform results in multiple new targets, an error will be returned.
func TransformIP(address net.IP, transforms []IpConversion) (net.IP, error) {
	ips, err := TransformIPToList(address, transforms)
	if err != nil {
		return nil, err
	}
	if len(ips) != 1 {
		return nil, errors.Errorf("Expect exactly one ip for TransformIP result. Got: %s", ips)
	}
	return ips[0], err
}

// TransformIPToList manipulates an net.IP based on a list of IpConversions. It can potentially expand one ip address into multiple addresses.
func TransformIPToList(address net.IP, transforms []IpConversion) ([]net.IP, error) {
	thisIP, err := ipToUint(address)
	if err != nil {
		return nil, err
	}
	for _, conv := range transforms {
		min, err := ipToUint(conv.Low)
		if err != nil {
			return nil, err
		}
		max, err := ipToUint(conv.High)
		if err != nil {
			return nil, err
		}
		if (thisIP >= min) && (thisIP <= max) {
			if len(conv.NewIPs) > 0 {
				return conv.NewIPs, nil
			}
			list := []net.IP{}
			for _, nb := range conv.NewBases {
				newbase, err := ipToUint(nb)
				if err != nil {
					return nil, err
				}
				list = append(list, UintToIP(newbase+(thisIP-min)))
			}
			return list, nil
		}
	}
	return []net.IP{address}, nil
}
