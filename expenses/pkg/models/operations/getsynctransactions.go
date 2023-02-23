package operations

type GetSyncTransactionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	SyncID    string `pathParam:"style=simple,explode=false,name=syncId"`
}

type GetSyncTransactionsQueryParams struct {
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
}

type GetSyncTransactionsRequest struct {
	PathParams  GetSyncTransactionsPathParams
	QueryParams GetSyncTransactionsQueryParams
}

type GetSyncTransactions200TextJSONHalLink struct {
	Href *string `json:"href,omitempty"`
}

type GetSyncTransactions200TextJSONResultsIntegrationTypeEnum string

const (
	GetSyncTransactions200TextJSONResultsIntegrationTypeEnumExpenses  GetSyncTransactions200TextJSONResultsIntegrationTypeEnum = "expenses"
	GetSyncTransactions200TextJSONResultsIntegrationTypeEnumBankfeeds GetSyncTransactions200TextJSONResultsIntegrationTypeEnum = "bankfeeds"
)

type GetSyncTransactions200TextJSONResultsStatusEnum string

const (
	GetSyncTransactions200TextJSONResultsStatusEnumUnknown         GetSyncTransactions200TextJSONResultsStatusEnum = "Unknown"
	GetSyncTransactions200TextJSONResultsStatusEnumPending         GetSyncTransactions200TextJSONResultsStatusEnum = "Pending"
	GetSyncTransactions200TextJSONResultsStatusEnumValidationError GetSyncTransactions200TextJSONResultsStatusEnum = "ValidationError"
	GetSyncTransactions200TextJSONResultsStatusEnumCompleted       GetSyncTransactions200TextJSONResultsStatusEnum = "Completed"
	GetSyncTransactions200TextJSONResultsStatusEnumPushError       GetSyncTransactions200TextJSONResultsStatusEnum = "PushError"
)

type GetSyncTransactions200TextJSONResults struct {
	IntegrationType *GetSyncTransactions200TextJSONResultsIntegrationTypeEnum `json:"integrationType,omitempty"`
	Message         *string                                                   `json:"message,omitempty"`
	Status          *GetSyncTransactions200TextJSONResultsStatusEnum          `json:"status,omitempty"`
	TransactionID   *string                                                   `json:"transactionId,omitempty"`
}

type GetSyncTransactions200TextJSON struct {
	Links        map[string]GetSyncTransactions200TextJSONHalLink `json:"links,omitempty"`
	PageNumber   int                                              `json:"pageNumber"`
	PageSize     int                                              `json:"pageSize"`
	Results      []GetSyncTransactions200TextJSONResults          `json:"results,omitempty"`
	TotalResults int                                              `json:"totalResults"`
}

type GetSyncTransactions200ApplicationJSONHalLink struct {
	Href *string `json:"href,omitempty"`
}

type GetSyncTransactions200ApplicationJSONResultsIntegrationTypeEnum string

const (
	GetSyncTransactions200ApplicationJSONResultsIntegrationTypeEnumExpenses  GetSyncTransactions200ApplicationJSONResultsIntegrationTypeEnum = "expenses"
	GetSyncTransactions200ApplicationJSONResultsIntegrationTypeEnumBankfeeds GetSyncTransactions200ApplicationJSONResultsIntegrationTypeEnum = "bankfeeds"
)

type GetSyncTransactions200ApplicationJSONResultsStatusEnum string

const (
	GetSyncTransactions200ApplicationJSONResultsStatusEnumUnknown         GetSyncTransactions200ApplicationJSONResultsStatusEnum = "Unknown"
	GetSyncTransactions200ApplicationJSONResultsStatusEnumPending         GetSyncTransactions200ApplicationJSONResultsStatusEnum = "Pending"
	GetSyncTransactions200ApplicationJSONResultsStatusEnumValidationError GetSyncTransactions200ApplicationJSONResultsStatusEnum = "ValidationError"
	GetSyncTransactions200ApplicationJSONResultsStatusEnumCompleted       GetSyncTransactions200ApplicationJSONResultsStatusEnum = "Completed"
	GetSyncTransactions200ApplicationJSONResultsStatusEnumPushError       GetSyncTransactions200ApplicationJSONResultsStatusEnum = "PushError"
)

type GetSyncTransactions200ApplicationJSONResults struct {
	IntegrationType *GetSyncTransactions200ApplicationJSONResultsIntegrationTypeEnum `json:"integrationType,omitempty"`
	Message         *string                                                          `json:"message,omitempty"`
	Status          *GetSyncTransactions200ApplicationJSONResultsStatusEnum          `json:"status,omitempty"`
	TransactionID   *string                                                          `json:"transactionId,omitempty"`
}

type GetSyncTransactions200ApplicationJSON struct {
	Links        map[string]GetSyncTransactions200ApplicationJSONHalLink `json:"links,omitempty"`
	PageNumber   int                                                     `json:"pageNumber"`
	PageSize     int                                                     `json:"pageSize"`
	Results      []GetSyncTransactions200ApplicationJSONResults          `json:"results,omitempty"`
	TotalResults int                                                     `json:"totalResults"`
}

type GetSyncTransactionsResponse struct {
	ContentType                                 string
	StatusCode                                  int
	GetSyncTransactions200ApplicationJSONObject *GetSyncTransactions200ApplicationJSON
	GetSyncTransactions200TextJSONObject        *GetSyncTransactions200TextJSON
	GetSyncTransactions200TextPlainObject       *string
}
