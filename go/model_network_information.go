/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type NetworkInformation struct {

	CurrentBlockIdentifier BlockIdentifier `json:"current_block_identifier"`

	// The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second. 
	CurrentBlockTimestamp int64 `json:"current_block_timestamp"`

	GenesisBlockIdentifier BlockIdentifier `json:"genesis_block_identifier"`

	Peers []Peer `json:"peers"`
}
