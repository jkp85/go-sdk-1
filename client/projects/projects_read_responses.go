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

// ProjectsReadReader is a Reader for the ProjectsRead structure.
type ProjectsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsReadOK creates a ProjectsReadOK with default headers values
func NewProjectsReadOK() *ProjectsReadOK {
	return &ProjectsReadOK{}
}

/*ProjectsReadOK handles this case with default header values.

Project retrieved
*/
type ProjectsReadOK struct {
	Payload *models.Project
}

func (o *ProjectsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{id}/][%d] projectsReadOK  %+v", 200, o.Payload)
}

func (o *ProjectsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsReadNotFound creates a ProjectsReadNotFound with default headers values
func NewProjectsReadNotFound() *ProjectsReadNotFound {
	return &ProjectsReadNotFound{}
}

/*ProjectsReadNotFound handles this case with default header values.

Project not found
*/
type ProjectsReadNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{id}/][%d] projectsReadNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
