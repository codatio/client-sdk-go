// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SyncSettings struct {
	// Unique identifier for your client in Codat.
	ClientID *string `json:"clientId,omitempty"`
	// Set to `True` if you want to override the default [sync settings](https://docs.codat.io/knowledge-base/advanced-sync-settings).
	OverridesDefaults *bool         `json:"overridesDefaults,omitempty"`
	Settings          []SyncSetting `json:"settings,omitempty"`
}

func (o *SyncSettings) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SyncSettings) GetOverridesDefaults() *bool {
	if o == nil {
		return nil
	}
	return o.OverridesDefaults
}

func (o *SyncSettings) GetSettings() []SyncSetting {
	if o == nil {
		return nil
	}
	return o.Settings
}
