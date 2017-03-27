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

// ServersOptionsResourcesReadReader is a Reader for the ServersOptionsResourcesRead structure.
type ServersOptionsResourcesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsResourcesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServersOptionsResourcesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewServersOptionsResourcesReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsResourcesReadOK creates a ServersOptionsResourcesReadOK with default headers values
func NewServersOptionsResourcesReadOK() *ServersOptionsResourcesReadOK {
	return &ServersOptionsResourcesReadOK{}
}

/*ServersOptionsResourcesReadOK handles this case with default header values.

EnvironmentResource retrieved
*/
type ServersOptionsResourcesReadOK struct {
	Payload *models.EnvironmentResource
}

func (o *ServersOptionsResourcesReadOK) Error() string {
	return fmt.Sprintf("[GET /{namespace}/servers/options/resources/{id}/][%d] serversOptionsResourcesReadOK  %+v", 200, o.Payload)
}

func (o *ServersOptionsResourcesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsResourcesReadNotFound creates a ServersOptionsResourcesReadNotFound with default headers values
func NewServersOptionsResourcesReadNotFound() *ServersOptionsResourcesReadNotFound {
	return &ServersOptionsResourcesReadNotFound{}
}

/*ServersOptionsResourcesReadNotFound handles this case with default header values.

EnvironmentResource not found
*/
type ServersOptionsResourcesReadNotFound struct {
	Payload *models.NotFound
}

func (o *ServersOptionsResourcesReadNotFound) Error() string {
	return fmt.Sprintf("[GET /{namespace}/servers/options/resources/{id}/][%d] serversOptionsResourcesReadNotFound  %+v", 404, o.Payload)
}

func (o *ServersOptionsResourcesReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
