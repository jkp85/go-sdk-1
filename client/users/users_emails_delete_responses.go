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

// UsersEmailsDeleteReader is a Reader for the UsersEmailsDelete structure.
type UsersEmailsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersEmailsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUsersEmailsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUsersEmailsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersEmailsDeleteNoContent creates a UsersEmailsDeleteNoContent with default headers values
func NewUsersEmailsDeleteNoContent() *UsersEmailsDeleteNoContent {
	return &UsersEmailsDeleteNoContent{}
}

/*UsersEmailsDeleteNoContent handles this case with default header values.

Email deleted
*/
type UsersEmailsDeleteNoContent struct {
}

func (o *UsersEmailsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{namespace}/users/{user_pk}/emails/{address}/][%d] usersEmailsDeleteNoContent ", 204)
}

func (o *UsersEmailsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersEmailsDeleteNotFound creates a UsersEmailsDeleteNotFound with default headers values
func NewUsersEmailsDeleteNotFound() *UsersEmailsDeleteNotFound {
	return &UsersEmailsDeleteNotFound{}
}

/*UsersEmailsDeleteNotFound handles this case with default header values.

Email not found
*/
type UsersEmailsDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *UsersEmailsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{namespace}/users/{user_pk}/emails/{address}/][%d] usersEmailsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *UsersEmailsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
