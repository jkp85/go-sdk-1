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

// UsersIntegrationsUpdateReader is a Reader for the UsersIntegrationsUpdate structure.
type UsersIntegrationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersIntegrationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersIntegrationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersIntegrationsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUsersIntegrationsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersIntegrationsUpdateOK creates a UsersIntegrationsUpdateOK with default headers values
func NewUsersIntegrationsUpdateOK() *UsersIntegrationsUpdateOK {
	return &UsersIntegrationsUpdateOK{}
}

/*UsersIntegrationsUpdateOK handles this case with default header values.

Integration updated
*/
type UsersIntegrationsUpdateOK struct {
	Payload *models.Integration
}

func (o *UsersIntegrationsUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/{user_id}/integrations/{id}/][%d] usersIntegrationsUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersIntegrationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersIntegrationsUpdateBadRequest creates a UsersIntegrationsUpdateBadRequest with default headers values
func NewUsersIntegrationsUpdateBadRequest() *UsersIntegrationsUpdateBadRequest {
	return &UsersIntegrationsUpdateBadRequest{}
}

/*UsersIntegrationsUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type UsersIntegrationsUpdateBadRequest struct {
	Payload *models.IntegrationError
}

func (o *UsersIntegrationsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/{user_id}/integrations/{id}/][%d] usersIntegrationsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersIntegrationsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersIntegrationsUpdateNotFound creates a UsersIntegrationsUpdateNotFound with default headers values
func NewUsersIntegrationsUpdateNotFound() *UsersIntegrationsUpdateNotFound {
	return &UsersIntegrationsUpdateNotFound{}
}

/*UsersIntegrationsUpdateNotFound handles this case with default header values.

Integration not found
*/
type UsersIntegrationsUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *UsersIntegrationsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/{user_id}/integrations/{id}/][%d] usersIntegrationsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UsersIntegrationsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
