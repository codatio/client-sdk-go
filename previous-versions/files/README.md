# Files

<!-- Start Codat Library Description -->
﻿Use Codat's Files API to upload your SMB customers' files.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/files
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
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
		Date:      files.String("2022-10-23T00:00:00.000Z"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Files](docs/sdks/files/README.md)

* [DownloadFiles](docs/sdks/files/README.md#downloadfiles) - Download all files for a company
* [ListFiles](docs/sdks/files/README.md#listfiles) - List all files uploaded by a company
* [UploadFiles](docs/sdks/files/README.md#uploadfiles) - Upload files for a company
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.codat.io` | None |

For example:


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
		files.WithServerIndex(0),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00.000Z"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

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
		files.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
		files.WithServerURL("https://api.codat.io"),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00.000Z"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

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
<!-- End Custom HTTP Client -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)