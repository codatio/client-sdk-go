// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/types"
)

type DataIntegrityByCount struct {
	// The percentage of records of the type specified in the route which have a match.
	MatchPercentage *types.Decimal `json:"matchPercentage,omitempty"`
	// The number of records of the type specified in the route which do have a match.
	Matched *types.Decimal `json:"matched,omitempty"`
	// The total of unmatched and matched.
	Total *types.Decimal `json:"total,omitempty"`
	// The number of records of the type specified in the route which don't have a match.
	Unmatched *types.Decimal `json:"unmatched,omitempty"`
}

func (o *DataIntegrityByCount) GetMatchPercentage() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.MatchPercentage
}

func (o *DataIntegrityByCount) GetMatched() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Matched
}

func (o *DataIntegrityByCount) GetTotal() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *DataIntegrityByCount) GetUnmatched() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Unmatched
}
