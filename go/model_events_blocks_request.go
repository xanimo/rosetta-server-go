/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// EventsBlocksRequest is utilized to fetch a sequence of BlockEvents indicating which blocks were added and removed from storage to reach the current state.
type EventsBlocksRequest struct {

	NetworkIdentifier *NetworkIdentifier `json:"network_identifier"`
	// offset is the offset into the event stream to sync events from. If this field is not populated, we return the limit events backwards from tip. If this is set to 0, we start from the beginning.
	Offset int64 `json:"offset,omitempty"`
	// limit is the maximum number of events to fetch in one call. The implementation may return <= limit events.
	Limit int64 `json:"limit,omitempty"`
}
