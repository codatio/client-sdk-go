package operations

import (
	"time"
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

// PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesCategorisedAccountAccountRef
// An object containing account reference data.
type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesCategorisedAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesCategorisedAccountModifiedDate struct {
	DetailType   *string    `json:"detailType,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Subtype      *string    `json:"subtype,omitempty"`
	Type         *string    `json:"type,omitempty"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesCategorisedAccount struct {
	AccountRef *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesCategorisedAccountAccountRef   `json:"accountRef,omitempty"`
	Confirmed  *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesCategorisedAccountModifiedDate `json:"confirmed,omitempty"`
	Suggested  *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesCategorisedAccountModifiedDate `json:"suggested,omitempty"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesResponse struct {
	CategorisedAccounts []PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesCategorisedAccount
	ContentType         string
	StatusCode          int
}
