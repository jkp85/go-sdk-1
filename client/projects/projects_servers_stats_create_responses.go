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

// ProjectsServersStatsCreateReader is a Reader for the ProjectsServersStatsCreate structure.
type ProjectsServersStatsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersStatsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersStatsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersStatsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersStatsCreateCreated creates a ProjectsServersStatsCreateCreated with default headers values
func NewProjectsServersStatsCreateCreated() *ProjectsServersStatsCreateCreated {
	return &ProjectsServersStatsCreateCreated{}
}

/*ProjectsServersStatsCreateCreated handles this case with default header values.

ServerStatistics created
*/
type ProjectsServersStatsCreateCreated struct {
	Payload *models.ServerStatistics
}

func (o *ProjectsServersStatsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/servers/{server_pk}/stats/][%d] projectsServersStatsCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsServersStatsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersStatsCreateBadRequest creates a ProjectsServersStatsCreateBadRequest with default headers values
func NewProjectsServersStatsCreateBadRequest() *ProjectsServersStatsCreateBadRequest {
	return &ProjectsServersStatsCreateBadRequest{}
}

/*ProjectsServersStatsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersStatsCreateBadRequest struct {
	Payload ProjectsServersStatsCreateBadRequestBody
}

func (o *ProjectsServersStatsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/servers/{server_pk}/stats/][%d] projectsServersStatsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersStatsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsServersStatsCreateBadRequestBody projects servers stats create bad request body
swagger:model ProjectsServersStatsCreateBadRequestBody
*/
type ProjectsServersStatsCreateBadRequestBody struct {

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// size field errors
	// Required: true
	Size []string `json:"size"`

	// start field errors
	// Required: true
	Start []string `json:"start"`

	// stop field errors
	// Required: true
	Stop []string `json:"stop"`
}

// Validate validates this projects servers stats create bad request body
func (o *ProjectsServersStatsCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStop(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsServersStatsCreateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsCreateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersStatsCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersStatsCreateBadRequestBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsCreateBadRequest"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersStatsCreateBadRequestBody) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsCreateBadRequest"+"."+"start", "body", o.Start); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersStatsCreateBadRequestBody) validateStop(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsCreateBadRequest"+"."+"stop", "body", o.Stop); err != nil {
		return err
	}

	return nil
}

/*ProjectsServersStatsCreateBody projects servers stats create body
swagger:model ProjectsServersStatsCreateBody
*/
type ProjectsServersStatsCreateBody struct {

	// size
	Size int64 `json:"size,omitempty"`

	// start
	Start string `json:"start,omitempty"`

	// stop
	Stop string `json:"stop,omitempty"`
}