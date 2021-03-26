/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.
type TransactionIdentifier struct {
	// Any transactions that are attributable only to a block (ex: a block event) should use the hash of the block as the identifier.
	Hash string `json:"hash"`
}
