// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// APIGatewayService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIGatewayService] method instead.
type APIGatewayService struct {
	Options            []option.RequestOption
	Configurations     *ConfigurationService
	Discovery          *DiscoveryService
	Operations         *OperationService
	Schemas            *SchemaService
	Settings           *SettingService
	UserSchemas        *UserSchemaService
	ExpressionTemplate *ExpressionTemplateService
}

// NewAPIGatewayService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIGatewayService(opts ...option.RequestOption) (r *APIGatewayService) {
	r = &APIGatewayService{}
	r.Options = opts
	r.Configurations = NewConfigurationService(opts...)
	r.Discovery = NewDiscoveryService(opts...)
	r.Operations = NewOperationService(opts...)
	r.Schemas = NewSchemaService(opts...)
	r.Settings = NewSettingService(opts...)
	r.UserSchemas = NewUserSchemaService(opts...)
	r.ExpressionTemplate = NewExpressionTemplateService(opts...)
	return
}
