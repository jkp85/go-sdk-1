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

// ServersOptionsResourcesUpdateReader is a Reader for the ServersOptionsResourcesUpdate structure.
type ServersOptionsResourcesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsResourcesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServersOptionsResourcesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServersOptionsResourcesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsResourcesUpdateOK creates a ServersOptionsResourcesUpdateOK with default headers values
func NewServersOptionsResourcesUpdateOK() *ServersOptionsResourcesUpdateOK {
	return &ServersOptionsResourcesUpdateOK{}
}

/*ServersOptionsResourcesUpdateOK handles this case with default header values.

EnvironmentResource updated
*/
type ServersOptionsResourcesUpdateOK struct {
	Payload *models.EnvironmentResource
}

func (o *ServersOptionsResourcesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/servers/options/resources/{id}/][%d] serversOptionsResourcesUpdateOK  %+v", 200, o.Payload)
}

func (o *ServersOptionsResourcesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsResourcesUpdateBadRequest creates a ServersOptionsResourcesUpdateBadRequest with default headers values
func NewServersOptionsResourcesUpdateBadRequest() *ServersOptionsResourcesUpdateBadRequest {
	return &ServersOptionsResourcesUpdateBadRequest{}
}

/*ServersOptionsResourcesUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ServersOptionsResourcesUpdateBadRequest struct {
}

func (o *ServersOptionsResourcesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/servers/options/resources/{id}/][%d] serversOptionsResourcesUpdateBadRequest ", 400)
}

func (o *ServersOptionsResourcesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ServersOptionsResourcesUpdateBody servers options resources update body
swagger:model ServersOptionsResourcesUpdateBody
*/
type ServersOptionsResourcesUpdateBody struct {

	// active
	Active bool `json:"active,omitempty"`

	// cpu
	// Required: true
	CPU *int64 `json:"cpu"`

	// memory
	// Required: true
	Memory *int64 `json:"memory"`

	// name
	// Required: true
	Name *string `json:"name"`
}
