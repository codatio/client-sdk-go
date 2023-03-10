package operations

import (
	"net/http"
	"time"
)

type GetTransferPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	TransferID   string `pathParam:"style=simple,explode=false,name=transferId"`
}

type GetTransferRequest struct {
	PathParams GetTransferPathParams
}

// GetTransferSourceModifiedDateContactRef
// The customer or supplier for the transfer, if available.
type GetTransferSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// GetTransferSourceModifiedDateTransferAccountAccountRef
// The account that the transfer is moving from or to.
type GetTransferSourceModifiedDateTransferAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetTransferSourceModifiedDateTransferAccount
// The details of the accounts the transfer is moving from.
type GetTransferSourceModifiedDateTransferAccount struct {
	AccountRef *GetTransferSourceModifiedDateTransferAccountAccountRef `json:"accountRef,omitempty"`
	Amount     *float64                                                `json:"amount,omitempty"`
	Currency   *string                                                 `json:"currency,omitempty"`
}

type GetTransferSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// GetTransferSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetTransferSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type GetTransferSourceModifiedDateTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetTransferSourceModifiedDate
// > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type GetTransferSourceModifiedDate struct {
	ContactRef           *GetTransferSourceModifiedDateContactRef            `json:"contactRef,omitempty"`
	Date                 *time.Time                                          `json:"date,omitempty"`
	DepositedRecordRefs  []string                                            `json:"depositedRecordRefs,omitempty"`
	Description          *string                                             `json:"description,omitempty"`
	From                 *GetTransferSourceModifiedDateTransferAccount       `json:"from,omitempty"`
	ID                   *string                                             `json:"id,omitempty"`
	Metadata             *GetTransferSourceModifiedDateMetadata              `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                          `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                                          `json:"sourceModifiedDate,omitempty"`
	SupplementalData     *GetTransferSourceModifiedDateSupplementalData      `json:"supplementalData,omitempty"`
	To                   *GetTransferSourceModifiedDateTransferAccount       `json:"to,omitempty"`
	TrackingCategoryRefs []GetTransferSourceModifiedDateTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}

type GetTransferResponse struct {
	ContentType        string
	SourceModifiedDate *GetTransferSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
