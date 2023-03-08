package operations

import (
	"net/http"
	"time"
)

type GetBankTransactionCategoryPathParams struct {
	CompanyID             string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID          string `pathParam:"style=simple,explode=false,name=connectionId"`
	TransactionCategoryID string `pathParam:"style=simple,explode=false,name=transactionCategoryId"`
}

type GetBankTransactionCategoryRequest struct {
	PathParams GetBankTransactionCategoryPathParams
}

type GetBankTransactionCategorySourceModifiedDateStatusEnum string

const (
	GetBankTransactionCategorySourceModifiedDateStatusEnumUnknown  GetBankTransactionCategorySourceModifiedDateStatusEnum = "Unknown"
	GetBankTransactionCategorySourceModifiedDateStatusEnumActive   GetBankTransactionCategorySourceModifiedDateStatusEnum = "Active"
	GetBankTransactionCategorySourceModifiedDateStatusEnumArchived GetBankTransactionCategorySourceModifiedDateStatusEnum = "Archived"
)

// GetBankTransactionCategorySourceModifiedDate
// The Banking Transaction Categories data type provides a list of hierarchical categories associated with a transaction for greater contextual meaning to transaction activity.
type GetBankTransactionCategorySourceModifiedDate struct {
	HasChildren        *bool                                                   `json:"hasChildren,omitempty"`
	ID                 string                                                  `json:"id"`
	ModifiedDate       *time.Time                                              `json:"modifiedDate,omitempty"`
	Name               string                                                  `json:"name"`
	ParentID           *string                                                 `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                              `json:"sourceModifiedDate,omitempty"`
	Status             *GetBankTransactionCategorySourceModifiedDateStatusEnum `json:"status,omitempty"`
}

type GetBankTransactionCategoryResponse struct {
	ContentType        string
	SourceModifiedDate *GetBankTransactionCategorySourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
