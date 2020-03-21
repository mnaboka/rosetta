/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type NetworkStatus struct {
	NetworkIdentifier PartialNetworkIdentifier `json:"network_identifier"`

	NetworkInformation NetworkInformation `json:"network_information"`
}
