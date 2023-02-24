/*
 * STACKIT DNS API
 *
 * This api provides dns
 *
 * API version: 1.0
 * Contact: stackit-dns@mail.schwarz
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// ResponseRRSet for rr set info.
type RrsetResponseRrSet struct {
	Message string `json:"message,omitempty"`
	Rrset *DomainRrSet `json:"rrset"`
}
