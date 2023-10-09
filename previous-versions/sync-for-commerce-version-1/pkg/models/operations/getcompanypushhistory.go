// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"net/http"
)

type GetCompanyPushHistoryRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `default:"100" queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (g GetCompanyPushHistoryRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetCompanyPushHistoryRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetCompanyPushHistoryRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCompanyPushHistoryRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetCompanyPushHistoryRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetCompanyPushHistoryRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetCompanyPushHistoryRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type GetCompanyPushHistoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Your `query` parameter was not correctly formed
	ErrorMessage *shared.ErrorMessage
	// OK
	PushOperations *shared.PushOperations
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCompanyPushHistoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCompanyPushHistoryResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetCompanyPushHistoryResponse) GetPushOperations() *shared.PushOperations {
	if o == nil {
		return nil
	}
	return o.PushOperations
}

func (o *GetCompanyPushHistoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCompanyPushHistoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
