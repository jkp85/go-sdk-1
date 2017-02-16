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

// ProjectsCollaboratorsPartialUpdateReader is a Reader for the ProjectsCollaboratorsPartialUpdate structure.
type ProjectsCollaboratorsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsCollaboratorsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsCollaboratorsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsCollaboratorsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProjectsCollaboratorsPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsCollaboratorsPartialUpdateOK creates a ProjectsCollaboratorsPartialUpdateOK with default headers values
func NewProjectsCollaboratorsPartialUpdateOK() *ProjectsCollaboratorsPartialUpdateOK {
	return &ProjectsCollaboratorsPartialUpdateOK{}
}

/*ProjectsCollaboratorsPartialUpdateOK handles this case with default header values.

Collaborator updated
*/
type ProjectsCollaboratorsPartialUpdateOK struct {
	Payload *models.Collaborator
}

func (o *ProjectsCollaboratorsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsCollaboratorsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collaborator)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCollaboratorsPartialUpdateBadRequest creates a ProjectsCollaboratorsPartialUpdateBadRequest with default headers values
func NewProjectsCollaboratorsPartialUpdateBadRequest() *ProjectsCollaboratorsPartialUpdateBadRequest {
	return &ProjectsCollaboratorsPartialUpdateBadRequest{}
}

/*ProjectsCollaboratorsPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsCollaboratorsPartialUpdateBadRequest struct {
}

func (o *ProjectsCollaboratorsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsPartialUpdateBadRequest ", 400)
}

func (o *ProjectsCollaboratorsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsCollaboratorsPartialUpdateNotFound creates a ProjectsCollaboratorsPartialUpdateNotFound with default headers values
func NewProjectsCollaboratorsPartialUpdateNotFound() *ProjectsCollaboratorsPartialUpdateNotFound {
	return &ProjectsCollaboratorsPartialUpdateNotFound{}
}

/*ProjectsCollaboratorsPartialUpdateNotFound handles this case with default header values.

Collaborator not found
*/
type ProjectsCollaboratorsPartialUpdateNotFound struct {
}

func (o *ProjectsCollaboratorsPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsPartialUpdateNotFound ", 404)
}

func (o *ProjectsCollaboratorsPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsCollaboratorsPartialUpdateBody projects collaborators partial update body
swagger:model ProjectsCollaboratorsPartialUpdateBody
*/
type ProjectsCollaboratorsPartialUpdateBody struct {

	// owner
	Owner bool `json:"owner,omitempty"`

	// user
	User *models.User `json:"user,omitempty"`
}
