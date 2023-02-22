package shared

import (
	"time"
)

type CodatPublicAPIModelsCompanyCompany struct {
	Created           *time.Time                                  `json:"created,omitempty"`
	CreatedByUserName *string                                     `json:"createdByUserName,omitempty"`
	DataConnections   []CodatPublicAPIModelsCompanyDataConnection `json:"dataConnections"`
	ID                string                                      `json:"id"`
	LastSync          *time.Time                                  `json:"lastSync,omitempty"`
	Name              string                                      `json:"name"`
	Platform          string                                      `json:"platform"`
	Redirect          string                                      `json:"redirect"`
}
