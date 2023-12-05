<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/221800355-0995e4ad-a386-4943-a4c2-e620341a5155.svg" media="(prefers-color-scheme: dark)">
        <img src="https://user-images.githubusercontent.com/6267663/221800359-b7f7776c-a44f-4384-8dd0-d9f7d5caef7d.svg">
    </picture>
    <h1>Codat Go</h1>
        <p><strong>The API for lending and embedded accounting automation</strong></p>
        <p>Codat makes it quicker and easier to build fintech products, like corporate cards, business dashboards, or SMB lending applications, that are integrated with the other systems small businesses use.</p>
    <a href="https://docs.codat.io/using-the-api/overview"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=4c2cec&style=for-the-badge" /></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

> **Beta Release**
>
> Exciting news! Our first Go library beta release is here, simplifying your development workflow with an effortless build process.

## Authentication

Codat uses API keys to control access to the API. 
Learn more about authentication and managing API keys [here](https://docs.codat.io/using-the-api/authentication).

You must keep the API key secret, so make sure it isn't available in publicly accessible areas, such as GitHub and client-side code.
Codat recommends the API key is only inserted at release time, and the number of people at your organization with access to your API key is minimised.

Codat expects the API key to be included in all API requests to the server, Base64 encoded within an 'Authorization' header.

```bash
Authorization: Basic BASE_64_ENCODED(API_KEY)
```

### Getting your Authorization Header

To get your authorization header from the [Codat Portal](https://app.codat.io):

1. In the navigation bar, click [**Developers > API keys**](https://app.codat.io/developers/api-keys).
2. In the **API Keys** section, copy your authorization header rather than the API key itself.

## Client Libraries

<!-- Start Codat Client Libraries -->
| Library | Description | `go get` |
| :- | :- | :- |
| **[Platform](https://github.com/codatio/client-sdk-go/tree/main/platform)** | Manage the building blocks of Codat, including companies, connections, and more. | `github.com/codatio/client-sdk-go/platform` |
| **[Bank Feeds](https://github.com/codatio/client-sdk-go/tree/main/bank-feeds)** | Set up bank feeds from accounts in your application to supported accounting platforms. | `github.com/codatio/client-sdk-go/bank-feeds` |
| **[Sync for Commerce](https://github.com/codatio/client-sdk-go/tree/main/sync-for-commerce)** | Push merchants' data from your ecommerce or point-of-sale (POS) platform into your merchants' accounting platform. | `github.com/codatio/client-sdk-go/sync-for-commerce` |
| **[Sync for Expenses](https://github.com/codatio/client-sdk-go/tree/main/sync-for-expenses)** | Push expenses to accounting platforms. | `github.com/codatio/client-sdk-go/sync-for-expenses` |
| **[Lending](https://github.com/codatio/client-sdk-go/tree/main/lending)** | Make credit decisions backed by enhanced financials, metrics, reports, and data integrity features. | `github.com/codatio/client-sdk-go/lending` |
| **[Sync for Payroll](https://github.com/codatio/client-sdk-go/tree/main/sync-for-payroll)** | Push payroll to accounting platforms. | `github.com/codatio/client-sdk-go/sync-for-payroll` |
| **[Sync for Payables](https://github.com/codatio/client-sdk-go/tree/main/sync-for-payables)** | Streamline your customers' accounts payable workflow. | `github.com/codatio/client-sdk-go/sync-for-payables` |

### Alternative products and versions

| Library | Description | `go get` |
| :- | :- | :- |
| **[Common](https://github.com/codatio/client-sdk-go/tree/main/previous-versions/common)** | Manage the building blocks of Codat, including companies, connections, and more. | `github.com/codatio/client-sdk-go/previous-versions/common` |
| **[Accounting](https://github.com/codatio/client-sdk-go/tree/main/previous-versions/accounting)** | Access standardized accounting data from our accounting integrations. | `github.com/codatio/client-sdk-go/previous-versions/accounting` |
| **[Banking](https://github.com/codatio/client-sdk-go/tree/main/previous-versions/banking)** | Access standardized banking data from our banking integrations. | `github.com/codatio/client-sdk-go/previous-versions/banking` |
| **[Commerce](https://github.com/codatio/client-sdk-go/tree/main/previous-versions/commerce)** | Access standardized commerce data from our commerce integrations. | `github.com/codatio/client-sdk-go/previous-versions/commerce` |
| **[Assess](https://github.com/codatio/client-sdk-go/tree/main/previous-versions/assess)** | Make credit decisions backed by enhanced financials, metrics, reports, and data integrity features. | `github.com/codatio/client-sdk-go/previous-versions/assess` |
| **[Sync for Commerce version 1](https://github.com/codatio/client-sdk-go/tree/main/previous-versions/sync-for-commerce-version-1)** | Push merchants' data from your ecommerce or point-of-sale (POS) platform into your merchants' accounting platform. | `github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1` |
| **[Sync for Expenses version 1](https://github.com/codatio/client-sdk-go/tree/main/previous-versions/sync-for-expenses-version-1)** | Push expenses to accounting platforms. | `github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1` |
| **[Files](https://github.com/codatio/client-sdk-go/tree/main/previous-versions/files)** | Use Codat's Files API to upload your SMB customers' files. | `github.com/codatio/client-sdk-go/previous-versions/files` |
<!-- End Codat Client Libraries -->

<!-- Start Codat Support Notes -->
<!-- End Codat Support Notes -->