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

// TeamsGroupsReplaceReader is a Reader for the TeamsGroupsReplace structure.
type TeamsGroupsReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsGroupsReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTeamsGroupsReplaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTeamsGroupsReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsGroupsReplaceOK creates a TeamsGroupsReplaceOK with default headers values
func NewTeamsGroupsReplaceOK() *TeamsGroupsReplaceOK {
	return &TeamsGroupsReplaceOK{}
}

/*TeamsGroupsReplaceOK handles this case with default header values.

Group replaced
*/
type TeamsGroupsReplaceOK struct {
	Payload *models.Group
}

func (o *TeamsGroupsReplaceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsGroupsReplaceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/teams/{team}/groups/{group}/][%d] teamsGroupsReplaceOK  %+v", 200, o.Payload)
}

func (o *TeamsGroupsReplaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamsGroupsReplaceBadRequest creates a TeamsGroupsReplaceBadRequest with default headers values
func NewTeamsGroupsReplaceBadRequest() *TeamsGroupsReplaceBadRequest {
	return &TeamsGroupsReplaceBadRequest{}
}

/*TeamsGroupsReplaceBadRequest handles this case with default header values.

Invalid data supplied
*/
type TeamsGroupsReplaceBadRequest struct {
	Payload *models.GroupError
}

func (o *TeamsGroupsReplaceBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsGroupsReplaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/teams/{team}/groups/{group}/][%d] teamsGroupsReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *TeamsGroupsReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
