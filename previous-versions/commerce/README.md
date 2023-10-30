# Commerce

<!-- Start Codat Library Description -->
﻿Codat's Commerce API enables you to pull up-date-date commerce data from several leading payments, point-of-sale, and eCommerce systems.
You can view your SMB customers' products, orders, payments, payouts, disputes, and more - all standardized to our Commerce data model.

<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/commerce
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
	"log"
)

func main() {
	s := commerce.New(
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.CompanyInfo.Get(ctx, operations.GetCompanyInfoRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CompanyInfo != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [CompanyInfo](docs/sdks/companyinfo/README.md)

* [Get](docs/sdks/companyinfo/README.md#get) - Get company info

### [Customers](docs/sdks/customers/README.md)

* [Get](docs/sdks/customers/README.md#get) - Get customer
* [List](docs/sdks/customers/README.md#list) - List customers

### [Disputes](docs/sdks/disputes/README.md)

* [Get](docs/sdks/disputes/README.md#get) - Get dispute
* [List](docs/sdks/disputes/README.md#list) - List disputes

### [Locations](docs/sdks/locations/README.md)

* [Get](docs/sdks/locations/README.md#get) - Get location
* [List](docs/sdks/locations/README.md#list) - List locations

### [Orders](docs/sdks/orders/README.md)

* [Get](docs/sdks/orders/README.md#get) - Get order
* [List](docs/sdks/orders/README.md#list) - List orders

### [Payments](docs/sdks/payments/README.md)

* [Get](docs/sdks/payments/README.md#get) - Get payment
* [GetMethod](docs/sdks/payments/README.md#getmethod) - Get payment method
* [List](docs/sdks/payments/README.md#list) - List payments
* [ListMethods](docs/sdks/payments/README.md#listmethods) - List payment methods

### [Products](docs/sdks/products/README.md)

* [Get](docs/sdks/products/README.md#get) - Get product
* [GetCategory](docs/sdks/products/README.md#getcategory) - Get product category
* [List](docs/sdks/products/README.md#list) - List products
* [ListCategories](docs/sdks/products/README.md#listcategories) - List product categories

### [TaxComponents](docs/sdks/taxcomponents/README.md)

* [Get](docs/sdks/taxcomponents/README.md#get) - Get tax component
* [List](docs/sdks/taxcomponents/README.md#list) - List tax components

### [Transactions](docs/sdks/transactions/README.md)

* [Get](docs/sdks/transactions/README.md#get) - Get transaction
* [List](docs/sdks/transactions/README.md#list) - List transactions
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.codat.io` | None |

For example:


```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
	"log"
)

func main() {
	s := commerce.New(
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
		commerce.WithServerIndex(0),
	)

	ctx := context.Background()
	res, err := s.CompanyInfo.Get(ctx, operations.GetCompanyInfoRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CompanyInfo != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:


```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
	"log"
)

func main() {
	s := commerce.New(
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
		commerce.WithServerURL("https://api.codat.io"),
	)

	ctx := context.Background()
	res, err := s.CompanyInfo.Get(ctx, operations.GetCompanyInfoRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CompanyInfo != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

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

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)