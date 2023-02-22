package shared

import (
	"time"
)

type CodatCommerceDataContractsModelsCommerceSyncCreateSyncResponse struct {
	CommerceSyncID       *string                                     `json:"commerceSyncId,omitempty"`
	CompanyID            *string                                     `json:"companyId,omitempty"`
	DataConnections      []CodatPublicAPIModelsCompanyDataConnection `json:"dataConnections,omitempty"`
	DataPushed           *bool                                       `json:"dataPushed,omitempty"`
	ErrorMessage         *string                                     `json:"errorMessage,omitempty"`
	SyncDateRangeUtc     *CodatInfrastructureDateDateRange           `json:"syncDateRangeUtc,omitempty"`
	SyncExceptionMessage *string                                     `json:"syncExceptionMessage,omitempty"`
	SyncStatus           *string                                     `json:"syncStatus,omitempty"`
	SyncStatusCode       *int                                        `json:"syncStatusCode,omitempty"`
	SyncUtc              *time.Time                                  `json:"syncUtc,omitempty"`
}
