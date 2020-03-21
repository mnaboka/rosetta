/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"errors"
)

// NetworkApiService is a service that implents the logic for the NetworkApiServicer
// This service should implement the business logic for every endpoint for the NetworkApi API.
// Include any external packages or services that will be required by this service.
type NetworkApiService struct {
}

// NewNetworkApiService creates a default api service
func NewNetworkApiService() NetworkApiServicer {
	return &NetworkApiService{}
}

// NetworkStatus - Get Network Status
func (s *NetworkApiService) NetworkStatus(networkStatusRequest NetworkStatusRequest) (interface{}, error) {
	// TODO - update NetworkStatus with the required logic for this service method.
	// Add api_network_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'NetworkStatus' not implemented")
}
