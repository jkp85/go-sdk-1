package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// BillingInvoiceItemsListReader is a Reader for the BillingInvoiceItemsList structure.
type BillingInvoiceItemsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingInvoiceItemsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBillingInvoiceItemsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingInvoiceItemsListOK creates a BillingInvoiceItemsListOK with default headers values
func NewBillingInvoiceItemsListOK() *BillingInvoiceItemsListOK {
	return &BillingInvoiceItemsListOK{}
}

/*BillingInvoiceItemsListOK handles this case with default header values.

InvoiceItem list.
*/
type BillingInvoiceItemsListOK struct {
	Payload models.BillingInvoiceItemsListOKBody
}

func (o *BillingInvoiceItemsListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *BillingInvoiceItemsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/billing/invoices/{invoice_id}/invoice-items/][%d] billingInvoiceItemsListOK  %+v", 200, o.Payload)
}

func (o *BillingInvoiceItemsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}