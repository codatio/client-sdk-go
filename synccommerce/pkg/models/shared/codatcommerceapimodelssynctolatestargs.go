package shared

import (
	"time"
)

type CodatCommerceAPIModelsSyncToLatestArgs struct {
	SyncTo *time.Time `json:"syncTo,omitempty"`
}
