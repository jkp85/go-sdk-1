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

// BillingCardsCreateReader is a Reader for the BillingCardsCreate structure.
type BillingCardsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingCardsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewBillingCardsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBillingCardsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingCardsCreateCreated creates a BillingCardsCreateCreated with default headers values
func NewBillingCardsCreateCreated() *BillingCardsCreateCreated {
	return &BillingCardsCreateCreated{}
}

/*BillingCardsCreateCreated handles this case with default header values.

Card created
*/
type BillingCardsCreateCreated struct {
	Payload *models.Card
}

func (o *BillingCardsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/billing/cards/][%d] billingCardsCreateCreated  %+v", 201, o.Payload)
}

func (o *BillingCardsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Card)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingCardsCreateBadRequest creates a BillingCardsCreateBadRequest with default headers values
func NewBillingCardsCreateBadRequest() *BillingCardsCreateBadRequest {
	return &BillingCardsCreateBadRequest{}
}

/*BillingCardsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type BillingCardsCreateBadRequest struct {
	Payload *models.CardError
}

func (o *BillingCardsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/billing/cards/][%d] billingCardsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *BillingCardsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
