// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UpdateConnectionStatus struct {
	// The current authorization status of the data connection.
	Status *DataConnectionStatus `json:"status,omitempty"`
}

func (o *UpdateConnectionStatus) GetStatus() *DataConnectionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
