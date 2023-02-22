package shared

import (
	"time"
)

type CodatDataContractsDatasetsCommerceTransaction struct {
	CreatedDate          *time.Time                                                     `json:"createdDate,omitempty"`
	Currency             *string                                                        `json:"currency,omitempty"`
	ID                   *string                                                        `json:"id,omitempty"`
	ModifiedDate         *time.Time                                                     `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                                                     `json:"sourceModifiedDate,omitempty"`
	SubType              *string                                                        `json:"subType,omitempty"`
	TotalAmount          *float64                                                       `json:"totalAmount,omitempty"`
	TransactionSourceRef *CodatDataContractsDatasetsCommerceTransactionSourceRef        `json:"transactionSourceRef,omitempty"`
	Type                 *CodatDataContractsDatasetsCommercePlatformTransactionTypeEnum `json:"type,omitempty"`
}
