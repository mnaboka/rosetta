/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type AccountBalanceRequest struct {
	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`

	AccountIdentifier AccountIdentifier `json:"account_identifier"`
}
