// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package accounting

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
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

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// CodatAccounting - Accounting API: A flexible API for pulling accounting data, normalized and aggregated from 20 accounting integrations.
//
// Standardize how you connect to your customers’ accounting software. View, create, update, and delete data in the same way for all the leading accounting platforms.
//
// [Read more...](https://docs.codat.io/accounting-api/overview)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
type CodatAccounting struct {
	// Account transactions
	AccountTransactions *AccountTransactions
	// Bank accounts
	BankAccounts *BankAccounts
	// Bank transactions for bank accounts
	BankAccountTransactions *BankAccountTransactions
	// Bills
	Bills *Bills
	// Customers
	Customers *Customers
	// Direct costs
	DirectCosts *DirectCosts
	// Direct incomes
	DirectIncomes *DirectIncomes
	// Invoices
	Invoices *Invoices
	// Item receipts
	ItemReceipts *ItemReceipts
	// Purchase orders
	PurchaseOrders *PurchaseOrders
	// Suppliers
	Suppliers *Suppliers
	// Transfers
	Transfers *Transfers
	// Bill credit notes
	BillCreditNotes *BillCreditNotes
	// Bill payments
	BillPayments *BillPayments
	// Accounts
	Accounts *Accounts
	// Credit notes
	CreditNotes *CreditNotes
	// Items
	Items *Items
	// Journal entries
	JournalEntries *JournalEntries
	// Journals
	Journals *Journals
	// Payments
	Payments *Payments
	// Reports
	Reports *Reports
	// Company info
	CompanyInfo *CompanyInfo
	// Payment methods
	PaymentMethods *PaymentMethods
	// Sales orders
	SalesOrders *SalesOrders
	// Tax rates
	TaxRates *TaxRates
	// Tracking categories
	TrackingCategories *TrackingCategories

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatAccounting)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatAccounting) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatAccounting) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatAccounting {
	sdk := &CodatAccounting{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "3.0.0",
			SDKVersion:        "0.31.0",
			GenVersion:        "2.214.3",
			UserAgent:         "speakeasy-sdk/go 0.31.0 2.214.3 3.0.0 github.com/codatio/client-sdk-go/previous-versions/accounting",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.AccountTransactions = newAccountTransactions(sdk.sdkConfiguration)

	sdk.BankAccounts = newBankAccounts(sdk.sdkConfiguration)

	sdk.BankAccountTransactions = newBankAccountTransactions(sdk.sdkConfiguration)

	sdk.Bills = newBills(sdk.sdkConfiguration)

	sdk.Customers = newCustomers(sdk.sdkConfiguration)

	sdk.DirectCosts = newDirectCosts(sdk.sdkConfiguration)

	sdk.DirectIncomes = newDirectIncomes(sdk.sdkConfiguration)

	sdk.Invoices = newInvoices(sdk.sdkConfiguration)

	sdk.ItemReceipts = newItemReceipts(sdk.sdkConfiguration)

	sdk.PurchaseOrders = newPurchaseOrders(sdk.sdkConfiguration)

	sdk.Suppliers = newSuppliers(sdk.sdkConfiguration)

	sdk.Transfers = newTransfers(sdk.sdkConfiguration)

	sdk.BillCreditNotes = newBillCreditNotes(sdk.sdkConfiguration)

	sdk.BillPayments = newBillPayments(sdk.sdkConfiguration)

	sdk.Accounts = newAccounts(sdk.sdkConfiguration)

	sdk.CreditNotes = newCreditNotes(sdk.sdkConfiguration)

	sdk.Items = newItems(sdk.sdkConfiguration)

	sdk.JournalEntries = newJournalEntries(sdk.sdkConfiguration)

	sdk.Journals = newJournals(sdk.sdkConfiguration)

	sdk.Payments = newPayments(sdk.sdkConfiguration)

	sdk.Reports = newReports(sdk.sdkConfiguration)

	sdk.CompanyInfo = newCompanyInfo(sdk.sdkConfiguration)

	sdk.PaymentMethods = newPaymentMethods(sdk.sdkConfiguration)

	sdk.SalesOrders = newSalesOrders(sdk.sdkConfiguration)

	sdk.TaxRates = newTaxRates(sdk.sdkConfiguration)

	sdk.TrackingCategories = newTrackingCategories(sdk.sdkConfiguration)

	return sdk
}
