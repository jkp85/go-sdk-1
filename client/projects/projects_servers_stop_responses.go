// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsServersStopReader is a Reader for the ProjectsServersStop structure.
type ProjectsServersStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersStopCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersStopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersStopCreated creates a ProjectsServersStopCreated with default headers values
func NewProjectsServersStopCreated() *ProjectsServersStopCreated {
	return &ProjectsServersStopCreated{}
}

/*ProjectsServersStopCreated handles this case with default header values.

Server stopped.
*/
type ProjectsServersStopCreated struct {
}

func (o *ProjectsServersStopCreated) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/projects/{project_id}/servers/{id}/stop/][%d] projectsServersStopCreated ", 201)
}

func (o *ProjectsServersStopCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsServersStopBadRequest creates a ProjectsServersStopBadRequest with default headers values
func NewProjectsServersStopBadRequest() *ProjectsServersStopBadRequest {
	return &ProjectsServersStopBadRequest{}
}

/*ProjectsServersStopBadRequest handles this case with default header values.

Invalid data supplied.
*/
type ProjectsServersStopBadRequest struct {
}

func (o *ProjectsServersStopBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/projects/{project_id}/servers/{id}/stop/][%d] projectsServersStopBadRequest ", 400)
}

func (o *ProjectsServersStopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
