/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type Transaction struct {
	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`

	Operations []Operation `json:"operations"`

	// Transactions that are related to other transactions (like a cross-shard transactioin) should include the `tranaction_identifier` of these transactions in the metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
