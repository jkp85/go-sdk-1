package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkp85/go-sdk/models"
)

// ServersOptionsTypesListReader is a Reader for the ServersOptionsTypesList structure.
type ServersOptionsTypesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsTypesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServersOptionsTypesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsTypesListOK creates a ServersOptionsTypesListOK with default headers values
func NewServersOptionsTypesListOK() *ServersOptionsTypesListOK {
	return &ServersOptionsTypesListOK{}
}

/*ServersOptionsTypesListOK handles this case with default header values.

EnvironmentType list
*/
type ServersOptionsTypesListOK struct {
	Payload []*models.EnvironmentType
}

func (o *ServersOptionsTypesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/servers/options/types/][%d] serversOptionsTypesListOK  %+v", 200, o.Payload)
}

func (o *ServersOptionsTypesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
