# CreateJournalEntry
Available in: `JournalEntries`

Posts a new journalEntry to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create journal entry model](https://docs.codat.io/accounting-api#/operations/get-create-journalEntries-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support creating journal entries.

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
    req := operations.CreateJournalEntryRequest{
        JournalEntry: &shared.JournalEntry{
            CreatedOn: codataccounting.String("2022-10-23T00:00:00Z"),
            Description: codataccounting.String("mollitia"),
            ID: codataccounting.String("a30b7b91-449a-4e69-8088-d418bb71804f"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("23d54393-5f37-47ac-9c9b-7e93b6a3c523"),
                        Name: codataccounting.String("Sharon Hayes"),
                    },
                    Currency: codataccounting.String("impedit"),
                    Description: codataccounting.String("dolor"),
                    NetAmount: 3078.67,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("deserunt"),
                                ID: codataccounting.String("b0ecb812-a661-4489-84a8-e9085075bc25"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("dolorem"),
                                ID: codataccounting.String("8253343f-b0a4-4e66-aa47-578d171e2941"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("totam"),
                                ID: codataccounting.String("18fc679b-6b2f-4253-99b8-55d015b62c8b"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("praesentium"),
                                ID: codataccounting.String("3a38a8a8-8c14-4420-8c2c-aeb1ae1ecf8c"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("34946bba-7a05-4a8b-8a9e-c5b3688cca36"),
                        Name: codataccounting.String("Sara Kuhic"),
                    },
                    Currency: codataccounting.String("iure"),
                    Description: codataccounting.String("sit"),
                    NetAmount: 8973.52,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("vel"),
                                ID: codataccounting.String("6e97e054-1033-447d-b8ff-2491145fab9e"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("veniam"),
                                ID: codataccounting.String("9a4af336-664e-4aa6-bf2f-f14e8c1b352a"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("porro"),
                                ID: codataccounting.String("cedacc52-2781-44ec-a016-bc41ea1342d4"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "104a25ef-71de-457a-91d6-14a4317692ea",
                Name: codataccounting.String("Lena Howell"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            PostedOn: codataccounting.String("2022-10-23T00:00:00Z"),
            RecordRef: &shared.InvoiceTo{
                DataType: codataccounting.String("vero"),
                ID: codataccounting.String("522b828a-9030-4660-b024-c79b4cc64c2b"),
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "culpa": map[string]interface{}{
                        "odit": "optio",
                    },
                },
            },
            UpdatedOn: codataccounting.String("2022-10-23T00:00:00Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(287244),
    }

    res, err := s.JournalEntries.CreateJournalEntry(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateJournalEntryResponse != nil {
        // handle response
    }
}
```