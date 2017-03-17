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

// ProjectsDataSourcesStopReader is a Reader for the ProjectsDataSourcesStop structure.
type ProjectsDataSourcesStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDataSourcesStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsDataSourcesStopCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsDataSourcesStopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsDataSourcesStopCreated creates a ProjectsDataSourcesStopCreated with default headers values
func NewProjectsDataSourcesStopCreated() *ProjectsDataSourcesStopCreated {
	return &ProjectsDataSourcesStopCreated{}
}

/*ProjectsDataSourcesStopCreated handles this case with default header values.

DataSource created
*/
type ProjectsDataSourcesStopCreated struct {
	Payload *models.DataSource
}

func (o *ProjectsDataSourcesStopCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/data-sources/{server}/stop/][%d] projectsDataSourcesStopCreated  %+v", 201, o.Payload)
}

func (o *ProjectsDataSourcesStopCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDataSourcesStopBadRequest creates a ProjectsDataSourcesStopBadRequest with default headers values
func NewProjectsDataSourcesStopBadRequest() *ProjectsDataSourcesStopBadRequest {
	return &ProjectsDataSourcesStopBadRequest{}
}

/*ProjectsDataSourcesStopBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsDataSourcesStopBadRequest struct {
}

func (o *ProjectsDataSourcesStopBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/data-sources/{server}/stop/][%d] projectsDataSourcesStopBadRequest ", 400)
}

func (o *ProjectsDataSourcesStopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsDataSourcesStopBody projects data sources stop body
swagger:model ProjectsDataSourcesStopBody
*/
type ProjectsDataSourcesStopBody struct {

	// server
	// Required: true
	Server *models.Server `json:"server"`
}
