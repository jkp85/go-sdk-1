package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// ServersOptionsTypesUpdateReader is a Reader for the ServersOptionsTypesUpdate structure.
type ServersOptionsTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServersOptionsTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServersOptionsTypesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsTypesUpdateOK creates a ServersOptionsTypesUpdateOK with default headers values
func NewServersOptionsTypesUpdateOK() *ServersOptionsTypesUpdateOK {
	return &ServersOptionsTypesUpdateOK{}
}

/*ServersOptionsTypesUpdateOK handles this case with default header values.

EnvironmentType updated
*/
type ServersOptionsTypesUpdateOK struct {
	Payload *models.EnvironmentType
}

func (o *ServersOptionsTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/servers/options/types/{id}/][%d] serversOptionsTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *ServersOptionsTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsTypesUpdateBadRequest creates a ServersOptionsTypesUpdateBadRequest with default headers values
func NewServersOptionsTypesUpdateBadRequest() *ServersOptionsTypesUpdateBadRequest {
	return &ServersOptionsTypesUpdateBadRequest{}
}

/*ServersOptionsTypesUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ServersOptionsTypesUpdateBadRequest struct {
	Payload ServersOptionsTypesUpdateBadRequestBody
}

func (o *ServersOptionsTypesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/servers/options/types/{id}/][%d] serversOptionsTypesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ServersOptionsTypesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ServersOptionsTypesUpdateBadRequestBody servers options types update bad request body
swagger:model ServersOptionsTypesUpdateBadRequestBody
*/
type ServersOptionsTypesUpdateBadRequestBody struct {

	// cmd field errors
	// Required: true
	Cmd []string `json:"cmd"`

	// container_path field errors
	// Required: true
	ContainerPath []string `json:"container_path"`

	// container_port field errors
	// Required: true
	ContainerPort []string `json:"container_port"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// image_name field errors
	// Required: true
	ImageName []string `json:"image_name"`

	// name field errors
	// Required: true
	Name []string `json:"name"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// work_dir field errors
	// Required: true
	WorkDir []string `json:"work_dir"`
}

// Validate validates this servers options types update bad request body
func (o *ServersOptionsTypesUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCmd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateContainerPath(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateContainerPort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateImageName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateWorkDir(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ServersOptionsTypesUpdateBadRequestBody) validateCmd(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsTypesUpdateBadRequest"+"."+"cmd", "body", o.Cmd); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsTypesUpdateBadRequestBody) validateContainerPath(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsTypesUpdateBadRequest"+"."+"container_path", "body", o.ContainerPath); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsTypesUpdateBadRequestBody) validateContainerPort(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsTypesUpdateBadRequest"+"."+"container_port", "body", o.ContainerPort); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsTypesUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsTypesUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsTypesUpdateBadRequestBody) validateImageName(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsTypesUpdateBadRequest"+"."+"image_name", "body", o.ImageName); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsTypesUpdateBadRequestBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsTypesUpdateBadRequest"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsTypesUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsTypesUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ServersOptionsTypesUpdateBadRequestBody) validateWorkDir(formats strfmt.Registry) error {

	if err := validate.Required("serversOptionsTypesUpdateBadRequest"+"."+"work_dir", "body", o.WorkDir); err != nil {
		return err
	}

	return nil
}

/*ServersOptionsTypesUpdateBody servers options types update body
swagger:model ServersOptionsTypesUpdateBody
*/
type ServersOptionsTypesUpdateBody struct {

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