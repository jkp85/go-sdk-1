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

// TeamsGroupsUpdateReader is a Reader for the TeamsGroupsUpdate structure.
type TeamsGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTeamsGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTeamsGroupsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTeamsGroupsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsGroupsUpdateOK creates a TeamsGroupsUpdateOK with default headers values
func NewTeamsGroupsUpdateOK() *TeamsGroupsUpdateOK {
	return &TeamsGroupsUpdateOK{}
}

/*TeamsGroupsUpdateOK handles this case with default header values.

Group updated
*/
type TeamsGroupsUpdateOK struct {
	Payload *models.Group
}

func (o *TeamsGroupsUpdateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/teams/{team}/groups/{group}/][%d] teamsGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *TeamsGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamsGroupsUpdateBadRequest creates a TeamsGroupsUpdateBadRequest with default headers values
func NewTeamsGroupsUpdateBadRequest() *TeamsGroupsUpdateBadRequest {
	return &TeamsGroupsUpdateBadRequest{}
}

/*TeamsGroupsUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type TeamsGroupsUpdateBadRequest struct {
	Payload *models.GroupError
}

func (o *TeamsGroupsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsGroupsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/teams/{team}/groups/{group}/][%d] teamsGroupsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TeamsGroupsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamsGroupsUpdateNotFound creates a TeamsGroupsUpdateNotFound with default headers values
func NewTeamsGroupsUpdateNotFound() *TeamsGroupsUpdateNotFound {
	return &TeamsGroupsUpdateNotFound{}
}

/*TeamsGroupsUpdateNotFound handles this case with default header values.

Group not found.
*/
type TeamsGroupsUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *TeamsGroupsUpdateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsGroupsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/teams/{team}/groups/{group}/][%d] teamsGroupsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *TeamsGroupsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
