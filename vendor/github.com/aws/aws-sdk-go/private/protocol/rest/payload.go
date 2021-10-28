package rest

import "reflect"

// PayloadMember returns the payload field member of i if there is one, or nil.
func PayloadMember(i interface{}) interface{} {
	if i == nil {
		return nil
	}

	v := reflect.ValueOf(i).Elem()
	if !v.IsValid() {
		return nil
	}
	if field, ok := v.Type().FieldByName("_"); ok {
		if payloadName := field.Tag.Get("payload"); payloadName != "" {
			field, _ := v.Type().FieldByName(payloadName)
			if field.Tag.Get("type") != "structure" {
				return nil
			}

			payload := v.FieldByName(payloadName)
			if payload.IsValid() || (payload.Kind() == reflect.Ptr && !payload.IsNil()) {
				return payload.Interface()
			}
		}
	}
	return nil
}

<<<<<<< HEAD
<<<<<<< HEAD
const nopayloadPayloadType = "nopayload"

// PayloadType returns the type of a payload field member of i if there is one,
// or "".
func PayloadType(i interface{}) string {
	v := reflect.Indirect(reflect.ValueOf(i))
	if !v.IsValid() {
		return ""
	}

	if field, ok := v.Type().FieldByName("_"); ok {
		if noPayload := field.Tag.Get(nopayloadPayloadType); noPayload != "" {
			return nopayloadPayloadType
		}

		if payloadName := field.Tag.Get("payload"); payloadName != "" {
			if member, ok := v.Type().FieldByName(payloadName); ok {
				return member.Tag.Get("type")
			}
		}
	}

||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// PayloadType returns the type of a payload field member of i if there is one, or "".
||||||| parent of 4d7e5ad26 (update vendored files)
// PayloadType returns the type of a payload field member of i if there is one, or "".
=======
const nopayloadPayloadType = "nopayload"

// PayloadType returns the type of a payload field member of i if there is one,
// or "".
>>>>>>> 4d7e5ad26 (update vendored files)
func PayloadType(i interface{}) string {
	v := reflect.Indirect(reflect.ValueOf(i))
	if !v.IsValid() {
		return ""
	}

	if field, ok := v.Type().FieldByName("_"); ok {
		if noPayload := field.Tag.Get(nopayloadPayloadType); noPayload != "" {
			return nopayloadPayloadType
		}

		if payloadName := field.Tag.Get("payload"); payloadName != "" {
			if member, ok := v.Type().FieldByName(payloadName); ok {
				return member.Tag.Get("type")
			}
		}
	}
<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======

>>>>>>> 4d7e5ad26 (update vendored files)
	return ""
}
