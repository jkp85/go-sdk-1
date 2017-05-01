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

// UsersEmailsCreateReader is a Reader for the UsersEmailsCreate structure.
type UsersEmailsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersEmailsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUsersEmailsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersEmailsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersEmailsCreateCreated creates a UsersEmailsCreateCreated with default headers values
func NewUsersEmailsCreateCreated() *UsersEmailsCreateCreated {
	return &UsersEmailsCreateCreated{}
}

/*UsersEmailsCreateCreated handles this case with default header values.

Email created
*/
type UsersEmailsCreateCreated struct {
	Payload *models.Email
}

func (o *UsersEmailsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /{namespace}/users/{user_pk}/emails/][%d] usersEmailsCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersEmailsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Email)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersEmailsCreateBadRequest creates a UsersEmailsCreateBadRequest with default headers values
func NewUsersEmailsCreateBadRequest() *UsersEmailsCreateBadRequest {
	return &UsersEmailsCreateBadRequest{}
}

/*UsersEmailsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type UsersEmailsCreateBadRequest struct {
	Payload UsersEmailsCreateBadRequestBody
}

func (o *UsersEmailsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /{namespace}/users/{user_pk}/emails/][%d] usersEmailsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersEmailsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UsersEmailsCreateBadRequestBody users emails create bad request body
swagger:model UsersEmailsCreateBadRequestBody
*/
type UsersEmailsCreateBadRequestBody struct {

	// address field errors
	// Required: true
	Address []string `json:"address"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// public field errors
	// Required: true
	Public []string `json:"public"`

	// unsubscribed field errors
	// Required: true
	Unsubscribed []string `json:"unsubscribed"`
}

// Validate validates this users emails create bad request body
func (o *UsersEmailsCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePublic(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateUnsubscribed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersEmailsCreateBadRequestBody) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("usersEmailsCreateBadRequest"+"."+"address", "body", o.Address); err != nil {
		return err
	}

	return nil
}

func (o *UsersEmailsCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("usersEmailsCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *UsersEmailsCreateBadRequestBody) validatePublic(formats strfmt.Registry) error {

	if err := validate.Required("usersEmailsCreateBadRequest"+"."+"public", "body", o.Public); err != nil {
		return err
	}

	return nil
}

func (o *UsersEmailsCreateBadRequestBody) validateUnsubscribed(formats strfmt.Registry) error {

	if err := validate.Required("usersEmailsCreateBadRequest"+"."+"unsubscribed", "body", o.Unsubscribed); err != nil {
		return err
	}

	return nil
}

/*UsersEmailsCreateBody users emails create body
swagger:model UsersEmailsCreateBody
*/
type UsersEmailsCreateBody struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// public
	Public bool `json:"public,omitempty"`

	// unsubscribed
	Unsubscribed bool `json:"unsubscribed,omitempty"`
}