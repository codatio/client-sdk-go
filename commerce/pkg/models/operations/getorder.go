// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
	"net/http"
)

type GetOrderRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for an order.
	OrderID string `pathParam:"style=simple,explode=false,name=orderId"`
}

// GetOrder409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetOrder409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetOrderResponse struct {
	ContentType string
	// OK
	Order       *shared.Order
	StatusCode  int
	RawResponse *http.Response
	// The data type's dataset has not been requested or is still syncing.
	GetOrder409ApplicationJSONObject *GetOrder409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
