# Commerce

<!-- Start Codat Library Description -->
﻿Codat's Commerce API enables you to pull up-date-date commerce data from several leading payments, point-of-sale, and eCommerce systems.
You can view your SMB customers' products, orders, payments, payouts, disputes, and more - all standardized to our Commerce data model.

<!-- End Codat Library Description -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/codatio/client-sdk-go/previous-versions/commerce
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
	commerce "github.com/codatio/client-sdk-go/previous-versions/commerce/v3"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := commerce.New(
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		CustomerID:   "7110701885",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Customer != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>


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

</details>
<!-- End Available Resources and Operations [operations] -->





<!-- Start Summary [summary] -->
## Summary

Commerce API: Codat's standardized API for accessing commerce data 

> ### New to Codat?
>
> Our Commerce API reference is relevant only to our existing clients.
> Please reach out to your Codat contact so that we can find the right product for you.

Codat's Commerce API allows you to access standardised data from over 11 commerce and POS systems.

Standardize how you connect to your customers’ payment, PoS, and eCommerce systems. Retrieve orders, payouts, payments, and product data in the same way for all the leading commerce software.

<!-- Start Codat Tags Table -->
## Endpoints

| Endpoints | Description |
| :- |:- |
| Customers | Retrieve standardized data from linked commerce software. |
| Disputes | Retrieve standardized data from linked commerce software. |
| Company info | Retrieve standardized data from linked commerce software. |
| Locations | Retrieve standardized data from linked commerce software. |
| Orders | Retrieve standardized data from linked commerce software. |
| Payments | Retrieve standardized data from linked commerce software. |
| Products | Retrieve standardized data from linked commerce software. |
| Tax components | Retrieve standardized data from linked commerce software. |
| Transactions | Retrieve standardized data from linked commerce software. |
<!-- End Codat Tags Table -->

[Read more...](https://docs.codat.io/commerce-api/overview)

[See our OpenAPI spec](https://github.com/codatio/oas)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
<!-- End Table of Contents [toc] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	commerce "github.com/codatio/client-sdk-go/previous-versions/commerce/v3"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	s := commerce.New(
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		CustomerID:   "7110701885",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
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
	if res.Customer != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	commerce "github.com/codatio/client-sdk-go/previous-versions/commerce/v3"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/retry"
	"log"
)

func main() {
	s := commerce.New(
		commerce.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		CustomerID:   "7110701885",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Customer != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Get` function may return the following errors:

| Error Type             | Status Code                            | Content Type     |
| ---------------------- | -------------------------------------- | ---------------- |
| sdkerrors.ErrorMessage | 401, 402, 403, 404, 409, 429, 500, 503 | application/json |
| sdkerrors.SDKError     | 4XX, 5XX                               | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	commerce "github.com/codatio/client-sdk-go/previous-versions/commerce/v3"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := commerce.New(
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		CustomerID:   "7110701885",
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

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	commerce "github.com/codatio/client-sdk-go/previous-versions/commerce/v3"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := commerce.New(
		commerce.WithServerURL("https://api.codat.io"),
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		CustomerID:   "7110701885",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Customer != nil {
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

| Name         | Type   | Scheme  |
| ------------ | ------ | ------- |
| `AuthHeader` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	commerce "github.com/codatio/client-sdk-go/previous-versions/commerce/v3"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := commerce.New(
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		CustomerID:   "7110701885",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Customer != nil {
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