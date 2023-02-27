package codatio

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/utils"
	"io"
	"net/http"
)

type excelReports struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newExcelReports(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *excelReports {
	return &excelReports{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetDataCompaniesCompanyIDAssessExcel - Request an Excel report for download
// Returns the status of the latest report requested.
func (s *excelReports) GetDataCompaniesCompanyIDAssessExcel(ctx context.Context, request operations.GetDataCompaniesCompanyIDAssessExcelRequest) (*operations.GetDataCompaniesCompanyIDAssessExcelResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/assess/excel", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDAssessExcelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetDataCompaniesCompanyIDAssessExcel200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetDataCompaniesCompanyIDAssessExcel200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketing - Get the marketing metrics from an accounting source for a given company.
// Request an Excel report for download.
func (s *excelReports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketing(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketingRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketingResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/accountingMetrics/marketing", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketingResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketing200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketing200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PostDataCompaniesCompanyIDAssessExcel - Request an Excel report for download
// Request an Excel report for download.
func (s *excelReports) PostDataCompaniesCompanyIDAssessExcel(ctx context.Context, request operations.PostDataCompaniesCompanyIDAssessExcelRequest) (*operations.PostDataCompaniesCompanyIDAssessExcelResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/assess/excel", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.PostDataCompaniesCompanyIDAssessExcelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostDataCompaniesCompanyIDAssessExcel200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PostDataCompaniesCompanyIDAssessExcel200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PostDataCompaniesCompanyIDAssessExcelDownload - Download generated excel report
// Download the Excel report to a local drive.
func (s *excelReports) PostDataCompaniesCompanyIDAssessExcelDownload(ctx context.Context, request operations.PostDataCompaniesCompanyIDAssessExcelDownloadRequest) (*operations.PostDataCompaniesCompanyIDAssessExcelDownloadResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/assess/excel/download", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.PostDataCompaniesCompanyIDAssessExcelDownloadResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/octet-stream`):
			out, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Body = out
		}
	}

	return res, nil
}
