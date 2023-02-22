package shared

type CustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}
