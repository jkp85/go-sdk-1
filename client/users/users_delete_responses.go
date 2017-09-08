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

// UsersDeleteReader is a Reader for the UsersDelete structure.
type UsersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUsersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUsersDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersDeleteNoContent creates a UsersDeleteNoContent with default headers values
func NewUsersDeleteNoContent() *UsersDeleteNoContent {
	return &UsersDeleteNoContent{}
}

/*UsersDeleteNoContent handles this case with default header values.

User deleted.
*/
type UsersDeleteNoContent struct {
}

func (o *UsersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/profiles/{user}/][%d] usersDeleteNoContent ", 204)
}

func (o *UsersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersDeleteNotFound creates a UsersDeleteNotFound with default headers values
func NewUsersDeleteNotFound() *UsersDeleteNotFound {
	return &UsersDeleteNotFound{}
}

/*UsersDeleteNotFound handles this case with default header values.

User not found
*/
type UsersDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *UsersDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/profiles/{user}/][%d] usersDeleteNotFound  %+v", 404, o.Payload)
}

func (o *UsersDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
