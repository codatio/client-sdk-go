// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package common

import (
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/utils"
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

// CodatCommon - Platform API: Platform API
// An API for the common components of all of Codat's products.
//
// These end points cover creating and managing your companies, data connections, and integrations.
//
// [Read about the building blocks of Codat...](https://docs.codat.io/core-concepts/companies)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
type CodatCommon struct {
	// Create and manage your Codat companies.
	Companies *companies
	// Manage your companies' data connections.
	Connections *connections
	// View and manage your available integrations in Codat.
	Integrations *integrations
	// View push options and get push statuses.
	PushData *pushData
	// Asynchronously retrieve data from an integration to refresh data in Codat.
	RefreshData *refreshData
	// Manage your Codat instance.
	Settings *settings
	// View and configure supplemental data for supported data types.
	SupplementalData *supplementalData
	// Manage webhooks, rules, and events.
	Webhooks *webhooks

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatCommon)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatCommon) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatCommon) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatCommon) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatCommon) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatCommon) {
		sdk.sdkConfiguration.Security = &security
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *CodatCommon) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatCommon {
	sdk := &CodatCommon{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "3.0.0",
			SDKVersion:        "0.27.0",
			GenVersion:        "2.129.1",
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

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.Integrations = newIntegrations(sdk.sdkConfiguration)

	sdk.PushData = newPushData(sdk.sdkConfiguration)

	sdk.RefreshData = newRefreshData(sdk.sdkConfiguration)

	sdk.Settings = newSettings(sdk.sdkConfiguration)

	sdk.SupplementalData = newSupplementalData(sdk.sdkConfiguration)

	sdk.Webhooks = newWebhooks(sdk.sdkConfiguration)

	return sdk
}
