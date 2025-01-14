// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DatasetStatusChangedErrorWebhookData struct {
	// Available data types
	DataType *DataType `json:"dataType,omitempty"`
	// Unique identifier for the dataset that completed its sync.
	DatasetID *string `json:"datasetId,omitempty"`
	// The current status of the dataset's sync.
	DatasetStatus *string `json:"datasetStatus,omitempty"`
}

func (o *DatasetStatusChangedErrorWebhookData) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *DatasetStatusChangedErrorWebhookData) GetDatasetID() *string {
	if o == nil {
		return nil
	}
	return o.DatasetID
}

func (o *DatasetStatusChangedErrorWebhookData) GetDatasetStatus() *string {
	if o == nil {
		return nil
	}
	return o.DatasetStatus
}
