package shared

import (
	"time"
)

type CodatInfrastructureDateDateRange struct {
	Finish *time.Time `json:"finish,omitempty"`
	Start  *time.Time `json:"start,omitempty"`
}
