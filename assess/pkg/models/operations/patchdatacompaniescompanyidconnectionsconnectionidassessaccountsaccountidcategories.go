package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesRequestBody struct {
	Confirmed shared.AccountCategory `json:"confirmed"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesRequest struct {
	PathParams PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesPathParams
	Request    *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesRequestBody `request:"mediaType=application/json"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesResponse struct {
	CategorisedAccount *shared.CategorisedAccount
	ContentType        string
	StatusCode         int
}
