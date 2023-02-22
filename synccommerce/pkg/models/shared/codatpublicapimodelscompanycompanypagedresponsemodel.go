package shared

// CodatPublicAPIModelsCompanyCompanyPagedResponseModel
// Used to represent what can be returned by an endpoint that supports paging.
// Usable with the [ProducesResponseType] attribute on a controller action.
type CodatPublicAPIModelsCompanyCompanyPagedResponseModel struct {
	Links        *CodatPublicAPIModelsCompanyCompanyPagedResponseLinksModel `json:"_links,omitempty"`
	PageNumber   *int                                                       `json:"pageNumber,omitempty"`
	PageSize     *int                                                       `json:"pageSize,omitempty"`
	Results      []CodatPublicAPIModelsCompanyCompany                       `json:"results,omitempty"`
	TotalResults *int                                                       `json:"totalResults,omitempty"`
}
