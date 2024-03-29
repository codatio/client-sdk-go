# AccountingDirectIncomes
(*AccountingDirectIncomes*)

## Overview

Direct incomes

### Available Operations

* [CreateAccountingDirectIncome](#createaccountingdirectincome) - Create direct income

## CreateAccountingDirectIncome

The *Create direct income* endpoint creates a new [direct income](https://docs.codat.io/accounting-api#/schemas/DirectIncome) for a given company's connection.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale. For example, cash sales of items to a customer, referral commissions, and service fee refunds are considered direct incomes.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create direct income model](https://docs.codat.io/accounting-api#/operations/get-create-directIncomes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/types"
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
    res, err := s.AccountingDirectIncomes.CreateAccountingDirectIncome(ctx, operations.CreateAccountingDirectIncomeRequest{
        AccountingDirectIncome: &shared.AccountingDirectIncome{
            ContactRef: &shared.ContactRef{
                DataType: shared.DataTypeInvoices.ToPointer(),
                ID: "<ID>",
            },
            Currency: "GBP",
            IssueDate: "2022-10-23T00:00:00Z",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{},
                    ItemRef: &shared.ItemReference{
                        ID: "<ID>",
                    },
                    Quantity: types.MustNewDecimalFromString("3642.55"),
                    TaxRateRef: &shared.TaxRateReference{},
                    TrackingCategoryRefs: []shared.TrackingCategoryRefs{
                        shared.TrackingCategoryRefs{
                            ID: "<ID>",
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("7950.78"),
                },
            },
            Metadata: &shared.Metadata{},
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
            PaymentAllocations: []shared.PaymentAllocationItems{
                shared.PaymentAllocationItems{
                    Allocation: shared.Allocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
                        Currency: syncforcommerceversion1.String("EUR"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{},
                        Currency: syncforcommerceversion1.String("USD"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
                    },
                },
            },
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
            SubTotal: types.MustNewDecimalFromString("5786.44"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "key": map[string]interface{}{
                        "key": "string",
                    },
                },
            },
            TaxAmount: types.MustNewDecimalFromString("2812.91"),
            TotalAmount: types.MustNewDecimalFromString("6636.11"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateDirectIncomeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.CreateAccountingDirectIncomeRequest](../../pkg/models/operations/createaccountingdirectincomerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |


### Response

**[*operations.CreateAccountingDirectIncomeResponse](../../pkg/models/operations/createaccountingdirectincomeresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |
