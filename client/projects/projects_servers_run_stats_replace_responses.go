package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// ProjectsServersRunStatsReplaceReader is a Reader for the ProjectsServersRunStatsReplace structure.
type ProjectsServersRunStatsReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersRunStatsReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersRunStatsReplaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersRunStatsReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersRunStatsReplaceOK creates a ProjectsServersRunStatsReplaceOK with default headers values
func NewProjectsServersRunStatsReplaceOK() *ProjectsServersRunStatsReplaceOK {
	return &ProjectsServersRunStatsReplaceOK{}
}

/*ProjectsServersRunStatsReplaceOK handles this case with default header values.

ServerRunStatistics updated.
*/
type ProjectsServersRunStatsReplaceOK struct {
	Payload *models.ServerRunStatistics
}

func (o *ProjectsServersRunStatsReplaceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsServersRunStatsReplaceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/projects/{project}/servers/{server}/run-stats/{id}/][%d] projectsServersRunStatsReplaceOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersRunStatsReplaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerRunStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersRunStatsReplaceBadRequest creates a ProjectsServersRunStatsReplaceBadRequest with default headers values
func NewProjectsServersRunStatsReplaceBadRequest() *ProjectsServersRunStatsReplaceBadRequest {
	return &ProjectsServersRunStatsReplaceBadRequest{}
}

/*ProjectsServersRunStatsReplaceBadRequest handles this case with default header values.

Invalid data supplied.
*/
type ProjectsServersRunStatsReplaceBadRequest struct {
	Payload *models.ServerRunStatisticsError
}

func (o *ProjectsServersRunStatsReplaceBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsServersRunStatsReplaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/projects/{project}/servers/{server}/run-stats/{id}/][%d] projectsServersRunStatsReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersRunStatsReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerRunStatisticsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
