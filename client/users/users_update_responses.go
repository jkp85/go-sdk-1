// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// UsersUpdateReader is a Reader for the UsersUpdate structure.
type UsersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUsersUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersUpdateOK creates a UsersUpdateOK with default headers values
func NewUsersUpdateOK() *UsersUpdateOK {
	return &UsersUpdateOK{}
}

/*UsersUpdateOK handles this case with default header values.

User updated.
*/
type UsersUpdateOK struct {
	Payload *models.User
}

func (o *UsersUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/profiles/{id}/][%d] usersUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUpdateBadRequest creates a UsersUpdateBadRequest with default headers values
func NewUsersUpdateBadRequest() *UsersUpdateBadRequest {
	return &UsersUpdateBadRequest{}
}

/*UsersUpdateBadRequest handles this case with default header values.

Invalid data supplied.
*/
type UsersUpdateBadRequest struct {
	Payload *models.UserError
}

func (o *UsersUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/profiles/{id}/][%d] usersUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUpdateNotFound creates a UsersUpdateNotFound with default headers values
func NewUsersUpdateNotFound() *UsersUpdateNotFound {
	return &UsersUpdateNotFound{}
}

/*UsersUpdateNotFound handles this case with default header values.

User not found.
*/
type UsersUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *UsersUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/profiles/{id}/][%d] usersUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UsersUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
