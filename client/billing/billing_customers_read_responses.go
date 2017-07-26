// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// BillingCustomersReadReader is a Reader for the BillingCustomersRead structure.
type BillingCustomersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingCustomersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBillingCustomersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewBillingCustomersReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingCustomersReadOK creates a BillingCustomersReadOK with default headers values
func NewBillingCustomersReadOK() *BillingCustomersReadOK {
	return &BillingCustomersReadOK{}
}

/*BillingCustomersReadOK handles this case with default header values.

Customer retrieved.
*/
type BillingCustomersReadOK struct {
	Payload *models.Customer
}

func (o *BillingCustomersReadOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/billing/customers/{id}/][%d] billingCustomersReadOK  %+v", 200, o.Payload)
}

func (o *BillingCustomersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingCustomersReadNotFound creates a BillingCustomersReadNotFound with default headers values
func NewBillingCustomersReadNotFound() *BillingCustomersReadNotFound {
	return &BillingCustomersReadNotFound{}
}

/*BillingCustomersReadNotFound handles this case with default header values.

Customer not found.
*/
type BillingCustomersReadNotFound struct {
	Payload *models.NotFound
}

func (o *BillingCustomersReadNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/billing/customers/{id}/][%d] billingCustomersReadNotFound  %+v", 404, o.Payload)
}

func (o *BillingCustomersReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
