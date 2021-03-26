/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// AccountCoinsRequest is utilized to make a request on the /account/coins endpoint.
type AccountCoinsRequest struct {

	NetworkIdentifier *NetworkIdentifier `json:"network_identifier"`

	AccountIdentifier *AccountIdentifier `json:"account_identifier"`
	// Include state from the mempool when looking up an account's unspent coins. Note, using this functionality breaks any guarantee of idempotency.
	IncludeMempool bool `json:"include_mempool"`
	// In some cases, the caller may not want to retrieve coins for all currencies for an AccountIdentifier. If the currencies field is populated, only coins for the specified currencies will be returned. If not populated, all unspent coins will be returned.
	Currencies []Currency `json:"currencies,omitempty"`
}
