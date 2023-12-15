# Accounting

<!-- Start Codat Library Description -->
﻿Codat's Accounting API is a flexible API for pulling and pushing up-to-date accounting data to your customer's accounting software.
It gives you a simple way to view, create, update adn delete data without having to worry about each platform's specific complexities.

<!-- End Codat Library Description -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/accounting
```
<!-- End SDK Installation [installation] -->

## Example Usage
<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"log"
)

func main() {
	s := accounting.New(
		accounting.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountTransactions.Get(ctx, operations.GetAccountTransactionRequest{
		AccountTransactionID: "string",
		CompanyID:            "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID:         "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountTransaction != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [AccountTransactions](docs/sdks/accounttransactions/README.md)

* [Get](docs/sdks/accounttransactions/README.md#get) - Get account transaction
* [List](docs/sdks/accounttransactions/README.md#list) - List account transactions

### [BankAccounts](docs/sdks/bankaccounts/README.md)

* [Create](docs/sdks/bankaccounts/README.md#create) - Create bank account
* [Get](docs/sdks/bankaccounts/README.md#get) - Get bank account
* [GetCreateUpdateModel](docs/sdks/bankaccounts/README.md#getcreateupdatemodel) - Get create/update bank account model
* [List](docs/sdks/bankaccounts/README.md#list) - List bank accounts
* [Update](docs/sdks/bankaccounts/README.md#update) - Update bank account

### [BankAccountTransactions](docs/sdks/bankaccounttransactions/README.md)

* [Create](docs/sdks/bankaccounttransactions/README.md#create) - Create bank account transactions
* [GetCreateModel](docs/sdks/bankaccounttransactions/README.md#getcreatemodel) - Get create bank account transactions model
* [List](docs/sdks/bankaccounttransactions/README.md#list) - List bank account transactions

### [Bills](docs/sdks/bills/README.md)

* [Create](docs/sdks/bills/README.md#create) - Create bill
* [Delete](docs/sdks/bills/README.md#delete) - Delete bill
* [DownloadAttachment](docs/sdks/bills/README.md#downloadattachment) - Download bill attachment
* [Get](docs/sdks/bills/README.md#get) - Get bill
* [GetAttachment](docs/sdks/bills/README.md#getattachment) - Get bill attachment
* [GetCreateUpdateModel](docs/sdks/bills/README.md#getcreateupdatemodel) - Get create/update bill model
* [List](docs/sdks/bills/README.md#list) - List bills
* [ListAttachments](docs/sdks/bills/README.md#listattachments) - List bill attachments
* [Update](docs/sdks/bills/README.md#update) - Update bill
* [UploadAttachment](docs/sdks/bills/README.md#uploadattachment) - Upload bill attachment

### [Customers](docs/sdks/customers/README.md)

* [Create](docs/sdks/customers/README.md#create) - Create customer
* [DownloadAttachment](docs/sdks/customers/README.md#downloadattachment) - Download customer attachment
* [Get](docs/sdks/customers/README.md#get) - Get customer
* [GetAttachment](docs/sdks/customers/README.md#getattachment) - Get customer attachment
* [GetCreateUpdateModel](docs/sdks/customers/README.md#getcreateupdatemodel) - Get create/update customer model
* [List](docs/sdks/customers/README.md#list) - List customers
* [ListAttachments](docs/sdks/customers/README.md#listattachments) - List customer attachments
* [Update](docs/sdks/customers/README.md#update) - Update customer

### [DirectCosts](docs/sdks/directcosts/README.md)

* [Create](docs/sdks/directcosts/README.md#create) - Create direct cost
* [DownloadAttachment](docs/sdks/directcosts/README.md#downloadattachment) - Download direct cost attachment
* [Get](docs/sdks/directcosts/README.md#get) - Get direct cost
* [GetAttachment](docs/sdks/directcosts/README.md#getattachment) - Get direct cost attachment
* [GetCreateModel](docs/sdks/directcosts/README.md#getcreatemodel) - Get create direct cost model
* [List](docs/sdks/directcosts/README.md#list) - List direct costs
* [ListAttachments](docs/sdks/directcosts/README.md#listattachments) - List direct cost attachments
* [UploadAttachment](docs/sdks/directcosts/README.md#uploadattachment) - Upload direct cost attachment

### [DirectIncomes](docs/sdks/directincomes/README.md)

* [Create](docs/sdks/directincomes/README.md#create) - Create direct income
* [DownloadAttachment](docs/sdks/directincomes/README.md#downloadattachment) - Download direct income attachment
* [Get](docs/sdks/directincomes/README.md#get) - Get direct income
* [GetAttachment](docs/sdks/directincomes/README.md#getattachment) - Get direct income attachment
* [GetCreateModel](docs/sdks/directincomes/README.md#getcreatemodel) - Get create direct income model
* [List](docs/sdks/directincomes/README.md#list) - List direct incomes
* [ListAttachments](docs/sdks/directincomes/README.md#listattachments) - List direct income attachments
* [UploadAttachment](docs/sdks/directincomes/README.md#uploadattachment) - Create direct income attachment

### [Invoices](docs/sdks/invoices/README.md)

* [Create](docs/sdks/invoices/README.md#create) - Create invoice
* [Delete](docs/sdks/invoices/README.md#delete) - Delete invoice
* [DownloadAttachment](docs/sdks/invoices/README.md#downloadattachment) - Download invoice attachment
* [DownloadPdf](docs/sdks/invoices/README.md#downloadpdf) - Get invoice as PDF
* [Get](docs/sdks/invoices/README.md#get) - Get invoice
* [GetAttachment](docs/sdks/invoices/README.md#getattachment) - Get invoice attachment
* [GetCreateUpdateModel](docs/sdks/invoices/README.md#getcreateupdatemodel) - Get create/update invoice model
* [List](docs/sdks/invoices/README.md#list) - List invoices
* [ListAttachments](docs/sdks/invoices/README.md#listattachments) - List invoice attachments
* [Update](docs/sdks/invoices/README.md#update) - Update invoice
* [UploadAttachment](docs/sdks/invoices/README.md#uploadattachment) - Upload invoice attachment

### [ItemReceipts](docs/sdks/itemreceipts/README.md)

* [Get](docs/sdks/itemreceipts/README.md#get) - Get item receipt
* [List](docs/sdks/itemreceipts/README.md#list) - List item receipts

### [PurchaseOrders](docs/sdks/purchaseorders/README.md)

* [Create](docs/sdks/purchaseorders/README.md#create) - Create purchase order
* [DownloadAttachment](docs/sdks/purchaseorders/README.md#downloadattachment) - Download purchase order attachment
* [DownloadPurchaseOrderPdf](docs/sdks/purchaseorders/README.md#downloadpurchaseorderpdf) - Download purchase order as PDF
* [Get](docs/sdks/purchaseorders/README.md#get) - Get purchase order
* [GetAttachment](docs/sdks/purchaseorders/README.md#getattachment) - Get purchase order attachment
* [GetCreateUpdateModel](docs/sdks/purchaseorders/README.md#getcreateupdatemodel) - Get create/update purchase order model
* [List](docs/sdks/purchaseorders/README.md#list) - List purchase orders
* [ListAttachments](docs/sdks/purchaseorders/README.md#listattachments) - List purchase order attachments
* [Update](docs/sdks/purchaseorders/README.md#update) - Update purchase order

### [Suppliers](docs/sdks/suppliers/README.md)

* [Create](docs/sdks/suppliers/README.md#create) - Create supplier
* [DownloadAttachment](docs/sdks/suppliers/README.md#downloadattachment) - Download supplier attachment
* [Get](docs/sdks/suppliers/README.md#get) - Get supplier
* [GetAttachment](docs/sdks/suppliers/README.md#getattachment) - Get supplier attachment
* [GetCreateUpdateModel](docs/sdks/suppliers/README.md#getcreateupdatemodel) - Get create/update supplier model
* [List](docs/sdks/suppliers/README.md#list) - List suppliers
* [ListAttachments](docs/sdks/suppliers/README.md#listattachments) - List supplier attachments
* [Update](docs/sdks/suppliers/README.md#update) - Update supplier

### [Transfers](docs/sdks/transfers/README.md)

* [Create](docs/sdks/transfers/README.md#create) - Create transfer
* [Get](docs/sdks/transfers/README.md#get) - Get transfer
* [GetCreateModel](docs/sdks/transfers/README.md#getcreatemodel) - Get create transfer model
* [List](docs/sdks/transfers/README.md#list) - List transfers
* [UploadAttachment](docs/sdks/transfers/README.md#uploadattachment) - Upload invoice attachment

### [BillCreditNotes](docs/sdks/billcreditnotes/README.md)

* [Create](docs/sdks/billcreditnotes/README.md#create) - Create bill credit note
* [Get](docs/sdks/billcreditnotes/README.md#get) - Get bill credit note
* [GetCreateUpdateModel](docs/sdks/billcreditnotes/README.md#getcreateupdatemodel) - Get create/update bill credit note model
* [List](docs/sdks/billcreditnotes/README.md#list) - List bill credit notes
* [Update](docs/sdks/billcreditnotes/README.md#update) - Update bill credit note
* [UploadAttachment](docs/sdks/billcreditnotes/README.md#uploadattachment) - Upload bill credit note attachment

### [BillPayments](docs/sdks/billpayments/README.md)

* [Create](docs/sdks/billpayments/README.md#create) - Create bill payments
* [Delete](docs/sdks/billpayments/README.md#delete) - Delete bill payment
* [Get](docs/sdks/billpayments/README.md#get) - Get bill payment
* [GetCreateModel](docs/sdks/billpayments/README.md#getcreatemodel) - Get create bill payment model
* [List](docs/sdks/billpayments/README.md#list) - List bill payments

### [Accounts](docs/sdks/accounts/README.md)

* [Create](docs/sdks/accounts/README.md#create) - Create account
* [Get](docs/sdks/accounts/README.md#get) - Get account
* [GetCreateModel](docs/sdks/accounts/README.md#getcreatemodel) - Get create account model
* [List](docs/sdks/accounts/README.md#list) - List accounts

### [CreditNotes](docs/sdks/creditnotes/README.md)

* [Create](docs/sdks/creditnotes/README.md#create) - Create credit note
* [Get](docs/sdks/creditnotes/README.md#get) - Get credit note
* [GetCreateUpdateModel](docs/sdks/creditnotes/README.md#getcreateupdatemodel) - Get create/update credit note model
* [List](docs/sdks/creditnotes/README.md#list) - List credit notes
* [Update](docs/sdks/creditnotes/README.md#update) - Update credit note

### [Items](docs/sdks/items/README.md)

* [Create](docs/sdks/items/README.md#create) - Create item
* [Get](docs/sdks/items/README.md#get) - Get item
* [GetCreateModel](docs/sdks/items/README.md#getcreatemodel) - Get create item model
* [List](docs/sdks/items/README.md#list) - List items

### [JournalEntries](docs/sdks/journalentries/README.md)

* [Create](docs/sdks/journalentries/README.md#create) - Create journal entry
* [Delete](docs/sdks/journalentries/README.md#delete) - Delete journal entry
* [Get](docs/sdks/journalentries/README.md#get) - Get journal entry
* [GetCreateModel](docs/sdks/journalentries/README.md#getcreatemodel) - Get create journal entry model
* [List](docs/sdks/journalentries/README.md#list) - List journal entries

### [Journals](docs/sdks/journals/README.md)

* [Create](docs/sdks/journals/README.md#create) - Create journal
* [Get](docs/sdks/journals/README.md#get) - Get journal
* [GetCreateModel](docs/sdks/journals/README.md#getcreatemodel) - Get create journal model
* [List](docs/sdks/journals/README.md#list) - List journals

### [Payments](docs/sdks/payments/README.md)

* [Create](docs/sdks/payments/README.md#create) - Create payment
* [Get](docs/sdks/payments/README.md#get) - Get payment
* [GetCreateModel](docs/sdks/payments/README.md#getcreatemodel) - Get create payment model
* [List](docs/sdks/payments/README.md#list) - List payments

### [Reports](docs/sdks/reports/README.md)

* [GetAgedCreditorsReport](docs/sdks/reports/README.md#getagedcreditorsreport) - Aged creditors report
* [GetAgedDebtorsReport](docs/sdks/reports/README.md#getageddebtorsreport) - Aged debtors report
* [GetBalanceSheet](docs/sdks/reports/README.md#getbalancesheet) - Get balance sheet
* [GetCashFlowStatement](docs/sdks/reports/README.md#getcashflowstatement) - Get cash flow statement
* [GetProfitAndLoss](docs/sdks/reports/README.md#getprofitandloss) - Get profit and loss
* [IsAgedCreditorsReportAvailable](docs/sdks/reports/README.md#isagedcreditorsreportavailable) - Aged creditors report available
* [IsAgedDebtorReportAvailable](docs/sdks/reports/README.md#isageddebtorreportavailable) - Aged debtors report available

### [CompanyInfo](docs/sdks/companyinfo/README.md)

* [Get](docs/sdks/companyinfo/README.md#get) - Get company info
* [Refresh](docs/sdks/companyinfo/README.md#refresh) - Refresh company info

### [PaymentMethods](docs/sdks/paymentmethods/README.md)

* [Get](docs/sdks/paymentmethods/README.md#get) - Get payment method
* [List](docs/sdks/paymentmethods/README.md#list) - List payment methods

### [SalesOrders](docs/sdks/salesorders/README.md)

* [Get](docs/sdks/salesorders/README.md#get) - Get sales order
* [List](docs/sdks/salesorders/README.md#list) - List sales orders

### [TaxRates](docs/sdks/taxrates/README.md)

* [Get](docs/sdks/taxrates/README.md#get) - Get tax rate
* [List](docs/sdks/taxrates/README.md#list) - List all tax rates

### [TrackingCategories](docs/sdks/trackingcategories/README.md)

* [Get](docs/sdks/trackingcategories/README.md#get) - Get tracking categories
* [List](docs/sdks/trackingcategories/README.md#list) - List tracking categories
<!-- End Available Resources and Operations [operations] -->





<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"log"
	"pkg/models/operations"
)

func main() {
	s := accounting.New(
		accounting.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountTransactions.Get(ctx, operations.GetAccountTransactionRequest{
		AccountTransactionID: "string",
		CompanyID:            "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID:         "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	}, operations.WithRetries(
		utils.RetryConfig{
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

	if res.AccountTransaction != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"log"
)

func main() {
	s := accounting.New(
		accounting.WithRetryConfig(
			utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		accounting.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountTransactions.Get(ctx, operations.GetAccountTransactionRequest{
		AccountTransactionID: "string",
		CompanyID:            "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID:         "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountTransaction != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"log"
)

func main() {
	s := accounting.New(
		accounting.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountTransactions.Get(ctx, operations.GetAccountTransactionRequest{
		AccountTransactionID: "string",
		CompanyID:            "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID:         "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
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
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"log"
)

func main() {
	s := accounting.New(
		accounting.WithServerIndex(0),
		accounting.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountTransactions.Get(ctx, operations.GetAccountTransactionRequest{
		AccountTransactionID: "string",
		CompanyID:            "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID:         "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountTransaction != nil {
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"log"
)

func main() {
	s := accounting.New(
		accounting.WithServerURL("https://api.codat.io"),
		accounting.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountTransactions.Get(ctx, operations.GetAccountTransactionRequest{
		AccountTransactionID: "string",
		CompanyID:            "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID:         "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountTransaction != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
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
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"log"
)

func main() {
	s := accounting.New(
		accounting.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountTransactions.Get(ctx, operations.GetAccountTransactionRequest{
		AccountTransactionID: "string",
		CompanyID:            "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID:         "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountTransaction != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

<!-- Start Codat Support Notes -->
### Support

If you encounter any challenges while utilizing our SDKs, please don't hesitate to reach out for assistance. 
You can raise any issues by contacting your dedicated Codat representative or reaching out to our [support team](mailto:support@codat.io).
We're here to help ensure a smooth experience for you.
<!-- End Codat Support Notes -->

<!-- Start Codat Generated By -->
### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
<!-- End Codat Generated By -->