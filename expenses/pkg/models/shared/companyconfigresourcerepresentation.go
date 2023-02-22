package shared

type CompanyConfigResourceRepresentationBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type CompanyConfigResourceRepresentationCustomer struct {
	ID *string `json:"id,omitempty"`
}

type CompanyConfigResourceRepresentationSupplier struct {
	ID *string `json:"id,omitempty"`
}

type CompanyConfigResourceRepresentation struct {
	BankAccount *CompanyConfigResourceRepresentationBankAccount `json:"bankAccount,omitempty"`
	Customer    *CompanyConfigResourceRepresentationCustomer    `json:"customer,omitempty"`
	Supplier    *CompanyConfigResourceRepresentationSupplier    `json:"supplier,omitempty"`
}
