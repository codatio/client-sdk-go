package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataIntegrityStatusForDataTypePathParams struct {
	CompanyID string                           `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  shared.DataIntegritydataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type GetDataIntegrityStatusForDataTypeRequest struct {
	PathParams GetDataIntegrityStatusForDataTypePathParams
}

type GetDataIntegrityStatusForDataType200ApplicationJSON struct {
	Metadata []shared.DataIntegrityStatus `json:"metadata,omitempty"`
}

type GetDataIntegrityStatusForDataTypeResponse struct {
	ContentType                                               string
	StatusCode                                                int
	GetDataIntegrityStatusForDataType200ApplicationJSONObject *GetDataIntegrityStatusForDataType200ApplicationJSON
}
