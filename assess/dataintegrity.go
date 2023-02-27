package codatio

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/utils"
	"net/http"
)

type dataIntegrity struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newDataIntegrity(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *dataIntegrity {
	return &dataIntegrity{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetails - Lists data integrity details for date type
// Gets record-by-record match results for a given company and datatype, optionally restricted by a Codat query string.
func (s *dataIntegrity) GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetails(ctx context.Context, request operations.GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsRequest) (*operations.GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/assess/dataTypes/{dataType}/dataIntegrity/details", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegrityDetailsLinks
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Links = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries - Get data integrity summary
// Gets match summary for a given company and datatype, optionally restricted by a Codat query string.
func (s *dataIntegrity) GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries(ctx context.Context, request operations.GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesRequest) (*operations.GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/assess/dataTypes/{dataType}/dataIntegrity/summaries", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetDataIntegrityStatusForDataType - Get data integrity status
// Gets match status for a given company and datatype.
func (s *dataIntegrity) GetDataIntegrityStatusForDataType(ctx context.Context, request operations.GetDataIntegrityStatusForDataTypeRequest) (*operations.GetDataIntegrityStatusForDataTypeResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/assess/dataTypes/{dataType}/dataIntegrity/status", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetDataIntegrityStatusForDataTypeResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetDataIntegrityStatusForDataType200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetDataIntegrityStatusForDataType200ApplicationJSONObject = out
		}
	}

	return res, nil
}
