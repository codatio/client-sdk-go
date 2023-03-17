package operations

import (
	"net/http"
	"time"
)

type GetAccountsForEnhancedProfitAndLossRequest struct {
	CompanyID       string `pathParam:"style=simple,explode=false,name=companyId"`
	NumberOfPeriods int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	ReportDate      string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetAccountsForEnhancedProfitAndLossEnhancedReportReportInfo struct {
	CompanyName   *string    `json:"companyName,omitempty"`
	Currency      *string    `json:"currency,omitempty"`
	GeneratedDate *time.Time `json:"generatedDate,omitempty"`
	ReportName    *string    `json:"reportName,omitempty"`
}

type GetAccountsForEnhancedProfitAndLossEnhancedReportReportItemsAccountCategoryLevels struct {
	Confidence *float64 `json:"confidence,omitempty"`
	LevelName  *string  `json:"levelName,omitempty"`
}

type GetAccountsForEnhancedProfitAndLossEnhancedReportReportItemsAccountCategory struct {
	Levels []GetAccountsForEnhancedProfitAndLossEnhancedReportReportItemsAccountCategoryLevels `json:"levels,omitempty"`
	Status *string                                                                             `json:"status,omitempty"`
}

type GetAccountsForEnhancedProfitAndLossEnhancedReportReportItems struct {
	AccountCategory *GetAccountsForEnhancedProfitAndLossEnhancedReportReportItemsAccountCategory `json:"accountCategory,omitempty"`
	AccountID       *string                                                                      `json:"accountId,omitempty"`
	AccountName     *string                                                                      `json:"accountName,omitempty"`
	Balance         *string                                                                      `json:"balance,omitempty"`
	Date            *time.Time                                                                   `json:"date,omitempty"`
}

type GetAccountsForEnhancedProfitAndLossEnhancedReport struct {
	ReportInfo  *GetAccountsForEnhancedProfitAndLossEnhancedReportReportInfo   `json:"reportInfo,omitempty"`
	ReportItems []GetAccountsForEnhancedProfitAndLossEnhancedReportReportItems `json:"reportItems,omitempty"`
}

type GetAccountsForEnhancedProfitAndLossResponse struct {
	ContentType    string
	EnhancedReport *GetAccountsForEnhancedProfitAndLossEnhancedReport
	StatusCode     int
	RawResponse    *http.Response
}
