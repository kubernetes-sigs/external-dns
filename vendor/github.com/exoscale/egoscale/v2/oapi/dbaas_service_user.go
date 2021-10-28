package oapi

import (
	"encoding/json"
	"time"
)

// UnmarshalJSON unmarshals a DbaasServiceUser structure into a temporary structure whose
// AccessCertNotValidAfterTime field of type string to be able to parse the original timestamp
// (ISO 8601) into a time.Time object, since json.Unmarshal() only supports RFC 3339 format.
func (u *DbaasServiceUser) UnmarshalJSON(data []byte) error {
	raw := struct {
		AccessCert                  *string                         `json:"access-cert,omitempty"`
		AccessCertNotValidAfterTime *string                         `json:"access-cert-not-valid-after-time,omitempty"`
		AccessControl               *DbaasServiceUserAccessControl  `json:"access-control,omitempty"`
		AccessKey                   *string                         `json:"access-key,omitempty"`
		Authentication              *DbaasServiceUserAuthentication `json:"authentication,omitempty"`
		Password                    *string                         `json:"password,omitempty"`
		Type                        string                          `json:"type"`
		Username                    string                          `json:"username"`
	}{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.AccessCertNotValidAfterTime != nil {
		accessCertNotValidAfterTime, err := time.Parse(iso8601Format, *raw.AccessCertNotValidAfterTime)
		if err != nil {
			return err
		}
		u.AccessCertNotValidAfterTime = &accessCertNotValidAfterTime
	}

	u.AccessCert = raw.AccessCert
	u.AccessControl = raw.AccessControl
	u.AccessKey = raw.AccessKey
	u.Authentication = raw.Authentication
	u.Password = raw.Password
	u.Type = raw.Type
	u.Username = raw.Username

	return nil
}

// MarshalJSON returns the JSON encoding of a DbaasServiceUser structure after having formatted the
// AccessCertNotValidAfterTime field in the original timestamp (ISO 8601), since time.MarshalJSON()
// only supports RFC 3339 format.
func (u *DbaasServiceUser) MarshalJSON() ([]byte, error) {
	raw := struct {
		AccessCert                  *string                         `json:"access-cert,omitempty"`
		AccessCertNotValidAfterTime *string                         `json:"access-cert-not-valid-after-time,omitempty"`
		AccessControl               *DbaasServiceUserAccessControl  `json:"access-control,omitempty"`
		AccessKey                   *string                         `json:"access-key,omitempty"`
		Authentication              *DbaasServiceUserAuthentication `json:"authentication,omitempty"`
		Password                    *string                         `json:"password,omitempty"`
		Type                        string                          `json:"type"`
		Username                    string                          `json:"username"`
	}{}

	if u.AccessCertNotValidAfterTime != nil {
		accessCertNotValidAfterTime := u.AccessCertNotValidAfterTime.Format(iso8601Format)
		raw.AccessCertNotValidAfterTime = &accessCertNotValidAfterTime
	}

	raw.AccessCert = u.AccessCert
	raw.AccessControl = u.AccessControl
	raw.AccessKey = u.AccessKey
	raw.Authentication = u.Authentication
	raw.Password = u.Password
	raw.Type = u.Type
	raw.Username = u.Username

	return json.Marshal(raw)
}
