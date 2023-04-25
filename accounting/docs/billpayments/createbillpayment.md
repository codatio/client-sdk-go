# CreateBillPayment
Available in: `BillPayments`

Posts a new bill payment to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create bill payment model](https://docs.codat.io/accounting-api#/operations/get-create-billPayments-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billPayments) for integrations that support creating bill payments.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateBillPaymentRequest{
        BillPayment: &shared.BillPayment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("825c1fc0-e115-4c80-bff9-18544ec42def"),
                Name: codataccounting.String("Gregg Torp"),
            },
            Currency: codataccounting.String("ab"),
            CurrencyRate: codataccounting.Float64(5734.44),
            Date: "2022-10-23T00:00:00Z",
            ID: codataccounting.String("77773e63-562a-47b4-88f0-5e3d48fdaf31"),
            Lines: []shared.BillPaymentLine{
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Amount: 6308.32,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(9979.95),
                            CurrencyRate: codataccounting.Float64(3632.14),
                            ID: codataccounting.String("fd94259c-0b36-4f25-aa94-4f3b756c11f6"),
                            Type: shared.BillPaymentLineLinkTypeEnumRefund,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("dolor"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("7a512624-3835-4bbc-85a2-3a45cefc5fde"),
                Name: codataccounting.String("Miss Linda Nader"),
            },
            Reference: codataccounting.String("quia"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "vel": map[string]interface{}{
                        "debitis": "ullam",
                        "architecto": "accusantium",
                        "perferendis": "veritatis",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "9c6dc5e3-4762-4799-bfbb-e6949fb2bb4e",
                SupplierName: codataccounting.String("quod"),
            },
            TotalAmount: codataccounting.Float64(6646.66),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(904440),
    }

    res, err := s.BillPayments.CreateBillPayment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillPaymentResponse != nil {
        // handle response
    }
}
```