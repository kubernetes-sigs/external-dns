package dnsv2

import (
	"encoding/hex"
	"fmt"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"net"
	//"sort"
	"strconv"
	"strings"
)

/*
{
  "metadata": {
    "zone": "example.com",
    "types": [
      "A"
    ],
    "page": 1,
    "pageSize": 25,
    "totalElements": 2
  },
  "recordsets": [
    {
      "name": "www.example.com",
      "type": "A",
      "ttl": 300,
      "rdata": [
        "10.0.0.2",
        "10.0.0.3"
      ]
    },
    {
      "name": "mail.example.com",
      "type": "A",
      "ttl": 300,
      "rdata": [
        "192.168.0.1",
        "192.168.0.2"
      ]
    }
  ]
}

*/

func FullIPv6(ip net.IP) string {
	dst := make([]byte, hex.EncodedLen(len(ip)))
	_ = hex.Encode(dst, ip)
	return string(dst[0:4]) + ":" +
		string(dst[4:8]) + ":" +
		string(dst[8:12]) + ":" +
		string(dst[12:16]) + ":" +
		string(dst[16:20]) + ":" +
		string(dst[20:24]) + ":" +
		string(dst[24:28]) + ":" +
		string(dst[28:])
}

func padvalue(str string) string {
	v_str := strings.Replace(str, "m", "", -1)
	v_float, err := strconv.ParseFloat(v_str, 32)
	if err != nil {
		return "FAIL"
	}
	v_result := fmt.Sprintf("%.2f", v_float)

	return v_result
}

// Used to pad coordinates to x.xxm format
func PadCoordinates(str string) string {

	s := strings.Split(str, " ")
	lat_d, lat_m, lat_s, lat_dir, long_d, long_m, long_s, long_dir, altitude, size, horiz_precision, vert_precision := s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]
	return lat_d + " " + lat_m + " " + lat_s + " " + lat_dir + " " + long_d + " " + long_m + " " + long_s + " " + long_dir + " " + padvalue(altitude) + "m " + padvalue(size) + "m " + padvalue(horiz_precision) + "m " + padvalue(vert_precision) + "m"

}

// Get single Recordset. Following convention for other single record CRUD operations, return a RecordBody.
func GetRecord(zone string, name string, record_type string) (*RecordBody, error) {

	record := &RecordBody{}

	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf("/config-dns/v2/zones/%s/names/%s/types/%s", zone, name, record_type),
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
		return nil, &RecordError{fieldName: name}
	} else {
		err = client.BodyJSON(res, record)
		if err != nil {
			return nil, err
		}
		return record, nil
	}
}

func GetRecordList(zone string, name string, record_type string) (*RecordSetResponse, error) {

	records := NewRecordSetResponse(name)

	req, err := client.NewRequest(
		Config,
		"GET",
		"/config-dns/v2/zones/"+zone+"/recordsets?types="+record_type+"&showAll=true",
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
		return nil, &RecordError{fieldName: name}
	} else {
		err = client.BodyJSON(res, records)
		if err != nil {
			return nil, err
		}
		return records, nil
	}
}

func GetRdata(zone string, name string, record_type string) ([]string, error) {
	records, err := GetRecordList(zone, name, record_type)
	if err != nil {
		return nil, err
	}

	var arrLength int
	for _, c := range records.Recordsets {
		if c.Name == name {
			arrLength = len(c.Rdata)
		}
	}

	rdata := make([]string, 0, arrLength)

	for _, r := range records.Recordsets {
		if r.Name == name {
			for _, i := range r.Rdata {
				str := i

				if record_type == "AAAA" {
					addr := net.ParseIP(str)
					result := FullIPv6(addr)
					str = result
				} else if record_type == "LOC" {
					str = PadCoordinates(str)
				}
				rdata = append(rdata, str)
			}
		}
	}
	return rdata, nil
}

func ProcessRdata(rdata []string, rtype string) []string {

	newrdata := make([]string, 0, len(rdata))
	for _, i := range rdata {
		str := i
		if rtype == "AAAA" {
			addr := net.ParseIP(str)
			result := FullIPv6(addr)
			str = result
		} else if rtype == "LOC" {
			str = PadCoordinates(str)
		}
		newrdata = append(newrdata, str)
	}
	return newrdata

}

// Utility method to parse RData in context of type. Return map of fields and values
func ParseRData(rtype string, rdata []string) map[string]interface{} {

	fieldMap := make(map[string]interface{}, 0)
	if len(rdata) == 0 {
		return fieldMap
	}
	newrdata := make([]string, 0, len(rdata))
	fieldMap["target"] = newrdata

	switch rtype {
	case "AFSDB":
		parts := strings.Split(rdata[0], " ")
		fieldMap["subtype"], _ = strconv.Atoi(parts[0])
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			newrdata = append(newrdata, parts[1])
		}
		fieldMap["target"] = newrdata

	case "DNSKEY":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["flags"], _ = strconv.Atoi(parts[0])
			fieldMap["protocol"], _ = strconv.Atoi(parts[1])
			fieldMap["algorithm"], _ = strconv.Atoi(parts[2])
			key := parts[3]
			// key can have whitespace
			if len(parts) > 4 {
				i := 4
				for i < len(parts) {
					key += " " + parts[i]
				}
			}
			fieldMap["key"] = key
			break
		}

	case "DS":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["keytag"], _ = strconv.Atoi(parts[0])
			fieldMap["digest_type"], _ = strconv.Atoi(parts[2])
			fieldMap["algorithm"], _ = strconv.Atoi(parts[1])
			dig := parts[3]
			// digest can have whitespace
			if len(parts) > 4 {
				i := 4
				for i < len(parts) {
					dig += " " + parts[i]
				}
			}
			fieldMap["digest"] = dig
			break
		}

	case "HINFO":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["hardware"] = parts[0]
			fieldMap["software"] = parts[1]
			break
		}
	/*
		// too many variations to calculate pri and increment
		case "MX":
			sort.Strings(rdata)
			parts := strings.Split(rdata[0], " ")
			fieldMap["priority"], _ = strconv.Atoi(parts[0])
			if len(rdata) > 1 {
				parts = strings.Split(rdata[1], " ")
				tpri, _ := strconv.Atoi(parts[0])
				fieldMap["priority_increment"] = tpri - fieldMap["priority"].(int)
			}
			for _, rcontent := range rdata {
				parts := strings.Split(rcontent, " ")
				newrdata = append(newrdata, parts[1])
			}
			fieldMap["target"] = newrdata
	*/

	case "NAPTR":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["order"], _ = strconv.Atoi(parts[0])
			fieldMap["preference"], _ = strconv.Atoi(parts[1])
			fieldMap["flagsnaptr"] = parts[2]
			fieldMap["service"] = parts[3]
			fieldMap["regexp"] = parts[4]
			fieldMap["replacement"] = parts[5]
			break
		}

	case "NSEC3":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["flags"], _ = strconv.Atoi(parts[1])
			fieldMap["algorithm"], _ = strconv.Atoi(parts[0])
			fieldMap["iterations"], _ = strconv.Atoi(parts[2])
			fieldMap["salt"] = parts[3]
			fieldMap["next_hashed_owner_name"] = parts[4]
			fieldMap["type_bitmaps"] = parts[5]
			break
		}

	case "NSEC3PARAM":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["flags"], _ = strconv.Atoi(parts[1])
			fieldMap["algorithm"], _ = strconv.Atoi(parts[0])
			fieldMap["iterations"], _ = strconv.Atoi(parts[2])
			fieldMap["salt"] = parts[3]
			break
		}

	case "RP":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["mailbox"] = parts[0]
			fieldMap["txt"] = parts[1]
			break
		}

	case "RRSIG":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["type_covered"] = parts[0]
			fieldMap["algorithm"], _ = strconv.Atoi(parts[1])
			fieldMap["labels"], _ = strconv.Atoi(parts[2])
			fieldMap["original_ttl"], _ = strconv.Atoi(parts[3])
			fieldMap["expiration"] = parts[4]
			fieldMap["inception"] = parts[5]
			fieldMap["signer"] = parts[7]
			fieldMap["keytag"], _ = strconv.Atoi(parts[6])
			sig := parts[8]
			// sig can have whitespace
			if len(parts) > 9 {
				i := 9
				for i < len(parts) {
					sig += " " + parts[i]
				}
			}
			fieldMap["signature"] = sig
			break
		}

	case "SRV":
		// pull out some fields
		parts := strings.Split(rdata[0], " ")
		fieldMap["priority"], _ = strconv.Atoi(parts[0])
		fieldMap["weight"], _ = strconv.Atoi(parts[1])
		fieldMap["port"], _ = strconv.Atoi(parts[2])
		// populate target
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			newrdata = append(newrdata, parts[3])
		}
		fieldMap["target"] = newrdata

	case "SSHFP":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["algorithm"], _ = strconv.Atoi(parts[0])
			fieldMap["fingerprint_type"], _ = strconv.Atoi(parts[1])
			fieldMap["fingerprint"] = parts[2]
			break
		}

	case "SOA":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["name_server"] = parts[0]
			fieldMap["email_address"] = parts[1]
			fieldMap["serial"], _ = strconv.Atoi(parts[2])
			fieldMap["refresh"], _ = strconv.Atoi(parts[3])
			fieldMap["retry"], _ = strconv.Atoi(parts[4])
			fieldMap["expiry"], _ = strconv.Atoi(parts[5])
			fieldMap["nxdomain_ttl"], _ = strconv.Atoi(parts[6])
			break
		}

	case "AKAMAITLC":
		parts := strings.Split(rdata[0], " ")
		fieldMap["answer_type"] = parts[0]
		fieldMap["dns_name"] = parts[1]

	case "SPF":
		for _, rcontent := range rdata {
			newrdata = append(newrdata, rcontent)
		}
		fieldMap["target"] = newrdata

	case "TXT":
		for _, rcontent := range rdata {
			newrdata = append(newrdata, rcontent)
		}
		fieldMap["target"] = newrdata

	case "AAAA":
		for _, i := range rdata {
			str := i
			addr := net.ParseIP(str)
			result := FullIPv6(addr)
			str = result
			newrdata = append(newrdata, str)
		}
		fieldMap["target"] = newrdata

	case "LOC":
		for _, i := range rdata {
			str := i
			str = PadCoordinates(str)
			newrdata = append(newrdata, str)
		}
		fieldMap["target"] = newrdata

	case "CERT":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			val, err := strconv.Atoi(parts[0])
			if err == nil {
				fieldMap["type_value"] = val
			} else {
				fieldMap["type_mnemonic"] = parts[0]
			}
			fieldMap["keytag"], _ = strconv.Atoi(parts[1])
			fieldMap["algorithm"], _ = strconv.Atoi(parts[2])
			fieldMap["certificate"] = parts[3]
			break
		}

	case "TLSA":
		for _, rcontent := range rdata {
			parts := strings.Split(rcontent, " ")
			fieldMap["usage"], _ = strconv.Atoi(parts[0])
			fieldMap["selector"], _ = strconv.Atoi(parts[1])
			fieldMap["match_type"], _ = strconv.Atoi(parts[2])
			fieldMap["certificate"] = parts[3]
			break
		}

	case "SVCB":
		for _, rcontent := range rdata {
			parts := strings.SplitN(rcontent, " ", 3)
			// has to be at least two fields.
			if len(parts) < 2 {
				break
			}
			fieldMap["svc_priority"], _ = strconv.Atoi(parts[0])
			fieldMap["target_name"] = parts[1]
			if len(parts) > 2 {
				fieldMap["svc_params"] = parts[2]
			}
			break
		}

	case "HTTPS":
		for _, rcontent := range rdata {
			parts := strings.SplitN(rcontent, " ", 3)
			// has to be at least two fields.
			if len(parts) < 2 {
				break
			}
			fieldMap["svc_priority"], _ = strconv.Atoi(parts[0])
			fieldMap["target_name"] = parts[1]
			if len(parts) > 2 {
				fieldMap["svc_params"] = parts[2]
			}
			break
		}

	default:
		for _, rcontent := range rdata {
			newrdata = append(newrdata, rcontent)
		}
		fieldMap["target"] = newrdata
	}

	return fieldMap

}
