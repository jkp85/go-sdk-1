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

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersUpdateOK creates a UsersUpdateOK with default headers values
func NewUsersUpdateOK() *UsersUpdateOK {
	return &UsersUpdateOK{}
}

/*UsersUpdateOK handles this case with default header values.

User updated
*/
type UsersUpdateOK struct {
	Payload *models.User
}

func (o *UsersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/users/{id}/][%d] usersUpdateOK  %+v", 200, o.Payload)
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

Invalid data supplied
*/
type UsersUpdateBadRequest struct {
}

func (o *UsersUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/users/{id}/][%d] usersUpdateBadRequest ", 400)
}

func (o *UsersUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*UsersUpdateBody users update body
swagger:model UsersUpdateBody
*/
type UsersUpdateBody struct {

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// password
	// Required: true
	Password *string `json:"password"`

	// profile
	// Required: true
	Profile interface{} `json:"profile"`

	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	// Required: true
	Username *string `json:"username"`
}
