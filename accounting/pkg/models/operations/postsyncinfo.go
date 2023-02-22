package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type PostSyncInfoPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PostSyncInfoSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostSyncInfoRequest struct {
	PathParams PostSyncInfoPathParams
	Security   PostSyncInfoSecurity
}

type PostSyncInfoResponse struct {
	ContentType string
	DataSet     *shared.DataSet
	StatusCode  int
}
