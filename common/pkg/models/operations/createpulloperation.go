package operations

import (
	"time"
)

type CreatePullOperationDataTypeEnum string

const (
	CreatePullOperationDataTypeEnumInvoices             CreatePullOperationDataTypeEnum = "invoices"
	CreatePullOperationDataTypeEnumAccounts             CreatePullOperationDataTypeEnum = "accounts"
	CreatePullOperationDataTypeEnumCommercePayments     CreatePullOperationDataTypeEnum = "commerce-payments"
	CreatePullOperationDataTypeEnumBankingAccounts      CreatePullOperationDataTypeEnum = "banking-accounts"
	CreatePullOperationDataTypeEnumCompany              CreatePullOperationDataTypeEnum = "company"
	CreatePullOperationDataTypeEnumProfitAndLoss        CreatePullOperationDataTypeEnum = "profitAndLoss"
	CreatePullOperationDataTypeEnumCommerceTransactions CreatePullOperationDataTypeEnum = "commerce-transactions"
	CreatePullOperationDataTypeEnumBills                CreatePullOperationDataTypeEnum = "bills"
	CreatePullOperationDataTypeEnumCustomers            CreatePullOperationDataTypeEnum = "customers"
)

type CreatePullOperationPathParams struct {
	CompanyID string                          `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  CreatePullOperationDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type CreatePullOperationQueryParams struct {
	ConnectionID *string `queryParam:"style=form,explode=true,name=connectionId"`
}

type CreatePullOperationRequest struct {
	PathParams  CreatePullOperationPathParams
	QueryParams CreatePullOperationQueryParams
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
	CreatePullOperation401ApplicationJSONObject *CreatePullOperation401ApplicationJSON
	CreatePullOperation404ApplicationJSONObject *CreatePullOperation404ApplicationJSON
}
