package shared

import (
	"time"
)

type Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnum string

const (
	Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnumUnknown            Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnum = "Unknown"
	Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnumCreated            Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnum = "Created"
	Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnumModified           Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnum = "Modified"
	Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnumDeleted            Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnum = "Deleted"
	Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnumAttachmentUploaded Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnum = "AttachmentUploaded"
)

type Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1Changes struct {
	AttachmentID *string                                                                                                                                                          `json:"attachmentId,omitempty"`
	RecordRef    *Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ChangesTypeEnum               `json:"type,omitempty"`
}

type Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnum string

const (
	Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnumPending  Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnum = "Pending"
	Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnumFailed   Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnum = "Failed"
	Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnumSuccess  Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnum = "Success"
	Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnumTimedOut Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnum = "TimedOut"
)

type Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1Validation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1Validation struct {
	Errors   []Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1ValidationValidationItem                  `json:"errors,omitempty"`
	Warnings []Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1 struct {
	Changes           []Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1Changes   `json:"changes,omitempty"`
	CompanyID         string                                                                                                                                        `json:"companyId"`
	CompletedOnUtc    *time.Time                                                                                                                                    `json:"completedOnUtc,omitempty"`
	DataConnectionKey string                                                                                                                                        `json:"dataConnectionKey"`
	DataType          *string                                                                                                                                       `json:"dataType,omitempty"`
	ErrorMessage      *string                                                                                                                                       `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                                                                                                        `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                                                                                                                     `json:"requestedOnUtc"`
	Status            Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1StatusEnum  `json:"status"`
	StatusCode        int                                                                                                                                           `json:"statusCode"`
	TimeoutInMinutes  *int                                                                                                                                          `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                                                                                                          `json:"timeoutInSeconds,omitempty"`
	Validation        *Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1Validation `json:"validation,omitempty"`
}
