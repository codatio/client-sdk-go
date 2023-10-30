// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package syncforpayroll

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/utils"
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

// CodatSyncPayroll - Sync for Payroll: The API for Sync for Payroll.
//
// Sync for Payroll is an API and a set of supporting tools built to help integrate your customers' payroll data from their HR and payroll platforms into their accounting platforms and to support its reconciliation.
//
// [Explore product](https://docs.codat.io/payroll/overview) | [See OpenAPI spec](https://github.com/codatio/oas)
//
// ---
//
// ## Endpoints
//
// | Endpoints            | Description                                                                                                |
// |:---------------------|:-----------------------------------------------------------------------------------------------------------|
// | Companies            | Create and manage your SMB users' companies.                                                               |
// | Connections          | Create new and manage existing data connections for a company.                                             |
// | Accounts             | Get, create, and update Accounts.                                                           |
// | Journal entries      | Get, create, and update Journal entries.                                                           |
// | Journals             | Get, create, and update Journals.                                                           |
// | Tracking categories  | Get, create, and update Tracking Categories for additional categorization of payroll components.                                                           |
// | Company info         | View company profile from the source platform.                                                             |
// | Manage data          | Control how data is retrieved from an integration.                                                         |
type CodatSyncPayroll struct {
	// Accounts
	Accounts *accounts
	// Create and manage your Codat companies.
	Companies *companies
	// View company information fetched from the source platform.
	CompanyInfo *companyInfo
	// Manage your companies' data connections.
	Connections *connections
	// Journal entries
	JournalEntries *journalEntries
	// Journals
	Journals *journals
	// Asynchronously retrieve data from an integration to refresh data in Codat.
	ManageData *manageData
	// Tracking categories
	TrackingCategories *trackingCategories

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatSyncPayroll)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatSyncPayroll) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatSyncPayroll) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatSyncPayroll) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatSyncPayroll) {
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
	return func(sdk *CodatSyncPayroll) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *CodatSyncPayroll) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *CodatSyncPayroll) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatSyncPayroll {
	sdk := &CodatSyncPayroll{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "3.0.0",
			SDKVersion:        "1.3.1",
			GenVersion:        "2.173.0",
			UserAgent:         "speakeasy-sdk/go 1.3.1 2.173.0 3.0.0 github.com/codatio/client-sdk-go/sync-for-payroll",
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

	sdk.Accounts = newAccounts(sdk.sdkConfiguration)

	sdk.Companies = newCompanies(sdk.sdkConfiguration)

	sdk.CompanyInfo = newCompanyInfo(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.JournalEntries = newJournalEntries(sdk.sdkConfiguration)

	sdk.Journals = newJournals(sdk.sdkConfiguration)

	sdk.ManageData = newManageData(sdk.sdkConfiguration)

	sdk.TrackingCategories = newTrackingCategories(sdk.sdkConfiguration)

	return sdk
}
