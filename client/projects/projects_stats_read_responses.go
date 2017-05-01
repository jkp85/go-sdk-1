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

// ProjectsStatsReadReader is a Reader for the ProjectsStatsRead structure.
type ProjectsStatsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsStatsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsStatsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsStatsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsStatsReadOK creates a ProjectsStatsReadOK with default headers values
func NewProjectsStatsReadOK() *ProjectsStatsReadOK {
	return &ProjectsStatsReadOK{}
}

/*ProjectsStatsReadOK handles this case with default header values.

ServerStatistics retrieved
*/
type ProjectsStatsReadOK struct {
	Payload *models.ServerStatistics
}

func (o *ProjectsStatsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/stats/{id}/][%d] projectsStatsReadOK  %+v", 200, o.Payload)
}

func (o *ProjectsStatsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsStatsReadNotFound creates a ProjectsStatsReadNotFound with default headers values
func NewProjectsStatsReadNotFound() *ProjectsStatsReadNotFound {
	return &ProjectsStatsReadNotFound{}
}

/*ProjectsStatsReadNotFound handles this case with default header values.

ServerStatistics not found
*/
type ProjectsStatsReadNotFound struct {
}

func (o *ProjectsStatsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/stats/{id}/][%d] projectsStatsReadNotFound ", 404)
}

func (o *ProjectsStatsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}