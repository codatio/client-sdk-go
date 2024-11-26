// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package syncforpayables

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/internal/hooks"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/retry"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/utils"
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

// CodatSyncPayables - Bill pay kit: The API reference for the Bill Pay kit.
//
// The bill pay kit is an API and a set of supporting tools designed to integrate a bill pay flow into your app as quickly as possible. It's ideal for facilitating essential bill payment processes within your SMB's accounting software.
//
// [Explore product](https://docs.codat.io/payables/bill-pay-kit) | [See OpenAPI spec](https://github.com/codatio/oas)
//
// ---
// ## Supported Integrations
//
// | Integration                   | Supported |
// |-------------------------------|-----------|
// | FreeAgent                     | Yes       |
// | QuickBooks Online             | Yes       |
// | Oracle NetSuite               | Yes       |
// | Xero                          | Yes       |
//
// ---
// <!-- Start Codat Tags Table -->
// ## Endpoints
//
// | Endpoints | Description |
// | :- |:- |
// | Companies | Create and manage your SMB users' companies. |
// | Connections | Create new and manage existing data connections for a company. |
// | Company information | View company profile from the source platform. |
// | Bills | Get, create, and update Bills. |
// | Bill payments | Get, create, and update Bill payments. |
// | Suppliers | Get, create, and update Suppliers. |
// | Bank accounts | Create a bank account for a given company's connection. |
// <!-- End Codat Tags Table -->
type CodatSyncPayables struct {
	// Create and manage your SMB users' companies.
	Companies *Companies
	// Create new and manage existing data connections for a company.
	Connections *Connections
	// View company profile from the source platform.
	CompanyInformation *CompanyInformation
	// Get, create, and update Bills.
	Bills *Bills
	// Get, create, and update Bill payments.
	BillPayments *BillPayments
	// Get, create, and update Suppliers.
	Suppliers *Suppliers
	// Create a bank account for a given company's connection.
	BankAccounts *BankAccounts

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatSyncPayables)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatSyncPayables) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatSyncPayables) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatSyncPayables) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatSyncPayables) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatSyncPayables) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *CodatSyncPayables) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *CodatSyncPayables) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *CodatSyncPayables) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatSyncPayables {
	sdk := &CodatSyncPayables{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "3.0.0",
			SDKVersion:        "5.0.0",
			GenVersion:        "2.463.0",
			UserAgent:         "speakeasy-sdk/go 5.0.0 2.463.0 3.0.0 github.com/codatio/client-sdk-go/sync-for-payables",
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

	sdk.CompanyInformation = newCompanyInformation(sdk.sdkConfiguration)

	sdk.Bills = newBills(sdk.sdkConfiguration)

	sdk.BillPayments = newBillPayments(sdk.sdkConfiguration)

	sdk.Suppliers = newSuppliers(sdk.sdkConfiguration)

	sdk.BankAccounts = newBankAccounts(sdk.sdkConfiguration)

	return sdk
}
