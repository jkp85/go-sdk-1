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

// ProjectsCreateReader is a Reader for the ProjectsCreate structure.
type ProjectsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsCreateCreated creates a ProjectsCreateCreated with default headers values
func NewProjectsCreateCreated() *ProjectsCreateCreated {
	return &ProjectsCreateCreated{}
}

/*ProjectsCreateCreated handles this case with default header values.

Project created
*/
type ProjectsCreateCreated struct {
	Payload *models.Project
}

func (o *ProjectsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/][%d] projectsCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCreateBadRequest creates a ProjectsCreateBadRequest with default headers values
func NewProjectsCreateBadRequest() *ProjectsCreateBadRequest {
	return &ProjectsCreateBadRequest{}
}

/*ProjectsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsCreateBadRequest struct {
	Payload ProjectsCreateBadRequestBody
}

func (o *ProjectsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/][%d] projectsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsCreateBadRequestBody projects create bad request body
swagger:model ProjectsCreateBadRequestBody
*/
type ProjectsCreateBadRequestBody struct {

	// description firld errors
	// Required: true
	Description []string `json:"description"`

	// id firld errors
	// Required: true
	ID []string `json:"id"`

	// last_updated firld errors
	// Required: true
	LastUpdated []string `json:"last_updated"`

	// name firld errors
	// Required: true
	Name []string `json:"name"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// private firld errors
	// Required: true
	Private []string `json:"private"`
}

// Validate validates this projects create bad request body
func (o *ProjectsCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateLastUpdated(formats); err != nil {
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

	if err := o.validatePrivate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsCreateBadRequestBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("projectsCreateBadRequest"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCreateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsCreateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCreateBadRequestBody) validateLastUpdated(formats strfmt.Registry) error {

	if err := validate.Required("projectsCreateBadRequest"+"."+"last_updated", "body", o.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCreateBadRequestBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("projectsCreateBadRequest"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCreateBadRequestBody) validatePrivate(formats strfmt.Registry) error {

	if err := validate.Required("projectsCreateBadRequest"+"."+"private", "body", o.Private); err != nil {
		return err
	}

	return nil
}

/*ProjectsCreateBody projects create body
swagger:model ProjectsCreateBody
*/
type ProjectsCreateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// private
	Private bool `json:"private,omitempty"`
}
