// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SyncFlowURL struct {
	// Sync flow URL.
	URL *string `json:"url,omitempty"`
}

func (o *SyncFlowURL) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
