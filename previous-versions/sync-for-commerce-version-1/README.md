# Sync for Commerce version 1

<!-- Start Codat Library Description -->
﻿Embedded accounting integrations for POS and eCommerce platforms.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
### Example

```go
package main

import (
	"context"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerceversion1.New(
		syncforcommerceversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowPreferences.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [SyncFlowPreferences](docs/sdks/syncflowpreferences/README.md)

* [GetConfigTextSyncFlow](docs/sdks/syncflowpreferences/README.md#getconfigtextsyncflow) - Retrieve preferences for text fields on sync flow
* [GetSyncFlowURL](docs/sdks/syncflowpreferences/README.md#getsyncflowurl) - Retrieve sync flow url
* [GetVisibleAccounts](docs/sdks/syncflowpreferences/README.md#getvisibleaccounts) - List visible accounts
* [UpdateConfigTextSyncFlow](docs/sdks/syncflowpreferences/README.md#updateconfigtextsyncflow) - Update preferences for text fields on sync flow
* [UpdateVisibleAccountsSyncFlow](docs/sdks/syncflowpreferences/README.md#updatevisibleaccountssyncflow) - Update the visible accounts on sync flow

### [Companies](docs/sdks/companies/README.md)

* [DeleteCompany](docs/sdks/companies/README.md#deletecompany) - Delete a company
* [GetCompany](docs/sdks/companies/README.md#getcompany) - Get company
* [UpdateCompany](docs/sdks/companies/README.md#updatecompany) - Update company

### [Connections](docs/sdks/connections/README.md)

* [DeleteConnection](docs/sdks/connections/README.md#deleteconnection) - Delete connection
* [GetConnection](docs/sdks/connections/README.md#getconnection) - Get connection
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [AccountingBankAccounts](docs/sdks/accountingbankaccounts/README.md)

* [GetAccountingBankAccount](docs/sdks/accountingbankaccounts/README.md#getaccountingbankaccount) - Get bank account
* [ListAccountingBankAccounts](docs/sdks/accountingbankaccounts/README.md#listaccountingbankaccounts) - List bank accounts

### [CommerceCustomers](docs/sdks/commercecustomers/README.md)

* [GetCommerceCustomer](docs/sdks/commercecustomers/README.md#getcommercecustomer) - Get customer
* [ListCommerceCustomers](docs/sdks/commercecustomers/README.md#listcommercecustomers) - List customers

### [CommerceCompanyInfo](docs/sdks/commercecompanyinfo/README.md)

* [GetCommerceCompanyInfo](docs/sdks/commercecompanyinfo/README.md#getcommercecompanyinfo) - Get company info

### [CommerceLocations](docs/sdks/commercelocations/README.md)

* [GetCommerceLocation](docs/sdks/commercelocations/README.md#getcommercelocation) - Get location
* [ListCommerceLocations](docs/sdks/commercelocations/README.md#listcommercelocations) - List locations

### [CommerceOrders](docs/sdks/commerceorders/README.md)

* [GetCommerceOrder](docs/sdks/commerceorders/README.md#getcommerceorder) - Get order
* [ListCommerceOrders](docs/sdks/commerceorders/README.md#listcommerceorders) - List orders

### [CommercePayments](docs/sdks/commercepayments/README.md)

* [GetCommercePayment](docs/sdks/commercepayments/README.md#getcommercepayment) - Get payment
* [GetMethod](docs/sdks/commercepayments/README.md#getmethod) - Get payment method
* [ListCommercePayments](docs/sdks/commercepayments/README.md#listcommercepayments) - List payments
* [ListMethods](docs/sdks/commercepayments/README.md#listmethods) - List payment methods

### [CommerceProducts](docs/sdks/commerceproducts/README.md)

* [GetCommerceProduct](docs/sdks/commerceproducts/README.md#getcommerceproduct) - Get product
* [ListCommerceProducts](docs/sdks/commerceproducts/README.md#listcommerceproducts) - List products

### [CommerceTransactions](docs/sdks/commercetransactions/README.md)

* [GetCommerceTransaction](docs/sdks/commercetransactions/README.md#getcommercetransaction) - Get transaction
* [ListCommerceTransactions](docs/sdks/commercetransactions/README.md#listcommercetransactions) - List transactions

### [AccountingAccounts](docs/sdks/accountingaccounts/README.md)

* [CreateAccountingAccount](docs/sdks/accountingaccounts/README.md#createaccountingaccount) - Create account
* [GetAccountingAccount](docs/sdks/accountingaccounts/README.md#getaccountingaccount) - Get account
* [ListAccountingAccounts](docs/sdks/accountingaccounts/README.md#listaccountingaccounts) - List accounts

### [AccountingCreditNotes](docs/sdks/accountingcreditnotes/README.md)

* [CreateAccountingCreditNote](docs/sdks/accountingcreditnotes/README.md#createaccountingcreditnote) - Create credit note

### [AccountingCustomers](docs/sdks/accountingcustomers/README.md)

* [CreateAccountingCustomer](docs/sdks/accountingcustomers/README.md#createaccountingcustomer) - Create customer

### [AccountingDirectIncomes](docs/sdks/accountingdirectincomes/README.md)

* [CreateAccountingDirectIncome](docs/sdks/accountingdirectincomes/README.md#createaccountingdirectincome) - Create direct income

### [AccountingInvoices](docs/sdks/accountinginvoices/README.md)

* [CreateAccountingInvoice](docs/sdks/accountinginvoices/README.md#createaccountinginvoice) - Create invoice

### [AccountingJournalEntries](docs/sdks/accountingjournalentries/README.md)

* [CreateAccountingJournalEntry](docs/sdks/accountingjournalentries/README.md#createaccountingjournalentry) - Create journal entry

### [AccountingPayments](docs/sdks/accountingpayments/README.md)

* [CreateAccountingPayment](docs/sdks/accountingpayments/README.md#createaccountingpayment) - Create payment

### [RefreshData](docs/sdks/refreshdata/README.md)

* [All](docs/sdks/refreshdata/README.md#all) - Refresh all data
* [ByDataType](docs/sdks/refreshdata/README.md#bydatatype) - Refresh data type
* [GetCompanyDataStatus](docs/sdks/refreshdata/README.md#getcompanydatastatus) - Get data status
* [GetPullOperation](docs/sdks/refreshdata/README.md#getpulloperation) - Get pull operation
* [ListPullOperations](docs/sdks/refreshdata/README.md#listpulloperations) - List pull operations

### [AccountingCompanyInfo](docs/sdks/accountingcompanyinfo/README.md)

* [GetAccountingCompanyInfo](docs/sdks/accountingcompanyinfo/README.md#getaccountingcompanyinfo) - Get company info
* [Refresh](docs/sdks/accountingcompanyinfo/README.md#refresh) - Refresh company info

### [PushData](docs/sdks/pushdata/README.md)

* [GetOperation](docs/sdks/pushdata/README.md#getoperation) - Get push operation
* [ListOperations](docs/sdks/pushdata/README.md#listoperations) - List push operations

### [Sync](docs/sdks/sync/README.md)

* [GetSyncStatus](docs/sdks/sync/README.md#getsyncstatus) - Get status for a company's syncs
* [RequestSync](docs/sdks/sync/README.md#requestsync) - Sync new
* [RequestSyncForDateRange](docs/sdks/sync/README.md#requestsyncfordaterange) - Sync range

### [Configuration](docs/sdks/configuration/README.md)

* [GetConfiguration](docs/sdks/configuration/README.md#getconfiguration) - Retrieve config preferences set for a company
* [SetConfiguration](docs/sdks/configuration/README.md#setconfiguration) - Create or update configuration

### [Integrations](docs/sdks/integrations/README.md)

* [GetIntegrationBranding](docs/sdks/integrations/README.md#getintegrationbranding) - Get branding for an integration
* [ListIntegrations](docs/sdks/integrations/README.md#listintegrations) - List integrations

### [CompanyManagement](docs/sdks/companymanagement/README.md)

* [CreateCompany](docs/sdks/companymanagement/README.md#createcompany) - Create sync for commerce company
* [CreateConnection](docs/sdks/companymanagement/README.md#createconnection) - Create connection
* [ListCompanies](docs/sdks/companymanagement/README.md#listcompanies) - List companies
* [ListConnections](docs/sdks/companymanagement/README.md#listconnections) - List data connections
* [UpdateConnection](docs/sdks/companymanagement/README.md#updateconnection) - Update data connection
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 401,402,403,429,500,503 | application/json        |
| sdkerrors.SDKError      | 400-600                 | */*                     |

### Example

```go
package main

import (
	"context"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerceversion1.New(
		syncforcommerceversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowPreferences.GetConfigTextSyncFlow(ctx)
	if err != nil {

		var e *sdkerrors.ErrorMessage
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.codat.io` | None |

#### Example

```go
package main

import (
	"context"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerceversion1.New(
		syncforcommerceversion1.WithServerIndex(0),
		syncforcommerceversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowPreferences.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerceversion1.New(
		syncforcommerceversion1.WithServerURL("https://api.codat.io"),
		syncforcommerceversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowPreferences.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Retries -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```go
package main

import (
	"context"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"log"
	"pkg/models/operations"
	"pkg/utils"
)

func main() {
	s := syncforcommerceversion1.New(
		syncforcommerceversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowPreferences.GetConfigTextSyncFlow(ctx, operations.WithRetries(utils.RetryConfig{
		Strategy: "backoff",
		Backoff: &utils.BackoffStrategy{
			InitialInterval: 1,
			MaxInterval:     50,
			Exponent:        1.1,
			MaxElapsedTime:  100,
		},
		RetryConnectionErrors: false,
	}))
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```go
package main

import (
	"context"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"log"
	"pkg/models/operations"
	"pkg/utils"
)

func main() {
	s := syncforcommerceversion1.New(
		syncforcommerceversion1.WithRetryConfig(utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}),
		syncforcommerceversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowPreferences.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```


<!-- End Retries -->



<!-- Start Authentication -->

## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `AuthHeader` | apiKey       | API key      |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerceversion1.New(
		syncforcommerceversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowPreferences.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)