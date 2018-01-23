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

// NotificationSettingsReadReader is a Reader for the NotificationSettingsRead structure.
type NotificationSettingsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationSettingsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNotificationSettingsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewNotificationSettingsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNotificationSettingsReadOK creates a NotificationSettingsReadOK with default headers values
func NewNotificationSettingsReadOK() *NotificationSettingsReadOK {
	return &NotificationSettingsReadOK{}
}

/*NotificationSettingsReadOK handles this case with default header values.

Global notification settings
*/
type NotificationSettingsReadOK struct {
	Payload models.NotificationSettingsReadOKBody
}

func (o *NotificationSettingsReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationSettingsReadOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/notifications/settings/][%d] notificationSettingsReadOK  %+v", 200, o.Payload)
}

func (o *NotificationSettingsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationSettingsReadNotFound creates a NotificationSettingsReadNotFound with default headers values
func NewNotificationSettingsReadNotFound() *NotificationSettingsReadNotFound {
	return &NotificationSettingsReadNotFound{}
}

/*NotificationSettingsReadNotFound handles this case with default header values.

Notification not found.
*/
type NotificationSettingsReadNotFound struct {
	Payload *models.NotFound
}

func (o *NotificationSettingsReadNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationSettingsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/notifications/settings/][%d] notificationSettingsReadNotFound  %+v", 404, o.Payload)
}

func (o *NotificationSettingsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
