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

// ProjectsDeleteReader is a Reader for the ProjectsDelete structure.
type ProjectsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsDeleteNoContent creates a ProjectsDeleteNoContent with default headers values
func NewProjectsDeleteNoContent() *ProjectsDeleteNoContent {
	return &ProjectsDeleteNoContent{}
}

/*ProjectsDeleteNoContent handles this case with default header values.

Project deleted.
*/
type ProjectsDeleteNoContent struct {
}

func (o *ProjectsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/{namespace}/projects/{project}/][%d] projectsDeleteNoContent ", 204)
}

func (o *ProjectsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsDeleteNotFound creates a ProjectsDeleteNotFound with default headers values
func NewProjectsDeleteNotFound() *ProjectsDeleteNotFound {
	return &ProjectsDeleteNotFound{}
}

/*ProjectsDeleteNotFound handles this case with default header values.

Project not found.
*/
type ProjectsDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/{namespace}/projects/{project}/][%d] projectsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
