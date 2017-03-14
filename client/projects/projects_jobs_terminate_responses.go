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

// ProjectsJobsTerminateReader is a Reader for the ProjectsJobsTerminate structure.
type ProjectsJobsTerminateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsJobsTerminateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsJobsTerminateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsJobsTerminateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsJobsTerminateCreated creates a ProjectsJobsTerminateCreated with default headers values
func NewProjectsJobsTerminateCreated() *ProjectsJobsTerminateCreated {
	return &ProjectsJobsTerminateCreated{}
}

/*ProjectsJobsTerminateCreated handles this case with default header values.

Job created
*/
type ProjectsJobsTerminateCreated struct {
	Payload *models.Job
}

func (o *ProjectsJobsTerminateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/jobs/{server}/terminate/][%d] projectsJobsTerminateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsJobsTerminateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsJobsTerminateBadRequest creates a ProjectsJobsTerminateBadRequest with default headers values
func NewProjectsJobsTerminateBadRequest() *ProjectsJobsTerminateBadRequest {
	return &ProjectsJobsTerminateBadRequest{}
}

/*ProjectsJobsTerminateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsJobsTerminateBadRequest struct {
}

func (o *ProjectsJobsTerminateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/jobs/{server}/terminate/][%d] projectsJobsTerminateBadRequest ", 400)
}

func (o *ProjectsJobsTerminateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsJobsTerminateBody projects jobs terminate body
swagger:model ProjectsJobsTerminateBody
*/
type ProjectsJobsTerminateBody struct {

	// method
	Method string `json:"method,omitempty"`

	// script
	// Required: true
	Script *string `json:"script"`

	// server
	// Required: true
	Server *models.Server `json:"server"`
}
