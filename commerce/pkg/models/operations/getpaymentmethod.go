// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
	"net/http"
)

type GetPaymentMethodRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a payment method.
	PaymentMethodID string `pathParam:"style=simple,explode=false,name=paymentMethodId"`
}

// GetPaymentMethod409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetPaymentMethod409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetPaymentMethodResponse struct {
	ContentType string
	// OK
	PaymentMethod *shared.PaymentMethod
	StatusCode    int
	RawResponse   *http.Response
	// The data type's dataset has not been requested or is still syncing.
	GetPaymentMethod409ApplicationJSONObject *GetPaymentMethod409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
