/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Blocks contain an array of Transactions that occurred at a particular BlockIdentifier. A hard requirement for blocks returned by Rosetta implementations is that they MUST be _inalterable_: once a client has requested and received a block identified by a specific BlockIndentifier, all future calls for that same BlockIdentifier must return the same block contents.
type Block struct {

	BlockIdentifier *BlockIdentifier `json:"block_identifier"`

	ParentBlockIdentifier *BlockIdentifier `json:"parent_block_identifier"`

	Timestamp int64 `json:"timestamp"`

	Transactions []Transaction `json:"transactions"`

	Metadata *interface{} `json:"metadata,omitempty"`
}
