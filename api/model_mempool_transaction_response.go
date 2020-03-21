/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type MempoolTransactionResponse struct {
	Transaction Transaction `json:"transaction"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
