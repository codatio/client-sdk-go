// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package assess

import (
	"bytes"
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/utils"
	"io"
	"net/http"
)

// ExcelReports - Downloadable reports
type ExcelReports struct {
	sdkConfiguration sdkConfiguration
}

func newExcelReports(sdkConfig sdkConfiguration) *ExcelReports {
	return &ExcelReports{
		sdkConfiguration: sdkConfig,
	}
}

// GenerateExcelReport - Generate Excel report
// Generate an Excel report which can subsequently be downloaded.
func (s *ExcelReports) GenerateExcelReport(ctx context.Context, request operations.GenerateExcelReportRequest, opts ...operations.Option) (*operations.GenerateExcelReportResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/assess/excel", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig == nil {
			retryConfig = &utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     60000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		} else {
			retryConfig = globalRetryConfig
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"408",
			"429",
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GenerateExcelReportResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.ExcelStatus
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ExcelStatus = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 402:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.ErrorMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// GetAccountingMarketingMetrics - Get marketing metrics report
// Get the marketing metrics from an accounting source for a given company.
//
// Request an Excel report for download.
func (s *ExcelReports) GetAccountingMarketingMetrics(ctx context.Context, request operations.GetAccountingMarketingMetricsRequest, opts ...operations.Option) (*operations.GetAccountingMarketingMetricsResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/accountingMetrics/marketing", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig == nil {
			retryConfig = &utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     60000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		} else {
			retryConfig = globalRetryConfig
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"408",
			"429",
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAccountingMarketingMetricsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Report = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 402:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.ErrorMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// GetExcelReport - Download Excel report
// Download the previously generated Excel report to a local drive.
func (s *ExcelReports) GetExcelReport(ctx context.Context, request operations.GetExcelReportRequest, opts ...operations.Option) (*operations.GetExcelReportResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/assess/excel/download", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/octet-stream;q=0")
	}

	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig == nil {
			retryConfig = &utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     60000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		} else {
			retryConfig = globalRetryConfig
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"408",
			"429",
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetExcelReportResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/octet-stream`):
			res.Body = rawBody
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 402:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.ErrorMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// GetExcelReportGenerationStatus - Get Excel report status
// Returns the status of the latest report requested.
func (s *ExcelReports) GetExcelReportGenerationStatus(ctx context.Context, request operations.GetExcelReportGenerationStatusRequest, opts ...operations.Option) (*operations.GetExcelReportGenerationStatusResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/assess/excel", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig == nil {
			retryConfig = &utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     60000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		} else {
			retryConfig = globalRetryConfig
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"408",
			"429",
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetExcelReportGenerationStatusResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.ExcelStatus
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ExcelStatus = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 402:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.ErrorMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
