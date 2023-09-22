// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package syncforexpensesversion1

import (
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
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
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// CodatSyncExpenses - Sync for Expenses (v1): The API for Sync for Expenses.
//
// Sync for Expenses is an API and a set of supporting tools. It has been built to
// enable corporate card and expense management platforms to provide high-quality
// integrations with multiple accounting platforms through a standardized API.
//
// [Read more...](https://docs.codat.io/sync-for-expenses/overview)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
//
// Not seeing what you expect? [See the main Sync for Commerce API](https://docs.codat.io/sync-for-commerce-api).
type CodatSyncExpenses struct {
	// Create and manage your Codat companies.
	Companies *companies
	// Companies sync configuration.
	Configuration *configuration
	// Create and manage partner expense connection.
	Connections *connections
	// Create expense datasets and upload receipts.
	Expenses *expenses
	// Mapping options for a companies expenses.
	MappingOptions *mappingOptions
	// Triggering a new sync of expenses to accounting software.
	Sync *sync
	// Check the status of ongoing or previous expense syncs.
	SyncStatus *syncStatus
	// Retrieve the status of transactions within a sync.
	TransactionStatus *transactionStatus

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatSyncExpenses)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		sdk.sdkConfiguration.Security = &security
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatSyncExpenses {
	sdk := &CodatSyncExpenses{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "prealpha",
			SDKVersion:        "0.24.1",
			GenVersion:        "2.125.1",
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

	sdk.Companies = newCompanies(sdk.sdkConfiguration)

	sdk.Configuration = newConfiguration(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.Expenses = newExpenses(sdk.sdkConfiguration)

	sdk.MappingOptions = newMappingOptions(sdk.sdkConfiguration)

	sdk.Sync = newSync(sdk.sdkConfiguration)

	sdk.SyncStatus = newSyncStatus(sdk.sdkConfiguration)

	sdk.TransactionStatus = newTransactionStatus(sdk.sdkConfiguration)

	return sdk
}
