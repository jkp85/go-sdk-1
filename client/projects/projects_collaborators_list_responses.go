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

// ProjectsCollaboratorsListReader is a Reader for the ProjectsCollaboratorsList structure.
type ProjectsCollaboratorsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsCollaboratorsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsCollaboratorsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsCollaboratorsListOK creates a ProjectsCollaboratorsListOK with default headers values
func NewProjectsCollaboratorsListOK() *ProjectsCollaboratorsListOK {
	return &ProjectsCollaboratorsListOK{}
}

/*ProjectsCollaboratorsListOK handles this case with default header values.

Collaborator list.
*/
type ProjectsCollaboratorsListOK struct {
	Payload []*models.Collaborator
}

func (o *ProjectsCollaboratorsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/projects/{project_id}/collaborators/][%d] projectsCollaboratorsListOK  %+v", 200, o.Payload)
}

func (o *ProjectsCollaboratorsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
