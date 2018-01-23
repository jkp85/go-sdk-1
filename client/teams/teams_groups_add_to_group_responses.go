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

// TeamsGroupsAddToGroupReader is a Reader for the TeamsGroupsAddToGroup structure.
type TeamsGroupsAddToGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsGroupsAddToGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTeamsGroupsAddToGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTeamsGroupsAddToGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTeamsGroupsAddToGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsGroupsAddToGroupOK creates a TeamsGroupsAddToGroupOK with default headers values
func NewTeamsGroupsAddToGroupOK() *TeamsGroupsAddToGroupOK {
	return &TeamsGroupsAddToGroupOK{}
}

/*TeamsGroupsAddToGroupOK handles this case with default header values.

User added to group.
*/
type TeamsGroupsAddToGroupOK struct {
	Payload *models.GroupUser
}

func (o *TeamsGroupsAddToGroupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsGroupsAddToGroupOK) Error() string {
	return fmt.Sprintf("[POST /v1/teams/{team}/groups/{group}/add/][%d] teamsGroupsAddToGroupOK  %+v", 200, o.Payload)
}

func (o *TeamsGroupsAddToGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamsGroupsAddToGroupBadRequest creates a TeamsGroupsAddToGroupBadRequest with default headers values
func NewTeamsGroupsAddToGroupBadRequest() *TeamsGroupsAddToGroupBadRequest {
	return &TeamsGroupsAddToGroupBadRequest{}
}

/*TeamsGroupsAddToGroupBadRequest handles this case with default header values.

Invalid data supplied
*/
type TeamsGroupsAddToGroupBadRequest struct {
	Payload *models.GroupUserError
}

func (o *TeamsGroupsAddToGroupBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsGroupsAddToGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/teams/{team}/groups/{group}/add/][%d] teamsGroupsAddToGroupBadRequest  %+v", 400, o.Payload)
}

func (o *TeamsGroupsAddToGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupUserError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamsGroupsAddToGroupNotFound creates a TeamsGroupsAddToGroupNotFound with default headers values
func NewTeamsGroupsAddToGroupNotFound() *TeamsGroupsAddToGroupNotFound {
	return &TeamsGroupsAddToGroupNotFound{}
}

/*TeamsGroupsAddToGroupNotFound handles this case with default header values.

Group not found
*/
type TeamsGroupsAddToGroupNotFound struct {
	Payload *models.NotFound
}

func (o *TeamsGroupsAddToGroupNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsGroupsAddToGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/teams/{team}/groups/{group}/add/][%d] teamsGroupsAddToGroupNotFound  %+v", 404, o.Payload)
}

func (o *TeamsGroupsAddToGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}