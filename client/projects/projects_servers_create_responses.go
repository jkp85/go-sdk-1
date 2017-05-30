package projects

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

// ProjectsServersCreateReader is a Reader for the ProjectsServersCreate structure.
type ProjectsServersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersCreateCreated creates a ProjectsServersCreateCreated with default headers values
func NewProjectsServersCreateCreated() *ProjectsServersCreateCreated {
	return &ProjectsServersCreateCreated{}
}

/*ProjectsServersCreateCreated handles this case with default header values.

Server created
*/
type ProjectsServersCreateCreated struct {
	Payload *models.Server
}

func (o *ProjectsServersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/servers/][%d] projectsServersCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsServersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersCreateBadRequest creates a ProjectsServersCreateBadRequest with default headers values
func NewProjectsServersCreateBadRequest() *ProjectsServersCreateBadRequest {
	return &ProjectsServersCreateBadRequest{}
}

/*ProjectsServersCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersCreateBadRequest struct {
	Payload ProjectsServersCreateBadRequestBody
}

func (o *ProjectsServersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/servers/][%d] projectsServersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsServersCreateBadRequestBody projects servers create bad request body
swagger:model ProjectsServersCreateBadRequestBody
*/
type ProjectsServersCreateBadRequestBody struct {

	// config field errors
	// Required: true
	Config []string `json:"config"`

	// connected field errors
	// Required: true
	Connected []string `json:"connected"`

	// created_at field errors
	// Required: true
	CreatedAt []string `json:"created_at"`

	// endpoint field errors
	// Required: true
	Endpoint []string `json:"endpoint"`

	// environment_resources field errors
	// Required: true
	EnvironmentResources []string `json:"environment_resources"`

	// host field errors
	// Required: true
	Host []string `json:"host"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// image_name field errors
	// Required: true
	ImageName []string `json:"image_name"`

	// logs_url field errors
	// Required: true
	LogsURL []string `json:"logs_url"`

	// name field errors
	// Required: true
	Name []string `json:"name"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// startup_script field errors
	// Required: true
	StartupScript []string `json:"startup_script"`

	// status field errors
	// Required: true
	Status []string `json:"status"`

	// status_url field errors
	// Required: true
	StatusURL []string `json:"status_url"`
}

// Validate validates this projects servers create bad request body
func (o *ProjectsServersCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateConnected(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEndpoint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEnvironmentResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateHost(formats); err != nil {
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

	if err := o.validateLogsURL(formats); err != nil {
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

	if err := o.validateStartupScript(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStatusURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"config", "body", o.Config); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateConnected(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"connected", "body", o.Connected); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"endpoint", "body", o.Endpoint); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateEnvironmentResources(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"environment_resources", "body", o.EnvironmentResources); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"host", "body", o.Host); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateImageName(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"image_name", "body", o.ImageName); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateLogsURL(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"logs_url", "body", o.LogsURL); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateStartupScript(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"startup_script", "body", o.StartupScript); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersCreateBadRequestBody) validateStatusURL(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersCreateBadRequest"+"."+"status_url", "body", o.StatusURL); err != nil {
		return err
	}

	return nil
}

/*ProjectsServersCreateBody projects servers create body
swagger:model ProjectsServersCreateBody
*/
type ProjectsServersCreateBody struct {

	// config
	Config interface{} `json:"config,omitempty"`

	// connected
	Connected []string `json:"connected"`

	// environment resources
	EnvironmentResources string `json:"environment_resources,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// image name
	ImageName string `json:"image_name,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// startup script
	StartupScript string `json:"startup_script,omitempty"`
}
