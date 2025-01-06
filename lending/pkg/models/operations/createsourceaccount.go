// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"net/http"
)

type CreateSourceAccountRequestBodyType string

const (
	CreateSourceAccountRequestBodyTypeSourceAccountV2Prototype CreateSourceAccountRequestBodyType = "sourceAccountV2Prototype"
	CreateSourceAccountRequestBodyTypeSourceAccountPrototype   CreateSourceAccountRequestBodyType = "sourceAccountPrototype"
)

type CreateSourceAccountRequestBody struct {
	SourceAccountV2Prototype *shared.SourceAccountV2Prototype `queryParam:"inline"`
	SourceAccountPrototype   *shared.SourceAccountPrototype   `queryParam:"inline"`

	Type CreateSourceAccountRequestBodyType
}

func CreateCreateSourceAccountRequestBodySourceAccountV2Prototype(sourceAccountV2Prototype shared.SourceAccountV2Prototype) CreateSourceAccountRequestBody {
	typ := CreateSourceAccountRequestBodyTypeSourceAccountV2Prototype

	return CreateSourceAccountRequestBody{
		SourceAccountV2Prototype: &sourceAccountV2Prototype,
		Type:                     typ,
	}
}

func CreateCreateSourceAccountRequestBodySourceAccountPrototype(sourceAccountPrototype shared.SourceAccountPrototype) CreateSourceAccountRequestBody {
	typ := CreateSourceAccountRequestBodyTypeSourceAccountPrototype

	return CreateSourceAccountRequestBody{
		SourceAccountPrototype: &sourceAccountPrototype,
		Type:                   typ,
	}
}

func (u *CreateSourceAccountRequestBody) UnmarshalJSON(data []byte) error {

	var sourceAccountPrototype shared.SourceAccountPrototype = shared.SourceAccountPrototype{}
	if err := utils.UnmarshalJSON(data, &sourceAccountPrototype, "", true, true); err == nil {
		u.SourceAccountPrototype = &sourceAccountPrototype
		u.Type = CreateSourceAccountRequestBodyTypeSourceAccountPrototype
		return nil
	}

	var sourceAccountV2Prototype shared.SourceAccountV2Prototype = shared.SourceAccountV2Prototype{}
	if err := utils.UnmarshalJSON(data, &sourceAccountV2Prototype, "", true, true); err == nil {
		u.SourceAccountV2Prototype = &sourceAccountV2Prototype
		u.Type = CreateSourceAccountRequestBodyTypeSourceAccountV2Prototype
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateSourceAccountRequestBody", string(data))
}

func (u CreateSourceAccountRequestBody) MarshalJSON() ([]byte, error) {
	if u.SourceAccountV2Prototype != nil {
		return utils.MarshalJSON(u.SourceAccountV2Prototype, "", true)
	}

	if u.SourceAccountPrototype != nil {
		return utils.MarshalJSON(u.SourceAccountPrototype, "", true)
	}

	return nil, errors.New("could not marshal union type CreateSourceAccountRequestBody: all fields are null")
}

type CreateSourceAccountRequest struct {
	RequestBody *CreateSourceAccountRequestBody `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *CreateSourceAccountRequest) GetRequestBody() *CreateSourceAccountRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateSourceAccountRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateSourceAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateSourceAccountResponseBodyType string

const (
	CreateSourceAccountResponseBodyTypeSourceAccountV2 CreateSourceAccountResponseBodyType = "SourceAccountV2"
	CreateSourceAccountResponseBodyTypeSourceAccount   CreateSourceAccountResponseBodyType = "SourceAccount"
)

// CreateSourceAccountResponseBody - Success
type CreateSourceAccountResponseBody struct {
	SourceAccountV2 *shared.SourceAccountV2 `queryParam:"inline"`
	SourceAccount   *shared.SourceAccount   `queryParam:"inline"`

	Type CreateSourceAccountResponseBodyType
}

func CreateCreateSourceAccountResponseBodySourceAccountV2(sourceAccountV2 shared.SourceAccountV2) CreateSourceAccountResponseBody {
	typ := CreateSourceAccountResponseBodyTypeSourceAccountV2

	return CreateSourceAccountResponseBody{
		SourceAccountV2: &sourceAccountV2,
		Type:            typ,
	}
}

func CreateCreateSourceAccountResponseBodySourceAccount(sourceAccount shared.SourceAccount) CreateSourceAccountResponseBody {
	typ := CreateSourceAccountResponseBodyTypeSourceAccount

	return CreateSourceAccountResponseBody{
		SourceAccount: &sourceAccount,
		Type:          typ,
	}
}

func (u *CreateSourceAccountResponseBody) UnmarshalJSON(data []byte) error {

	var sourceAccount shared.SourceAccount = shared.SourceAccount{}
	if err := utils.UnmarshalJSON(data, &sourceAccount, "", true, true); err == nil {
		u.SourceAccount = &sourceAccount
		u.Type = CreateSourceAccountResponseBodyTypeSourceAccount
		return nil
	}

	var sourceAccountV2 shared.SourceAccountV2 = shared.SourceAccountV2{}
	if err := utils.UnmarshalJSON(data, &sourceAccountV2, "", true, true); err == nil {
		u.SourceAccountV2 = &sourceAccountV2
		u.Type = CreateSourceAccountResponseBodyTypeSourceAccountV2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateSourceAccountResponseBody", string(data))
}

func (u CreateSourceAccountResponseBody) MarshalJSON() ([]byte, error) {
	if u.SourceAccountV2 != nil {
		return utils.MarshalJSON(u.SourceAccountV2, "", true)
	}

	if u.SourceAccount != nil {
		return utils.MarshalJSON(u.SourceAccount, "", true)
	}

	return nil, errors.New("could not marshal union type CreateSourceAccountResponseBody: all fields are null")
}

type CreateSourceAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	OneOf *CreateSourceAccountResponseBody
}

func (o *CreateSourceAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSourceAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSourceAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSourceAccountResponse) GetOneOf() *CreateSourceAccountResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
