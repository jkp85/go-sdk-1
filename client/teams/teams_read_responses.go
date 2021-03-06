package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// TeamsReadReader is a Reader for the TeamsRead structure.
type TeamsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTeamsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewTeamsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsReadOK creates a TeamsReadOK with default headers values
func NewTeamsReadOK() *TeamsReadOK {
	return &TeamsReadOK{}
}

/*TeamsReadOK handles this case with default header values.

Team retrieved successfully.
*/
type TeamsReadOK struct {
	Payload *models.Team
}

func (o *TeamsReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsReadOK) Error() string {
	return fmt.Sprintf("[GET /v1/teams/{team}/][%d] teamsReadOK  %+v", 200, o.Payload)
}

func (o *TeamsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamsReadNotFound creates a TeamsReadNotFound with default headers values
func NewTeamsReadNotFound() *TeamsReadNotFound {
	return &TeamsReadNotFound{}
}

/*TeamsReadNotFound handles this case with default header values.

Team not found.
*/
type TeamsReadNotFound struct {
	Payload *models.NotFound
}

func (o *TeamsReadNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/teams/{team}/][%d] teamsReadNotFound  %+v", 404, o.Payload)
}

func (o *TeamsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
