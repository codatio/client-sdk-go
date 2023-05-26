// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
	"net/http"
)

type ListLocationsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

// ListLocations409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type ListLocations409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListLocationsResponse struct {
	ContentType string
	// OK
	LocationsResponse *shared.LocationsResponse
	StatusCode        int
	RawResponse       *http.Response
	// The data type's dataset has not been requested or is still syncing.
	ListLocations409ApplicationJSONObject *ListLocations409ApplicationJSON
	// Your `query` parameter was not correctly formed
	Schema *shared.Schema
}
