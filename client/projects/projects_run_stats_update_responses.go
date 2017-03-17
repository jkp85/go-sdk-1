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

// ProjectsRunStatsUpdateReader is a Reader for the ProjectsRunStatsUpdate structure.
type ProjectsRunStatsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsRunStatsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsRunStatsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsRunStatsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsRunStatsUpdateOK creates a ProjectsRunStatsUpdateOK with default headers values
func NewProjectsRunStatsUpdateOK() *ProjectsRunStatsUpdateOK {
	return &ProjectsRunStatsUpdateOK{}
}

/*ProjectsRunStatsUpdateOK handles this case with default header values.

ServerRunStatistics updated
*/
type ProjectsRunStatsUpdateOK struct {
	Payload *models.ServerRunStatistics
}

func (o *ProjectsRunStatsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/run-stats/{id}/][%d] projectsRunStatsUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsRunStatsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerRunStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsRunStatsUpdateBadRequest creates a ProjectsRunStatsUpdateBadRequest with default headers values
func NewProjectsRunStatsUpdateBadRequest() *ProjectsRunStatsUpdateBadRequest {
	return &ProjectsRunStatsUpdateBadRequest{}
}

/*ProjectsRunStatsUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsRunStatsUpdateBadRequest struct {
}

func (o *ProjectsRunStatsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/run-stats/{id}/][%d] projectsRunStatsUpdateBadRequest ", 400)
}

func (o *ProjectsRunStatsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsRunStatsUpdateBody projects run stats update body
swagger:model ProjectsRunStatsUpdateBody
*/
type ProjectsRunStatsUpdateBody struct {

	// exit code
	ExitCode int64 `json:"exit_code,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// stacktrace
	Stacktrace string `json:"stacktrace,omitempty"`

	// start
	Start string `json:"start,omitempty"`

	// stop
	Stop string `json:"stop,omitempty"`
}
