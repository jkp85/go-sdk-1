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

// BillingInvoicesReadReader is a Reader for the BillingInvoicesRead structure.
type BillingInvoicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingInvoicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBillingInvoicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewBillingInvoicesReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingInvoicesReadOK creates a BillingInvoicesReadOK with default headers values
func NewBillingInvoicesReadOK() *BillingInvoicesReadOK {
	return &BillingInvoicesReadOK{}
}

/*BillingInvoicesReadOK handles this case with default header values.

Invoice retrieved.
*/
type BillingInvoicesReadOK struct {
	Payload *models.Invoice
}

func (o *BillingInvoicesReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *BillingInvoicesReadOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/billing/invoices/{id}/][%d] billingInvoicesReadOK  %+v", 200, o.Payload)
}

func (o *BillingInvoicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Invoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingInvoicesReadNotFound creates a BillingInvoicesReadNotFound with default headers values
func NewBillingInvoicesReadNotFound() *BillingInvoicesReadNotFound {
	return &BillingInvoicesReadNotFound{}
}

/*BillingInvoicesReadNotFound handles this case with default header values.

Invoice not found.
*/
type BillingInvoicesReadNotFound struct {
	Payload *models.NotFound
}

func (o *BillingInvoicesReadNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BillingInvoicesReadNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/billing/invoices/{id}/][%d] billingInvoicesReadNotFound  %+v", 404, o.Payload)
}

func (o *BillingInvoicesReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
