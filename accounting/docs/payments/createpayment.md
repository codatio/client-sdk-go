# CreatePayment
Available in: `Payments`

Posts a new payment to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create payment model](https://docs.codat.io/accounting-api#/operations/get-create-payments-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating payments.

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
    req := operations.CreatePaymentRequest{
        Payment: &shared.Payment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("60629451-4c3d-4b9c-a9f3-8bd2be878703"),
                Name: codataccounting.String("Cora Ferry"),
            },
            Currency: codataccounting.String("provident"),
            CurrencyRate: codataccounting.Float64(6660.49),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("mollitia"),
                ID: "8465a328-3279-4b71-9d1c-ea673d86e3b3",
            },
            Date: "2022-10-23T00:00:00Z",
            ID: codataccounting.String("5e49a313-5778-4ce5-8cac-b0e3ea975045"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Amount: 6257.74,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(9916.87),
                            CurrencyRate: codataccounting.Float64(4113.08),
                            ID: codataccounting.String("3b215186-ab5e-43a0-a261-4315d1568299"),
                            Type: shared.PaymentLinkTypeEnumDiscount,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(3778.77),
                            CurrencyRate: codataccounting.Float64(1003.06),
                            ID: codataccounting.String("afc7186f-f20b-47a7-bdf4-0ca0d7657c16"),
                            Type: shared.PaymentLinkTypeEnumInvoice,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(698.78),
                            CurrencyRate: codataccounting.Float64(7166.01),
                            ID: codataccounting.String("bf055271-b251-41dd-a06d-d1b28272bc9c"),
                            Type: shared.PaymentLinkTypeEnumUnlinked,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(1694.68),
                            CurrencyRate: codataccounting.Float64(1257.69),
                            ID: codataccounting.String("1697b188-0fcb-4b2b-93c1-5f670bd17848"),
                            Type: shared.PaymentLinkTypeEnumUnlinked,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Amount: 1228.58,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(3516.07),
                            CurrencyRate: codataccounting.Float64(2234.48),
                            ID: codataccounting.String("eeb3b6e2-41c3-4109-9836-63c66dcbb7df"),
                            Type: shared.PaymentLinkTypeEnumCreditNote,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(7852.92),
                            CurrencyRate: codataccounting.Float64(7427.38),
                            ID: codataccounting.String("09c8b408-e071-4377-8de4-fee101d9780a"),
                            Type: shared.PaymentLinkTypeEnumUnlinked,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Amount: 434.45,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(2602.42),
                            CurrencyRate: codataccounting.Float64(4839.58),
                            ID: codataccounting.String("b95040d6-c8b2-4a5f-8022-07e4048f9000"),
                            Type: shared.PaymentLinkTypeEnumPayment,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(9351.45),
                            CurrencyRate: codataccounting.Float64(8479.43),
                            ID: codataccounting.String("290278eb-4ae9-4d64-961e-91500323b2c0"),
                            Type: shared.PaymentLinkTypeEnumRefund,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(7176.59),
                            CurrencyRate: codataccounting.Float64(6050.35),
                            ID: codataccounting.String("24771f56-69e5-4b7e-8762-6649d84eb9e4"),
                            Type: shared.PaymentLinkTypeEnumPaymentOnAccount,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(9525.87),
                            CurrencyRate: codataccounting.Float64(8696.62),
                            ID: codataccounting.String("2276e0b8-8fb8-47d6-ba5b-6e8dbf812f83"),
                            Type: shared.PaymentLinkTypeEnumPaymentOnAccount,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("inventore"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("ca6a9ffc-5619-429c-8a95-60a1395918da"),
                Name: codataccounting.String("Hope Gusikowski"),
            },
            Reference: codataccounting.String("nihil"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "officiis": map[string]interface{}{
                        "quisquam": "asperiores",
                    },
                    "praesentium": map[string]interface{}{
                        "inventore": "ab",
                        "dolore": "amet",
                        "nulla": "officia",
                        "natus": "nesciunt",
                    },
                    "eaque": map[string]interface{}{
                        "nobis": "magni",
                        "nihil": "laborum",
                        "aut": "voluptatum",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(6542.99),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(969433),
    }

    res, err := s.Payments.CreatePayment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePaymentResponse != nil {
        // handle response
    }
}
```