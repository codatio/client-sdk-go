package codatio

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/utils"
	"net/http"
)

type reports struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newReports(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *reports {
	return &reports{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccounts - Enhanced Balance Sheet Accounts
// Gets a list of accounts with account categories per statement period, specific to balance sheet
func (s *reports) GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccounts(ctx context.Context, request operations.GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsRequest) (*operations.GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/reports/enhancedBalanceSheet/accounts", request.PathParams)

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

	res := &operations.GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EnhancedReport
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.EnhancedReport = out
		}
	}

	return res, nil
}

// GetCompaniesCompanyIDReportsEnhancedCashFlowTransactions - Get enhanced cash flow report
// Gets a list of banking transactions and their categories.
func (s *reports) GetCompaniesCompanyIDReportsEnhancedCashFlowTransactions(ctx context.Context, request operations.GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsRequest) (*operations.GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/reports/enhancedCashFlow/transactions", request.PathParams)

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

	res := &operations.GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EnhancedCashFlowTransactions
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.EnhancedCashFlowTransactions = out
		}
	}

	return res, nil
}

// GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccounts - Enhanced Profit and Loss Accounts
// Gets a list of accounts with account categories per statement period, specific to profit and loss
func (s *reports) GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccounts(ctx context.Context, request operations.GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsRequest) (*operations.GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/reports/enhancedProfitAndLoss/accounts", request.PathParams)

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

	res := &operations.GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EnhancedReport
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.EnhancedReport = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsCustomerRetention - Get the customer retention metrics for a specific company.
// Gets the customer retention metrics for a specific company connection, over one or more periods of time.
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsCustomerRetention(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsCustomerRetentionRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsCustomerRetentionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/commerceMetrics/customerRetention", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsCustomerRetentionResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Report = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsLifetimeValue - Get the lifetime value metric for a specific company.
// Gets the lifetime value metric for a specific company connection, over one or more periods of time.
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsLifetimeValue(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsLifetimeValueRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsLifetimeValueResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/commerceMetrics/lifetimeValue", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsLifetimeValueResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Report = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrders - Get order information for a specific company
// Gets the order information for a specific company connection, over one or more periods of time.
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrders(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrdersRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrdersResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/commerceMetrics/orders", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrdersResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Report = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefunds - Get the refunds information for a specific company
// Gets the refunds information for a specific company connection, over one or more periods of time.
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefunds(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefundsRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefundsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/commerceMetrics/refunds", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefundsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Report = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue - Commerce Revenue Metrics
// Get the revenue and revenue growth for a specific company connection, over one or more periods of time.
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenue(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenueRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenueResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/commerceMetrics/revenue", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRevenueResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Report = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheet - Enhanced Balance Sheet
// Gets a fully categorized balance sheet statement for a given company, over one or more period(s).
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheet(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheetRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheetResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/enhancedBalanceSheet", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Report = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLoss - Enhanced Profit and Loss
// Gets a fully categorized profit and loss statement for a given company, over one or more period(s).
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLoss(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLossRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLossResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/enhancedProfitAndLoss", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLossResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Report = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics - List finanicial metrics
// Gets all the available financial metrics for a given company, over one or more periods.
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/financialMetrics", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsMrr - Get key metrics for subscription revenue
// Gets key metrics for subscription revenue.
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsMrr(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsMrrRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsMrrResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/subscriptions/mrr", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsMrrResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Report = out
		}
	}

	return res, nil
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsProcess - Request production of key subscription revenue metrics
// Request production of key subscription revenue metrics.
func (s *reports) GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsProcess(ctx context.Context, request operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsProcessRequest) (*operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsProcessResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/data/companies/{companyId}/connections/{connectionId}/assess/subscriptions/process", request.PathParams)

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

	res := &operations.GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsProcessResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Report
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Report = out
		}
	}

	return res, nil
}
