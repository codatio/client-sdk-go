// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatcommerce

import (
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

// CodatCommerce - Codat's standardized API for accessing commerce data
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

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*CodatCommerce)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk._serverURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatCommerce) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk._defaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatCommerce) {
		sdk._security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatCommerce {
	sdk := &CodatCommerce{
		_language:   "go",
		_sdkVersion: "0.16.1",
		_genVersion: "2.32.7",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {
		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}
	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.CompanyInfo = newCompanyInfo(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Customers = newCustomers(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Disputes = newDisputes(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Locations = newLocations(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Orders = newOrders(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Payments = newPayments(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Products = newProducts(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.TaxComponents = newTaxComponents(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Transactions = newTransactions(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
