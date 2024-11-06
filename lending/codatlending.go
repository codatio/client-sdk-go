// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package lending

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/v5/internal/hooks"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/retry"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production
	"https://api.codat.io",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// CodatLending - Lending API: Our Lending API helps you make smarter credit decisions on small businesses by enabling you to pull your customers' latest data from accounting, banking, and commerce software they are already using. It also includes features to help providers verify the accuracy of data and process it more efficiently.
//
// The Lending API is built on top of the latest accounting, commerce, and banking data, providing you with the most important data points you need to get a full picture of SMB creditworthiness and make a comprehensive assessment of your customers.
//
// [Explore product](https://docs.codat.io/lending/overview) | [See OpenAPI spec](https://github.com/codatio/oas)
//
// <!-- Start Codat Tags Table -->
// ## Endpoints
//
// | Endpoints | Description |
// | :- |:- |
// | Companies | Create and manage your SMB users' companies. |
// | Connections | Create new and manage existing data connections for a company. |
// | Bank statements | Retrieve banking data from linked bank accounts. |
// | Sales | Retrieve standardized sales data from a linked commerce software. |
// | Financial statements | Financial data and reports from a linked accounting software. |
// | Liabilities | Debt and other liabilities. |
// | Accounts payable | Data from a linked accounting software representing money the business owes money to its suppliers. |
// | Accounts receivable | Data from a linked accounting software representing money owed to the business for sold goods or services. |
// | Transactions | Data from a linked accounting software representing transactions. |
// | Company info | View company information fetched from the source platform. |
// | Data integrity | Match mutable accounting data with immutable banking data to increase confidence in financial data. |
// | Excel reports | Download reports in Excel format. |
// | Manage data | Control how data is retrieved from an integration. |
// | File upload | Endpoints to manage uploaded files. |
// | Loan writeback | Implement the [loan writeback](https://docs.codat.io/lending/guides/loan-writeback/introduction) procedure in your lending process to maintain an accurate position of a loan during the entire lending cycle. |
// <!-- End Codat Tags Table -->
type CodatLending struct {
	// Create and manage your SMB users' companies.
	Companies *Companies
	// Create new and manage existing data connections for a company.
	Connections   *Connections
	LoanWriteback *LoanWriteback
	// Retrieve banking data from linked bank accounts.
	BankStatements *BankStatements
	Transactions   *Transactions
	// Access bank transactions from an accounting software.
	AccountingBankData *CodatLendingAccountingBankData
	Banking            *Banking
	AccountsPayable    *AccountsPayable
	Sales              *Sales
	// View company information fetched from the source platform.
	CompanyInfo        *CompanyInfo
	AccountsReceivable *AccountsReceivable
	// Endpoints to manage uploaded files.
	FileUpload          *FileUpload
	FinancialStatements *FinancialStatements
	ManageData          *ManageData
	// Endpoints to manage generation of reports
	ManageReports *ManageReports
	// Debt and other liabilities.
	Liabilities *Liabilities
	// Match mutable accounting data with immutable banking data to increase confidence in financial data.
	DataIntegrity *DataIntegrity
	// Download reports in Excel format.
	ExcelReports *ExcelReports

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatLending)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatLending) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatLending) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatLending) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatLending) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatLending) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *CodatLending) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *CodatLending) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *CodatLending) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatLending {
	sdk := &CodatLending{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "3.0.0",
			SDKVersion:        "5.5.0",
			GenVersion:        "2.451.0",
			UserAgent:         "speakeasy-sdk/go 5.5.0 2.451.0 3.0.0 github.com/codatio/client-sdk-go/lending",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Companies = newCompanies(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.LoanWriteback = newLoanWriteback(sdk.sdkConfiguration)

	sdk.BankStatements = newBankStatements(sdk.sdkConfiguration)

	sdk.Transactions = newTransactions(sdk.sdkConfiguration)

	sdk.AccountingBankData = newCodatLendingAccountingBankData(sdk.sdkConfiguration)

	sdk.Banking = newBanking(sdk.sdkConfiguration)

	sdk.AccountsPayable = newAccountsPayable(sdk.sdkConfiguration)

	sdk.Sales = newSales(sdk.sdkConfiguration)

	sdk.CompanyInfo = newCompanyInfo(sdk.sdkConfiguration)

	sdk.AccountsReceivable = newAccountsReceivable(sdk.sdkConfiguration)

	sdk.FileUpload = newFileUpload(sdk.sdkConfiguration)

	sdk.FinancialStatements = newFinancialStatements(sdk.sdkConfiguration)

	sdk.ManageData = newManageData(sdk.sdkConfiguration)

	sdk.ManageReports = newManageReports(sdk.sdkConfiguration)

	sdk.Liabilities = newLiabilities(sdk.sdkConfiguration)

	sdk.DataIntegrity = newDataIntegrity(sdk.sdkConfiguration)

	sdk.ExcelReports = newExcelReports(sdk.sdkConfiguration)

	return sdk
}
