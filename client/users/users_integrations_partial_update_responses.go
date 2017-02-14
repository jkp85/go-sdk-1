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

// UsersIntegrationsPartialUpdateReader is a Reader for the UsersIntegrationsPartialUpdate structure.
type UsersIntegrationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersIntegrationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersIntegrationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersIntegrationsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUsersIntegrationsPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersIntegrationsPartialUpdateOK creates a UsersIntegrationsPartialUpdateOK with default headers values
func NewUsersIntegrationsPartialUpdateOK() *UsersIntegrationsPartialUpdateOK {
	return &UsersIntegrationsPartialUpdateOK{}
}

/*UsersIntegrationsPartialUpdateOK handles this case with default header values.

Integration updated
*/
type UsersIntegrationsPartialUpdateOK struct {
	Payload *models.Integration
}

func (o *UsersIntegrationsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersIntegrationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersIntegrationsPartialUpdateBadRequest creates a UsersIntegrationsPartialUpdateBadRequest with default headers values
func NewUsersIntegrationsPartialUpdateBadRequest() *UsersIntegrationsPartialUpdateBadRequest {
	return &UsersIntegrationsPartialUpdateBadRequest{}
}

/*UsersIntegrationsPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type UsersIntegrationsPartialUpdateBadRequest struct {
}

func (o *UsersIntegrationsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsPartialUpdateBadRequest ", 400)
}

func (o *UsersIntegrationsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersIntegrationsPartialUpdateNotFound creates a UsersIntegrationsPartialUpdateNotFound with default headers values
func NewUsersIntegrationsPartialUpdateNotFound() *UsersIntegrationsPartialUpdateNotFound {
	return &UsersIntegrationsPartialUpdateNotFound{}
}

/*UsersIntegrationsPartialUpdateNotFound handles this case with default header values.

Integration not found
*/
type UsersIntegrationsPartialUpdateNotFound struct {
}

func (o *UsersIntegrationsPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsPartialUpdateNotFound ", 404)
}

func (o *UsersIntegrationsPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*UsersIntegrationsPartialUpdateBody users integrations partial update body
swagger:model UsersIntegrationsPartialUpdateBody
*/
type UsersIntegrationsPartialUpdateBody struct {

	// integration email
	IntegrationEmail string `json:"integration_email,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// scopes
	Scopes string `json:"scopes,omitempty"`
}
