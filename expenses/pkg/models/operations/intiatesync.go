package operations

type IntiateSyncPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type IntiateSyncRequestBody struct {
	DatasetIds []string `json:"datasetIds"`
}

type IntiateSyncRequest struct {
	PathParams IntiateSyncPathParams
	Request    *IntiateSyncRequestBody `request:"mediaType=application/json"`
}

type IntiateSync200ApplicationJSON struct {
	SyncID *string `json:"syncId,omitempty"`
}

type IntiateSyncResponse struct {
	ContentType                         string
	StatusCode                          int
	IntiateSync200ApplicationJSONObject *IntiateSync200ApplicationJSON
}
