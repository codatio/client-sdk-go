// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType string

const (
	Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionTypeArray     Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType = "Array"
	Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionTypeObject    Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType = "Object"
	Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionTypeString    Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType = "String"
	Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionTypeNumber    Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType = "Number"
	Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionTypeBoolean   Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType = "Boolean"
	Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionTypeDateTime  Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType = "DateTime"
	Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionTypeFile      Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType = "File"
	Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionTypeMultiPart Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType = "MultiPart"
)

func (e Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType) ToPointer() *Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType {
	return &e
}

func (e *Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType: %v", v)
	}
}
