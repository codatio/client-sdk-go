# AccountingCustomers
(*AccountingCustomers*)

## Overview

Customers

### Available Operations

* [CreateAccountingCustomer](#createaccountingcustomer) - Create customer

## CreateAccountingCustomer

The *Create customer* endpoint creates a new [customer](https://docs.codat.io/accounting-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"log"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingCustomers.CreateAccountingCustomer(ctx, operations.CreateAccountingCustomerRequest{
        AccountingCustomer: &shared.AccountingCustomer{
            Addresses: []shared.Items1{
                shared.Items1{
                    Type: shared.AccountingAddressTypeDelivery,
                },
            },
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items1{
                        Type: shared.AccountingAddressTypeUnknown,
                    },
                    ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
                    Phone: []shared.PhoneNumberItems{
                        shared.PhoneNumberItems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            DefaultCurrency: syncforcommerceversion1.String("GBP"),
            Metadata: &shared.Metadata{},
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
            Status: shared.CustomerStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "key": map[string]interface{}{
                        "key": "string",
                    },
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateAccountingCustomerRequest](../../pkg/models/operations/createaccountingcustomerrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.CreateAccountingCustomerResponse](../../pkg/models/operations/createaccountingcustomerresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |
