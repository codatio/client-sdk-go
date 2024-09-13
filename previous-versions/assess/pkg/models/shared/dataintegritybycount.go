// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type DataIntegrityByCount struct {
	// The percentage of records of the type specified in the route which have a match.
	MatchPercentage *decimal.Big `decimal:"number" json:"matchPercentage,omitempty"`
	// The number of records of the type specified in the route which do have a match.
	Matched *decimal.Big `decimal:"number" json:"matched,omitempty"`
	// The total of unmatched and matched.
	Total *decimal.Big `decimal:"number" json:"total,omitempty"`
	// The number of records of the type specified in the route which don't have a match.
	Unmatched *decimal.Big `decimal:"number" json:"unmatched,omitempty"`
}

func (d DataIntegrityByCount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DataIntegrityByCount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DataIntegrityByCount) GetMatchPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.MatchPercentage
}

func (o *DataIntegrityByCount) GetMatched() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Matched
}

func (o *DataIntegrityByCount) GetTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *DataIntegrityByCount) GetUnmatched() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Unmatched
}
