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

// ProjectsServersStatsDeleteReader is a Reader for the ProjectsServersStatsDelete structure.
type ProjectsServersStatsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersStatsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsServersStatsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsServersStatsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersStatsDeleteNoContent creates a ProjectsServersStatsDeleteNoContent with default headers values
func NewProjectsServersStatsDeleteNoContent() *ProjectsServersStatsDeleteNoContent {
	return &ProjectsServersStatsDeleteNoContent{}
}

/*ProjectsServersStatsDeleteNoContent handles this case with default header values.

ServerStatistics deleted
*/
type ProjectsServersStatsDeleteNoContent struct {
}

func (o *ProjectsServersStatsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{namespace}/projects/{project_pk}/servers/{server_pk}/stats/{id}/][%d] projectsServersStatsDeleteNoContent ", 204)
}

func (o *ProjectsServersStatsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsServersStatsDeleteNotFound creates a ProjectsServersStatsDeleteNotFound with default headers values
func NewProjectsServersStatsDeleteNotFound() *ProjectsServersStatsDeleteNotFound {
	return &ProjectsServersStatsDeleteNotFound{}
}

/*ProjectsServersStatsDeleteNotFound handles this case with default header values.

ServerStatistics not found
*/
type ProjectsServersStatsDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsServersStatsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{namespace}/projects/{project_pk}/servers/{server_pk}/stats/{id}/][%d] projectsServersStatsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsServersStatsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
