// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DataIntegrityStatuses struct {
	Metadata []DataIntegrityStatus `json:"metadata,omitempty"`
}

func (o *DataIntegrityStatuses) GetMetadata() []DataIntegrityStatus {
	if o == nil {
		return nil
	}
	return o.Metadata
}
