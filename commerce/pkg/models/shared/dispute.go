package shared

import (
	"time"
)

type DisputeRecordRef struct {
	ID   string      `json:"id"`
	Type interface{} `json:"type"`
}

type DisputeStatusEnum string

const (
	DisputeStatusEnumWon                     DisputeStatusEnum = "Won"
	DisputeStatusEnumLost                    DisputeStatusEnum = "Lost"
	DisputeStatusEnumAccepted                DisputeStatusEnum = "Accepted"
	DisputeStatusEnumProcessing              DisputeStatusEnum = "Processing"
	DisputeStatusEnumChargeRefunded          DisputeStatusEnum = "ChargeRefunded"
	DisputeStatusEnumEvidenceRequired        DisputeStatusEnum = "EvidenceRequired"
	DisputeStatusEnumInquiryEvidenceRequired DisputeStatusEnum = "InquiryEvidenceRequired"
	DisputeStatusEnumInquiryProcessing       DisputeStatusEnum = "InquiryProcessing"
	DisputeStatusEnumInquiryClosed           DisputeStatusEnum = "InquiryClosed"
	DisputeStatusEnumWaitingThirdParty       DisputeStatusEnum = "WaitingThirdParty"
	DisputeStatusEnumUnknown                 DisputeStatusEnum = "Unknown"
)

type Dispute struct {
	CreatedDate          *time.Time         `json:"createdDate,omitempty"`
	Currency             string             `json:"currency"`
	DisputedTransactions *DisputeRecordRef  `json:"disputedTransactions,omitempty"`
	DueDate              *time.Time         `json:"dueDate,omitempty"`
	ID                   string             `json:"id"`
	ModifiedDate         *time.Time         `json:"modifiedDate,omitempty"`
	Reason               *string            `json:"reason,omitempty"`
	SourceModifiedDate   *time.Time         `json:"sourceModifiedDate,omitempty"`
	Status               *DisputeStatusEnum `json:"status,omitempty"`
	TotalAmount          *float64           `json:"totalAmount,omitempty"`
}
