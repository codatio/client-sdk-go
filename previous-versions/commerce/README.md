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

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
)

func main() {
    s := codatcommerce.New(
        codatcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CompanyInfo.Get(ctx, operations.GetCompanyInfoRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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