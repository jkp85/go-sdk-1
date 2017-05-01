package users

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

// UsersIntegrationsPartialUpdateReader is a Reader for the UsersIntegrationsPartialUpdate structure.
type UsersIntegrationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersIntegrationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersIntegrationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersIntegrationsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUsersIntegrationsPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersIntegrationsPartialUpdateOK creates a UsersIntegrationsPartialUpdateOK with default headers values
func NewUsersIntegrationsPartialUpdateOK() *UsersIntegrationsPartialUpdateOK {
	return &UsersIntegrationsPartialUpdateOK{}
}

/*UsersIntegrationsPartialUpdateOK handles this case with default header values.

Integration updated
*/
type UsersIntegrationsPartialUpdateOK struct {
	Payload *models.Integration
}

func (o *UsersIntegrationsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersIntegrationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersIntegrationsPartialUpdateBadRequest creates a UsersIntegrationsPartialUpdateBadRequest with default headers values
func NewUsersIntegrationsPartialUpdateBadRequest() *UsersIntegrationsPartialUpdateBadRequest {
	return &UsersIntegrationsPartialUpdateBadRequest{}
}

/*UsersIntegrationsPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type UsersIntegrationsPartialUpdateBadRequest struct {
	Payload UsersIntegrationsPartialUpdateBadRequestBody
}

func (o *UsersIntegrationsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersIntegrationsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersIntegrationsPartialUpdateNotFound creates a UsersIntegrationsPartialUpdateNotFound with default headers values
func NewUsersIntegrationsPartialUpdateNotFound() *UsersIntegrationsPartialUpdateNotFound {
	return &UsersIntegrationsPartialUpdateNotFound{}
}

/*UsersIntegrationsPartialUpdateNotFound handles this case with default header values.

Integration not found
*/
type UsersIntegrationsPartialUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *UsersIntegrationsPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UsersIntegrationsPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UsersIntegrationsPartialUpdateBadRequestBody users integrations partial update bad request body
swagger:model UsersIntegrationsPartialUpdateBadRequestBody
*/
type UsersIntegrationsPartialUpdateBadRequestBody struct {

	// extra_data field errors
	// Required: true
	ExtraData []string `json:"extra_data"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// provider field errors
	// Required: true
	Provider []string `json:"provider"`
}

// Validate validates this users integrations partial update bad request body
func (o *UsersIntegrationsPartialUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExtraData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateProvider(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersIntegrationsPartialUpdateBadRequestBody) validateExtraData(formats strfmt.Registry) error {

	if err := validate.Required("usersIntegrationsPartialUpdateBadRequest"+"."+"extra_data", "body", o.ExtraData); err != nil {
		return err
	}

	return nil
}

func (o *UsersIntegrationsPartialUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("usersIntegrationsPartialUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *UsersIntegrationsPartialUpdateBadRequestBody) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("usersIntegrationsPartialUpdateBadRequest"+"."+"provider", "body", o.Provider); err != nil {
		return err
	}

	return nil
}

/*UsersIntegrationsPartialUpdateBody users integrations partial update body
swagger:model UsersIntegrationsPartialUpdateBody
*/
type UsersIntegrationsPartialUpdateBody struct {

	// extra data
	ExtraData string `json:"extra_data,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`
}