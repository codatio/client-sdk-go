// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type ListDirectIncomeAttachmentsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a direct income
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

func (o *ListDirectIncomeAttachmentsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListDirectIncomeAttachmentsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListDirectIncomeAttachmentsRequest) GetDirectIncomeID() string {
	if o == nil {
		return ""
	}
	return o.DirectIncomeID
}

type ListDirectIncomeAttachmentsResponse struct {
	// Success
	AttachmentsDataset *shared.AttachmentsDataset
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

func (o *ListDirectIncomeAttachmentsResponse) GetAttachmentsDataset() *shared.AttachmentsDataset {
	if o == nil {
		return nil
	}
	return o.AttachmentsDataset
}

func (o *ListDirectIncomeAttachmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDirectIncomeAttachmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDirectIncomeAttachmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListDirectIncomeAttachmentsResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
