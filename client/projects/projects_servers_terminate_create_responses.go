package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsServersTerminateCreateReader is a Reader for the ProjectsServersTerminateCreate structure.
type ProjectsServersTerminateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersTerminateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersTerminateCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersTerminateCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersTerminateCreateCreated creates a ProjectsServersTerminateCreateCreated with default headers values
func NewProjectsServersTerminateCreateCreated() *ProjectsServersTerminateCreateCreated {
	return &ProjectsServersTerminateCreateCreated{}
}

/*ProjectsServersTerminateCreateCreated handles this case with default header values.

Terminate
*/
type ProjectsServersTerminateCreateCreated struct {
}

func (o *ProjectsServersTerminateCreateCreated) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/servers/{id}/terminate/][%d] projectsServersTerminateCreateCreated ", 201)
}

func (o *ProjectsServersTerminateCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsServersTerminateCreateBadRequest creates a ProjectsServersTerminateCreateBadRequest with default headers values
func NewProjectsServersTerminateCreateBadRequest() *ProjectsServersTerminateCreateBadRequest {
	return &ProjectsServersTerminateCreateBadRequest{}
}

/*ProjectsServersTerminateCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersTerminateCreateBadRequest struct {
}

func (o *ProjectsServersTerminateCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/servers/{id}/terminate/][%d] projectsServersTerminateCreateBadRequest ", 400)
}

func (o *ProjectsServersTerminateCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
