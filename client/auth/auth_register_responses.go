package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// AuthRegisterReader is a Reader for the AuthRegister structure.
type AuthRegisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthRegisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAuthRegisterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAuthRegisterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthRegisterCreated creates a AuthRegisterCreated with default headers values
func NewAuthRegisterCreated() *AuthRegisterCreated {
	return &AuthRegisterCreated{}
}

/*AuthRegisterCreated handles this case with default header values.

User created
*/
type AuthRegisterCreated struct {
	Payload *models.User
}

func (o *AuthRegisterCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *AuthRegisterCreated) Error() string {
	return fmt.Sprintf("[POST /auth/register/][%d] authRegisterCreated  %+v", 201, o.Payload)
}

func (o *AuthRegisterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRegisterBadRequest creates a AuthRegisterBadRequest with default headers values
func NewAuthRegisterBadRequest() *AuthRegisterBadRequest {
	return &AuthRegisterBadRequest{}
}

/*AuthRegisterBadRequest handles this case with default header values.

Invalid data supplied
*/
type AuthRegisterBadRequest struct {
	Payload *models.UserError
}

func (o *AuthRegisterBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *AuthRegisterBadRequest) Error() string {
	return fmt.Sprintf("[POST /auth/register/][%d] authRegisterBadRequest  %+v", 400, o.Payload)
}

func (o *AuthRegisterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
