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

// ProjectsFilesCreateReader is a Reader for the ProjectsFilesCreate structure.
type ProjectsFilesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsFilesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsFilesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsFilesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsFilesCreateCreated creates a ProjectsFilesCreateCreated with default headers values
func NewProjectsFilesCreateCreated() *ProjectsFilesCreateCreated {
	return &ProjectsFilesCreateCreated{}
}

/*ProjectsFilesCreateCreated handles this case with default header values.

File created
*/
type ProjectsFilesCreateCreated struct {
	Payload *models.File
}

func (o *ProjectsFilesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/files/][%d] projectsFilesCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsFilesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.File)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsFilesCreateBadRequest creates a ProjectsFilesCreateBadRequest with default headers values
func NewProjectsFilesCreateBadRequest() *ProjectsFilesCreateBadRequest {
	return &ProjectsFilesCreateBadRequest{}
}

/*ProjectsFilesCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsFilesCreateBadRequest struct {
	Payload ProjectsFilesCreateBadRequestBody
}

func (o *ProjectsFilesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/files/][%d] projectsFilesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsFilesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsFilesCreateBadRequestBody projects files create bad request body
swagger:model ProjectsFilesCreateBadRequestBody
*/
type ProjectsFilesCreateBadRequestBody struct {

	// author field errors
	// Required: true
	Author []string `json:"author"`

	// content field errors
	// Required: true
	Content []string `json:"content"`

	// encoding field errors
	// Required: true
	Encoding []string `json:"encoding"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// path field errors
	// Required: true
	Path []string `json:"path"`

	// project field errors
	// Required: true
	Project []string `json:"project"`

	// public field errors
	// Required: true
	Public []string `json:"public"`

	// size field errors
	// Required: true
	Size []string `json:"size"`
}

// Validate validates this projects files create bad request body
func (o *ProjectsFilesCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEncoding(formats); err != nil {
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

	if err := o.validatePath(formats); err != nil {
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

	if err := o.validateSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsFilesCreateBadRequestBody) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesCreateBadRequest"+"."+"author", "body", o.Author); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesCreateBadRequestBody) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesCreateBadRequest"+"."+"content", "body", o.Content); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesCreateBadRequestBody) validateEncoding(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesCreateBadRequest"+"."+"encoding", "body", o.Encoding); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesCreateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesCreateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesCreateBadRequestBody) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesCreateBadRequest"+"."+"path", "body", o.Path); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesCreateBadRequestBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesCreateBadRequest"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesCreateBadRequestBody) validatePublic(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesCreateBadRequest"+"."+"public", "body", o.Public); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesCreateBadRequestBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesCreateBadRequest"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	return nil
}

/*ProjectsFilesCreateBody projects files create body
swagger:model ProjectsFilesCreateBody
*/
type ProjectsFilesCreateBody struct {

	// author
	// Required: true
	Author *string `json:"author"`

	// content
	// Required: true
	Content *string `json:"content"`

	// encoding
	// Required: true
	Encoding *string `json:"encoding"`

	// path
	// Required: true
	Path *string `json:"path"`

	// project
	// Required: true
	Project *string `json:"project"`

	// public
	Public bool `json:"public,omitempty"`
}
