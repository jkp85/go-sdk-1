package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
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

func (o *UsersEmailsCreateCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersEmailsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/users/{user}/emails/][%d] usersEmailsCreateCreated  %+v", 201, o.Payload)
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
	Payload *models.EmailError
}

func (o *UsersEmailsCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersEmailsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/users/{user}/emails/][%d] usersEmailsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersEmailsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
