// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/retry"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var errRequestFailed = errors.New("request failed")

// Deprecated: Use retry.BackoffStrategy instead.
type BackoffStrategy = retry.BackoffStrategy

// Deprecated: Use retry.Config instead.
type RetryConfig = retry.Config

type Retries struct {
	Config      *retry.Config
	StatusCodes []string
}

func Retry(ctx context.Context, r Retries, action func() (*http.Response, error)) (*http.Response, error) {
	switch r.Config.Strategy {
	case "backoff":
		if r.Config.Backoff == nil {
			return action()
		}

		config := backoff.NewExponentialBackOff()
		config.InitialInterval = time.Duration(r.Config.Backoff.InitialInterval) * time.Millisecond
		config.MaxInterval = time.Duration(r.Config.Backoff.MaxInterval) * time.Millisecond
		config.Multiplier = r.Config.Backoff.Exponent
		config.MaxElapsedTime = time.Duration(r.Config.Backoff.MaxElapsedTime) * time.Millisecond
		config.Reset()

		var resp *http.Response

		err := backoff.Retry(func() error {
			if resp != nil {
				resp.Body.Close()
			}

			select {
			case <-ctx.Done():
				return backoff.Permanent(ctx.Err())
			default:
			}

			res, err := action()
			if err != nil {
				urlError := new(url.Error)
				if errors.As(err, &urlError) {
					if (urlError.Temporary() || urlError.Timeout()) && r.Config.RetryConnectionErrors {
						return err
					}
				}

				return backoff.Permanent(err)
			}
			resp = res
			if res == nil {
				return fmt.Errorf("no response")
			}

			for _, code := range r.StatusCodes {
				if strings.Contains(strings.ToUpper(code), "X") {
					codeRange, err := strconv.Atoi(code[:1])
					if err != nil {
						continue
					}

					s := res.StatusCode / 100

					if s >= codeRange && s < codeRange+1 {
						return errRequestFailed
					}
				} else {
					parsedCode, err := strconv.Atoi(code)
					if err != nil {
						continue
					}

					if res.StatusCode == parsedCode {
						return errRequestFailed
					}
				}
			}

			resp = res

			return nil
		}, config)
		if err != nil && !errors.Is(err, errRequestFailed) {
			return nil, err
		}

		return resp, nil
	default:
		return action()
	}
}
