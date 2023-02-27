package operations

type GetSyncFlowURLPathParams struct {
	AccountingKey string `pathParam:"style=simple,explode=false,name=accountingKey"`
	CommerceKey   string `pathParam:"style=simple,explode=false,name=commerceKey"`
}

type GetSyncFlowURLQueryParams struct {
	MerchantIdentifier *string `queryParam:"style=form,explode=true,name=merchantIdentifier"`
}

type GetSyncFlowURLRequest struct {
	PathParams  GetSyncFlowURLPathParams
	QueryParams GetSyncFlowURLQueryParams
}

type GetSyncFlowURLResponse struct {
	ContentType string
	StatusCode  int
}
