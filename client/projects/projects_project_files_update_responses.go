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

// ProjectsProjectFilesUpdateReader is a Reader for the ProjectsProjectFilesUpdate structure.
type ProjectsProjectFilesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectFilesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsProjectFilesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsProjectFilesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsProjectFilesUpdateOK creates a ProjectsProjectFilesUpdateOK with default headers values
func NewProjectsProjectFilesUpdateOK() *ProjectsProjectFilesUpdateOK {
	return &ProjectsProjectFilesUpdateOK{}
}

/*ProjectsProjectFilesUpdateOK handles this case with default header values.

ProjectFile updated
*/
type ProjectsProjectFilesUpdateOK struct {
	Payload *models.ProjectFile
}

func (o *ProjectsProjectFilesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/projects/{project_pk}/project_files/{id}/][%d] projectsProjectFilesUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsProjectFilesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsProjectFilesUpdateBadRequest creates a ProjectsProjectFilesUpdateBadRequest with default headers values
func NewProjectsProjectFilesUpdateBadRequest() *ProjectsProjectFilesUpdateBadRequest {
	return &ProjectsProjectFilesUpdateBadRequest{}
}

/*ProjectsProjectFilesUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsProjectFilesUpdateBadRequest struct {
	Payload ProjectsProjectFilesUpdateBadRequestBody
}

func (o *ProjectsProjectFilesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/projects/{project_pk}/project_files/{id}/][%d] projectsProjectFilesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsProjectFilesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsProjectFilesUpdateBadRequestBody projects project files update bad request body
swagger:model ProjectsProjectFilesUpdateBadRequestBody
*/
type ProjectsProjectFilesUpdateBadRequestBody struct {

	// file field errors
	// Required: true
	File []string `json:"file"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// project field errors
	// Required: true
	Project []string `json:"project"`

	// public field errors
	// Required: true
	Public []string `json:"public"`
}

// Validate validates this projects project files update bad request body
func (o *ProjectsProjectFilesUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFile(formats); err != nil {
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

	if err := o.validateProject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePublic(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsProjectFilesUpdateBadRequestBody) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("projectsProjectFilesUpdateBadRequest"+"."+"file", "body", o.File); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsProjectFilesUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsProjectFilesUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsProjectFilesUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsProjectFilesUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsProjectFilesUpdateBadRequestBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("projectsProjectFilesUpdateBadRequest"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsProjectFilesUpdateBadRequestBody) validatePublic(formats strfmt.Registry) error {

	if err := validate.Required("projectsProjectFilesUpdateBadRequest"+"."+"public", "body", o.Public); err != nil {
		return err
	}

	return nil
}