package operations

import (
	"net/http"
	"time"
)

type CreatePaymentPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreatePaymentQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// CreatePaymentSourceModifiedDateAccountRef
// Account the payment is recorded against in the accounting platform.
type CreatePaymentSourceModifiedDateAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreatePaymentSourceModifiedDateCustomerRef
// Customer the payment is recorded against in the accounting platform.
type CreatePaymentSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreatePaymentSourceModifiedDateLinesLinksTypeEnum string

const (
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumUnknown          CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "Unknown"
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumUnlinked         CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "Unlinked"
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumInvoice          CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "Invoice"
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumCreditNote       CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "CreditNote"
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumOther            CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "Other"
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumRefund           CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "Refund"
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumPayment          CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "Payment"
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumPaymentOnAccount CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "PaymentOnAccount"
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumManualJournal    CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "ManualJournal"
	CreatePaymentSourceModifiedDateLinesLinksTypeEnumDiscount         CreatePaymentSourceModifiedDateLinesLinksTypeEnum = "Discount"
)

type CreatePaymentSourceModifiedDateLinesLinks struct {
	Amount       *float64                                          `json:"amount,omitempty"`
	CurrencyRate *float64                                          `json:"currencyRate,omitempty"`
	ID           *string                                           `json:"id,omitempty"`
	Type         CreatePaymentSourceModifiedDateLinesLinksTypeEnum `json:"type"`
}

type CreatePaymentSourceModifiedDateLines struct {
	AllocatedOnDate *time.Time                                  `json:"allocatedOnDate,omitempty"`
	Amount          float64                                     `json:"amount"`
	Links           []CreatePaymentSourceModifiedDateLinesLinks `json:"links,omitempty"`
}

type CreatePaymentSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreatePaymentSourceModifiedDatePaymentMethodRef
// The Payment Method to which the payment is linked in the accounting platform.
type CreatePaymentSourceModifiedDatePaymentMethodRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreatePaymentSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreatePaymentSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreatePaymentSourceModifiedDate
// > **Payments or bill payments?**
// >
// >  In Codat, payments represent accounts receivable only. For accounts payable, see [bill payments](https://docs.codat.io/accounting-api#/schemas/BillPayment). These include [bills](https://docs.codat.io/accounting-api#/schemas/Bill) and credit notes against bills.
//
// > View the coverage for payments in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Payments include all accounts receivable transaction data. This includes [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) and [credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote).
//
// A payment in Codat usually represents an allocation of money within any customer accounts receivable account. This includes, but is not strictly limited to:
//
// - A payment made against an invoice, like a credit card, cheque, or cash payment.
// - An allocation of a customer's credit note, either to an invoice or maybe a refund.
// - A payment made directly to that accounts receivable account. This might be an overpayment or a prepayment. It might also be the refund of a payment made directly to an accounts receivable account.
//
// Depending on the payments allowed by the underlying accounting package, some payment types may be combined. Please see the [Example data](#section-example-data) below for more details.
//
// In Codat, a payment contains details of:
//
// - When the payment was recorded in the accounting system.
// - How much it is for and in what currency that amount is in.
// - Who the payment was _paid by_ – the _customer_.
// - The payment method used.
// - The breakdown of the types of payments – the _line items_.
//
// Payments is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
//
// ## Payment types
//
// ## Payment of an invoice
//
// A payment paying a single invoice has one entry in its `lines` array. This **line** has the following properties:
//
// - An _amount_ that indicates the amount of the invoice that was paid. This is always positive.
// - A **links** array containing one element with the following properties:
//   - A **type** that indicates the type of **link**, in this case an `Invoice`.
//   - An **id** that contains the ID of the invoice that was paid.
//   - An **amount** for the link. The sum of the **line.amount** and the **links.amount** must equal `0`.
//
// The **amount** field on the **line** equals the **totalAmount** on the payment.
//
// ## Payment of multiple invoices
//
// A single payment can pay multiple invoices. This can be represented in one of two formats depending on how the customer keeps their books:
//
// - The payment has multiple entries in its **lines** array, one for each invoice that is paid. Each line follows the example and rules described in [Payment of an invoice](#payment-of-an-invoice).
// - The payment has a line with multiple links to each invoice. This occurs when the proportion of the original payment allocated to each invoice is not available.
//
// Each **line** has the same properties as those described in [Payment of an invoice](#payment-of-an-invoice), with the **amount** indicating how much of the payment was allocated to the invoice. The sum of line amounts equals the **totalAmount** on the payment.
//
// ## Payments and refunds on account
//
// A payment on account, that is a payment that doesn’t pay a specific invoice, has one entry in its lines array. The **line** has the following properties:
//
// - A **totalAmount** that indicates the amount paid by a customer or refunded to them by a company. A payment to the customer is always negative. A refund is always positive.
// - A **links** array containing one element with the following properties:
// - A **type** that indicates the type of link. For a payment this is `PaymentOnAccount`. For a refund this is `Refund`.
// - The **id** containing the ID of the customer.
// - The **amount** for the link is `0` – the **totalAmount** _or_ the amount of the payment or refund.
//
// It is possible to have a payment that is part _on account_ and part _allocated_ to an invoice. Each line should follow the examples above.
//
// ## Using a credit note to pay an invoice
//
// The payment of an invoice using a credit note has one entry in its **lines** array. This **line** has the following properties:
//
// - An **amount** that indicates the amount of money moved, which in this case is `0`, as the credit note and invoice allocation must balance each other.
// - A **links** array containing two elements:
//   - The first **link** has:
//   - A **type** that indicates the type of **link**, in this case an `Invoice`.
//   - An **id** that contains the ID of the invoice that was paid.
//   - The second **link** has:
//   - A **type** that indicates the type of **link**, in this case a `CreditNote`.
//   - An **id** that contains the ID of the credit note used by this payment.
//
// The **amount** field on the **line** equals the **totalAmount** on the payment.
//
// ## Refunding a credit note
//
// A payment refunding a credit note has one entry in its **lines** array. This **line** has the following properties:
//
// - An **amount** that indicates the amount of the credit note that was refunded. This is always negative for a refund.
// - A **links** array that contains one element with the following properties:
//   - A **type** that indicates the type of **link**, in this case a `CreditNote`.
//   - An **id** that contains the ID of the credit note that was refunded.
//
// The **totalAmount** field on the payment equals the **amount** field of the **line**. These are both negative, as this is money leaving accounts receivable.
//
// ## Refunding a payment
//
// If a payment is refunded, for example, if a customer overpaid an invoice and the overpayment is returned to the customer, there are two payment records:
//
// - One for the incoming over payment.
// - Another for the outgoing refund.
//
// The payment issuing the refund has a negative **totalAmount**. This payment also has one entry in its lines array with the following properties:
//
// - An **amount** that indicates the amount that was refunded. This is always negative.
// - A **links** array that contains one element with the following properties:
//   - A **type** that indicates the type of **link**, in this case a `Payment`.
//   - An **id** that contains the ID of the payment that was refunded.
//
// The **amount** field on the **line** equals the **totalAmount** on the payment and is negative, as this is money leaving accounts receivable.
//
// The payment that was refunded has a line where the **amount** is positive and the type of the link is `Refund`. This payment may have several entries in its **lines** array if it was used to partly pay an invoice.
//
// For example: A £1,050 payment on a £1,000 invoice with a refund of £50 has two lines:
//
// - One for £1,000 linked to the invoice that was paid.
// - Another for £50 linked to the payment that refunded the overpayment with a** type** of `Refund` and an ID that corresponds to the payment.
//
// The **line** linked to the payment has the following properties:
//
// - An **amount** that indicates the amount that was refunded. This is positive as its money that was added to accounts receivable. It's balanced out by the negative amount of the refund.
// - A **links** array containing one element with the following properties:
//   - A **type** that indicates the type of **link**, in this case a `Refund`.
//   - An **id** that contains the ID of the payment that refunded this line.
//
// > 📘 Support for linked payments
// >
// > Not all accounting packages support linking payments in this way. In some platforms, you may see a payment on account and a refund on account.
//
// ## Foreign currencies
//
// There are two types of currency rate that are included in the payments data type:
//
// Payment currency rate:
//
// - Base currency of the accounts receivable account.
// - Foreign currency of the payment.
//
// Payment line link currency rate:
//
// - Base currency of the item the link represents.
// - Foreign currency of the payment.
//
// These two rates allow the calculation of currency loss or gain for any of the transactions affected by the payment lines. The second rate is used when a payment is applied to an item in a currency that doesn't match either:
//
// - The base currency for the accounts receivable account.
// - The currency of the item.
//
// ```json Currency rate example
//
//	{
//	    "id": "123",
//	    "note": ""
//	    "totalAmount": 99.99,
//	    "currency": "GBP",
//	    "lines": [
//	        {
//	            "amount": 99.99,
//	            "links": [
//	                {
//	                    "type": "Invoice",
//	                    "id": "178",
//	                    "amount": -50,
//	                    "currencyRate":  1.9998,
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ## Example data
//
// > 📘 Object properties
// >
// > For the sake of brevity, the examples here may omit properties from objects. For the full object definition, see [Payments](https://api.codat.io/swagger/index.html#/Payments).
//
// ## Simple examples
//
// ```json Payment for invoice
//
//	{
//	    "totalAmount": 1000,
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Allocation of credit note
//
//	{
//	    "totalAmount": 0,
//	    "lines": [
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Payment of invoice and payment on account
//
//	{
//	    "totalAmount": 2000,
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "y",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Refund of credit note
//
//	{
//	    "totalAmount": -1000,
//	    "lines": [
//	        {
//	            "amount" : -1000,
//	            "links" : [
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Refund on accounts receivable account
//
//	{
//	    "totalAmount": -1000,
//	    "lines": [
//	        {
//	            "amount" : -1000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Linked refund on accounts receivable account
//
//	{
//	    "id" : "payment-001",
//	    "totalAmount": 1000,
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Refund",
//	                    "id" : "refund-001",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
//	{
//	    "id" : "refund-001",
//	    "totalAmount": -1000,
//	    "lines": [
//	        {
//	            "amount" : -1000,
//	            "links" : [
//	                {
//	                    "type" : "Payment",
//	                    "id" : "payment-001",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Using a credit note and cash to pay an invoice
//
//	{
//	    "totalAmount": 250,
//	    "lines": [
//	        {
//	            "amount": 0,
//	            "links": [
//	                {
//	                    "type": "Invoice",
//	                    "id": "x",
//	                    "amount": -750
//	                },
//	                {
//	                    "type": "CreditNote",
//	                    "id": "y",
//	                    "amount": 750
//	                }
//	            ]
//	        },
//	        {
//	            "amount": 250,
//	            "links": [
//	                {
//	                    "type": "Invoice",
//	                    "id": "x",
//	                    "amount": -250
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ## Complex examples
//
// ```json Use two credit notes and 1000 in to "bank" (cash, cheque etc.) to pay invoice
//
//	{
//	    "totalAmount": 1000,
//	    "lines": [
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "z",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Pay an invoice with two credit notes and cash, with 1000 left "on account"
//
//	{
//	    "totalAmount": 2000,
//	    "lines": [
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "z",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "customer-001",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Two credit notes pay two invoices with no allocation amount specified
//
//	{
//	    "totalAmount": 0,
//	    "lines": [
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "w",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "z",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Two credit notes and cash pay three invoices with no allocation amount specified, and refund cash
//
//	{
//	    "totalAmount": 2000,
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "w",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "u",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "z",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Refund",
//	                    "id" : "refund-001",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
//	{
//	    "id" : "refund-001",
//	    "totalAmount": -1000,
//	    "lines": [
//	        {
//	            "amount" : -1000,
//	            "links" : [
//	                {
//	                    "type" : "Payment",
//	                    "id" : "payment-001",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// In this example, a payment on account is used to pay the same invoice in January and again in February.
//
// ```json January
//
//	{
//	    "id": "001",
//	    "totalAmount": 5000,
//	    "date" : "1901-01-01",
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "Invoice-x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 4000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "PaymentOnAccount-y",
//	                    "amount" : -4000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json February
//
//	{
//	    "id": "001",
//	    "totalAmount": 5000,
//	    "date" : "1901-02-01",
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "Invoice-x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "Invoice-y",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 3000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "PaymentOnAccount-y",
//	                    "amount" : -3000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Two credit notes and some cash pay two invoices with no allocations specified
//
//	{
//	    "totalAmount": 500,
//	    "lines": [
//	        {
//	            "amount": 500,
//	            "links": [{
//	                    "type": "Invoice",
//	                    "id": "a",
//	                    "amount": -1000
//	                }, {
//	                    "type": "Invoice",
//	                    "id": "b",
//	                    "amount": -1000
//	                }, {
//	                    "type": "CreditNote",
//	                    "id": "y",
//	                    "amount": 750
//	                },{
//	                    "type": "CreditNote",
//	                    "id": "z",
//	                    "amount": 750
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
type CreatePaymentSourceModifiedDate struct {
	AccountRef         *CreatePaymentSourceModifiedDateAccountRef       `json:"accountRef,omitempty"`
	Currency           *string                                          `json:"currency,omitempty"`
	CurrencyRate       *float64                                         `json:"currencyRate,omitempty"`
	CustomerRef        *CreatePaymentSourceModifiedDateCustomerRef      `json:"customerRef,omitempty"`
	Date               time.Time                                        `json:"date"`
	ID                 *string                                          `json:"id,omitempty"`
	Lines              []CreatePaymentSourceModifiedDateLines           `json:"lines,omitempty"`
	Metadata           *CreatePaymentSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                       `json:"modifiedDate,omitempty"`
	Note               *string                                          `json:"note,omitempty"`
	PaymentMethodRef   *CreatePaymentSourceModifiedDatePaymentMethodRef `json:"paymentMethodRef,omitempty"`
	Reference          *string                                          `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                       `json:"sourceModifiedDate,omitempty"`
	SupplementalData   *CreatePaymentSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TotalAmount        *float64                                         `json:"totalAmount,omitempty"`
}

type CreatePaymentRequest struct {
	PathParams  CreatePaymentPathParams
	QueryParams CreatePaymentQueryParams
	Request     *CreatePaymentSourceModifiedDate `request:"mediaType=application/json"`
}

type CreatePayment200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreatePayment200ApplicationJSONChangesTypeEnum string

const (
	CreatePayment200ApplicationJSONChangesTypeEnumUnknown            CreatePayment200ApplicationJSONChangesTypeEnum = "Unknown"
	CreatePayment200ApplicationJSONChangesTypeEnumCreated            CreatePayment200ApplicationJSONChangesTypeEnum = "Created"
	CreatePayment200ApplicationJSONChangesTypeEnumModified           CreatePayment200ApplicationJSONChangesTypeEnum = "Modified"
	CreatePayment200ApplicationJSONChangesTypeEnumDeleted            CreatePayment200ApplicationJSONChangesTypeEnum = "Deleted"
	CreatePayment200ApplicationJSONChangesTypeEnumAttachmentUploaded CreatePayment200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreatePayment200ApplicationJSONChanges struct {
	AttachmentID *string                                                       `json:"attachmentId,omitempty"`
	RecordRef    *CreatePayment200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreatePayment200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreatePayment200ApplicationJSONSourceModifiedDateAccountRef
// Account the payment is recorded against in the accounting platform.
type CreatePayment200ApplicationJSONSourceModifiedDateAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreatePayment200ApplicationJSONSourceModifiedDateCustomerRef
// Customer the payment is recorded against in the accounting platform.
type CreatePayment200ApplicationJSONSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum string

const (
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumUnknown          CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "Unknown"
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumUnlinked         CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "Unlinked"
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumInvoice          CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "Invoice"
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumCreditNote       CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "CreditNote"
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumOther            CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "Other"
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumRefund           CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "Refund"
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumPayment          CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "Payment"
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumPaymentOnAccount CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "PaymentOnAccount"
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumManualJournal    CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "ManualJournal"
	CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnumDiscount         CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum = "Discount"
)

type CreatePayment200ApplicationJSONSourceModifiedDateLinesLinks struct {
	Amount       *float64                                                            `json:"amount,omitempty"`
	CurrencyRate *float64                                                            `json:"currencyRate,omitempty"`
	ID           *string                                                             `json:"id,omitempty"`
	Type         CreatePayment200ApplicationJSONSourceModifiedDateLinesLinksTypeEnum `json:"type"`
}

type CreatePayment200ApplicationJSONSourceModifiedDateLines struct {
	AllocatedOnDate *time.Time                                                    `json:"allocatedOnDate,omitempty"`
	Amount          float64                                                       `json:"amount"`
	Links           []CreatePayment200ApplicationJSONSourceModifiedDateLinesLinks `json:"links,omitempty"`
}

type CreatePayment200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreatePayment200ApplicationJSONSourceModifiedDatePaymentMethodRef
// The Payment Method to which the payment is linked in the accounting platform.
type CreatePayment200ApplicationJSONSourceModifiedDatePaymentMethodRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreatePayment200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreatePayment200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreatePayment200ApplicationJSONSourceModifiedDate
// > **Payments or bill payments?**
// >
// >  In Codat, payments represent accounts receivable only. For accounts payable, see [bill payments](https://docs.codat.io/accounting-api#/schemas/BillPayment). These include [bills](https://docs.codat.io/accounting-api#/schemas/Bill) and credit notes against bills.
//
// > View the coverage for payments in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Payments include all accounts receivable transaction data. This includes [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) and [credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote).
//
// A payment in Codat usually represents an allocation of money within any customer accounts receivable account. This includes, but is not strictly limited to:
//
// - A payment made against an invoice, like a credit card, cheque, or cash payment.
// - An allocation of a customer's credit note, either to an invoice or maybe a refund.
// - A payment made directly to that accounts receivable account. This might be an overpayment or a prepayment. It might also be the refund of a payment made directly to an accounts receivable account.
//
// Depending on the payments allowed by the underlying accounting package, some payment types may be combined. Please see the [Example data](#section-example-data) below for more details.
//
// In Codat, a payment contains details of:
//
// - When the payment was recorded in the accounting system.
// - How much it is for and in what currency that amount is in.
// - Who the payment was _paid by_ – the _customer_.
// - The payment method used.
// - The breakdown of the types of payments – the _line items_.
//
// Payments is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
//
// ## Payment types
//
// ## Payment of an invoice
//
// A payment paying a single invoice has one entry in its `lines` array. This **line** has the following properties:
//
// - An _amount_ that indicates the amount of the invoice that was paid. This is always positive.
// - A **links** array containing one element with the following properties:
//   - A **type** that indicates the type of **link**, in this case an `Invoice`.
//   - An **id** that contains the ID of the invoice that was paid.
//   - An **amount** for the link. The sum of the **line.amount** and the **links.amount** must equal `0`.
//
// The **amount** field on the **line** equals the **totalAmount** on the payment.
//
// ## Payment of multiple invoices
//
// A single payment can pay multiple invoices. This can be represented in one of two formats depending on how the customer keeps their books:
//
// - The payment has multiple entries in its **lines** array, one for each invoice that is paid. Each line follows the example and rules described in [Payment of an invoice](#payment-of-an-invoice).
// - The payment has a line with multiple links to each invoice. This occurs when the proportion of the original payment allocated to each invoice is not available.
//
// Each **line** has the same properties as those described in [Payment of an invoice](#payment-of-an-invoice), with the **amount** indicating how much of the payment was allocated to the invoice. The sum of line amounts equals the **totalAmount** on the payment.
//
// ## Payments and refunds on account
//
// A payment on account, that is a payment that doesn’t pay a specific invoice, has one entry in its lines array. The **line** has the following properties:
//
// - A **totalAmount** that indicates the amount paid by a customer or refunded to them by a company. A payment to the customer is always negative. A refund is always positive.
// - A **links** array containing one element with the following properties:
// - A **type** that indicates the type of link. For a payment this is `PaymentOnAccount`. For a refund this is `Refund`.
// - The **id** containing the ID of the customer.
// - The **amount** for the link is `0` – the **totalAmount** _or_ the amount of the payment or refund.
//
// It is possible to have a payment that is part _on account_ and part _allocated_ to an invoice. Each line should follow the examples above.
//
// ## Using a credit note to pay an invoice
//
// The payment of an invoice using a credit note has one entry in its **lines** array. This **line** has the following properties:
//
// - An **amount** that indicates the amount of money moved, which in this case is `0`, as the credit note and invoice allocation must balance each other.
// - A **links** array containing two elements:
//   - The first **link** has:
//   - A **type** that indicates the type of **link**, in this case an `Invoice`.
//   - An **id** that contains the ID of the invoice that was paid.
//   - The second **link** has:
//   - A **type** that indicates the type of **link**, in this case a `CreditNote`.
//   - An **id** that contains the ID of the credit note used by this payment.
//
// The **amount** field on the **line** equals the **totalAmount** on the payment.
//
// ## Refunding a credit note
//
// A payment refunding a credit note has one entry in its **lines** array. This **line** has the following properties:
//
// - An **amount** that indicates the amount of the credit note that was refunded. This is always negative for a refund.
// - A **links** array that contains one element with the following properties:
//   - A **type** that indicates the type of **link**, in this case a `CreditNote`.
//   - An **id** that contains the ID of the credit note that was refunded.
//
// The **totalAmount** field on the payment equals the **amount** field of the **line**. These are both negative, as this is money leaving accounts receivable.
//
// ## Refunding a payment
//
// If a payment is refunded, for example, if a customer overpaid an invoice and the overpayment is returned to the customer, there are two payment records:
//
// - One for the incoming over payment.
// - Another for the outgoing refund.
//
// The payment issuing the refund has a negative **totalAmount**. This payment also has one entry in its lines array with the following properties:
//
// - An **amount** that indicates the amount that was refunded. This is always negative.
// - A **links** array that contains one element with the following properties:
//   - A **type** that indicates the type of **link**, in this case a `Payment`.
//   - An **id** that contains the ID of the payment that was refunded.
//
// The **amount** field on the **line** equals the **totalAmount** on the payment and is negative, as this is money leaving accounts receivable.
//
// The payment that was refunded has a line where the **amount** is positive and the type of the link is `Refund`. This payment may have several entries in its **lines** array if it was used to partly pay an invoice.
//
// For example: A £1,050 payment on a £1,000 invoice with a refund of £50 has two lines:
//
// - One for £1,000 linked to the invoice that was paid.
// - Another for £50 linked to the payment that refunded the overpayment with a** type** of `Refund` and an ID that corresponds to the payment.
//
// The **line** linked to the payment has the following properties:
//
// - An **amount** that indicates the amount that was refunded. This is positive as its money that was added to accounts receivable. It's balanced out by the negative amount of the refund.
// - A **links** array containing one element with the following properties:
//   - A **type** that indicates the type of **link**, in this case a `Refund`.
//   - An **id** that contains the ID of the payment that refunded this line.
//
// > 📘 Support for linked payments
// >
// > Not all accounting packages support linking payments in this way. In some platforms, you may see a payment on account and a refund on account.
//
// ## Foreign currencies
//
// There are two types of currency rate that are included in the payments data type:
//
// Payment currency rate:
//
// - Base currency of the accounts receivable account.
// - Foreign currency of the payment.
//
// Payment line link currency rate:
//
// - Base currency of the item the link represents.
// - Foreign currency of the payment.
//
// These two rates allow the calculation of currency loss or gain for any of the transactions affected by the payment lines. The second rate is used when a payment is applied to an item in a currency that doesn't match either:
//
// - The base currency for the accounts receivable account.
// - The currency of the item.
//
// ```json Currency rate example
//
//	{
//	    "id": "123",
//	    "note": ""
//	    "totalAmount": 99.99,
//	    "currency": "GBP",
//	    "lines": [
//	        {
//	            "amount": 99.99,
//	            "links": [
//	                {
//	                    "type": "Invoice",
//	                    "id": "178",
//	                    "amount": -50,
//	                    "currencyRate":  1.9998,
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ## Example data
//
// > 📘 Object properties
// >
// > For the sake of brevity, the examples here may omit properties from objects. For the full object definition, see [Payments](https://api.codat.io/swagger/index.html#/Payments).
//
// ## Simple examples
//
// ```json Payment for invoice
//
//	{
//	    "totalAmount": 1000,
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Allocation of credit note
//
//	{
//	    "totalAmount": 0,
//	    "lines": [
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Payment of invoice and payment on account
//
//	{
//	    "totalAmount": 2000,
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "y",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Refund of credit note
//
//	{
//	    "totalAmount": -1000,
//	    "lines": [
//	        {
//	            "amount" : -1000,
//	            "links" : [
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Refund on accounts receivable account
//
//	{
//	    "totalAmount": -1000,
//	    "lines": [
//	        {
//	            "amount" : -1000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Linked refund on accounts receivable account
//
//	{
//	    "id" : "payment-001",
//	    "totalAmount": 1000,
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Refund",
//	                    "id" : "refund-001",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
//	{
//	    "id" : "refund-001",
//	    "totalAmount": -1000,
//	    "lines": [
//	        {
//	            "amount" : -1000,
//	            "links" : [
//	                {
//	                    "type" : "Payment",
//	                    "id" : "payment-001",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Using a credit note and cash to pay an invoice
//
//	{
//	    "totalAmount": 250,
//	    "lines": [
//	        {
//	            "amount": 0,
//	            "links": [
//	                {
//	                    "type": "Invoice",
//	                    "id": "x",
//	                    "amount": -750
//	                },
//	                {
//	                    "type": "CreditNote",
//	                    "id": "y",
//	                    "amount": 750
//	                }
//	            ]
//	        },
//	        {
//	            "amount": 250,
//	            "links": [
//	                {
//	                    "type": "Invoice",
//	                    "id": "x",
//	                    "amount": -250
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ## Complex examples
//
// ```json Use two credit notes and 1000 in to "bank" (cash, cheque etc.) to pay invoice
//
//	{
//	    "totalAmount": 1000,
//	    "lines": [
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "z",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Pay an invoice with two credit notes and cash, with 1000 left "on account"
//
//	{
//	    "totalAmount": 2000,
//	    "lines": [
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "z",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "customer-001",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Two credit notes pay two invoices with no allocation amount specified
//
//	{
//	    "totalAmount": 0,
//	    "lines": [
//	        {
//	            "amount" : 0,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "w",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "z",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Two credit notes and cash pay three invoices with no allocation amount specified, and refund cash
//
//	{
//	    "totalAmount": 2000,
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "w",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "x",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "u",
//	                    "amount" : -1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "y",
//	                    "amount" : 1000
//	                },
//	                {
//	                    "type" : "CreditNote",
//	                    "id" : "z",
//	                    "amount" : 1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Refund",
//	                    "id" : "refund-001",
//	                    "amount" : -1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
//	{
//	    "id" : "refund-001",
//	    "totalAmount": -1000,
//	    "lines": [
//	        {
//	            "amount" : -1000,
//	            "links" : [
//	                {
//	                    "type" : "Payment",
//	                    "id" : "payment-001",
//	                    "amount" : 1000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// In this example, a payment on account is used to pay the same invoice in January and again in February.
//
// ```json January
//
//	{
//	    "id": "001",
//	    "totalAmount": 5000,
//	    "date" : "1901-01-01",
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "Invoice-x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 4000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "PaymentOnAccount-y",
//	                    "amount" : -4000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json February
//
//	{
//	    "id": "001",
//	    "totalAmount": 5000,
//	    "date" : "1901-02-01",
//	    "lines": [
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "Invoice-x",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 1000,
//	            "links" : [
//	                {
//	                    "type" : "Invoice",
//	                    "id" : "Invoice-y",
//	                    "amount" : -1000
//	                }
//	            ]
//	        },
//	        {
//	            "amount" : 3000,
//	            "links" : [
//	                {
//	                    "type" : "PaymentOnAccount",
//	                    "id" : "PaymentOnAccount-y",
//	                    "amount" : -3000
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
//
// ```json Two credit notes and some cash pay two invoices with no allocations specified
//
//	{
//	    "totalAmount": 500,
//	    "lines": [
//	        {
//	            "amount": 500,
//	            "links": [{
//	                    "type": "Invoice",
//	                    "id": "a",
//	                    "amount": -1000
//	                }, {
//	                    "type": "Invoice",
//	                    "id": "b",
//	                    "amount": -1000
//	                }, {
//	                    "type": "CreditNote",
//	                    "id": "y",
//	                    "amount": 750
//	                },{
//	                    "type": "CreditNote",
//	                    "id": "z",
//	                    "amount": 750
//	                }
//	            ]
//	        }
//	    ]
//	}
//
// ```
type CreatePayment200ApplicationJSONSourceModifiedDate struct {
	AccountRef         *CreatePayment200ApplicationJSONSourceModifiedDateAccountRef       `json:"accountRef,omitempty"`
	Currency           *string                                                            `json:"currency,omitempty"`
	CurrencyRate       *float64                                                           `json:"currencyRate,omitempty"`
	CustomerRef        *CreatePayment200ApplicationJSONSourceModifiedDateCustomerRef      `json:"customerRef,omitempty"`
	Date               time.Time                                                          `json:"date"`
	ID                 *string                                                            `json:"id,omitempty"`
	Lines              []CreatePayment200ApplicationJSONSourceModifiedDateLines           `json:"lines,omitempty"`
	Metadata           *CreatePayment200ApplicationJSONSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                         `json:"modifiedDate,omitempty"`
	Note               *string                                                            `json:"note,omitempty"`
	PaymentMethodRef   *CreatePayment200ApplicationJSONSourceModifiedDatePaymentMethodRef `json:"paymentMethodRef,omitempty"`
	Reference          *string                                                            `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                         `json:"sourceModifiedDate,omitempty"`
	SupplementalData   *CreatePayment200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TotalAmount        *float64                                                           `json:"totalAmount,omitempty"`
}

type CreatePayment200ApplicationJSONStatusEnum string

const (
	CreatePayment200ApplicationJSONStatusEnumPending  CreatePayment200ApplicationJSONStatusEnum = "Pending"
	CreatePayment200ApplicationJSONStatusEnumFailed   CreatePayment200ApplicationJSONStatusEnum = "Failed"
	CreatePayment200ApplicationJSONStatusEnumSuccess  CreatePayment200ApplicationJSONStatusEnum = "Success"
	CreatePayment200ApplicationJSONStatusEnumTimedOut CreatePayment200ApplicationJSONStatusEnum = "TimedOut"
)

type CreatePayment200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreatePayment200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreatePayment200ApplicationJSONValidation struct {
	Errors   []CreatePayment200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreatePayment200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreatePayment200ApplicationJSON struct {
	Changes           []CreatePayment200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                             `json:"companyId"`
	CompletedOnUtc    *time.Time                                         `json:"completedOnUtc,omitempty"`
	Data              *CreatePayment200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                             `json:"dataConnectionKey"`
	DataType          *string                                            `json:"dataType,omitempty"`
	ErrorMessage      *string                                            `json:"errorMessage,omitempty"`
	PushOperationKey  string                                             `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                          `json:"requestedOnUtc"`
	Status            CreatePayment200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                `json:"statusCode"`
	TimeoutInMinutes  *int                                               `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                               `json:"timeoutInSeconds,omitempty"`
	Validation        *CreatePayment200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreatePaymentResponse struct {
	ContentType                           string
	StatusCode                            int
	RawResponse                           *http.Response
	CreatePayment200ApplicationJSONObject *CreatePayment200ApplicationJSON
}
