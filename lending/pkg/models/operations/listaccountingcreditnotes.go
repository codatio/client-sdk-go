// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"net/http"
)

type ListAccountingCreditNotesRequest struct {
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

func (l ListAccountingCreditNotesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountingCreditNotesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountingCreditNotesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListAccountingCreditNotesRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListAccountingCreditNotesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListAccountingCreditNotesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListAccountingCreditNotesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type ListAccountingCreditNotesResponse struct {
	// Success
	AccountingCreditNotes *shared.AccountingCreditNotes
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountingCreditNotesResponse) GetAccountingCreditNotes() *shared.AccountingCreditNotes {
	if o == nil {
		return nil
	}
	return o.AccountingCreditNotes
}

func (o *ListAccountingCreditNotesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountingCreditNotesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountingCreditNotesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
