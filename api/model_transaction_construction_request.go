/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type TransactionConstructionRequest struct {
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`

	AccountIdentifier AccountIdentifier `json:"account_identifier"`

	// Some blockchains require different metadata for different types of transaction construction (ex: delegation versus a transfer).  Instead of requiring a blockchain node to return all possible types of metadata for construction (which may require multiple node fetches), the client can specify a `method` to limit the metadata returned to only the subset required.
	Method string `json:"method,omitempty"`
}
