// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v2/pkg/models/shared"
	"net/http"
)

type GetAccountingCashFlowStatementRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Number of months defining the period of interest.
	PeriodLength int `queryParam:"style=form,explode=true,name=periodLength"`
	// Number of periods with `periodLength` to compare.
	PeriodsToCompare int     `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *string `queryParam:"style=form,explode=true,name=startMonth"`
}

func (o *GetAccountingCashFlowStatementRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingCashFlowStatementRequest) GetPeriodLength() int {
	if o == nil {
		return 0
	}
	return o.PeriodLength
}

func (o *GetAccountingCashFlowStatementRequest) GetPeriodsToCompare() int {
	if o == nil {
		return 0
	}
	return o.PeriodsToCompare
}

func (o *GetAccountingCashFlowStatementRequest) GetStartMonth() *string {
	if o == nil {
		return nil
	}
	return o.StartMonth
}

type GetAccountingCashFlowStatementResponse struct {
	// Success
	AccountingCashFlowStatement *shared.AccountingCashFlowStatement
	ContentType                 string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingCashFlowStatementResponse) GetAccountingCashFlowStatement() *shared.AccountingCashFlowStatement {
	if o == nil {
		return nil
	}
	return o.AccountingCashFlowStatement
}

func (o *GetAccountingCashFlowStatementResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingCashFlowStatementResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingCashFlowStatementResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingCashFlowStatementResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
