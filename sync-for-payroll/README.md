# Sync for Payroll

<!-- Start Codat Library Description -->
Push payroll to accounting platforms.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/sync-for-payroll
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/types"
	"log"
)

func main() {
	s := syncforpayroll.New(
		syncforpayroll.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
		Account: &shared.Account{
			Currency:               syncforpayroll.String("USD"),
			CurrentBalance:         types.MustNewDecimalFromString("0"),
			Description:            syncforpayroll.String("Invoices the business has issued but has not yet collected payment on."),
			FullyQualifiedCategory: syncforpayroll.String("Asset.Current"),
			FullyQualifiedName:     syncforpayroll.String("Cash On Hand"),
			ID:                     syncforpayroll.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
			Metadata:               &shared.AccountMetadata{},
			ModifiedDate:           syncforpayroll.String("2022-10-23T00:00:00.000Z"),
			Name:                   syncforpayroll.String("Accounts Receivable"),
			NominalCode:            syncforpayroll.String("610"),
			SourceModifiedDate:     syncforpayroll.String("2022-10-23T00:00:00.000Z"),
			Status:                 shared.AccountStatusActive.ToPointer(),
			SupplementalData: &shared.SupplementalData{
				Content: map[string]map[string]interface{}{
					"Money": map[string]interface{}{
						"blue": "shred",
					},
				},
			},
			Type: shared.AccountTypeAsset.ToPointer(),
			ValidDatatypeLinks: []shared.AccountValidDataTypeLinks{
				shared.AccountValidDataTypeLinks{
					Links: []string{
						"abnormally",
					},
				},
			},
		},
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAccountResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Accounts](docs/sdks/accounts/README.md)

* [Create](docs/sdks/accounts/README.md#create) - Create account
* [Get](docs/sdks/accounts/README.md#get) - Get account
* [GetCreateModel](docs/sdks/accounts/README.md#getcreatemodel) - Get create account model
* [List](docs/sdks/accounts/README.md#list) - List accounts

### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [CompanyInfo](docs/sdks/companyinfo/README.md)

* [GetAccountingProfile](docs/sdks/companyinfo/README.md#getaccountingprofile) - Get company accounting profile

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [JournalEntries](docs/sdks/journalentries/README.md)

* [Create](docs/sdks/journalentries/README.md#create) - Create journal entry
* [Delete](docs/sdks/journalentries/README.md#delete) - Delete journal entry
* [Get](docs/sdks/journalentries/README.md#get) - Get journal entry
* [GetCreateModel](docs/sdks/journalentries/README.md#getcreatemodel) - Get create journal entry model
* [List](docs/sdks/journalentries/README.md#list) - List journal entries

### [Journals](docs/sdks/journals/README.md)

* [Create](docs/sdks/journals/README.md#create) - Create journal
* [Get](docs/sdks/journals/README.md#get) - Get journal
* [GetCreateModel](docs/sdks/journals/README.md#getcreatemodel) - Get create journal model
* [List](docs/sdks/journals/README.md#list) - List journals

### [ManageData](docs/sdks/managedata/README.md)

* [GetDataStatus](docs/sdks/managedata/README.md#getdatastatus) - Get data status
* [GetPullOperation](docs/sdks/managedata/README.md#getpulloperation) - Get pull operation
* [GetPushOperation](docs/sdks/managedata/README.md#getpushoperation) - Get push operation
* [List](docs/sdks/managedata/README.md#list) - List push operations
* [ListPullOperations](docs/sdks/managedata/README.md#listpulloperations) - List pull operations
* [RefreshAllDataTypes](docs/sdks/managedata/README.md#refreshalldatatypes) - Refresh all data
* [RefreshDataType](docs/sdks/managedata/README.md#refreshdatatype) - Refresh data type

### [TrackingCategories](docs/sdks/trackingcategories/README.md)

* [Get](docs/sdks/trackingcategories/README.md#get) - Get tracking categories
* [List](docs/sdks/trackingcategories/README.md#list) - List tracking categories
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->



<!-- End Dev Containers -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)