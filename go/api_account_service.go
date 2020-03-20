/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

// AccountApiService is a service that implents the logic for the AccountApiServicer
// This service should implement the business logic for every endpoint for the AccountApi API.
// Include any external packages or services that will be required by this service.
type AccountApiService struct {
}

// NewAccountApiService creates a default api service
func NewAccountApiService() AccountApiServicer {
	return &AccountApiService{}
}

// AccountBalance - Get an Account Balance
func (s *AccountApiService) AccountBalance(accountBalanceRequest AccountBalanceRequest) (interface{}, error) {
	// TODO - update AccountBalance with the required logic for this service method.
	// Add api_account_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'AccountBalance' not implemented")
}
