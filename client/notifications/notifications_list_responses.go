package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// NotificationsListReader is a Reader for the NotificationsList structure.
type NotificationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNotificationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNotificationsListOK creates a NotificationsListOK with default headers values
func NewNotificationsListOK() *NotificationsListOK {
	return &NotificationsListOK{}
}

/*NotificationsListOK handles this case with default header values.

List of notifications
*/
type NotificationsListOK struct {
	Payload models.NotificationsListOKBody
}

func (o *NotificationsListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/notifications/][%d] notificationsListOK  %+v", 200, o.Payload)
}

func (o *NotificationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
