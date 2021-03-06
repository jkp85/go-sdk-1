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

// ServiceTriggerDeleteReader is a Reader for the ServiceTriggerDelete structure.
type ServiceTriggerDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceTriggerDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewServiceTriggerDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewServiceTriggerDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceTriggerDeleteNoContent creates a ServiceTriggerDeleteNoContent with default headers values
func NewServiceTriggerDeleteNoContent() *ServiceTriggerDeleteNoContent {
	return &ServiceTriggerDeleteNoContent{}
}

/*ServiceTriggerDeleteNoContent handles this case with default header values.

ServerAction deleted
*/
type ServiceTriggerDeleteNoContent struct {
}

func (o *ServiceTriggerDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/{namespace}/projects/{project}/servers/{server}/triggers/{trigger}/][%d] serviceTriggerDeleteNoContent ", 204)
}

func (o *ServiceTriggerDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceTriggerDeleteNotFound creates a ServiceTriggerDeleteNotFound with default headers values
func NewServiceTriggerDeleteNotFound() *ServiceTriggerDeleteNotFound {
	return &ServiceTriggerDeleteNotFound{}
}

/*ServiceTriggerDeleteNotFound handles this case with default header values.

ServerAction not found
*/
type ServiceTriggerDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *ServiceTriggerDeleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ServiceTriggerDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/{namespace}/projects/{project}/servers/{server}/triggers/{trigger}/][%d] serviceTriggerDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ServiceTriggerDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
