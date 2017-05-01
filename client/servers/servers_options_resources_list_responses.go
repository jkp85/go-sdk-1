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

// ServersOptionsResourcesListReader is a Reader for the ServersOptionsResourcesList structure.
type ServersOptionsResourcesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsResourcesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServersOptionsResourcesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsResourcesListOK creates a ServersOptionsResourcesListOK with default headers values
func NewServersOptionsResourcesListOK() *ServersOptionsResourcesListOK {
	return &ServersOptionsResourcesListOK{}
}

/*ServersOptionsResourcesListOK handles this case with default header values.

EnvironmentResource list
*/
type ServersOptionsResourcesListOK struct {
	Payload []*models.EnvironmentResource
}

func (o *ServersOptionsResourcesListOK) Error() string {
	return fmt.Sprintf("[GET /{namespace}/servers/options/resources/][%d] serversOptionsResourcesListOK  %+v", 200, o.Payload)
}

func (o *ServersOptionsResourcesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}