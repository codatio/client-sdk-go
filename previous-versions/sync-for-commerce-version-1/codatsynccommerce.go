// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package syncforcommerceversion1

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
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

// CodatSyncCommerce - Sync for Commerce (v1): The API for Sync for Commerce V1.
//
// Sync for Commerce automatically replicates and reconciles sales data from a merchant’s source PoS, Payments, and eCommerce systems into their accounting software. This eliminates manual processing by merchants and transforms their ability to run and grow their business.
//
// [Read More...](https://docs.codat.io/commerce/overview)
//
// Not seeing what you expect? [See the main Sync for Commerce API](https://docs.codat.io/sync-for-commerce-api).
type CodatSyncCommerce struct {
	// Configure preferences for any given Sync for Commerce company using sync flow.
	SyncFlowPreferences *SyncFlowPreferences
	// Create and manage your Codat companies.
	Companies *Companies
	// Manage your companies' data connections.
	Connections *Connections
	// Bank accounts
	AccountingBankAccounts *AccountingBankAccounts
	// Retrieve standardized data from linked commerce platforms.
	CommerceCustomers *CommerceCustomers
	// Retrieve standardized data from linked commerce platforms.
	CommerceCompanyInfo *CommerceCompanyInfo
	// Retrieve standardized data from linked commerce platforms.
	CommerceLocations *CommerceLocations
	// Retrieve standardized data from linked commerce platforms.
	CommerceOrders *CommerceOrders
	// Retrieve standardized data from linked commerce platforms.
	CommercePayments *CommercePayments
	// Retrieve standardized data from linked commerce platforms.
	CommerceProducts *CommerceProducts
	// Retrieve standardized data from linked commerce platforms.
	CommerceTransactions *CommerceTransactions
	// Accounts
	AccountingAccounts *AccountingAccounts
	// Credit notes
	AccountingCreditNotes *AccountingCreditNotes
	// Customers
	AccountingCustomers *AccountingCustomers
	// Direct incomes
	AccountingDirectIncomes *AccountingDirectIncomes
	// Invoices
	AccountingInvoices *AccountingInvoices
	// Journal entries
	AccountingJournalEntries *AccountingJournalEntries
	// Payments
	AccountingPayments *AccountingPayments
	// Asynchronously retrieve data from an integration to refresh data in Codat.
	RefreshData *RefreshData
	// Company info
	AccountingCompanyInfo *AccountingCompanyInfo
	// View push options and get push statuses.
	PushData *PushData
	// Initiate a sync of Sync for Commerce company data into their respective accounting software.
	Sync *Sync
	// Expressively configure preferences for any given Sync for Commerce company.
	Configuration *Configuration
	// View useful information about codat's integrations.
	Integrations *Integrations
	// Create new and manage existing Sync for Commerce companies.
	CompanyManagement *CompanyManagement

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatSyncCommerce)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatSyncCommerce) {
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
	return func(sdk *CodatSyncCommerce) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatSyncCommerce {
	sdk := &CodatSyncCommerce{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.1",
			SDKVersion:        "0.22.0",
			GenVersion:        "2.214.3",
			UserAgent:         "speakeasy-sdk/go 0.22.0 2.214.3 1.1 github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1",
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

	sdk.SyncFlowPreferences = newSyncFlowPreferences(sdk.sdkConfiguration)

	sdk.Companies = newCompanies(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.AccountingBankAccounts = newAccountingBankAccounts(sdk.sdkConfiguration)

	sdk.CommerceCustomers = newCommerceCustomers(sdk.sdkConfiguration)

	sdk.CommerceCompanyInfo = newCommerceCompanyInfo(sdk.sdkConfiguration)

	sdk.CommerceLocations = newCommerceLocations(sdk.sdkConfiguration)

	sdk.CommerceOrders = newCommerceOrders(sdk.sdkConfiguration)

	sdk.CommercePayments = newCommercePayments(sdk.sdkConfiguration)

	sdk.CommerceProducts = newCommerceProducts(sdk.sdkConfiguration)

	sdk.CommerceTransactions = newCommerceTransactions(sdk.sdkConfiguration)

	sdk.AccountingAccounts = newAccountingAccounts(sdk.sdkConfiguration)

	sdk.AccountingCreditNotes = newAccountingCreditNotes(sdk.sdkConfiguration)

	sdk.AccountingCustomers = newAccountingCustomers(sdk.sdkConfiguration)

	sdk.AccountingDirectIncomes = newAccountingDirectIncomes(sdk.sdkConfiguration)

	sdk.AccountingInvoices = newAccountingInvoices(sdk.sdkConfiguration)

	sdk.AccountingJournalEntries = newAccountingJournalEntries(sdk.sdkConfiguration)

	sdk.AccountingPayments = newAccountingPayments(sdk.sdkConfiguration)

	sdk.RefreshData = newRefreshData(sdk.sdkConfiguration)

	sdk.AccountingCompanyInfo = newAccountingCompanyInfo(sdk.sdkConfiguration)

	sdk.PushData = newPushData(sdk.sdkConfiguration)

	sdk.Sync = newSync(sdk.sdkConfiguration)

	sdk.Configuration = newConfiguration(sdk.sdkConfiguration)

	sdk.Integrations = newIntegrations(sdk.sdkConfiguration)

	sdk.CompanyManagement = newCompanyManagement(sdk.sdkConfiguration)

	return sdk
}
