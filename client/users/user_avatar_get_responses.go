package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserAvatarGetReader is a Reader for the UserAvatarGet structure.
type UserAvatarGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAvatarGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserAvatarGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserAvatarGetOK creates a UserAvatarGetOK with default headers values
func NewUserAvatarGetOK() *UserAvatarGetOK {
	return &UserAvatarGetOK{}
}

/*UserAvatarGetOK handles this case with default header values.

User avatar
*/
type UserAvatarGetOK struct {
}

func (o *UserAvatarGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/{user}/avatar/][%d] userAvatarGetOK ", 200)
}

func (o *UserAvatarGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
