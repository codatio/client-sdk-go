package operations

import (
	"net/http"
)

type CreateBillAttachmentsRequest struct {
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateBillAttachmentsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
