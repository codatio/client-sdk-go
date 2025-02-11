// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CommerceReportError struct {
	// Additional details on the error.
	Details map[string][]string `json:"details,omitempty"`
	// Message returned by error.
	Message *string `json:"message,omitempty"`
	// The type of error.
	Type *string `json:"type,omitempty"`
}

func (o *CommerceReportError) GetDetails() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *CommerceReportError) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CommerceReportError) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
