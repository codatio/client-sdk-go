package operations

import (
	"time"
)

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsQueryParams struct {
	NumberOfPeriods int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	ReportDate      string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsRequest struct {
	PathParams  GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsPathParams
	QueryParams GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsQueryParams
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReportReportInfo struct {
	CompanyName   *string    `json:"companyName,omitempty"`
	Currency      *string    `json:"currency,omitempty"`
	GeneratedDate *time.Time `json:"generatedDate,omitempty"`
	ReportName    *string    `json:"reportName,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReportReportItemsAccountCategoryLevels struct {
	Confidence *float64 `json:"confidence,omitempty"`
	LevelName  *string  `json:"levelName,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReportReportItemsAccountCategory struct {
	Levels []GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReportReportItemsAccountCategoryLevels `json:"levels,omitempty"`
	Status *string                                                                                                   `json:"status,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReportReportItems struct {
	AccountCategory *GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReportReportItemsAccountCategory `json:"accountCategory,omitempty"`
	AccountID       *string                                                                                            `json:"accountId,omitempty"`
	AccountName     *string                                                                                            `json:"accountName,omitempty"`
	Balance         *string                                                                                            `json:"balance,omitempty"`
	Date            *time.Time                                                                                         `json:"date,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReport struct {
	ReportInfo  *GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReportReportInfo   `json:"reportInfo,omitempty"`
	ReportItems []GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReportReportItems `json:"reportItems,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsResponse struct {
	ContentType    string
	EnhancedReport *GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsEnhancedReport
	StatusCode     int
}
