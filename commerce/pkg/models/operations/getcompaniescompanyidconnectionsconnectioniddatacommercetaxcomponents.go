package operations

import (
	"time"
)

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponentsRequest struct {
	PathParams GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponentsPathParams
}

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSONSourceModifiedDate struct {
	ID                 string     `json:"id"`
	IsCompound         *bool      `json:"isCompound,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               string     `json:"name"`
	Rate               *float32   `json:"rate,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSON struct {
	TaxComponents []GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSONSourceModifiedDate `json:"taxComponents,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponentsResponse struct {
	ContentType                                                                                   string
	StatusCode                                                                                    int
	GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSONObject *GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSON
}
