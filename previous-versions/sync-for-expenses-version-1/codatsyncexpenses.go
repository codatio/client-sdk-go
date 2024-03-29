// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package syncforexpensesversion1

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/internal/hooks"
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
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
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
// Not seeing what you expect? [See the main Sync for Expenses API](https://docs.codat.io/sync-for-expenses-api).
type CodatSyncExpenses struct {
	// Create and manage your Codat companies.
	Companies *Companies
	// Create and manage partner expense connection.
	Connections *Connections
	// Companies sync configuration.
	Configuration *Configuration
	// Create expense datasets and upload receipts.
	Expenses *Expenses
	// Mapping options for a companies expenses.
	MappingOptions *MappingOptions
	// Triggering a new sync of expenses to accounting software.
	Sync *Sync
	// Check the status of ongoing or previous expense syncs.
	SyncStatus *SyncStatus
	// Retrieve the status of transactions within a sync.
	TransactionStatus *TransactionStatus

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
		sdk.sdkConfiguration.Client = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
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
			SDKVersion:        "0.27.0",
			GenVersion:        "2.286.2",
			UserAgent:         "speakeasy-sdk/go 0.27.0 2.286.2 prealpha github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1",
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

	sdk.Configuration = newConfiguration(sdk.sdkConfiguration)

	sdk.Expenses = newExpenses(sdk.sdkConfiguration)

	sdk.MappingOptions = newMappingOptions(sdk.sdkConfiguration)

	sdk.Sync = newSync(sdk.sdkConfiguration)

	sdk.SyncStatus = newSyncStatus(sdk.sdkConfiguration)

	sdk.TransactionStatus = newTransactionStatus(sdk.sdkConfiguration)

	return sdk
}
