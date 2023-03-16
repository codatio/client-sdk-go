package codatio

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
	"github.com/codatio/client-sdk-go/expenses/pkg/utils"
	"net/http"
)

type transactionStatus struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newTransactionStatus(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *transactionStatus {
	return &transactionStatus{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetSyncTransaction - Get Sync Transaction
// Gets the status of a transaction for a sync
func (s *transactionStatus) GetSyncTransaction(ctx context.Context, request operations.GetSyncTransactionRequest) (*operations.GetSyncTransactionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/sync/expenses/syncs/{syncId}/transactions/{transactionId}", request.PathParams, nil)

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

	res := &operations.GetSyncTransactionResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []operations.GetSyncTransaction200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetSyncTransaction200ApplicationJSONObjects = out
		}
	}

	return res, nil
}

// GetSyncTransactions - Get Sync transactions
// Get's the transactions and status for a sync
func (s *transactionStatus) GetSyncTransactions(ctx context.Context, request operations.GetSyncTransactionsRequest) (*operations.GetSyncTransactionsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/sync/expenses/syncs/{syncId}/transactions", request.PathParams, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams, nil); err != nil {
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

	res := &operations.GetSyncTransactionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetSyncTransactions200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetSyncTransactions200ApplicationJSONObject = out
		}
	}

	return res, nil
}
