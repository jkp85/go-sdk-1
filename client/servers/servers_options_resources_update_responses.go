package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// ServersOptionsResourcesUpdateReader is a Reader for the ServersOptionsResourcesUpdate structure.
type ServersOptionsResourcesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsResourcesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServersOptionsResourcesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServersOptionsResourcesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsResourcesUpdateOK creates a ServersOptionsResourcesUpdateOK with default headers values
func NewServersOptionsResourcesUpdateOK() *ServersOptionsResourcesUpdateOK {
	return &ServersOptionsResourcesUpdateOK{}
}

/*ServersOptionsResourcesUpdateOK handles this case with default header values.

EnvironmentResource updated
*/
type ServersOptionsResourcesUpdateOK struct {
	Payload *models.EnvironmentResource
}

func (o *ServersOptionsResourcesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/servers/options/resources/{id}/][%d] serversOptionsResourcesUpdateOK  %+v", 200, o.Payload)
}

func (o *ServersOptionsResourcesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsResourcesUpdateBadRequest creates a ServersOptionsResourcesUpdateBadRequest with default headers values
func NewServersOptionsResourcesUpdateBadRequest() *ServersOptionsResourcesUpdateBadRequest {
	return &ServersOptionsResourcesUpdateBadRequest{}
}

/*ServersOptionsResourcesUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ServersOptionsResourcesUpdateBadRequest struct {
	Payload ServersOptionsResourcesUpdateBadRequestBody
}

func (o *ServersOptionsResourcesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/servers/options/resources/{id}/][%d] serversOptionsResourcesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ServersOptionsResourcesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ServersOptionsResourcesUpdateBadRequestBody servers options resources update bad request body
swagger:model ServersOptionsResourcesUpdateBadRequestBody
*/
type ServersOptionsResourcesUpdateBadRequestBody struct {

	// active field errors
	// Required: true
	Active []string `json:"active"`

	// cpu field errors
	// Required: true
	CPU []string `json:"cpu"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// memory field errors
	// Required: true
	Memory []string `json:"memory"`

	// name field errors
	// Required: true
	Name []string `json:"name"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`
}

// Validate validates this servers options resources update bad request body
func (o *ServersOptionsResourcesUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActive(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCPU(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMemory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServersOptionsResourcesUpdateBadRequestBody) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesUpdateBadRequest"+"."+"active", "body", o.Active); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesUpdateBadRequestBody) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesUpdateBadRequest"+"."+"cpu", "body", o.CPU); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesUpdateBadRequestBody) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesUpdateBadRequest"+"."+"memory", "body", o.Memory); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesUpdateBadRequestBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesUpdateBadRequest"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

/*ServersOptionsResourcesUpdateBody servers options resources update body
swagger:model ServersOptionsResourcesUpdateBody
*/
type ServersOptionsResourcesUpdateBody struct {

	// active
	Active bool `json:"active,omitempty"`

	// cpu
	// Required: true
	CPU *int64 `json:"cpu"`

	// memory
	// Required: true
	Memory *int64 `json:"memory"`

	// name
	// Required: true
	Name *string `json:"name"`
}