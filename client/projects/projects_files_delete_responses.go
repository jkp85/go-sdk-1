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

// ProjectsFilesDeleteReader is a Reader for the ProjectsFilesDelete structure.
type ProjectsFilesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsFilesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsFilesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsFilesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsFilesDeleteNoContent creates a ProjectsFilesDeleteNoContent with default headers values
func NewProjectsFilesDeleteNoContent() *ProjectsFilesDeleteNoContent {
	return &ProjectsFilesDeleteNoContent{}
}

/*ProjectsFilesDeleteNoContent handles this case with default header values.

File deleted
*/
type ProjectsFilesDeleteNoContent struct {
}

func (o *ProjectsFilesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{namespace}/projects/{project_pk}/files/{id}/][%d] projectsFilesDeleteNoContent ", 204)
}

func (o *ProjectsFilesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsFilesDeleteNotFound creates a ProjectsFilesDeleteNotFound with default headers values
func NewProjectsFilesDeleteNotFound() *ProjectsFilesDeleteNotFound {
	return &ProjectsFilesDeleteNotFound{}
}

/*ProjectsFilesDeleteNotFound handles this case with default header values.

File not found
*/
type ProjectsFilesDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsFilesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{namespace}/projects/{project_pk}/files/{id}/][%d] projectsFilesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsFilesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
