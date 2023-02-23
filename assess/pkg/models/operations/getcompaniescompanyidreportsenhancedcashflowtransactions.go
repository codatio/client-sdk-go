package operations

import (
	"time"
)

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsQueryParams struct {
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsRequest struct {
	PathParams  GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsPathParams
	QueryParams GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsQueryParams
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSourcesAccountsReportSourceRef struct {
	SourceType *string `json:"sourceType,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSourcesAccounts struct {
	AccountName     *string                                                                                                                 `json:"accountName,omitempty"`
	AccountProvider *string                                                                                                                 `json:"accountProvider,omitempty"`
	AccountType     *string                                                                                                                 `json:"accountType,omitempty"`
	Currency        *string                                                                                                                 `json:"currency,omitempty"`
	CurrentBalance  *float64                                                                                                                `json:"currentBalance,omitempty"`
	PlatformName    *string                                                                                                                 `json:"platformName,omitempty"`
	SourceRef       *GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSourcesAccountsReportSourceRef `json:"sourceRef,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSources struct {
	Accounts []GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSourcesAccounts `json:"accounts,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportInfo struct {
	CompanyName   *string `json:"companyName,omitempty"`
	GeneratedDate *string `json:"generatedDate,omitempty"`
	PageNumber    *int64  `json:"pageNumber,omitempty"`
	PageSize      *int64  `json:"pageSize,omitempty"`
	ReportName    *string `json:"reportName,omitempty"`
	TotalResults  *int64  `json:"totalResults,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactionsReportSourceRef struct {
	SourceType *string `json:"sourceType,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactionsTransactionCategory struct {
	Confidence *float64 `json:"confidence,omitempty"`
	Levels     []string `json:"levels,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactions struct {
	Amount              *float64                                                                                                                        `json:"amount,omitempty"`
	Currency            *string                                                                                                                         `json:"currency,omitempty"`
	Date                *time.Time                                                                                                                      `json:"date,omitempty"`
	Description         *string                                                                                                                         `json:"description,omitempty"`
	ID                  *string                                                                                                                         `json:"id,omitempty"`
	SourceRef           *GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactionsReportSourceRef     `json:"sourceRef,omitempty"`
	TransactionCategory *GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactionsTransactionCategory `json:"transactionCategory,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItems struct {
	Transactions []GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactions `json:"transactions,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactions struct {
	DataSources []GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSources `json:"dataSources,omitempty"`
	ReportInfo  *GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportInfo   `json:"reportInfo,omitempty"`
	ReportItems []GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItems `json:"reportItems,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsResponse struct {
	ContentType                  string
	EnhancedCashFlowTransactions *GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsEnhancedCashFlowTransactions
	StatusCode                   int
}
