package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetPurchaseOrderPathParams struct {
	CompanyID       string `pathParam:"style=simple,explode=false,name=companyId"`
	PurchaseOrderID string `pathParam:"style=simple,explode=false,name=purchaseOrderId"`
}

type GetPurchaseOrderSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetPurchaseOrderRequest struct {
	PathParams GetPurchaseOrderPathParams
	Security   GetPurchaseOrderSecurity
}

type GetPurchaseOrderResponse struct {
	ContentType   string
	PurchaseOrder *shared.PurchaseOrder
	StatusCode    int
}
