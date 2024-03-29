// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SupplementalDataConfigurationSupplementalDataSourceConfiguration - The client's defined name for the object.
type SupplementalDataConfigurationSupplementalDataSourceConfiguration struct {
	// The underlying endpoint of the source system which the configuration is targeting.
	DataSource *string `json:"dataSource,omitempty"`
	// The additional properties that are required when pulling records.
	PullData map[string]string `json:"pullData,omitempty"`
	// The additional properties that are required to create and/or update records.
	PushData map[string]string `json:"pushData,omitempty"`
}

func (o *SupplementalDataConfigurationSupplementalDataSourceConfiguration) GetDataSource() *string {
	if o == nil {
		return nil
	}
	return o.DataSource
}

func (o *SupplementalDataConfigurationSupplementalDataSourceConfiguration) GetPullData() map[string]string {
	if o == nil {
		return nil
	}
	return o.PullData
}

func (o *SupplementalDataConfigurationSupplementalDataSourceConfiguration) GetPushData() map[string]string {
	if o == nil {
		return nil
	}
	return o.PushData
}

type SupplementalDataConfiguration struct {
	SupplementalDataConfig map[string]SupplementalDataConfigurationSupplementalDataSourceConfiguration `json:"supplementalDataConfig,omitempty"`
}

func (o *SupplementalDataConfiguration) GetSupplementalDataConfig() map[string]SupplementalDataConfigurationSupplementalDataSourceConfiguration {
	if o == nil {
		return nil
	}
	return o.SupplementalDataConfig
}
