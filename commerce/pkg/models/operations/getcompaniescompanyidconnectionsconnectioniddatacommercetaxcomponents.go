package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponentsRequest struct {
	PathParams GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponentsPathParams
}

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSON struct {
	TaxComponents []shared.TaxComponent `json:"taxComponents,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponentsResponse struct {
	ContentType                                                                                   string
	StatusCode                                                                                    int
	GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSONObject *GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSON
}
