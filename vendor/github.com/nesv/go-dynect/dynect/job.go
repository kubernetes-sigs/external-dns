package dynect

type JobData struct {
	Status   string         `json:"status"`
	Data     interface{}    `json:"data"`
	ID       int            `json:"job_id"`
	Messages []MessageBlock `json:"msgs"`
}
