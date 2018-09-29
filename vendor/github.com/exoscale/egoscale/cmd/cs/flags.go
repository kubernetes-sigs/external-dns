package main

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"

	"github.com/exoscale/egoscale"
)

type uint8Generic struct {
	value *uint8
}

func (g *uint8Generic) Set(value string) error {
	v, err := strconv.ParseUint(value, 10, 8)
	if err == nil {
		*g.value = uint8(v)
	}

	return err
}

func (g *uint8Generic) String() string {
	if g.value == nil {
		return ""
	}
	return strconv.FormatUint(uint64(*g.value), 10)
}

type uint16Generic struct {
	value *uint16
}

func (g *uint16Generic) Set(value string) error {
	v, err := strconv.ParseUint(value, 10, 16)
	if err == nil {
		*g.value = uint16(v)
	}

	return err
}

func (g *uint16Generic) String() string {
	if g.value == nil {
		return ""
	}
	return strconv.FormatUint(uint64(*g.value), 10)
}

type int16Generic struct {
	value *int16
}

func (g *int16Generic) Set(value string) error {
	v, err := strconv.ParseInt(value, 10, 16)
	if err == nil {
		*g.value = int16(v)
	}

	return err
}

func (g *int16Generic) String() string {
	return strconv.FormatInt(int64(*g.value), 10)
}

type boolPtrGeneric struct {
	value **bool
}

func (g *boolPtrGeneric) Set(value string) error {
	v, err := strconv.ParseBool(value)
	if err == nil {
		*(g.value) = &v
	}

	return err
}

func (g *boolPtrGeneric) String() string {
	if g.value == nil || *(g.value) == nil {
		return "unset"
	}
	return strconv.FormatBool(**(g.value))
}

type ipGeneric struct {
	value *net.IP
}

func (g *ipGeneric) Set(value string) error {
	*(g.value) = net.ParseIP(value)
	if *(g.value) == nil {
		return fmt.Errorf("not a valid IP address, got %s", value)
	}
	return nil
}

func (g *ipGeneric) String() string {
	if g.value == nil || *(g.value) == nil {
		return ""
	}

	return (*(g.value)).String()
}

type cidrGeneric struct {
	value **egoscale.CIDR
}

func (g *cidrGeneric) Set(value string) error {
	cidr, err := egoscale.ParseCIDR(value)
	if err != nil {
		return fmt.Errorf("not a valid CIDR, got %s", value)
	}
	*(g.value) = cidr
	return nil
}

func (g *cidrGeneric) String() string {
	if g.value != nil && *(g.value) != nil {
		return (*(g.value)).String()
	}
	return ""
}

type cidrListGeneric struct {
	value *[]egoscale.CIDR
}

func (g *cidrListGeneric) Set(value string) error {
	m := g.value
	if *m == nil {
		n := make([]egoscale.CIDR, 0)
		*m = n
	}

	values := strings.Split(value, ",")

	for _, value := range values {
		cidr, err := egoscale.ParseCIDR(value)
		if err != nil {
			return err
		}
		*m = append(*m, *cidr)
	}
	return nil
}

func (g *cidrListGeneric) String() string {
	m := g.value
	if m == nil || *m == nil {
		return ""
	}
	vs := make([]string, 0, len(*m))
	for _, v := range *m {
		vs = append(vs, v.String())
	}

	return strings.Join(vs, ",")
}

type uuidGeneric struct {
	value **egoscale.UUID
}

func (g *uuidGeneric) Set(value string) error {
	uuid, err := egoscale.ParseUUID(value)
	if err != nil {
		return fmt.Errorf("not a valid UUID, got %s", value)
	}
	*(g.value) = uuid
	return nil
}

func (g *uuidGeneric) String() string {
	if g.value != nil && *(g.value) != nil {
		return (*(g.value)).String()
	}
	return ""
}

type intTypeGeneric struct {
	addr    interface{}
	value   int64
	base    int
	bitSize int
	typ     reflect.Type
}

func (g *intTypeGeneric) Set(value string) error {
	v, err := strconv.ParseInt(value, g.base, g.bitSize)
	if err != nil {
		return err
	}

	tv := reflect.ValueOf(g.addr)
	var fv reflect.Value
	switch g.bitSize {
	case 8:
		val := (int8)(v)
		fv = reflect.ValueOf(&val)
	case 16:
		val := (int16)(v)
		fv = reflect.ValueOf(&val)
	case 32:
		val := (int)(v)
		fv = reflect.ValueOf(&val)
	case 64:
		fv = reflect.ValueOf(&v)
	}
	tv.Elem().Set(fv.Convert(reflect.PtrTo(g.typ)).Elem())

	g.value = v

	return nil
}

func (g *intTypeGeneric) String() string {
	if g.addr != nil {
		return strconv.FormatInt(g.value, g.base)
	}
	return ""
}

type stringerTypeGeneric struct {
	addr  interface{}
	value string
	typ   reflect.Type
}

func (g *stringerTypeGeneric) Set(value string) error {
	tv := reflect.ValueOf(g.addr)
	fv := reflect.ValueOf(&value)

	tv.Elem().Set(fv.Convert(reflect.PtrTo(g.typ)).Elem())

	g.value = value

	return nil
}

func (g *stringerTypeGeneric) String() string {
	return g.value
}

type mapGeneric struct {
	value *map[string]string
}

func (g *mapGeneric) Set(value string) error {
	m := g.value
	if *m == nil {
		n := make(map[string]string)
		*m = n
	}

	values := strings.SplitN(value, "=", 2)
	if len(values) != 2 {
		return fmt.Errorf("not a valid key=value content, got %s", value)
	}

	(*m)[values[0]] = values[1]
	return nil
}

func (g *mapGeneric) String() string {
	m := g.value
	if *m == nil {
		return ""
	}
	vs := make([]string, 0, len(*m))
	for k, v := range *m {
		vs = append(vs, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(vs, ",")
}

type tagGeneric struct {
	value *[]egoscale.ResourceTag
}

func (g *tagGeneric) Set(value string) error {
	m := g.value
	if *m == nil {
		n := make([]egoscale.ResourceTag, 0)
		*m = n
	}

	values := strings.SplitN(value, "=", 2)
	if len(values) != 2 {
		return fmt.Errorf("not a valid key=value content, got %s", value)
	}

	*m = append(*m, egoscale.ResourceTag{
		Key:   values[0],
		Value: values[1],
	})
	return nil
}

func (g *tagGeneric) String() string {
	m := g.value
	if m == nil || *m == nil {
		return ""
	}
	vs := make([]string, 0, len(*m))
	for _, v := range *m {
		vs = append(vs, fmt.Sprintf("%s=%s", v.Key, v.Value))
	}

	return strings.Join(vs, ",")
}
