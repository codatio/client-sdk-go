package operations

import (
	"net/http"
)

type GetSyncFlowURLRequest struct {
	AccountingKey      string  `pathParam:"style=simple,explode=false,name=accountingKey"`
	CommerceKey        string  `pathParam:"style=simple,explode=false,name=commerceKey"`
	MerchantIdentifier *string `queryParam:"style=form,explode=true,name=merchantIdentifier"`
}

type GetSyncFlowURLResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
