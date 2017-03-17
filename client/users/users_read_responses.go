package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkp85/go-sdk/models"
)

// UsersReadReader is a Reader for the UsersRead structure.
type UsersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUsersReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersReadOK creates a UsersReadOK with default headers values
func NewUsersReadOK() *UsersReadOK {
	return &UsersReadOK{}
}

/*UsersReadOK handles this case with default header values.

User retrieved
*/
type UsersReadOK struct {
	Payload *models.User
}

func (o *UsersReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/users/{id}/][%d] usersReadOK  %+v", 200, o.Payload)
}

func (o *UsersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersReadNotFound creates a UsersReadNotFound with default headers values
func NewUsersReadNotFound() *UsersReadNotFound {
	return &UsersReadNotFound{}
}

/*UsersReadNotFound handles this case with default header values.

User not found
*/
type UsersReadNotFound struct {
	Payload *models.NotFound
}

func (o *UsersReadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/users/{id}/][%d] usersReadNotFound  %+v", 404, o.Payload)
}

func (o *UsersReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
