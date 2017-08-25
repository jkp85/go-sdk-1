// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// ServersOptionsServerSizeCreateReader is a Reader for the ServersOptionsServerSizeCreate structure.
type ServersOptionsServerSizeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsServerSizeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewServersOptionsServerSizeCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServersOptionsServerSizeCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsServerSizeCreateCreated creates a ServersOptionsServerSizeCreateCreated with default headers values
func NewServersOptionsServerSizeCreateCreated() *ServersOptionsServerSizeCreateCreated {
	return &ServersOptionsServerSizeCreateCreated{}
}

/*ServersOptionsServerSizeCreateCreated handles this case with default header values.

ServerSize created. This operation is available only to super users.
*/
type ServersOptionsServerSizeCreateCreated struct {
	Payload *models.ServerSize
}

func (o *ServersOptionsServerSizeCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/servers/options/server-size/][%d] serversOptionsServerSizeCreateCreated  %+v", 201, o.Payload)
}

func (o *ServersOptionsServerSizeCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerSize)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsServerSizeCreateBadRequest creates a ServersOptionsServerSizeCreateBadRequest with default headers values
func NewServersOptionsServerSizeCreateBadRequest() *ServersOptionsServerSizeCreateBadRequest {
	return &ServersOptionsServerSizeCreateBadRequest{}
}

/*ServersOptionsServerSizeCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ServersOptionsServerSizeCreateBadRequest struct {
	Payload *models.ServerSizeError
}

func (o *ServersOptionsServerSizeCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/servers/options/server-size/][%d] serversOptionsServerSizeCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ServersOptionsServerSizeCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerSizeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
