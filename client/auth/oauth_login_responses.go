package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OauthLoginReader is a Reader for the OauthLogin structure.
type OauthLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OauthLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 302:
		result := NewOauthLoginFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOauthLoginFound creates a OauthLoginFound with default headers values
func NewOauthLoginFound() *OauthLoginFound {
	return &OauthLoginFound{}
}

/*OauthLoginFound handles this case with default header values.

Redirect to backend auth page
*/
type OauthLoginFound struct {
}

func (o *OauthLoginFound) Error() string {
	return fmt.Sprintf("[GET /auth/login/{provider}/][%d] oauthLoginFound ", 302)
}

func (o *OauthLoginFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
