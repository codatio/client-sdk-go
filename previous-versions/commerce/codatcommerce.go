// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package commerce

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/internal/hooks"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/retry"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/utils"
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

// CodatCommerce - Commerce API: Codat's standardized API for accessing commerce data
// > ### New to Codat?
// >
// > Our Commerce API reference is relevant only to our existing clients.
// > Please reach out to your Codat contact so that we can find the right product for you.
//
// Codat's Commerce API allows you to access standardised data from over 11 commerce and POS systems.
//
// Standardize how you connect to your customers’ payment, PoS, and eCommerce systems. Retrieve orders, payouts, payments, and product data in the same way for all the leading commerce software.
//
// <!-- Start Codat Tags Table -->
// ## Endpoints
//
// | Endpoints | Description |
// | :- |:- |
// | Customers | Retrieve standardized data from linked commerce software. |
// | Disputes | Retrieve standardized data from linked commerce software. |
// | Company info | Retrieve standardized data from linked commerce software. |
// | Locations | Retrieve standardized data from linked commerce software. |
// | Orders | Retrieve standardized data from linked commerce software. |
// | Payments | Retrieve standardized data from linked commerce software. |
// | Products | Retrieve standardized data from linked commerce software. |
// | Tax components | Retrieve standardized data from linked commerce software. |
// | Transactions | Retrieve standardized data from linked commerce software. |
// <!-- End Codat Tags Table -->
//
// [Read more...](https://docs.codat.io/commerce-api/overview)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
type CodatCommerce struct {
	// Retrieve standardized data from linked commerce software.
	Customers *Customers
	// Retrieve standardized data from linked commerce software.
	Disputes *Disputes
	// Retrieve standardized data from linked commerce software.
	CompanyInfo *CompanyInfo
	// Retrieve standardized data from linked commerce software.
	Locations *Locations
	// Retrieve standardized data from linked commerce software.
	Orders *Orders
	// Retrieve standardized data from linked commerce software.
	Payments *Payments
	// Retrieve standardized data from linked commerce software.
	Products *Products
	// Retrieve standardized data from linked commerce software.
	TaxComponents *TaxComponents
	// Retrieve standardized data from linked commerce software.
	Transactions *Transactions

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatCommerce)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatCommerce) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatCommerce) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatCommerce {
	sdk := &CodatCommerce{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "3.0.0",
			SDKVersion:        "3.0.0",
			GenVersion:        "2.463.0",
			UserAgent:         "speakeasy-sdk/go 3.0.0 2.463.0 3.0.0 github.com/codatio/client-sdk-go/previous-versions/commerce",
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

	sdk.Customers = newCustomers(sdk.sdkConfiguration)

	sdk.Disputes = newDisputes(sdk.sdkConfiguration)

	sdk.CompanyInfo = newCompanyInfo(sdk.sdkConfiguration)

	sdk.Locations = newLocations(sdk.sdkConfiguration)

	sdk.Orders = newOrders(sdk.sdkConfiguration)

	sdk.Payments = newPayments(sdk.sdkConfiguration)

	sdk.Products = newProducts(sdk.sdkConfiguration)

	sdk.TaxComponents = newTaxComponents(sdk.sdkConfiguration)

	sdk.Transactions = newTransactions(sdk.sdkConfiguration)

	return sdk
}
