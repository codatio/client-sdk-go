// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// StartUploadSessionRequestDataType - A key for a Codat data type.
type StartUploadSessionRequestDataType string

const (
	StartUploadSessionRequestDataTypeBankingAccounts     StartUploadSessionRequestDataType = "banking-accounts"
	StartUploadSessionRequestDataTypeBankingTransactions StartUploadSessionRequestDataType = "banking-transactions"
)

func (e StartUploadSessionRequestDataType) ToPointer() *StartUploadSessionRequestDataType {
	return &e
}
func (e *StartUploadSessionRequestDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "banking-accounts":
		fallthrough
	case "banking-transactions":
		*e = StartUploadSessionRequestDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StartUploadSessionRequestDataType: %v", v)
	}
}

type StartUploadSessionRequest struct {
	// A key for a Codat data type.
	DataType *StartUploadSessionRequestDataType `json:"dataType,omitempty"`
}

func (o *StartUploadSessionRequest) GetDataType() *StartUploadSessionRequestDataType {
	if o == nil {
		return nil
	}
	return o.DataType
}
