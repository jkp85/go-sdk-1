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

// ProjectsServersUpdateReader is a Reader for the ProjectsServersUpdate structure.
type ProjectsServersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProjectsServersUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersUpdateOK creates a ProjectsServersUpdateOK with default headers values
func NewProjectsServersUpdateOK() *ProjectsServersUpdateOK {
	return &ProjectsServersUpdateOK{}
}

/*ProjectsServersUpdateOK handles this case with default header values.

Server updated
*/
type ProjectsServersUpdateOK struct {
	Payload *models.Server
}

func (o *ProjectsServersUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project_id}/servers/{id}/][%d] projectsServersUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersUpdateBadRequest creates a ProjectsServersUpdateBadRequest with default headers values
func NewProjectsServersUpdateBadRequest() *ProjectsServersUpdateBadRequest {
	return &ProjectsServersUpdateBadRequest{}
}

/*ProjectsServersUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersUpdateBadRequest struct {
	Payload *models.ServerData
}

func (o *ProjectsServersUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project_id}/servers/{id}/][%d] projectsServersUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersUpdateNotFound creates a ProjectsServersUpdateNotFound with default headers values
func NewProjectsServersUpdateNotFound() *ProjectsServersUpdateNotFound {
	return &ProjectsServersUpdateNotFound{}
}

/*ProjectsServersUpdateNotFound handles this case with default header values.

Server not found
*/
type ProjectsServersUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsServersUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project_id}/servers/{id}/][%d] projectsServersUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsServersUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
