package operations

import (
	"net/http"
)

type CreateBillAttachmentsPathParams struct {
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateBillAttachmentsRequest struct {
	PathParams CreateBillAttachmentsPathParams
}

type CreateBillAttachmentsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
