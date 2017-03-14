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

// ProjectsModelsCreateReader is a Reader for the ProjectsModelsCreate structure.
type ProjectsModelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsModelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsModelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsModelsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsModelsCreateCreated creates a ProjectsModelsCreateCreated with default headers values
func NewProjectsModelsCreateCreated() *ProjectsModelsCreateCreated {
	return &ProjectsModelsCreateCreated{}
}

/*ProjectsModelsCreateCreated handles this case with default header values.

Model created
*/
type ProjectsModelsCreateCreated struct {
	Payload *models.Model
}

func (o *ProjectsModelsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/models/][%d] projectsModelsCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsModelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsModelsCreateBadRequest creates a ProjectsModelsCreateBadRequest with default headers values
func NewProjectsModelsCreateBadRequest() *ProjectsModelsCreateBadRequest {
	return &ProjectsModelsCreateBadRequest{}
}

/*ProjectsModelsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsModelsCreateBadRequest struct {
}

func (o *ProjectsModelsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/models/][%d] projectsModelsCreateBadRequest ", 400)
}

func (o *ProjectsModelsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsModelsCreateBody projects models create body
swagger:model ProjectsModelsCreateBody
*/
type ProjectsModelsCreateBody struct {

	// method
	Method string `json:"method,omitempty"`

	// script
	// Required: true
	Script *string `json:"script"`

	// server
	// Required: true
	Server *models.Server `json:"server"`
}
