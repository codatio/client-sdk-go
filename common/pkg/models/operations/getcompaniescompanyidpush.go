package operations

import (
	"net/http"
	"time"
)

type GetCompaniesCompanyIDPushPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDPushQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type GetCompaniesCompanyIDPushRequest struct {
	PathParams  GetCompaniesCompanyIDPushPathParams
	QueryParams GetCompaniesCompanyIDPushQueryParams
}

type GetCompaniesCompanyIDPushLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetCompaniesCompanyIDPushLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetCompaniesCompanyIDPushLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetCompaniesCompanyIDPushLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetCompaniesCompanyIDPushLinksLinks struct {
	Current  GetCompaniesCompanyIDPushLinksLinksCurrent   `json:"current"`
	Next     *GetCompaniesCompanyIDPushLinksLinksNext     `json:"next,omitempty"`
	Previous *GetCompaniesCompanyIDPushLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetCompaniesCompanyIDPushLinksLinksSelf      `json:"self"`
}

type GetCompaniesCompanyIDPushLinksResultsChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type GetCompaniesCompanyIDPushLinksResultsChangesTypeEnum string

const (
	GetCompaniesCompanyIDPushLinksResultsChangesTypeEnumUnknown            GetCompaniesCompanyIDPushLinksResultsChangesTypeEnum = "Unknown"
	GetCompaniesCompanyIDPushLinksResultsChangesTypeEnumCreated            GetCompaniesCompanyIDPushLinksResultsChangesTypeEnum = "Created"
	GetCompaniesCompanyIDPushLinksResultsChangesTypeEnumModified           GetCompaniesCompanyIDPushLinksResultsChangesTypeEnum = "Modified"
	GetCompaniesCompanyIDPushLinksResultsChangesTypeEnumDeleted            GetCompaniesCompanyIDPushLinksResultsChangesTypeEnum = "Deleted"
	GetCompaniesCompanyIDPushLinksResultsChangesTypeEnumAttachmentUploaded GetCompaniesCompanyIDPushLinksResultsChangesTypeEnum = "AttachmentUploaded"
)

type GetCompaniesCompanyIDPushLinksResultsChanges struct {
	AttachmentID *string                                                             `json:"attachmentId,omitempty"`
	RecordRef    *GetCompaniesCompanyIDPushLinksResultsChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *GetCompaniesCompanyIDPushLinksResultsChangesTypeEnum               `json:"type,omitempty"`
}

type GetCompaniesCompanyIDPushLinksResultsStatusEnum string

const (
	GetCompaniesCompanyIDPushLinksResultsStatusEnumPending  GetCompaniesCompanyIDPushLinksResultsStatusEnum = "Pending"
	GetCompaniesCompanyIDPushLinksResultsStatusEnumFailed   GetCompaniesCompanyIDPushLinksResultsStatusEnum = "Failed"
	GetCompaniesCompanyIDPushLinksResultsStatusEnumSuccess  GetCompaniesCompanyIDPushLinksResultsStatusEnum = "Success"
	GetCompaniesCompanyIDPushLinksResultsStatusEnumTimedOut GetCompaniesCompanyIDPushLinksResultsStatusEnum = "TimedOut"
)

type GetCompaniesCompanyIDPushLinksResultsValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// GetCompaniesCompanyIDPushLinksResultsValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type GetCompaniesCompanyIDPushLinksResultsValidation struct {
	Errors   []GetCompaniesCompanyIDPushLinksResultsValidationValidationItem `json:"errors,omitempty"`
	Warnings []GetCompaniesCompanyIDPushLinksResultsValidationValidationItem `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDPushLinksResults struct {
	Changes           []GetCompaniesCompanyIDPushLinksResultsChanges   `json:"changes,omitempty"`
	CompanyID         string                                           `json:"companyId"`
	CompletedOnUtc    *time.Time                                       `json:"completedOnUtc,omitempty"`
	DataConnectionKey string                                           `json:"dataConnectionKey"`
	DataType          *string                                          `json:"dataType,omitempty"`
	ErrorMessage      *string                                          `json:"errorMessage,omitempty"`
	PushOperationKey  string                                           `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                        `json:"requestedOnUtc"`
	Status            GetCompaniesCompanyIDPushLinksResultsStatusEnum  `json:"status"`
	StatusCode        int                                              `json:"statusCode"`
	TimeoutInMinutes  *int                                             `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                             `json:"timeoutInSeconds,omitempty"`
	Validation        *GetCompaniesCompanyIDPushLinksResultsValidation `json:"validation,omitempty"`
}

// GetCompaniesCompanyIDPushLinks
// Codat's Paging Model
type GetCompaniesCompanyIDPushLinks struct {
	Links        GetCompaniesCompanyIDPushLinksLinks     `json:"_links"`
	PageNumber   int64                                   `json:"pageNumber"`
	PageSize     int64                                   `json:"pageSize"`
	Results      []GetCompaniesCompanyIDPushLinksResults `json:"results,omitempty"`
	TotalResults int64                                   `json:"totalResults"`
}

type GetCompaniesCompanyIDPushResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *GetCompaniesCompanyIDPushLinks
}
