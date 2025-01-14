// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"net/http"
)

type GetCommerceOrderRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for an order.
	OrderID string `pathParam:"style=simple,explode=false,name=orderId"`
}

func (o *GetCommerceOrderRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCommerceOrderRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCommerceOrderRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

type GetCommerceOrderResponse struct {
	// OK
	CommerceOrder *shared.CommerceOrder
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCommerceOrderResponse) GetCommerceOrder() *shared.CommerceOrder {
	if o == nil {
		return nil
	}
	return o.CommerceOrder
}

func (o *GetCommerceOrderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCommerceOrderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCommerceOrderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
