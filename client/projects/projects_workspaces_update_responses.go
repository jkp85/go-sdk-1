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

// ProjectsWorkspacesUpdateReader is a Reader for the ProjectsWorkspacesUpdate structure.
type ProjectsWorkspacesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsWorkspacesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsWorkspacesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsWorkspacesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsWorkspacesUpdateOK creates a ProjectsWorkspacesUpdateOK with default headers values
func NewProjectsWorkspacesUpdateOK() *ProjectsWorkspacesUpdateOK {
	return &ProjectsWorkspacesUpdateOK{}
}

/*ProjectsWorkspacesUpdateOK handles this case with default header values.

Workspace updated
*/
type ProjectsWorkspacesUpdateOK struct {
	Payload *models.Workspace
}

func (o *ProjectsWorkspacesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/workspaces/{server}/][%d] projectsWorkspacesUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsWorkspacesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workspace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsWorkspacesUpdateBadRequest creates a ProjectsWorkspacesUpdateBadRequest with default headers values
func NewProjectsWorkspacesUpdateBadRequest() *ProjectsWorkspacesUpdateBadRequest {
	return &ProjectsWorkspacesUpdateBadRequest{}
}

/*ProjectsWorkspacesUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsWorkspacesUpdateBadRequest struct {
}

func (o *ProjectsWorkspacesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/workspaces/{server}/][%d] projectsWorkspacesUpdateBadRequest ", 400)
}

func (o *ProjectsWorkspacesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsWorkspacesUpdateBody projects workspaces update body
swagger:model ProjectsWorkspacesUpdateBody
*/
type ProjectsWorkspacesUpdateBody struct {

	// server
	// Required: true
	Server *models.Server `json:"server"`
}