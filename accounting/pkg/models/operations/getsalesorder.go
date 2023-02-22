package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetSalesOrderPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	SalesOrderID string `pathParam:"style=simple,explode=false,name=salesOrderId"`
}

type GetSalesOrderSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetSalesOrderRequest struct {
	PathParams GetSalesOrderPathParams
	Security   GetSalesOrderSecurity
}

type GetSalesOrderResponse struct {
	ContentType string
	SalesOrder  *shared.SalesOrder
	StatusCode  int
}
