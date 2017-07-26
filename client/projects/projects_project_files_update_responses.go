// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

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

	case 404:
		result := NewProjectsProjectFilesUpdateNotFound()
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

ProjectFile updated.
*/
type ProjectsProjectFilesUpdateOK struct {
	Payload *models.ProjectFile
}

func (o *ProjectsProjectFilesUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project_id}/project_files/{id}/][%d] projectsProjectFilesUpdateOK  %+v", 200, o.Payload)
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

Invalid data supplied.
*/
type ProjectsProjectFilesUpdateBadRequest struct {
	Payload *models.ProjectFileError
}

func (o *ProjectsProjectFilesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project_id}/project_files/{id}/][%d] projectsProjectFilesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsProjectFilesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectFileError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsProjectFilesUpdateNotFound creates a ProjectsProjectFilesUpdateNotFound with default headers values
func NewProjectsProjectFilesUpdateNotFound() *ProjectsProjectFilesUpdateNotFound {
	return &ProjectsProjectFilesUpdateNotFound{}
}

/*ProjectsProjectFilesUpdateNotFound handles this case with default header values.

ProjectFile not found.
*/
type ProjectsProjectFilesUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsProjectFilesUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project_id}/project_files/{id}/][%d] projectsProjectFilesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsProjectFilesUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
