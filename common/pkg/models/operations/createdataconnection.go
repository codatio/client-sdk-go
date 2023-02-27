package operations

type CreateDataConnectionPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type CreateDataConnectionRequest struct {
	PathParams CreateDataConnectionPathParams
}

type CreateDataConnectionResponse struct {
	ContentType string
	StatusCode  int
}
