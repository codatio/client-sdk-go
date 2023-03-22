// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type GetTransferRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	TransferID   string `pathParam:"style=simple,explode=false,name=transferId"`
}

// GetTransferSourceModifiedDateContactRef - The customer or supplier for the transfer, if available.
type GetTransferSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// GetTransferSourceModifiedDateTransferAccountAccountRef - The account that the transfer is moving from or to.
type GetTransferSourceModifiedDateTransferAccountAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

// GetTransferSourceModifiedDateTransferAccount - The details of the accounts the transfer is moving from.
type GetTransferSourceModifiedDateTransferAccount struct {
	// The account that the transfer is moving from or to.
	AccountRef *GetTransferSourceModifiedDateTransferAccountAccountRef `json:"accountRef,omitempty"`
	// The amount transferred between accounts.
	Amount *float64 `json:"amount,omitempty"`
	// ISO currency code recorded for the transfer in the accounting platform.
	Currency *string `json:"currency,omitempty"`
}

type GetTransferSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// GetTransferSourceModifiedDateSupplementalData - Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetTransferSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetTransferSourceModifiedDateTrackingCategoryRefs - References a category against which the item is tracked.
type GetTransferSourceModifiedDateTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetTransferSourceModifiedDate - > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type GetTransferSourceModifiedDate struct {
	// The customer or supplier for the transfer, if available.
	ContactRef *GetTransferSourceModifiedDateContactRef `json:"contactRef,omitempty"`
	// The day on which the transfer was made.
	Date                *time.Time `json:"date,omitempty"`
	DepositedRecordRefs []string   `json:"depositedRecordRefs,omitempty"`
	// Description of the transfer.
	Description *string `json:"description,omitempty"`
	// The details of the accounts the transfer is moving from.
	From *GetTransferSourceModifiedDateTransferAccount `json:"from,omitempty"`
	// Unique identifier for the transfer.
	ID       *string                                `json:"id,omitempty"`
	Metadata *GetTransferSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *GetTransferSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	// The details of the accounts the transfer is moving to.
	To *GetTransferSourceModifiedDateTransferAccount `json:"to,omitempty"`
	// Reference to the tracking categories this transfer is being tracked against.
	TrackingCategoryRefs []GetTransferSourceModifiedDateTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}

type GetTransferResponse struct {
	ContentType string
	// Success
	SourceModifiedDate *GetTransferSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
