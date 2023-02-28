# github.com/codatio/client-sdk-go/commerce

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/commerce
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/commerce"
    "github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
    "github.com/codatio/client-sdk-go/commerce/pkg/models/operations"
)

func main() {
    opts := []codatio.SDKOption{
        codatio.WithSecurity(
            shared.Security{
                APIKey: shared.SchemeAPIKey{
                    APIKey: "YOUR_API_KEY_HERE",
                },
            },
        ),
    }

    s := codatio.New(opts...)
    
    req := operations.GetCommerceInfoRequest{
        PathParams: operations.GetCommerceInfoPathParams{
            CompanyID: "unde",
            ConnectionID: "deserunt",
        },
    }

    ctx := context.Background()
    res, err := s.CompanyInfo.GetCommerceInfo(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SourceModifiedDate != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### CompanyInfo

* `GetCommerceInfo` - Get company info

### Customers

* `ListCommerceCustomers` - List customers

### Disputes

* `ListCommerceDisputes` - List disputes

### Locations

* `ListCommerceLocations` - List locations

### Orders

* `ListCommerceOrders` - List orders

### Payments

* `ListCommercePaymentMethods` - List payment methods
* `ListCommercePayments` - List payments

### Products

* `ListCommerceProductCategories` - List product categories
* `ListCommerceProducts` - List products

### TaxComponents

* `GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents` - List tax components

### Transactions

* `ListCommerceTransactions` - List transactions
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)