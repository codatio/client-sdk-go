package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetItemPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	ItemID    string `pathParam:"style=simple,explode=false,name=itemId"`
}

type GetItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetItemRequest struct {
	PathParams GetItemPathParams
	Security   GetItemSecurity
}

type GetItemResponse struct {
	ContentType string
	Item        *shared.Item
	StatusCode  int
}
