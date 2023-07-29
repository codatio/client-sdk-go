// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetPurchaseOrderRequest struct {
	CompanyID       string `pathParam:"style=simple,explode=false,name=companyId"`
	PurchaseOrderID string `pathParam:"style=simple,explode=false,name=purchaseOrderId"`
}

func (o *GetPurchaseOrderRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetPurchaseOrderRequest) GetPurchaseOrderID() string {
	if o == nil {
		return ""
	}
	return o.PurchaseOrderID
}

type GetPurchaseOrderResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// Success
	PurchaseOrder *shared.PurchaseOrder
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetPurchaseOrderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPurchaseOrderResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetPurchaseOrderResponse) GetPurchaseOrder() *shared.PurchaseOrder {
	if o == nil {
		return nil
	}
	return o.PurchaseOrder
}

func (o *GetPurchaseOrderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPurchaseOrderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
