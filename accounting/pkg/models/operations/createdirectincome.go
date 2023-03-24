// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

// CreateDirectIncomeSourceModifiedDateContactRef - A customer or supplier associated with the direct income.
type CreateDirectIncomeSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateDirectIncomeSourceModifiedDateLineItemsAccountRef - Reference to the account to which the line item is linked.
type CreateDirectIncomeSourceModifiedDateLineItemsAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncomeSourceModifiedDateLineItemsItemRef - Reference to the product, service type, or inventory item to which the direct cost is linked.
type CreateDirectIncomeSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncomeSourceModifiedDateLineItemsTaxRateRef - Reference to the tax rate to which the line item is linked.
type CreateDirectIncomeSourceModifiedDateLineItemsTaxRateRef struct {
	// Applicable tax rate.
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	// 'id' from the 'taxRates' data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the 'taxRates' data type.
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncomeSourceModifiedDateLineItemsTrackingCategoryRefs - References a category against which the item is tracked.
type CreateDirectIncomeSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectIncomeSourceModifiedDateLineItems struct {
	// Reference to the account to which the line item is linked.
	AccountRef *CreateDirectIncomeSourceModifiedDateLineItemsAccountRef `json:"accountRef,omitempty"`
	// A user-friendly name of the goods or services.
	Description *string `json:"description,omitempty"`
	// Discount amount for the line before tax.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// Discount percentage for the line before tax.
	DiscountPercentage *float64 `json:"discountPercentage,omitempty"`
	// Reference to the product, service type, or inventory item to which the direct cost is linked.
	ItemRef *CreateDirectIncomeSourceModifiedDateLineItemsItemRef `json:"itemRef,omitempty"`
	// The number of units of goods or services received.
	//
	// Note: If the platform does not provide this information, the quantity will be mapped as 1.
	Quantity float64 `json:"quantity"`
	// The amount of the line, inclusive of discounts, but exclusive of tax.
	SubTotal *float64 `json:"subTotal,omitempty"`
	// The amount of tax for the line.
	// Note: If the platform does not provide this information, the quantity will be mapped as 0.00.
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Reference to the tax rate to which the line item is linked.
	TaxRateRef *CreateDirectIncomeSourceModifiedDateLineItemsTaxRateRef `json:"taxRateRef,omitempty"`
	// The total amount of the line, including tax.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// An array of categories against which this direct cost is tracked.
	TrackingCategoryRefs []CreateDirectIncomeSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	// The price of each unit of goods or services.
	// Note: If the platform does not provide this information, the unit amount will be mapped to the total amount.
	UnitAmount float64 `json:"unitAmount"`
}

type CreateDirectIncomeSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateDirectIncomeSourceModifiedDatePaymentAllocationsAllocation struct {
	// The date the payment was allocated.
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	// The currency of the transaction.
	Currency *string `json:"currency,omitempty"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// The total amount that has been allocated.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

// CreateDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef - The account that the allocated payment is made from or to.
type CreateDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

type CreateDirectIncomeSourceModifiedDatePaymentAllocationsPayment struct {
	// The account that the allocated payment is made from or to.
	AccountRef *CreateDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	// Currency the payment has been made in.
	Currency *string `json:"currency,omitempty"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// Identifier of the allocated payment.
	ID *string `json:"id,omitempty"`
	// Notes attached to the allocated payment.
	Note *string `json:"note,omitempty"`
	// The date the payment was paid.
	PaidOnDate *time.Time `json:"paidOnDate,omitempty"`
	// Reference to the allocated payment.
	Reference *string `json:"reference,omitempty"`
	// Total amount that was paid.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

type CreateDirectIncomeSourceModifiedDatePaymentAllocations struct {
	Allocation CreateDirectIncomeSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateDirectIncomeSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

// CreateDirectIncomeSourceModifiedDateSupplementalData - Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateDirectIncomeSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateDirectIncomeSourceModifiedDate - > **Language tip:**  Direct incomes may also be referred to as **Receive transactions**, **Receive money transactions**, **Sales receipts**, or **Cash sales** in various accounting platforms.
//
// > View the coverage for direct incomes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Direct incomes include:
//
// - Selling an item directly to a contact, and receiving payment at the point of the sale.
// - Refunding an item in cash to a contact.
// - Depositing money into a bank account.
//
// From the Direct Incomes endpoints, you can:
//
// - [Get a list of all direct incomes for a specific company](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes)
// - [Get a single direct income for a specific company and connection](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes__directIncomeId_)
// - [Add a new direct income to a specific company's accounting package](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/post_companies__companyId__connections__connectionId__push_directIncomes)
//
// Direct incomes is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type CreateDirectIncomeSourceModifiedDate struct {
	// A customer or supplier associated with the direct income.
	ContactRef *CreateDirectIncomeSourceModifiedDateContactRef `json:"contactRef,omitempty"`
	// The currency of the direct income.
	Currency string `json:"currency"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// Identifier of the direct income, unique for the company.
	ID *string `json:"id,omitempty"`
	// The date of the direct income as recorded in the accounting platform.
	IssueDate time.Time `json:"issueDate"`
	// An array of line items.
	LineItems []CreateDirectIncomeSourceModifiedDateLineItems `json:"lineItems"`
	Metadata  *CreateDirectIncomeSourceModifiedDateMetadata   `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate       *time.Time                                               `json:"modifiedDate,omitempty"`
	Note               *string                                                  `json:"note,omitempty"`
	PaymentAllocations []CreateDirectIncomeSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	// User-friendly reference for the direct income.
	Reference *string `json:"reference,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
	// The total amount of the direct incomes, excluding any taxes.
	SubTotal float64 `json:"subTotal"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *CreateDirectIncomeSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	// The total amount of tax on the direct incomes.
	TaxAmount float64 `json:"taxAmount"`
	// The amount of the direct incomes, inclusive of tax.
	TotalAmount float64 `json:"totalAmount"`
}

type CreateDirectIncomeRequest struct {
	RequestBody      *CreateDirectIncomeSourceModifiedDate `request:"mediaType=application/json"`
	CompanyID        string                                `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                                `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int                                  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateDirectIncome200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateDirectIncome200ApplicationJSONChangesTypeEnum string

const (
	CreateDirectIncome200ApplicationJSONChangesTypeEnumUnknown            CreateDirectIncome200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateDirectIncome200ApplicationJSONChangesTypeEnumCreated            CreateDirectIncome200ApplicationJSONChangesTypeEnum = "Created"
	CreateDirectIncome200ApplicationJSONChangesTypeEnumModified           CreateDirectIncome200ApplicationJSONChangesTypeEnum = "Modified"
	CreateDirectIncome200ApplicationJSONChangesTypeEnumDeleted            CreateDirectIncome200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateDirectIncome200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateDirectIncome200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateDirectIncome200ApplicationJSONChanges struct {
	AttachmentID *string                                                            `json:"attachmentId,omitempty"`
	RecordRef    *CreateDirectIncome200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateDirectIncome200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateContactRef - A customer or supplier associated with the direct income.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsAccountRef - Reference to the account to which the line item is linked.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsItemRef - Reference to the product, service type, or inventory item to which the direct cost is linked.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef - Reference to the tax rate to which the line item is linked.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	// Applicable tax rate.
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	// 'id' from the 'taxRates' data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the 'taxRates' data type.
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs - References a category against which the item is tracked.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItems struct {
	// Reference to the account to which the line item is linked.
	AccountRef *CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsAccountRef `json:"accountRef,omitempty"`
	// A user-friendly name of the goods or services.
	Description *string `json:"description,omitempty"`
	// Discount amount for the line before tax.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// Discount percentage for the line before tax.
	DiscountPercentage *float64 `json:"discountPercentage,omitempty"`
	// Reference to the product, service type, or inventory item to which the direct cost is linked.
	ItemRef *CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsItemRef `json:"itemRef,omitempty"`
	// The number of units of goods or services received.
	//
	// Note: If the platform does not provide this information, the quantity will be mapped as 1.
	Quantity float64 `json:"quantity"`
	// The amount of the line, inclusive of discounts, but exclusive of tax.
	SubTotal *float64 `json:"subTotal,omitempty"`
	// The amount of tax for the line.
	// Note: If the platform does not provide this information, the quantity will be mapped as 0.00.
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Reference to the tax rate to which the line item is linked.
	TaxRateRef *CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef `json:"taxRateRef,omitempty"`
	// The total amount of the line, including tax.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// An array of categories against which this direct cost is tracked.
	TrackingCategoryRefs []CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	// The price of each unit of goods or services.
	// Note: If the platform does not provide this information, the unit amount will be mapped to the total amount.
	UnitAmount float64 `json:"unitAmount"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	// The date the payment was allocated.
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	// The currency of the transaction.
	Currency *string `json:"currency,omitempty"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// The total amount that has been allocated.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef - The account that the allocated payment is made from or to.
type CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	// The account that the allocated payment is made from or to.
	AccountRef *CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	// Currency the payment has been made in.
	Currency *string `json:"currency,omitempty"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// Identifier of the allocated payment.
	ID *string `json:"id,omitempty"`
	// Notes attached to the allocated payment.
	Note *string `json:"note,omitempty"`
	// The date the payment was paid.
	PaidOnDate *time.Time `json:"paidOnDate,omitempty"`
	// Reference to the allocated payment.
	Reference *string `json:"reference,omitempty"`
	// Total amount that was paid.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateSupplementalData - Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDate - > **Language tip:**  Direct incomes may also be referred to as **Receive transactions**, **Receive money transactions**, **Sales receipts**, or **Cash sales** in various accounting platforms.
//
// > View the coverage for direct incomes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Direct incomes include:
//
// - Selling an item directly to a contact, and receiving payment at the point of the sale.
// - Refunding an item in cash to a contact.
// - Depositing money into a bank account.
//
// From the Direct Incomes endpoints, you can:
//
// - [Get a list of all direct incomes for a specific company](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes)
// - [Get a single direct income for a specific company and connection](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes__directIncomeId_)
// - [Add a new direct income to a specific company's accounting package](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/post_companies__companyId__connections__connectionId__push_directIncomes)
//
// Direct incomes is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type CreateDirectIncome200ApplicationJSONSourceModifiedDate struct {
	// A customer or supplier associated with the direct income.
	ContactRef *CreateDirectIncome200ApplicationJSONSourceModifiedDateContactRef `json:"contactRef,omitempty"`
	// The currency of the direct income.
	Currency string `json:"currency"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// Identifier of the direct income, unique for the company.
	ID *string `json:"id,omitempty"`
	// The date of the direct income as recorded in the accounting platform.
	IssueDate time.Time `json:"issueDate"`
	// An array of line items.
	LineItems []CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItems `json:"lineItems"`
	Metadata  *CreateDirectIncome200ApplicationJSONSourceModifiedDateMetadata   `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate       *time.Time                                                                 `json:"modifiedDate,omitempty"`
	Note               *string                                                                    `json:"note,omitempty"`
	PaymentAllocations []CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	// User-friendly reference for the direct income.
	Reference *string `json:"reference,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
	// The total amount of the direct incomes, excluding any taxes.
	SubTotal float64 `json:"subTotal"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *CreateDirectIncome200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	// The total amount of tax on the direct incomes.
	TaxAmount float64 `json:"taxAmount"`
	// The amount of the direct incomes, inclusive of tax.
	TotalAmount float64 `json:"totalAmount"`
}

// CreateDirectIncome200ApplicationJSONStatusEnum - The status of the push operation.
type CreateDirectIncome200ApplicationJSONStatusEnum string

const (
	CreateDirectIncome200ApplicationJSONStatusEnumPending  CreateDirectIncome200ApplicationJSONStatusEnum = "Pending"
	CreateDirectIncome200ApplicationJSONStatusEnumFailed   CreateDirectIncome200ApplicationJSONStatusEnum = "Failed"
	CreateDirectIncome200ApplicationJSONStatusEnumSuccess  CreateDirectIncome200ApplicationJSONStatusEnum = "Success"
	CreateDirectIncome200ApplicationJSONStatusEnumTimedOut CreateDirectIncome200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateDirectIncome200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateDirectIncome200ApplicationJSONValidation - A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateDirectIncome200ApplicationJSONValidation struct {
	Errors   []CreateDirectIncome200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateDirectIncome200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

// CreateDirectIncome200ApplicationJSON - Success
type CreateDirectIncome200ApplicationJSON struct {
	Changes []CreateDirectIncome200ApplicationJSONChanges `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
	// The datetime when the push was completed, null if Pending.
	CompletedOnUtc *time.Time `json:"completedOnUtc,omitempty"`
	// > **Language tip:**  Direct incomes may also be referred to as **Receive transactions**, **Receive money transactions**, **Sales receipts**, or **Cash sales** in various accounting platforms.
	//
	// > View the coverage for direct incomes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes" target="_blank">Data coverage explorer</a>.
	//
	// ## Overview
	//
	// Direct incomes include:
	//
	// - Selling an item directly to a contact, and receiving payment at the point of the sale.
	// - Refunding an item in cash to a contact.
	// - Depositing money into a bank account.
	//
	// From the Direct Incomes endpoints, you can:
	//
	// - [Get a list of all direct incomes for a specific company](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes)
	// - [Get a single direct income for a specific company and connection](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes__directIncomeId_)
	// - [Add a new direct income to a specific company's accounting package](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/post_companies__companyId__connections__connectionId__push_directIncomes)
	//
	// Direct incomes is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
	//
	Data *CreateDirectIncome200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionKey string `json:"dataConnectionKey"`
	// The type of data being pushed, eg invoices, customers.
	DataType     *string `json:"dataType,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationKey string `json:"pushOperationKey"`
	// The datetime when the push was requested.
	RequestedOnUtc time.Time `json:"requestedOnUtc"`
	// The status of the push operation.
	Status           CreateDirectIncome200ApplicationJSONStatusEnum `json:"status"`
	StatusCode       int                                            `json:"statusCode"`
	TimeoutInMinutes *int                                           `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds *int                                           `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *CreateDirectIncome200ApplicationJSONValidation `json:"validation,omitempty"`
}

type CreateDirectIncomeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	CreateDirectIncome200ApplicationJSONObject *CreateDirectIncome200ApplicationJSON
}
