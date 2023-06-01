// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
	"net/http"
)

type GetTaxComponentRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a tax component.
	TaxID string `pathParam:"style=simple,explode=false,name=taxId"`
}

// GetTaxComponent409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetTaxComponent409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetTaxComponentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	TaxComponent *shared.TaxComponent
	// The data type's dataset has not been requested or is still syncing.
	GetTaxComponent409ApplicationJSONObject *GetTaxComponent409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
