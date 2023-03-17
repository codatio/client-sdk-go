package operations

import (
	"net/http"
	"time"
)

type GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponentsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
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
	RawResponse                                                                                   *http.Response
	GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSONObject *GetCompaniesCompanyIDConnectionsConnectionIDDataCommerceTaxComponents200ApplicationJSON
}
