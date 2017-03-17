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

	"github.com/jkp85/go-sdk/models"
)

// ServersOptionsResourcesPartialUpdateReader is a Reader for the ServersOptionsResourcesPartialUpdate structure.
type ServersOptionsResourcesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsResourcesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServersOptionsResourcesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServersOptionsResourcesPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewServersOptionsResourcesPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsResourcesPartialUpdateOK creates a ServersOptionsResourcesPartialUpdateOK with default headers values
func NewServersOptionsResourcesPartialUpdateOK() *ServersOptionsResourcesPartialUpdateOK {
	return &ServersOptionsResourcesPartialUpdateOK{}
}

/*ServersOptionsResourcesPartialUpdateOK handles this case with default header values.

EnvironmentResource updated
*/
type ServersOptionsResourcesPartialUpdateOK struct {
	Payload *models.EnvironmentResource
}

func (o *ServersOptionsResourcesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/servers/options/resources/{id}/][%d] serversOptionsResourcesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ServersOptionsResourcesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsResourcesPartialUpdateBadRequest creates a ServersOptionsResourcesPartialUpdateBadRequest with default headers values
func NewServersOptionsResourcesPartialUpdateBadRequest() *ServersOptionsResourcesPartialUpdateBadRequest {
	return &ServersOptionsResourcesPartialUpdateBadRequest{}
}

/*ServersOptionsResourcesPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ServersOptionsResourcesPartialUpdateBadRequest struct {
	Payload ServersOptionsResourcesPartialUpdateBadRequestBody
}

func (o *ServersOptionsResourcesPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/servers/options/resources/{id}/][%d] serversOptionsResourcesPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ServersOptionsResourcesPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsResourcesPartialUpdateNotFound creates a ServersOptionsResourcesPartialUpdateNotFound with default headers values
func NewServersOptionsResourcesPartialUpdateNotFound() *ServersOptionsResourcesPartialUpdateNotFound {
	return &ServersOptionsResourcesPartialUpdateNotFound{}
}

/*ServersOptionsResourcesPartialUpdateNotFound handles this case with default header values.

EnvironmentResource not found
*/
type ServersOptionsResourcesPartialUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *ServersOptionsResourcesPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/servers/options/resources/{id}/][%d] serversOptionsResourcesPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ServersOptionsResourcesPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ServersOptionsResourcesPartialUpdateBadRequestBody servers options resources partial update bad request body
swagger:model ServersOptionsResourcesPartialUpdateBadRequestBody
*/
type ServersOptionsResourcesPartialUpdateBadRequestBody struct {

	// active firld errors
	// Required: true
	Active []string `json:"active"`

	// cpu firld errors
	// Required: true
	CPU []string `json:"cpu"`

	// id firld errors
	// Required: true
	ID []string `json:"id"`

	// memory firld errors
	// Required: true
	Memory []string `json:"memory"`

	// name firld errors
	// Required: true
	Name []string `json:"name"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`
}

// Validate validates this servers options resources partial update bad request body
func (o *ServersOptionsResourcesPartialUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *ServersOptionsResourcesPartialUpdateBadRequestBody) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesPartialUpdateBadRequest"+"."+"active", "body", o.Active); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesPartialUpdateBadRequestBody) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesPartialUpdateBadRequest"+"."+"cpu", "body", o.CPU); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesPartialUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesPartialUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesPartialUpdateBadRequestBody) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesPartialUpdateBadRequest"+"."+"memory", "body", o.Memory); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesPartialUpdateBadRequestBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesPartialUpdateBadRequest"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsResourcesPartialUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsResourcesPartialUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

/*ServersOptionsResourcesPartialUpdateBody servers options resources partial update body
swagger:model ServersOptionsResourcesPartialUpdateBody
*/
type ServersOptionsResourcesPartialUpdateBody struct {

	// active
	Active bool `json:"active,omitempty"`

	// cpu
	CPU int64 `json:"cpu,omitempty"`

	// memory
	Memory int64 `json:"memory,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}
