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

// ProjectsWorkspacesCreateReader is a Reader for the ProjectsWorkspacesCreate structure.
type ProjectsWorkspacesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsWorkspacesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsWorkspacesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsWorkspacesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsWorkspacesCreateCreated creates a ProjectsWorkspacesCreateCreated with default headers values
func NewProjectsWorkspacesCreateCreated() *ProjectsWorkspacesCreateCreated {
	return &ProjectsWorkspacesCreateCreated{}
}

/*ProjectsWorkspacesCreateCreated handles this case with default header values.

Workspace created
*/
type ProjectsWorkspacesCreateCreated struct {
	Payload *models.Workspace
}

func (o *ProjectsWorkspacesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/workspaces/][%d] projectsWorkspacesCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsWorkspacesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workspace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsWorkspacesCreateBadRequest creates a ProjectsWorkspacesCreateBadRequest with default headers values
func NewProjectsWorkspacesCreateBadRequest() *ProjectsWorkspacesCreateBadRequest {
	return &ProjectsWorkspacesCreateBadRequest{}
}

/*ProjectsWorkspacesCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsWorkspacesCreateBadRequest struct {
}

func (o *ProjectsWorkspacesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/workspaces/][%d] projectsWorkspacesCreateBadRequest ", 400)
}

func (o *ProjectsWorkspacesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsWorkspacesCreateBody projects workspaces create body
swagger:model ProjectsWorkspacesCreateBody
*/
type ProjectsWorkspacesCreateBody struct {

	// server
	// Required: true
	Server *models.Server `json:"server"`
}