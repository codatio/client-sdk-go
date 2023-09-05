// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"net/http"
)

type GetAccountingProfitAndLossRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Number of months defining the period of interest.
	PeriodLength int `queryParam:"style=form,explode=true,name=periodLength"`
	// Number of periods with `periodLength` to compare.
	PeriodsToCompare int     `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *string `queryParam:"style=form,explode=true,name=startMonth"`
}

func (o *GetAccountingProfitAndLossRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingProfitAndLossRequest) GetPeriodLength() int {
	if o == nil {
		return 0
	}
	return o.PeriodLength
}

func (o *GetAccountingProfitAndLossRequest) GetPeriodsToCompare() int {
	if o == nil {
		return 0
	}
	return o.PeriodsToCompare
}

func (o *GetAccountingProfitAndLossRequest) GetStartMonth() *string {
	if o == nil {
		return nil
	}
	return o.StartMonth
}

type GetAccountingProfitAndLossResponse struct {
	// Success
	AccountingProfitAndLossReport *shared.AccountingProfitAndLossReport
	ContentType                   string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingProfitAndLossResponse) GetAccountingProfitAndLossReport() *shared.AccountingProfitAndLossReport {
	if o == nil {
		return nil
	}
	return o.AccountingProfitAndLossReport
}

func (o *GetAccountingProfitAndLossResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingProfitAndLossResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingProfitAndLossResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingProfitAndLossResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}