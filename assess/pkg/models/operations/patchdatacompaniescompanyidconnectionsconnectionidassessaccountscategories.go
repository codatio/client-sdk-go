package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesRequestBodyCategoriesAccountRef struct {
	ID string `json:"id"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesRequestBodyCategories struct {
	AccountRef *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesRequestBodyCategoriesAccountRef `json:"accountRef,omitempty"`
	Confirmed  *interface{}                                                                                               `json:"confirmed,omitempty"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesRequestBody struct {
	Categories []PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesRequestBodyCategories `json:"categories,omitempty"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesRequest struct {
	PathParams PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesPathParams
	Request    *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesRequestBody `request:"mediaType=application/json"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesResponse struct {
	CategorisedAccounts []shared.CategorisedAccount
	ContentType         string
	StatusCode          int
}
