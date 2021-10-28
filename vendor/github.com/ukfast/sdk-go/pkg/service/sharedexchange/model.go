//go:generate go run ../../gen/model_response/main.go -package sharedexchange -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package sharedexchange -source model.go -destination model_paginated_generated.go

package sharedexchange

import "github.com/ukfast/sdk-go/pkg/connection"

// Domain represents an Shared Exchange domain
// +genie:model_response
// +genie:model_paginated
type Domain struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	Version   string              `json:"version"`
	CreatedAt connection.DateTime `json:"created_at"`
}
