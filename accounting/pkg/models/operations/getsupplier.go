package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetSupplierPathParams struct {
	CompanyID  string `pathParam:"style=simple,explode=false,name=companyId"`
	SupplierID string `pathParam:"style=simple,explode=false,name=supplierId"`
}

type GetSupplierSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetSupplierRequest struct {
	PathParams GetSupplierPathParams
	Security   GetSupplierSecurity
}

type GetSupplierResponse struct {
	ContentType string
	StatusCode  int
	Supplier    *shared.Supplier
}
