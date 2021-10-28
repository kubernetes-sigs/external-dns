//go:generate go run ../../gen/model_response/main.go -package ecloudflex -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package ecloudflex -source model.go -destination model_paginated_generated.go

package ecloudflex

import "github.com/ukfast/sdk-go/pkg/connection"

// Project represents an eCloud Flex project
// +genie:model_response
// +genie:model_paginated
type Project struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	CreatedAt connection.DateTime `json:"created_at"`
}
