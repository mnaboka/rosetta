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
	"encoding/json"
	"net/http"
	"strings"
	// "github.com/gorilla/mux"
)

// A NetworkApiController binds http requests to an api service and writes the service results to the http response
type NetworkApiController struct {
	service NetworkApiServicer
}

// NewNetworkApiController creates a default api controller
func NewNetworkApiController(s NetworkApiServicer) Router {
	return &NetworkApiController{service: s}
}

// Routes returns all of the api route for the NetworkApiController
func (c *NetworkApiController) Routes() Routes {
	return Routes{
		{
			"NetworkStatus",
			strings.ToUpper("Post"),
			"/network/status",
			c.NetworkStatus,
		},
	}
}

// NetworkStatus - Get Network Status
func (c *NetworkApiController) NetworkStatus(w http.ResponseWriter, r *http.Request) {
	networkStatusRequest := &NetworkStatusRequest{}
	if err := json.NewDecoder(r.Body).Decode(&networkStatusRequest); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.NetworkStatus(*networkStatusRequest)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
