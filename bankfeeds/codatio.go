package codatio

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/utils"
	"net/http"
	"time"
)

var ServerList = []string{
	"https://api.codat.io",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// SDK Documentation: Bank Feeds API enables your SMB users to set up bank feeds from accounts in your application to supported accounting platforms.
//
// A bank feed is a connection between a source bank account—in your application—and a target bank account in a supported accounting package.
//
// [Read more...](https://docs.codat.io/bank-feeds-api/overview)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
type Codatio struct {
	BankAccountTransactions *bankAccountTransactions
	BankFeedAccounts        *bankFeedAccounts

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*Codatio)

func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Codatio) {
		sdk._serverURL = serverURL
	}
}

func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Codatio) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Codatio) {
		sdk._defaultClient = client
	}
}

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Codatio) {
		sdk._security = &security
	}
}

func New(opts ...SDKOption) *Codatio {
	sdk := &Codatio{
		_language:   "go",
		_sdkVersion: "0.6.2",
		_genVersion: "1.12.0",
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

	sdk.BankAccountTransactions = newBankAccountTransactions(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.BankFeedAccounts = newBankFeedAccounts(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
