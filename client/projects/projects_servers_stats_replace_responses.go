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

// ProjectsServersStatsReplaceReader is a Reader for the ProjectsServersStatsReplace structure.
type ProjectsServersStatsReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersStatsReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersStatsReplaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersStatsReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersStatsReplaceOK creates a ProjectsServersStatsReplaceOK with default headers values
func NewProjectsServersStatsReplaceOK() *ProjectsServersStatsReplaceOK {
	return &ProjectsServersStatsReplaceOK{}
}

/*ProjectsServersStatsReplaceOK handles this case with default header values.

ServerStatistics updated
*/
type ProjectsServersStatsReplaceOK struct {
	Payload *models.ServerStatistics
}

func (o *ProjectsServersStatsReplaceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsServersStatsReplaceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/projects/{project}/servers/{server}/stats/{id}/][%d] projectsServersStatsReplaceOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersStatsReplaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersStatsReplaceBadRequest creates a ProjectsServersStatsReplaceBadRequest with default headers values
func NewProjectsServersStatsReplaceBadRequest() *ProjectsServersStatsReplaceBadRequest {
	return &ProjectsServersStatsReplaceBadRequest{}
}

/*ProjectsServersStatsReplaceBadRequest handles this case with default header values.

Invalid data supplied.
*/
type ProjectsServersStatsReplaceBadRequest struct {
	Payload *models.ServerStatisticsError
}

func (o *ProjectsServersStatsReplaceBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsServersStatsReplaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/projects/{project}/servers/{server}/stats/{id}/][%d] projectsServersStatsReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersStatsReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerStatisticsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
