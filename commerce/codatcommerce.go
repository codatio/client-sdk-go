// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatcommerce

import (
	"fmt"
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/commerce/pkg/utils"
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
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// CodatCommerce - Commerce API: Codat's standardized API for accessing commerce data
// Codat's Commerce API allows you to access standardised data from over 11 commerce and POS systems.
//
// Standardize how you connect to your customers’ payment, PoS, and eCommerce systems. Retrieve orders, payouts, payments, and product data in the same way for all the leading commerce platforms.
//
// [Read more...](https://docs.codat.io/commerce-api/overview)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
type CodatCommerce struct {
	// CompanyInfo - Retrieve standardized data from linked commerce platforms.
	CompanyInfo *companyInfo
	// Customers - Retrieve standardized data from linked commerce platforms.
	Customers *customers
	// Disputes - Retrieve standardized data from linked commerce platforms.
	Disputes *disputes
	// Locations - Retrieve standardized data from linked commerce platforms.
	Locations *locations
	// Orders - Retrieve standardized data from linked commerce platforms.
	Orders *orders
	// Payments - Retrieve standardized data from linked commerce platforms.
	Payments *payments
	// Products - Retrieve standardized data from linked commerce platforms.
	Products *products
	// TaxComponents - Retrieve standardized data from linked commerce platforms.
	TaxComponents *taxComponents
	// Transactions - Retrieve standardized data from linked commerce platforms.
	Transactions *transactions

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
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatCommerce {
	sdk := &CodatCommerce{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "3.0.0",
			SDKVersion:        "1.0.0",
			GenVersion:        "2.81.1",
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

	sdk.CompanyInfo = newCompanyInfo(sdk.sdkConfiguration)

	sdk.Customers = newCustomers(sdk.sdkConfiguration)

	sdk.Disputes = newDisputes(sdk.sdkConfiguration)

	sdk.Locations = newLocations(sdk.sdkConfiguration)

	sdk.Orders = newOrders(sdk.sdkConfiguration)

	sdk.Payments = newPayments(sdk.sdkConfiguration)

	sdk.Products = newProducts(sdk.sdkConfiguration)

	sdk.TaxComponents = newTaxComponents(sdk.sdkConfiguration)

	sdk.Transactions = newTransactions(sdk.sdkConfiguration)

	return sdk
}
