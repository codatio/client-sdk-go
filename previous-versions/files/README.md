# Files

<!-- Start Codat Library Description -->
﻿Use Codat's Files API to upload your SMB customers' files.
<!-- End Codat Library Description -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/files
```
<!-- End SDK Installation [installation] -->

## Example Usage
<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"log"
)

func main() {
	s := files.New(
		files.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00Z"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Files](docs/sdks/files/README.md)

* [DownloadFiles](docs/sdks/files/README.md#downloadfiles) - Download all files for a company
* [ListFiles](docs/sdks/files/README.md#listfiles) - List all files uploaded by a company
* [UploadFiles](docs/sdks/files/README.md#uploadfiles) - Upload files for a company
<!-- End Available Resources and Operations [operations] -->





<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/utils"
	"log"
	"pkg/models/operations"
)

func main() {
	s := files.New(
		files.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00Z"),
	}, operations.WithRetries(
		utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/utils"
	"log"
)

func main() {
	s := files.New(
		files.WithRetryConfig(
			utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		files.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00Z"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.Schema                    | 400,401,402,404,429,500,503         | application/json                    |
| sdkerrors.DownloadFilesErrorMessage | 403                                 | application/json                    |
| sdkerrors.SDKError                  | 400-600                             | */*                                 |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"log"
)

func main() {
	s := files.New(
		files.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00Z"),
	})
	if err != nil {

		var e *sdkerrors.Schema
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.DownloadFilesErrorMessage
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.codat.io` | None |

#### Example

```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"log"
)

func main() {
	s := files.New(
		files.WithServerIndex(0),
		files.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00Z"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"log"
)

func main() {
	s := files.New(
		files.WithServerURL("https://api.codat.io"),
		files.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00Z"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `AuthHeader` | apiKey       | API key      |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"log"
)

func main() {
	s := files.New(
		files.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00Z"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

<!-- Start Codat Support Notes -->
### Support

If you encounter any challenges while utilizing our SDKs, please don't hesitate to reach out for assistance. 
You can raise any issues by contacting your dedicated Codat representative or reaching out to our [support team](mailto:support@codat.io).
We're here to help ensure a smooth experience for you.
<!-- End Codat Support Notes -->

<!-- Start Codat Generated By -->
### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
<!-- End Codat Generated By -->