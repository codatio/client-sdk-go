package shared

type DataTypeEnum string

const (
	DataTypeEnumInvoices             DataTypeEnum = "invoices"
	DataTypeEnumAccounts             DataTypeEnum = "accounts"
	DataTypeEnumCommercePayments     DataTypeEnum = "commerce-payments"
	DataTypeEnumBankingAccounts      DataTypeEnum = "banking-accounts"
	DataTypeEnumCompany              DataTypeEnum = "company"
	DataTypeEnumProfitAndLoss        DataTypeEnum = "profitAndLoss"
	DataTypeEnumCommerceTransactions DataTypeEnum = "commerce-transactions"
	DataTypeEnumBills                DataTypeEnum = "bills"
	DataTypeEnumCustomers            DataTypeEnum = "customers"
)
