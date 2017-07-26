// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// TriggersCreateReader is a Reader for the TriggersCreate structure.
type TriggersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewTriggersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTriggersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTriggersCreateCreated creates a TriggersCreateCreated with default headers values
func NewTriggersCreateCreated() *TriggersCreateCreated {
	return &TriggersCreateCreated{}
}

/*TriggersCreateCreated handles this case with default header values.

Trigger created
*/
type TriggersCreateCreated struct {
	Payload *models.Trigger
}

func (o *TriggersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/triggers/][%d] triggersCreateCreated  %+v", 201, o.Payload)
}

func (o *TriggersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Trigger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggersCreateBadRequest creates a TriggersCreateBadRequest with default headers values
func NewTriggersCreateBadRequest() *TriggersCreateBadRequest {
	return &TriggersCreateBadRequest{}
}

/*TriggersCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type TriggersCreateBadRequest struct {
	Payload *models.TriggerError
}

func (o *TriggersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/triggers/][%d] triggersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *TriggersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TriggerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
