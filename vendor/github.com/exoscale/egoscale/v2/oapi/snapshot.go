package oapi

import (
	"encoding/json"
	"time"
)

// UnmarshalJSON unmarshals a Snapshot structure into a temporary structure whose "CreatedAt" field of type
// string to be able to parse the original timestamp (ISO 8601) into a time.Time object, since json.Unmarshal()
// only supports RFC 3339 format.
func (t *Snapshot) UnmarshalJSON(data []byte) error {
	raw := struct {
		CreatedAt *string `json:"created-at,omitempty"`
		Export    *struct {
			Md5sum       *string `json:"md5sum,omitempty"`
			PresignedUrl *string `json:"presigned-url,omitempty"` // nolint:revive
		} `json:"export,omitempty"`
		Id       *string        `json:"id,omitempty"` // nolint:revive
		Instance *Instance      `json:"instance,omitempty"`
		Name     *string        `json:"name,omitempty"`
		Size     *int64         `json:"size,omitempty"`
		State    *SnapshotState `json:"state,omitempty"`
	}{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.CreatedAt != nil {
		createdAt, err := time.Parse(iso8601Format, *raw.CreatedAt)
		if err != nil {
			return err
		}
		t.CreatedAt = &createdAt
	}

	t.Export = raw.Export
	t.Id = raw.Id
	t.Instance = raw.Instance
	t.Name = raw.Name
	t.Size = raw.Size
	t.State = raw.State

	return nil
}

// MarshalJSON returns the JSON encoding of a Snapshot structure after having formatted the CreatedAt field
// in the original timestamp (ISO 8601), since time.MarshalJSON() only supports RFC 3339 format.
func (t *Snapshot) MarshalJSON() ([]byte, error) {
	raw := struct {
		CreatedAt *string `json:"created-at,omitempty"`
		Export    *struct {
			Md5sum       *string `json:"md5sum,omitempty"`
			PresignedUrl *string `json:"presigned-url,omitempty"` // nolint:revive
		} `json:"export,omitempty"`
		Id       *string        `json:"id,omitempty"` // nolint:revive
		Instance *Instance      `json:"instance,omitempty"`
		Name     *string        `json:"name,omitempty"`
		Size     *int64         `json:"size,omitempty"`
		State    *SnapshotState `json:"state,omitempty"`
	}{}

	if t.CreatedAt != nil {
		createdAt := t.CreatedAt.Format(iso8601Format)
		raw.CreatedAt = &createdAt
	}

	raw.Export = t.Export
	raw.Id = t.Id
	raw.Instance = t.Instance
	raw.Name = t.Name
	raw.Size = t.Size
	raw.State = t.State

	return json.Marshal(raw)
}
