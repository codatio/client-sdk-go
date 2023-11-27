// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/utils"
	"net/http"
)

type UpdateBillRequest struct {
	Bill *shared.Bill `request:"mediaType=application/json"`
	// Unique identifier for a bill.
	BillID string `pathParam:"style=simple,explode=false,name=billId"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// When updating data in the destination platform Codat checks the `sourceModifiedDate` against the `lastupdated` date from the accounting platform, if they're different Codat will return an error suggesting you should initiate another pull of the data. If this is set to `true` then the update will override this check.
	ForceUpdate *bool `default:"false" queryParam:"style=form,explode=true,name=forceUpdate"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (u UpdateBillRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateBillRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateBillRequest) GetBill() *shared.Bill {
	if o == nil {
		return nil
	}
	return o.Bill
}

func (o *UpdateBillRequest) GetBillID() string {
	if o == nil {
		return ""
	}
	return o.BillID
}

func (o *UpdateBillRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateBillRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateBillRequest) GetForceUpdate() *bool {
	if o == nil {
		return nil
	}
	return o.ForceUpdate
}

func (o *UpdateBillRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type UpdateBillResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	UpdateBillResponse *shared.UpdateBillResponse
}

func (o *UpdateBillResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateBillResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateBillResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateBillResponse) GetUpdateBillResponse() *shared.UpdateBillResponse {
	if o == nil {
		return nil
	}
	return o.UpdateBillResponse
}
