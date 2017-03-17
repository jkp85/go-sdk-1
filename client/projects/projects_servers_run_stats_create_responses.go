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

	"github.com/jkp85/go-sdk/models"
)

// ProjectsServersRunStatsCreateReader is a Reader for the ProjectsServersRunStatsCreate structure.
type ProjectsServersRunStatsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersRunStatsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersRunStatsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersRunStatsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersRunStatsCreateCreated creates a ProjectsServersRunStatsCreateCreated with default headers values
func NewProjectsServersRunStatsCreateCreated() *ProjectsServersRunStatsCreateCreated {
	return &ProjectsServersRunStatsCreateCreated{}
}

/*ProjectsServersRunStatsCreateCreated handles this case with default header values.

ServerRunStatistics created
*/
type ProjectsServersRunStatsCreateCreated struct {
	Payload *models.ServerRunStatistics
}

func (o *ProjectsServersRunStatsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/run-stats/][%d] projectsServersRunStatsCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsServersRunStatsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerRunStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersRunStatsCreateBadRequest creates a ProjectsServersRunStatsCreateBadRequest with default headers values
func NewProjectsServersRunStatsCreateBadRequest() *ProjectsServersRunStatsCreateBadRequest {
	return &ProjectsServersRunStatsCreateBadRequest{}
}

/*ProjectsServersRunStatsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersRunStatsCreateBadRequest struct {
	Payload ProjectsServersRunStatsCreateBadRequestBody
}

func (o *ProjectsServersRunStatsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/run-stats/][%d] projectsServersRunStatsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersRunStatsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsServersRunStatsCreateBadRequestBody projects servers run stats create bad request body
swagger:model ProjectsServersRunStatsCreateBadRequestBody
*/
type ProjectsServersRunStatsCreateBadRequestBody struct {

	// exit_code firld errors
	// Required: true
	ExitCode []string `json:"exit_code"`

	// id firld errors
	// Required: true
	ID []string `json:"id"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// size firld errors
	// Required: true
	Size []string `json:"size"`

	// stacktrace firld errors
	// Required: true
	Stacktrace []string `json:"stacktrace"`

	// start firld errors
	// Required: true
	Start []string `json:"start"`

	// stop firld errors
	// Required: true
	Stop []string `json:"stop"`
}

// Validate validates this projects servers run stats create bad request body
func (o *ProjectsServersRunStatsCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExitCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

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

	if err := o.validateStacktrace(formats); err != nil {
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

func (o *ProjectsServersRunStatsCreateBadRequestBody) validateExitCode(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersRunStatsCreateBadRequest"+"."+"exit_code", "body", o.ExitCode); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersRunStatsCreateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersRunStatsCreateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersRunStatsCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersRunStatsCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersRunStatsCreateBadRequestBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersRunStatsCreateBadRequest"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersRunStatsCreateBadRequestBody) validateStacktrace(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersRunStatsCreateBadRequest"+"."+"stacktrace", "body", o.Stacktrace); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersRunStatsCreateBadRequestBody) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersRunStatsCreateBadRequest"+"."+"start", "body", o.Start); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersRunStatsCreateBadRequestBody) validateStop(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersRunStatsCreateBadRequest"+"."+"stop", "body", o.Stop); err != nil {
		return err
	}

	return nil
}

/*ProjectsServersRunStatsCreateBody projects servers run stats create body
swagger:model ProjectsServersRunStatsCreateBody
*/
type ProjectsServersRunStatsCreateBody struct {

	// exit code
	ExitCode int64 `json:"exit_code,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// stacktrace
	Stacktrace string `json:"stacktrace,omitempty"`

	// start
	Start string `json:"start,omitempty"`

	// stop
	Stop string `json:"stop,omitempty"`
}
