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

// ProjectsJobsReadReader is a Reader for the ProjectsJobsRead structure.
type ProjectsJobsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsJobsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsJobsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsJobsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsJobsReadOK creates a ProjectsJobsReadOK with default headers values
func NewProjectsJobsReadOK() *ProjectsJobsReadOK {
	return &ProjectsJobsReadOK{}
}

/*ProjectsJobsReadOK handles this case with default header values.

Job retrieved
*/
type ProjectsJobsReadOK struct {
	Payload *models.Job
}

func (o *ProjectsJobsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/jobs/{server}/][%d] projectsJobsReadOK  %+v", 200, o.Payload)
}

func (o *ProjectsJobsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsJobsReadNotFound creates a ProjectsJobsReadNotFound with default headers values
func NewProjectsJobsReadNotFound() *ProjectsJobsReadNotFound {
	return &ProjectsJobsReadNotFound{}
}

/*ProjectsJobsReadNotFound handles this case with default header values.

Job not found
*/
type ProjectsJobsReadNotFound struct {
}

func (o *ProjectsJobsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/jobs/{server}/][%d] projectsJobsReadNotFound ", 404)
}

func (o *ProjectsJobsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}