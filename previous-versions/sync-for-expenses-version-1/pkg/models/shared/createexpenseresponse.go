// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateExpenseResponse struct {
	// Unique id of dataset created
	DatasetID *string `json:"datasetId,omitempty"`
}

func (o *CreateExpenseResponse) GetDatasetID() *string {
	if o == nil {
		return nil
	}
	return o.DatasetID
}
