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

// ProjectsFilesReadReader is a Reader for the ProjectsFilesRead structure.
type ProjectsFilesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsFilesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsFilesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsFilesReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsFilesReadOK creates a ProjectsFilesReadOK with default headers values
func NewProjectsFilesReadOK() *ProjectsFilesReadOK {
	return &ProjectsFilesReadOK{}
}

/*ProjectsFilesReadOK handles this case with default header values.

File retrieved
*/
type ProjectsFilesReadOK struct {
	Payload *models.File
}

func (o *ProjectsFilesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/files/{id}/][%d] projectsFilesReadOK  %+v", 200, o.Payload)
}

func (o *ProjectsFilesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.File)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsFilesReadNotFound creates a ProjectsFilesReadNotFound with default headers values
func NewProjectsFilesReadNotFound() *ProjectsFilesReadNotFound {
	return &ProjectsFilesReadNotFound{}
}

/*ProjectsFilesReadNotFound handles this case with default header values.

File not found
*/
type ProjectsFilesReadNotFound struct {
}

func (o *ProjectsFilesReadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/files/{id}/][%d] projectsFilesReadNotFound ", 404)
}

func (o *ProjectsFilesReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
