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

// NotificationSettingsEntityReadReader is a Reader for the NotificationSettingsEntityRead structure.
type NotificationSettingsEntityReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationSettingsEntityReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNotificationSettingsEntityReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewNotificationSettingsEntityReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNotificationSettingsEntityReadOK creates a NotificationSettingsEntityReadOK with default headers values
func NewNotificationSettingsEntityReadOK() *NotificationSettingsEntityReadOK {
	return &NotificationSettingsEntityReadOK{}
}

/*NotificationSettingsEntityReadOK handles this case with default header values.

Global notification settings
*/
type NotificationSettingsEntityReadOK struct {
	Payload models.NotificationSettingsEntityReadOKBody
}

func (o *NotificationSettingsEntityReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationSettingsEntityReadOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/notifications/settings/entity/{entity}][%d] notificationSettingsEntityReadOK  %+v", 200, o.Payload)
}

func (o *NotificationSettingsEntityReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationSettingsEntityReadNotFound creates a NotificationSettingsEntityReadNotFound with default headers values
func NewNotificationSettingsEntityReadNotFound() *NotificationSettingsEntityReadNotFound {
	return &NotificationSettingsEntityReadNotFound{}
}

/*NotificationSettingsEntityReadNotFound handles this case with default header values.

Notification not found.
*/
type NotificationSettingsEntityReadNotFound struct {
	Payload *models.NotFound
}

func (o *NotificationSettingsEntityReadNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationSettingsEntityReadNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/notifications/settings/entity/{entity}][%d] notificationSettingsEntityReadNotFound  %+v", 404, o.Payload)
}

func (o *NotificationSettingsEntityReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
