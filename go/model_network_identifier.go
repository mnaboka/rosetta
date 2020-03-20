/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NetworkIdentifier - The `network_identifier` specifies which network a particular object is associated with. 
type NetworkIdentifier struct {

	Blockchain string `json:"blockchain"`

	Network string `json:"network"`

	SubNetworkIdentifier SubNetworkIdentifier `json:"sub_network_identifier,omitempty"`
}
