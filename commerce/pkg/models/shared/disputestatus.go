// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DisputeStatus - Current status of the dispute
type DisputeStatus string

const (
	DisputeStatusWon                     DisputeStatus = "Won"
	DisputeStatusLost                    DisputeStatus = "Lost"
	DisputeStatusAccepted                DisputeStatus = "Accepted"
	DisputeStatusProcessing              DisputeStatus = "Processing"
	DisputeStatusChargeRefunded          DisputeStatus = "ChargeRefunded"
	DisputeStatusEvidenceRequired        DisputeStatus = "EvidenceRequired"
	DisputeStatusInquiryEvidenceRequired DisputeStatus = "InquiryEvidenceRequired"
	DisputeStatusInquiryProcessing       DisputeStatus = "InquiryProcessing"
	DisputeStatusInquiryClosed           DisputeStatus = "InquiryClosed"
	DisputeStatusWaitingThirdParty       DisputeStatus = "WaitingThirdParty"
	DisputeStatusUnknown                 DisputeStatus = "Unknown"
)

func (e DisputeStatus) ToPointer() *DisputeStatus {
	return &e
}

func (e *DisputeStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Won":
		fallthrough
	case "Lost":
		fallthrough
	case "Accepted":
		fallthrough
	case "Processing":
		fallthrough
	case "ChargeRefunded":
		fallthrough
	case "EvidenceRequired":
		fallthrough
	case "InquiryEvidenceRequired":
		fallthrough
	case "InquiryProcessing":
		fallthrough
	case "InquiryClosed":
		fallthrough
	case "WaitingThirdParty":
		fallthrough
	case "Unknown":
		*e = DisputeStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DisputeStatus: %v", v)
	}
}
