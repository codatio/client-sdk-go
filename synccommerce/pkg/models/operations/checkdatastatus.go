package operations

type CheckDataStatusPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	DatasetID string `pathParam:"style=simple,explode=false,name=datasetId"`
}

type CheckDataStatusRequest struct {
	PathParams CheckDataStatusPathParams
}

type CheckDataStatusResponse struct {
	ContentType string
	StatusCode  int
}
