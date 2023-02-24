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

// RRSetPost for rr set info.
type RrsetRrSetPost struct {
	// user comment
	Comment string `json:"comment,omitempty"`
	// name of the record which should be a valid domain according to rfc1035 Section 2.3.4
	Name string `json:"name"`
	// records
	Records []RrsetRecordPost `json:"records"`
	// time to live
	Ttl int32 `json:"ttl"`
	// record set type
	Type_ string `json:"type"`
}
