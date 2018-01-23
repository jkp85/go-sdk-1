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

// TeamsGroupsRemoveFromGroupReader is a Reader for the TeamsGroupsRemoveFromGroup structure.
type TeamsGroupsRemoveFromGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsGroupsRemoveFromGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTeamsGroupsRemoveFromGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTeamsGroupsRemoveFromGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTeamsGroupsRemoveFromGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsGroupsRemoveFromGroupOK creates a TeamsGroupsRemoveFromGroupOK with default headers values
func NewTeamsGroupsRemoveFromGroupOK() *TeamsGroupsRemoveFromGroupOK {
	return &TeamsGroupsRemoveFromGroupOK{}
}

/*TeamsGroupsRemoveFromGroupOK handles this case with default header values.

User removed from group.
*/
type TeamsGroupsRemoveFromGroupOK struct {
}

func (o *TeamsGroupsRemoveFromGroupOK) Error() string {
	return fmt.Sprintf("[POST /v1/teams/{team}/groups/{group}/remove/][%d] teamsGroupsRemoveFromGroupOK ", 200)
}

func (o *TeamsGroupsRemoveFromGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamsGroupsRemoveFromGroupBadRequest creates a TeamsGroupsRemoveFromGroupBadRequest with default headers values
func NewTeamsGroupsRemoveFromGroupBadRequest() *TeamsGroupsRemoveFromGroupBadRequest {
	return &TeamsGroupsRemoveFromGroupBadRequest{}
}

/*TeamsGroupsRemoveFromGroupBadRequest handles this case with default header values.

Invalid data supplied
*/
type TeamsGroupsRemoveFromGroupBadRequest struct {
}

func (o *TeamsGroupsRemoveFromGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/teams/{team}/groups/{group}/remove/][%d] teamsGroupsRemoveFromGroupBadRequest ", 400)
}

func (o *TeamsGroupsRemoveFromGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamsGroupsRemoveFromGroupNotFound creates a TeamsGroupsRemoveFromGroupNotFound with default headers values
func NewTeamsGroupsRemoveFromGroupNotFound() *TeamsGroupsRemoveFromGroupNotFound {
	return &TeamsGroupsRemoveFromGroupNotFound{}
}

/*TeamsGroupsRemoveFromGroupNotFound handles this case with default header values.

Group not found
*/
type TeamsGroupsRemoveFromGroupNotFound struct {
	Payload *models.NotFound
}

func (o *TeamsGroupsRemoveFromGroupNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsGroupsRemoveFromGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/teams/{team}/groups/{group}/remove/][%d] teamsGroupsRemoveFromGroupNotFound  %+v", 404, o.Payload)
}

func (o *TeamsGroupsRemoveFromGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}