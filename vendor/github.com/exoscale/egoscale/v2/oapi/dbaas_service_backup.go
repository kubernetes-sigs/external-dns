package oapi

import (
	"encoding/json"
	"time"
)

// UnmarshalJSON unmarshals a DbaasServiceBackup structure into a temporary structure whose BackupTime field of
// type string to be able to parse the original timestamp (ISO 8601) into a time.Time object, since
// json.Unmarshal() only supports RFC 3339 format.
func (b *DbaasServiceBackup) UnmarshalJSON(data []byte) error {
	raw := struct {
		BackupName string `json:"backup-name"`
		BackupTime string `json:"backup-time"`
		DataSize   int64  `json:"data-size"`
	}{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	backupTime, err := time.Parse(iso8601Format, raw.BackupTime)
	if err != nil {
		return err
	}
	b.BackupTime = backupTime

	b.BackupName = raw.BackupName
	b.DataSize = raw.DataSize

	return nil
}

// MarshalJSON returns the JSON encoding of a DbaasServiceBackup structure after having formatted the BackupTime
// field in the original timestamp (ISO 8601), since time.MarshalJSON() only supports RFC 3339 format.
func (b *DbaasServiceBackup) MarshalJSON() ([]byte, error) {
	raw := struct {
		BackupName string `json:"backup-name"`
		BackupTime string `json:"backup-time"`
		DataSize   int64  `json:"data-size"`
	}{}

	raw.BackupName = b.BackupName
	raw.BackupTime = b.BackupTime.Format(iso8601Format)
	raw.DataSize = b.DataSize

	return json.Marshal(raw)
}
