// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UpdateConnection struct {
	// The current authorization status of the data connection.
	Status *DataConnectionStatus `json:"status,omitempty"`
}

func (o *UpdateConnection) GetStatus() *DataConnectionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
