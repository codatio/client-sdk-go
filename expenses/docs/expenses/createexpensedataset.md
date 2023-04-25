# CreateExpenseDataset
Available in: `Expenses`

Create an expense transaction

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateExpenseDatasetRequest{
        CreateExpenseRequest: &shared.CreateExpenseRequest{
            Items: []shared.ExpenseTransaction{
                shared.ExpenseTransaction{
                    Currency: "GBP",
                    CurrencyRate: codatsyncexpenses.Float64(5928.45),
                    ID: "4d7c6929-7770-412b-91bb-44d3bc71d111",
                    IssueDate: "2022-10-23T00:00:00Z",
                    Lines: []shared.ExpenseTransactionLine{
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: 110.42,
                            TaxAmount: 14.43,
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: 110.42,
                            TaxAmount: 14.43,
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: 110.42,
                            TaxAmount: 14.43,
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                    },
                    MerchantName: codatsyncexpenses.String("Amazon UK"),
                    Notes: codatsyncexpenses.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
                    Type: shared.ExpenseTransactionTypeEnumPayment,
                },
                shared.ExpenseTransaction{
                    Currency: "GBP",
                    CurrencyRate: codatsyncexpenses.Float64(5448.83),
                    ID: "4d7c6929-7770-412b-91bb-44d3bc71d111",
                    IssueDate: "2022-10-23T00:00:00Z",
                    Lines: []shared.ExpenseTransactionLine{
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: 110.42,
                            TaxAmount: 14.43,
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: 110.42,
                            TaxAmount: 14.43,
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: 110.42,
                            TaxAmount: 14.43,
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: 110.42,
                            TaxAmount: 14.43,
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                    },
                    MerchantName: codatsyncexpenses.String("Amazon UK"),
                    Notes: codatsyncexpenses.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
                    Type: shared.ExpenseTransactionTypeEnumPayment,
                },
                shared.ExpenseTransaction{
                    Currency: "GBP",
                    CurrencyRate: codatsyncexpenses.Float64(4375.87),
                    ID: "4d7c6929-7770-412b-91bb-44d3bc71d111",
                    IssueDate: "2022-10-23T00:00:00Z",
                    Lines: []shared.ExpenseTransactionLine{
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: 110.42,
                            TaxAmount: 14.43,
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: 110.42,
                            TaxAmount: 14.43,
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                    },
                    MerchantName: codatsyncexpenses.String("Amazon UK"),
                    Notes: codatsyncexpenses.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
                    Type: shared.ExpenseTransactionTypeEnumPayment,
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Expenses.CreateExpenseDataset(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateExpenseResponse != nil {
        // handle response
    }
}
```