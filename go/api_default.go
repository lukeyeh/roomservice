/*
 * roomservice
 *
 * An implementation of joinable rooms
 *
 * API version: 0.0.1
 * Contact: decline@umass.edu
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{service: s}
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"GetRoomByCode",
			strings.ToUpper("Get"),
			"/v1/room/{code}",
			c.GetRoomByCode,
		},
	}
}

// GetRoomByCode -
func (c *DefaultApiController) GetRoomByCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]
	result, err := c.service.GetRoomByCode(code)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
