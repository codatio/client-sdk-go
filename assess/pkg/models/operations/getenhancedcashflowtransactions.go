package operations

import (
	"net/http"
	"time"
)

type GetEnhancedCashFlowTransactionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetEnhancedCashFlowTransactionsQueryParams struct {
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetEnhancedCashFlowTransactionsRequest struct {
	PathParams  GetEnhancedCashFlowTransactionsPathParams
	QueryParams GetEnhancedCashFlowTransactionsQueryParams
}

type GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSourcesAccountsReportSourceRef struct {
	SourceType *string `json:"sourceType,omitempty"`
}

type GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSourcesAccounts struct {
	AccountName     *string                                                                                        `json:"accountName,omitempty"`
	AccountProvider *string                                                                                        `json:"accountProvider,omitempty"`
	AccountType     *string                                                                                        `json:"accountType,omitempty"`
	Currency        *string                                                                                        `json:"currency,omitempty"`
	CurrentBalance  *float64                                                                                       `json:"currentBalance,omitempty"`
	PlatformName    *string                                                                                        `json:"platformName,omitempty"`
	SourceRef       *GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSourcesAccountsReportSourceRef `json:"sourceRef,omitempty"`
}

type GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSources struct {
	Accounts []GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSourcesAccounts `json:"accounts,omitempty"`
}

type GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportInfo struct {
	CompanyName   *string `json:"companyName,omitempty"`
	GeneratedDate *string `json:"generatedDate,omitempty"`
	PageNumber    *int64  `json:"pageNumber,omitempty"`
	PageSize      *int64  `json:"pageSize,omitempty"`
	ReportName    *string `json:"reportName,omitempty"`
	TotalResults  *int64  `json:"totalResults,omitempty"`
}

type GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactionsReportSourceRef struct {
	SourceType *string `json:"sourceType,omitempty"`
}

type GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactionsTransactionCategory struct {
	Confidence *float64 `json:"confidence,omitempty"`
	Levels     []string `json:"levels,omitempty"`
}

type GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactions struct {
	Amount              *float64                                                                                               `json:"amount,omitempty"`
	Currency            *string                                                                                                `json:"currency,omitempty"`
	Date                *time.Time                                                                                             `json:"date,omitempty"`
	Description         *string                                                                                                `json:"description,omitempty"`
	ID                  *string                                                                                                `json:"id,omitempty"`
	SourceRef           *GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactionsReportSourceRef     `json:"sourceRef,omitempty"`
	TransactionCategory *GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactionsTransactionCategory `json:"transactionCategory,omitempty"`
}

type GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItems struct {
	Transactions []GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItemsTransactions `json:"transactions,omitempty"`
}

type GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactions struct {
	DataSources []GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsDataSources `json:"dataSources,omitempty"`
	ReportInfo  *GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportInfo   `json:"reportInfo,omitempty"`
	ReportItems []GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactionsReportItems `json:"reportItems,omitempty"`
}

type GetEnhancedCashFlowTransactionsResponse struct {
	ContentType                  string
	EnhancedCashFlowTransactions *GetEnhancedCashFlowTransactionsEnhancedCashFlowTransactions
	StatusCode                   int
	RawResponse                  *http.Response
}
