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

// ServersOptionsTypesCreateReader is a Reader for the ServersOptionsTypesCreate structure.
type ServersOptionsTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewServersOptionsTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServersOptionsTypesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsTypesCreateCreated creates a ServersOptionsTypesCreateCreated with default headers values
func NewServersOptionsTypesCreateCreated() *ServersOptionsTypesCreateCreated {
	return &ServersOptionsTypesCreateCreated{}
}

/*ServersOptionsTypesCreateCreated handles this case with default header values.

EnvironmentType created
*/
type ServersOptionsTypesCreateCreated struct {
	Payload *models.EnvironmentType
}

func (o *ServersOptionsTypesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/servers/options/types/][%d] serversOptionsTypesCreateCreated  %+v", 201, o.Payload)
}

func (o *ServersOptionsTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsTypesCreateBadRequest creates a ServersOptionsTypesCreateBadRequest with default headers values
func NewServersOptionsTypesCreateBadRequest() *ServersOptionsTypesCreateBadRequest {
	return &ServersOptionsTypesCreateBadRequest{}
}

/*ServersOptionsTypesCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ServersOptionsTypesCreateBadRequest struct {
}

func (o *ServersOptionsTypesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/servers/options/types/][%d] serversOptionsTypesCreateBadRequest ", 400)
}

func (o *ServersOptionsTypesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ServersOptionsTypesCreateBody servers options types create body
swagger:model ServersOptionsTypesCreateBody
*/
type ServersOptionsTypesCreateBody struct {

	// cmd
	Cmd string `json:"cmd,omitempty"`

	// container path
	ContainerPath string `json:"container_path,omitempty"`

	// container port
	ContainerPort int64 `json:"container_port,omitempty"`

	// image name
	// Required: true
	ImageName *string `json:"image_name"`

	// name
	// Required: true
	Name *string `json:"name"`

	// work dir
	WorkDir string `json:"work_dir,omitempty"`
}
