package awsutil

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// StringValue returns the string representation of a value.
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//
// Deprecated: Use Prettify instead.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//
// Deprecated: Use Prettify instead.
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
func StringValue(i interface{}) string {
	var buf bytes.Buffer
	stringValue(reflect.ValueOf(i), 0, &buf)
	return buf.String()
}

func stringValue(v reflect.Value, indent int, buf *bytes.Buffer) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		buf.WriteString("{\n")

		for i := 0; i < v.Type().NumField(); i++ {
			ft := v.Type().Field(i)
			fv := v.Field(i)

			if ft.Name[0:1] == strings.ToLower(ft.Name[0:1]) {
				continue // ignore unexported fields
			}
			if (fv.Kind() == reflect.Ptr || fv.Kind() == reflect.Slice) && fv.IsNil() {
				continue // ignore unset fields
			}

			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(ft.Name + ": ")

			if tag := ft.Tag.Get("sensitive"); tag == "true" {
				buf.WriteString("<sensitive>")
			} else {
				stringValue(fv, indent+2, buf)
			}

			buf.WriteString(",\n")
		}

		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")
	case reflect.Slice:
		nl, id, id2 := "", "", ""
		if v.Len() > 3 {
			nl, id, id2 = "\n", strings.Repeat(" ", indent), strings.Repeat(" ", indent+2)
		}
		buf.WriteString("[" + nl)
		for i := 0; i < v.Len(); i++ {
			buf.WriteString(id2)
			stringValue(v.Index(i), indent+2, buf)

			if i < v.Len()-1 {
				buf.WriteString("," + nl)
			}
		}

		buf.WriteString(nl + id + "]")
	case reflect.Map:
		buf.WriteString("{\n")

		for i, k := range v.MapKeys() {
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(k.String() + ": ")
			stringValue(v.MapIndex(k), indent+2, buf)

			if i < v.Len()-1 {
				buf.WriteString(",\n")
			}
		}

		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")
	default:
		format := "%v"
		switch v.Interface().(type) {
		case string:
			format = "%q"
		}
		fmt.Fprintf(buf, format, v.Interface())
	}
}
