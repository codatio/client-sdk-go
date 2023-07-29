# PaymentLinkType

Types of payment line links, either:  
`Unknown`  
`Unlinked` - Not used  
`Invoice` - ID refers to the invoice  
`CreditNote` - ID refers to the credit note  
`Refund` - ID refers to the sibling payment  
`Payment` - ID refers to the sibling payment  
`PaymentOnAccount` - ID refers to the customer  
`Other` - ID refers to the customer  
`Manual Journal`  
`Discount` - ID refers to the payment


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PaymentLinkTypeUnknown`          | Unknown                           |
| `PaymentLinkTypeUnlinked`         | Unlinked                          |
| `PaymentLinkTypeInvoice`          | Invoice                           |
| `PaymentLinkTypeCreditNote`       | CreditNote                        |
| `PaymentLinkTypeOther`            | Other                             |
| `PaymentLinkTypeRefund`           | Refund                            |
| `PaymentLinkTypePayment`          | Payment                           |
| `PaymentLinkTypePaymentOnAccount` | PaymentOnAccount                  |
| `PaymentLinkTypeManualJournal`    | ManualJournal                     |
| `PaymentLinkTypeDiscount`         | Discount                          |