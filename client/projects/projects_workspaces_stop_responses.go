package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkp85/go-sdk/models"
)

// ProjectsWorkspacesStopReader is a Reader for the ProjectsWorkspacesStop structure.
type ProjectsWorkspacesStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsWorkspacesStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsWorkspacesStopCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsWorkspacesStopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsWorkspacesStopCreated creates a ProjectsWorkspacesStopCreated with default headers values
func NewProjectsWorkspacesStopCreated() *ProjectsWorkspacesStopCreated {
	return &ProjectsWorkspacesStopCreated{}
}

/*ProjectsWorkspacesStopCreated handles this case with default header values.

Workspace created
*/
type ProjectsWorkspacesStopCreated struct {
	Payload *models.Workspace
}

func (o *ProjectsWorkspacesStopCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/workspaces/{server}/stop/][%d] projectsWorkspacesStopCreated  %+v", 201, o.Payload)
}

func (o *ProjectsWorkspacesStopCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workspace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsWorkspacesStopBadRequest creates a ProjectsWorkspacesStopBadRequest with default headers values
func NewProjectsWorkspacesStopBadRequest() *ProjectsWorkspacesStopBadRequest {
	return &ProjectsWorkspacesStopBadRequest{}
}

/*ProjectsWorkspacesStopBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsWorkspacesStopBadRequest struct {
}

func (o *ProjectsWorkspacesStopBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/workspaces/{server}/stop/][%d] projectsWorkspacesStopBadRequest ", 400)
}

func (o *ProjectsWorkspacesStopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsWorkspacesStopBody projects workspaces stop body
swagger:model ProjectsWorkspacesStopBody
*/
type ProjectsWorkspacesStopBody struct {

	// server
	// Required: true
	Server *models.Server `json:"server"`
}
