# Bank Feeds API

Bank Feeds API enables your SMB users to set up bank feeds from accounts in your application to supported accounting platforms.

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/bankfeeds
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccountTransactions.Create(ctx, operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: codatbankfeeds.String("corrupti"),
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 7151.9,
                    Balance: 8442.66,
                    ClearedOnDate: codatbankfeeds.String("unde"),
                    Counterparty: codatbankfeeds.String("nulla"),
                    Description: codatbankfeeds.String("corrupti"),
                    ID: codatbankfeeds.String("d69a674e-0f46-47cc-8796-ed151a05dfc2"),
                    ModifiedDate: codatbankfeeds.String("at"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("at"),
                    SourceModifiedDate: codatbankfeeds.String("maiores"),
                    TransactionType: shared.BankTransactionTypeAtm,
                },
                shared.BankTransactionLine{
                    Amount: 7991.59,
                    Balance: 8009.11,
                    ClearedOnDate: codatbankfeeds.String("esse"),
                    Counterparty: codatbankfeeds.String("totam"),
                    Description: codatbankfeeds.String("porro"),
                    ID: codatbankfeeds.String("a1ba928f-c816-4742-8b73-9205929396fe"),
                    ModifiedDate: codatbankfeeds.String("fuga"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("in"),
                    SourceModifiedDate: codatbankfeeds.String("corporis"),
                    TransactionType: shared.BankTransactionTypeCheck,
                },
                shared.BankTransactionLine{
                    Amount: 4370.32,
                    Balance: 9023.49,
                    ClearedOnDate: codatbankfeeds.String("quidem"),
                    Counterparty: codatbankfeeds.String("architecto"),
                    Description: codatbankfeeds.String("ipsa"),
                    ID: codatbankfeeds.String("faaa2352-c595-4590-baff-1a3a2fa94677"),
                    ModifiedDate: codatbankfeeds.String("velit"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("error"),
                    SourceModifiedDate: codatbankfeeds.String("quia"),
                    TransactionType: shared.BankTransactionTypeSerChg,
                },
            },
        },
        AccountID: "vitae",
        AllowSyncOnPushComplete: codatbankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatbankfeeds.Int(674752),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankTransactionsResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [BankAccountTransactions](docs/bankaccounttransactions/README.md)

* [Create](docs/bankaccounttransactions/README.md#create) - Create bank transactions
* [Get](docs/bankaccounttransactions/README.md#get) - List push options for bank account bank transactions
* [List](docs/bankaccounttransactions/README.md#list) - List bank transactions for bank account

### [BankFeedAccounts](docs/bankfeedaccounts/README.md)

* [Create](docs/bankfeedaccounts/README.md#create) - Create bank feed bank accounts
* [Get](docs/bankfeedaccounts/README.md#get) - List bank feed bank accounts
* [Update](docs/bankfeedaccounts/README.md#update) - Update bank feed bank account

### [Companies](docs/companies/README.md)

* [Create](docs/companies/README.md#create) - Create company
* [Delete](docs/companies/README.md#delete) - Delete a company
* [Get](docs/companies/README.md#get) - Get company
* [List](docs/companies/README.md#list) - List companies
* [Update](docs/companies/README.md#update) - Update company

### [Connections](docs/connections/README.md)

* [Create](docs/connections/README.md#create) - Create connection
* [Delete](docs/connections/README.md#delete) - Delete connection
* [Get](docs/connections/README.md#get) - Get connection
* [List](docs/connections/README.md#list) - List connections
* [Proxy](docs/connections/README.md#proxy) - A proxy or passthrough endpoint used to query unsupported third party endpoints.
* [UnlinkConnection](docs/connections/README.md#unlinkconnection) - Unlink connection
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
