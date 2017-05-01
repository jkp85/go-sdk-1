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

// UsersEmailsPartialUpdateReader is a Reader for the UsersEmailsPartialUpdate structure.
type UsersEmailsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersEmailsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersEmailsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersEmailsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUsersEmailsPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersEmailsPartialUpdateOK creates a UsersEmailsPartialUpdateOK with default headers values
func NewUsersEmailsPartialUpdateOK() *UsersEmailsPartialUpdateOK {
	return &UsersEmailsPartialUpdateOK{}
}

/*UsersEmailsPartialUpdateOK handles this case with default header values.

Email updated
*/
type UsersEmailsPartialUpdateOK struct {
	Payload *models.Email
}

func (o *UsersEmailsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/users/{user_pk}/emails/{address}/][%d] usersEmailsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersEmailsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Email)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersEmailsPartialUpdateBadRequest creates a UsersEmailsPartialUpdateBadRequest with default headers values
func NewUsersEmailsPartialUpdateBadRequest() *UsersEmailsPartialUpdateBadRequest {
	return &UsersEmailsPartialUpdateBadRequest{}
}

/*UsersEmailsPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type UsersEmailsPartialUpdateBadRequest struct {
	Payload UsersEmailsPartialUpdateBadRequestBody
}

func (o *UsersEmailsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/users/{user_pk}/emails/{address}/][%d] usersEmailsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersEmailsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersEmailsPartialUpdateNotFound creates a UsersEmailsPartialUpdateNotFound with default headers values
func NewUsersEmailsPartialUpdateNotFound() *UsersEmailsPartialUpdateNotFound {
	return &UsersEmailsPartialUpdateNotFound{}
}

/*UsersEmailsPartialUpdateNotFound handles this case with default header values.

Email not found
*/
type UsersEmailsPartialUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *UsersEmailsPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/users/{user_pk}/emails/{address}/][%d] usersEmailsPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UsersEmailsPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UsersEmailsPartialUpdateBadRequestBody users emails partial update bad request body
swagger:model UsersEmailsPartialUpdateBadRequestBody
*/
type UsersEmailsPartialUpdateBadRequestBody struct {

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

// Validate validates this users emails partial update bad request body
func (o *UsersEmailsPartialUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *UsersEmailsPartialUpdateBadRequestBody) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("usersEmailsPartialUpdateBadRequest"+"."+"address", "body", o.Address); err != nil {
		return err
	}

	return nil
}

func (o *UsersEmailsPartialUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("usersEmailsPartialUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *UsersEmailsPartialUpdateBadRequestBody) validatePublic(formats strfmt.Registry) error {

	if err := validate.Required("usersEmailsPartialUpdateBadRequest"+"."+"public", "body", o.Public); err != nil {
		return err
	}

	return nil
}

func (o *UsersEmailsPartialUpdateBadRequestBody) validateUnsubscribed(formats strfmt.Registry) error {

	if err := validate.Required("usersEmailsPartialUpdateBadRequest"+"."+"unsubscribed", "body", o.Unsubscribed); err != nil {
		return err
	}

	return nil
}

/*UsersEmailsPartialUpdateBody users emails partial update body
swagger:model UsersEmailsPartialUpdateBody
*/
type UsersEmailsPartialUpdateBody struct {

	// address
	Address string `json:"address,omitempty"`

	// public
	Public bool `json:"public,omitempty"`

	// unsubscribed
	Unsubscribed bool `json:"unsubscribed,omitempty"`
}