// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
	"net/http"
)

type GetProductRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a product.
	ProductID string `pathParam:"style=simple,explode=false,name=productId"`
}

// GetProduct409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetProduct409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetProductResponse struct {
	ContentType string
	// OK
	Product     *shared.Product
	StatusCode  int
	RawResponse *http.Response
	// The data type's dataset has not been requested or is still syncing.
	GetProduct409ApplicationJSONObject *GetProduct409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
