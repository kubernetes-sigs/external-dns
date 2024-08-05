package oapi

// NewReference creates an oapi reference (inlined type)
func NewReference(command, id, link *string) *struct {
	Command *string `json:"command,omitempty"`
	Id      *string `json:"id,omitempty"` // revive:disable-line
	Link    *string `json:"link,omitempty"`
} {
	return &struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}{
		Id:      id,
		Command: command,
		Link:    link,
	}
}
