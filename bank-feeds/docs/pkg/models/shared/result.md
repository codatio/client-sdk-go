# Result


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Error`                                                                 | **string*                                                               | :heavy_minus_sign:                                                      | The error description for the attempted creation of the source account. | A bank account already exists with the same Id                          |
| `StatusCode`                                                            | **string*                                                               | :heavy_minus_sign:                                                      | The error status code for the attempted creation of the source account. | 409                                                                     |