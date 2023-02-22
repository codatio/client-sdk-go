package shared

type TransactionMetadataPagedResponse struct {
	Links        map[string]HalLink    `json:"links,omitempty"`
	PageNumber   int                   `json:"pageNumber"`
	PageSize     int                   `json:"pageSize"`
	Results      []TransactionMetadata `json:"results,omitempty"`
	TotalResults int                   `json:"totalResults"`
}
