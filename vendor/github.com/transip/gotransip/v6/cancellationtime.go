package gotransip

// CancellationTime represents the possible ways of canceling a contract
type CancellationTime string

// CancellationRequest is used to generate a json body that contains the endTime property
// the endTime could either be 'end' or 'immediately'
type CancellationRequest struct {
	EndTime CancellationTime `json:"endTime"`
}

var (
	// CancellationTimeEnd specifies to cancel the contract when the contract was
	// due to end anyway
	CancellationTimeEnd CancellationTime = "end"
	// CancellationTimeImmediately specifies to cancel the contract immediately
	CancellationTimeImmediately CancellationTime = "immediately"
)
