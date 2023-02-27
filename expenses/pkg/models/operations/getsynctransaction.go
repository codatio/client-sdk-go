package operations

type GetSyncTransactionPathParams struct {
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	SyncID        string `pathParam:"style=simple,explode=false,name=syncId"`
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

type GetSyncTransactionRequest struct {
	PathParams GetSyncTransactionPathParams
}

type GetSyncTransaction200TextJSONIntegrationTypeEnum string

const (
	GetSyncTransaction200TextJSONIntegrationTypeEnumExpenses  GetSyncTransaction200TextJSONIntegrationTypeEnum = "expenses"
	GetSyncTransaction200TextJSONIntegrationTypeEnumBankfeeds GetSyncTransaction200TextJSONIntegrationTypeEnum = "bankfeeds"
)

type GetSyncTransaction200TextJSONStatusEnum string

const (
	GetSyncTransaction200TextJSONStatusEnumUnknown         GetSyncTransaction200TextJSONStatusEnum = "Unknown"
	GetSyncTransaction200TextJSONStatusEnumPending         GetSyncTransaction200TextJSONStatusEnum = "Pending"
	GetSyncTransaction200TextJSONStatusEnumValidationError GetSyncTransaction200TextJSONStatusEnum = "ValidationError"
	GetSyncTransaction200TextJSONStatusEnumCompleted       GetSyncTransaction200TextJSONStatusEnum = "Completed"
	GetSyncTransaction200TextJSONStatusEnumPushError       GetSyncTransaction200TextJSONStatusEnum = "PushError"
)

type GetSyncTransaction200TextJSON struct {
	IntegrationType *GetSyncTransaction200TextJSONIntegrationTypeEnum `json:"integrationType,omitempty"`
	Message         *string                                           `json:"message,omitempty"`
	Status          *GetSyncTransaction200TextJSONStatusEnum          `json:"status,omitempty"`
	TransactionID   *string                                           `json:"transactionId,omitempty"`
}

type GetSyncTransaction200ApplicationJSONIntegrationTypeEnum string

const (
	GetSyncTransaction200ApplicationJSONIntegrationTypeEnumExpenses  GetSyncTransaction200ApplicationJSONIntegrationTypeEnum = "expenses"
	GetSyncTransaction200ApplicationJSONIntegrationTypeEnumBankfeeds GetSyncTransaction200ApplicationJSONIntegrationTypeEnum = "bankfeeds"
)

type GetSyncTransaction200ApplicationJSONStatusEnum string

const (
	GetSyncTransaction200ApplicationJSONStatusEnumUnknown         GetSyncTransaction200ApplicationJSONStatusEnum = "Unknown"
	GetSyncTransaction200ApplicationJSONStatusEnumPending         GetSyncTransaction200ApplicationJSONStatusEnum = "Pending"
	GetSyncTransaction200ApplicationJSONStatusEnumValidationError GetSyncTransaction200ApplicationJSONStatusEnum = "ValidationError"
	GetSyncTransaction200ApplicationJSONStatusEnumCompleted       GetSyncTransaction200ApplicationJSONStatusEnum = "Completed"
	GetSyncTransaction200ApplicationJSONStatusEnumPushError       GetSyncTransaction200ApplicationJSONStatusEnum = "PushError"
)

type GetSyncTransaction200ApplicationJSON struct {
	IntegrationType *GetSyncTransaction200ApplicationJSONIntegrationTypeEnum `json:"integrationType,omitempty"`
	Message         *string                                                  `json:"message,omitempty"`
	Status          *GetSyncTransaction200ApplicationJSONStatusEnum          `json:"status,omitempty"`
	TransactionID   *string                                                  `json:"transactionId,omitempty"`
}

type GetSyncTransactionResponse struct {
	ContentType                                 string
	StatusCode                                  int
	GetSyncTransaction200ApplicationJSONObjects []GetSyncTransaction200ApplicationJSON
	GetSyncTransaction200TextJSONObjects        []GetSyncTransaction200TextJSON
	GetSyncTransaction200TextPlainArray         *string
}
