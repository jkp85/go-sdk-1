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

// UsersIntegrationsUpdateReader is a Reader for the UsersIntegrationsUpdate structure.
type UsersIntegrationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersIntegrationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersIntegrationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersIntegrationsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersIntegrationsUpdateOK creates a UsersIntegrationsUpdateOK with default headers values
func NewUsersIntegrationsUpdateOK() *UsersIntegrationsUpdateOK {
	return &UsersIntegrationsUpdateOK{}
}

/*UsersIntegrationsUpdateOK handles this case with default header values.

Integration updated
*/
type UsersIntegrationsUpdateOK struct {
	Payload *models.Integration
}

func (o *UsersIntegrationsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersIntegrationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersIntegrationsUpdateBadRequest creates a UsersIntegrationsUpdateBadRequest with default headers values
func NewUsersIntegrationsUpdateBadRequest() *UsersIntegrationsUpdateBadRequest {
	return &UsersIntegrationsUpdateBadRequest{}
}

/*UsersIntegrationsUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type UsersIntegrationsUpdateBadRequest struct {
	Payload UsersIntegrationsUpdateBadRequestBody
}

func (o *UsersIntegrationsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersIntegrationsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UsersIntegrationsUpdateBadRequestBody users integrations update bad request body
swagger:model UsersIntegrationsUpdateBadRequestBody
*/
type UsersIntegrationsUpdateBadRequestBody struct {

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

// Validate validates this users integrations update bad request body
func (o *UsersIntegrationsUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *UsersIntegrationsUpdateBadRequestBody) validateExtraData(formats strfmt.Registry) error {

	if err := validate.Required("usersIntegrationsUpdateBadRequest"+"."+"extra_data", "body", o.ExtraData); err != nil {
		return err
	}

	return nil
}

func (o *UsersIntegrationsUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("usersIntegrationsUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *UsersIntegrationsUpdateBadRequestBody) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("usersIntegrationsUpdateBadRequest"+"."+"provider", "body", o.Provider); err != nil {
		return err
	}

	return nil
}

/*UsersIntegrationsUpdateBody users integrations update body
swagger:model UsersIntegrationsUpdateBody
*/
type UsersIntegrationsUpdateBody struct {

	// extra data
	ExtraData string `json:"extra_data,omitempty"`

	// provider
	// Required: true
	Provider *string `json:"provider"`
}
