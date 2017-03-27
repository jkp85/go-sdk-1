package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AuthInvalidateSessionsCreateReader is a Reader for the AuthInvalidateSessionsCreate structure.
type AuthInvalidateSessionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthInvalidateSessionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAuthInvalidateSessionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAuthInvalidateSessionsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthInvalidateSessionsCreateCreated creates a AuthInvalidateSessionsCreateCreated with default headers values
func NewAuthInvalidateSessionsCreateCreated() *AuthInvalidateSessionsCreateCreated {
	return &AuthInvalidateSessionsCreateCreated{}
}

/*AuthInvalidateSessionsCreateCreated handles this case with default header values.

Invalidate Sessions
*/
type AuthInvalidateSessionsCreateCreated struct {
}

func (o *AuthInvalidateSessionsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /auth/invalidate-sessions/][%d] authInvalidateSessionsCreateCreated ", 201)
}

func (o *AuthInvalidateSessionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthInvalidateSessionsCreateBadRequest creates a AuthInvalidateSessionsCreateBadRequest with default headers values
func NewAuthInvalidateSessionsCreateBadRequest() *AuthInvalidateSessionsCreateBadRequest {
	return &AuthInvalidateSessionsCreateBadRequest{}
}

/*AuthInvalidateSessionsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type AuthInvalidateSessionsCreateBadRequest struct {
}

func (o *AuthInvalidateSessionsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /auth/invalidate-sessions/][%d] authInvalidateSessionsCreateBadRequest ", 400)
}

func (o *AuthInvalidateSessionsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
