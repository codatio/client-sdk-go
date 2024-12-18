// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"net/http"
)

type GetTransactionCategoryRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// The unique identifier for a banking transaction category
	TransactionCategoryID string `pathParam:"style=simple,explode=false,name=transactionCategoryId"`
}

func (o *GetTransactionCategoryRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetTransactionCategoryRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetTransactionCategoryRequest) GetTransactionCategoryID() string {
	if o == nil {
		return ""
	}
	return o.TransactionCategoryID
}

type GetTransactionCategoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	TransactionCategory *shared.TransactionCategory
}

func (o *GetTransactionCategoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransactionCategoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransactionCategoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransactionCategoryResponse) GetTransactionCategory() *shared.TransactionCategory {
	if o == nil {
		return nil
	}
	return o.TransactionCategory
}
