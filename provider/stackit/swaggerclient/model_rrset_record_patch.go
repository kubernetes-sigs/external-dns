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

// RecordPatch for record patch in record set.
type RrsetRecordPatch struct {
	Action string `json:"action"`
	// records
	Records []RrsetRecordPost `json:"records"`
}
