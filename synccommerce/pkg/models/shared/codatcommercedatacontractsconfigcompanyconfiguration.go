package shared

type CodatCommerceDataContractsConfigCompanyConfiguration struct {
	AccountingSoftwareCompanyName *string                                        `json:"accountingSoftwareCompanyName,omitempty"`
	CompanyID                     *string                                        `json:"companyId,omitempty"`
	Configuration                 *CodatCommerceDataContractsConfigConfiguration `json:"configuration,omitempty"`
	Configured                    *bool                                          `json:"configured,omitempty"`
	Enabled                       *bool                                          `json:"enabled,omitempty"`
	Schedule                      *CodatCommerceDataContractsConfigSchedule      `json:"schedule,omitempty"`
}
