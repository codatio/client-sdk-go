package operations

import (
	"net/http"
	"time"
)

type CreatePullOperationDataTypeEnum string

const (
	CreatePullOperationDataTypeEnumAccountTransactions          CreatePullOperationDataTypeEnum = "accountTransactions"
	CreatePullOperationDataTypeEnumBalanceSheet                 CreatePullOperationDataTypeEnum = "balanceSheet"
	CreatePullOperationDataTypeEnumBankAccounts                 CreatePullOperationDataTypeEnum = "bankAccounts"
	CreatePullOperationDataTypeEnumBankTransactions             CreatePullOperationDataTypeEnum = "bankTransactions"
	CreatePullOperationDataTypeEnumBillCreditNotes              CreatePullOperationDataTypeEnum = "billCreditNotes"
	CreatePullOperationDataTypeEnumBillPayments                 CreatePullOperationDataTypeEnum = "billPayments"
	CreatePullOperationDataTypeEnumBills                        CreatePullOperationDataTypeEnum = "bills"
	CreatePullOperationDataTypeEnumCashFlowStatement            CreatePullOperationDataTypeEnum = "cashFlowStatement"
	CreatePullOperationDataTypeEnumChartOfAccounts              CreatePullOperationDataTypeEnum = "chartOfAccounts"
	CreatePullOperationDataTypeEnumCompany                      CreatePullOperationDataTypeEnum = "company"
	CreatePullOperationDataTypeEnumCreditNotes                  CreatePullOperationDataTypeEnum = "creditNotes"
	CreatePullOperationDataTypeEnumCustomers                    CreatePullOperationDataTypeEnum = "customers"
	CreatePullOperationDataTypeEnumDirectCosts                  CreatePullOperationDataTypeEnum = "directCosts"
	CreatePullOperationDataTypeEnumDirectIncomes                CreatePullOperationDataTypeEnum = "directIncomes"
	CreatePullOperationDataTypeEnumInvoices                     CreatePullOperationDataTypeEnum = "invoices"
	CreatePullOperationDataTypeEnumItems                        CreatePullOperationDataTypeEnum = "items"
	CreatePullOperationDataTypeEnumJournalEntries               CreatePullOperationDataTypeEnum = "journalEntries"
	CreatePullOperationDataTypeEnumJournals                     CreatePullOperationDataTypeEnum = "journals"
	CreatePullOperationDataTypeEnumPaymentMethods               CreatePullOperationDataTypeEnum = "paymentMethods"
	CreatePullOperationDataTypeEnumPayments                     CreatePullOperationDataTypeEnum = "payments"
	CreatePullOperationDataTypeEnumProfitAndLoss                CreatePullOperationDataTypeEnum = "profitAndLoss"
	CreatePullOperationDataTypeEnumProjects                     CreatePullOperationDataTypeEnum = "projects"
	CreatePullOperationDataTypeEnumPurchaseOrders               CreatePullOperationDataTypeEnum = "purchaseOrders"
	CreatePullOperationDataTypeEnumSalesOrders                  CreatePullOperationDataTypeEnum = "salesOrders"
	CreatePullOperationDataTypeEnumSuppliers                    CreatePullOperationDataTypeEnum = "suppliers"
	CreatePullOperationDataTypeEnumTaxRates                     CreatePullOperationDataTypeEnum = "taxRates"
	CreatePullOperationDataTypeEnumTrackingCategories           CreatePullOperationDataTypeEnum = "trackingCategories"
	CreatePullOperationDataTypeEnumTransfers                    CreatePullOperationDataTypeEnum = "transfers"
	CreatePullOperationDataTypeEnumBankingAccountBalances       CreatePullOperationDataTypeEnum = "banking-accountBalances"
	CreatePullOperationDataTypeEnumBankingAccounts              CreatePullOperationDataTypeEnum = "banking-accounts"
	CreatePullOperationDataTypeEnumBankingTransactionCategories CreatePullOperationDataTypeEnum = "banking-transactionCategories"
	CreatePullOperationDataTypeEnumBankingTransactions          CreatePullOperationDataTypeEnum = "banking-transactions"
	CreatePullOperationDataTypeEnumCommerceCompanyInfo          CreatePullOperationDataTypeEnum = "commerce-companyInfo"
	CreatePullOperationDataTypeEnumCommerceCustomers            CreatePullOperationDataTypeEnum = "commerce-customers"
	CreatePullOperationDataTypeEnumCommerceDisputes             CreatePullOperationDataTypeEnum = "commerce-disputes"
	CreatePullOperationDataTypeEnumCommerceLocations            CreatePullOperationDataTypeEnum = "commerce-locations"
	CreatePullOperationDataTypeEnumCommerceOrders               CreatePullOperationDataTypeEnum = "commerce-orders"
	CreatePullOperationDataTypeEnumCommercePaymentMethods       CreatePullOperationDataTypeEnum = "commerce-paymentMethods"
	CreatePullOperationDataTypeEnumCommercePayments             CreatePullOperationDataTypeEnum = "commerce-payments"
	CreatePullOperationDataTypeEnumCommerceProductCategories    CreatePullOperationDataTypeEnum = "commerce-productCategories"
	CreatePullOperationDataTypeEnumCommerceProducts             CreatePullOperationDataTypeEnum = "commerce-products"
	CreatePullOperationDataTypeEnumCommerceTaxComponents        CreatePullOperationDataTypeEnum = "commerce-taxComponents"
	CreatePullOperationDataTypeEnumCommerceTransactions         CreatePullOperationDataTypeEnum = "commerce-transactions"
)

type CreatePullOperationRequest struct {
	CompanyID    string                          `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID *string                         `queryParam:"style=form,explode=true,name=connectionId"`
	DataType     CreatePullOperationDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type CreatePullOperation404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type CreatePullOperation401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type CreatePullOperationPullOperationStatusEnum string

const (
	CreatePullOperationPullOperationStatusEnumInitial            CreatePullOperationPullOperationStatusEnum = "Initial"
	CreatePullOperationPullOperationStatusEnumQueued             CreatePullOperationPullOperationStatusEnum = "Queued"
	CreatePullOperationPullOperationStatusEnumFetching           CreatePullOperationPullOperationStatusEnum = "Fetching"
	CreatePullOperationPullOperationStatusEnumMapQueued          CreatePullOperationPullOperationStatusEnum = "MapQueued"
	CreatePullOperationPullOperationStatusEnumMapping            CreatePullOperationPullOperationStatusEnum = "Mapping"
	CreatePullOperationPullOperationStatusEnumComplete           CreatePullOperationPullOperationStatusEnum = "Complete"
	CreatePullOperationPullOperationStatusEnumFetchError         CreatePullOperationPullOperationStatusEnum = "FetchError"
	CreatePullOperationPullOperationStatusEnumMapError           CreatePullOperationPullOperationStatusEnum = "MapError"
	CreatePullOperationPullOperationStatusEnumInternalError      CreatePullOperationPullOperationStatusEnum = "InternalError"
	CreatePullOperationPullOperationStatusEnumProcessingQueued   CreatePullOperationPullOperationStatusEnum = "ProcessingQueued"
	CreatePullOperationPullOperationStatusEnumProcessing         CreatePullOperationPullOperationStatusEnum = "Processing"
	CreatePullOperationPullOperationStatusEnumProcessingError    CreatePullOperationPullOperationStatusEnum = "ProcessingError"
	CreatePullOperationPullOperationStatusEnumValidationQueued   CreatePullOperationPullOperationStatusEnum = "ValidationQueued"
	CreatePullOperationPullOperationStatusEnumValidating         CreatePullOperationPullOperationStatusEnum = "Validating"
	CreatePullOperationPullOperationStatusEnumValidationError    CreatePullOperationPullOperationStatusEnum = "ValidationError"
	CreatePullOperationPullOperationStatusEnumAuthError          CreatePullOperationPullOperationStatusEnum = "AuthError"
	CreatePullOperationPullOperationStatusEnumCancelled          CreatePullOperationPullOperationStatusEnum = "Cancelled"
	CreatePullOperationPullOperationStatusEnumRouting            CreatePullOperationPullOperationStatusEnum = "Routing"
	CreatePullOperationPullOperationStatusEnumRoutingError       CreatePullOperationPullOperationStatusEnum = "RoutingError"
	CreatePullOperationPullOperationStatusEnumNotSupported       CreatePullOperationPullOperationStatusEnum = "NotSupported"
	CreatePullOperationPullOperationStatusEnumRateLimitError     CreatePullOperationPullOperationStatusEnum = "RateLimitError"
	CreatePullOperationPullOperationStatusEnumPermissionsError   CreatePullOperationPullOperationStatusEnum = "PermissionsError"
	CreatePullOperationPullOperationStatusEnumPrerequisiteNotMet CreatePullOperationPullOperationStatusEnum = "PrerequisiteNotMet"
)

// CreatePullOperationPullOperation
// Information about a queued, in progress or completed pull operation.
// *Formally called `dataset`*
type CreatePullOperationPullOperation struct {
	CompanyID    string                                     `json:"companyId"`
	ConnectionID string                                     `json:"connectionId"`
	DataType     string                                     `json:"dataType"`
	ID           string                                     `json:"id"`
	IsCompleted  bool                                       `json:"isCompleted"`
	IsErrored    bool                                       `json:"isErrored"`
	Progress     int64                                      `json:"progress"`
	Requested    time.Time                                  `json:"requested"`
	Status       CreatePullOperationPullOperationStatusEnum `json:"status"`
}

type CreatePullOperationResponse struct {
	ContentType                                 string
	PullOperation                               *CreatePullOperationPullOperation
	StatusCode                                  int
	RawResponse                                 *http.Response
	CreatePullOperation401ApplicationJSONObject *CreatePullOperation401ApplicationJSON
	CreatePullOperation404ApplicationJSONObject *CreatePullOperation404ApplicationJSON
}
