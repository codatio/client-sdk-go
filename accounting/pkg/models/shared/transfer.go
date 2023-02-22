package shared

import (
	"time"
)

// TransferContactRef
// The customer or supplier for the transfer, if available.
type TransferContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// TransferTransferAccount
// The details of the accounts the transfer is moving to.
type TransferTransferAccount struct {
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	Amount     *float64    `json:"amount,omitempty"`
	Currency   *string     `json:"currency,omitempty"`
}

type TransferTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// Transfer
// > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/docs/datamodel-accounting-account-transactions).
type Transfer struct {
	ContactRef           *TransferContactRef            `json:"contactRef,omitempty"`
	Date                 *time.Time                     `json:"date,omitempty"`
	DepositedRecordRefs  []string                       `json:"depositedRecordRefs,omitempty"`
	Description          *string                        `json:"description,omitempty"`
	From                 *To                            `json:"from,omitempty"`
	ID                   *string                        `json:"id,omitempty"`
	Metadata             *Metadata                      `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                     `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                     `json:"sourceModifiedDate,omitempty"`
	SupplementalData     *SupplementalData              `json:"supplementalData,omitempty"`
	To                   *TransferTransferAccount       `json:"to,omitempty"`
	TrackingCategoryRefs []TransferTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}
