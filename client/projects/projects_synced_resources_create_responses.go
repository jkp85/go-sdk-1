package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// ProjectsSyncedResourcesCreateReader is a Reader for the ProjectsSyncedResourcesCreate structure.
type ProjectsSyncedResourcesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsSyncedResourcesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsSyncedResourcesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsSyncedResourcesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsSyncedResourcesCreateCreated creates a ProjectsSyncedResourcesCreateCreated with default headers values
func NewProjectsSyncedResourcesCreateCreated() *ProjectsSyncedResourcesCreateCreated {
	return &ProjectsSyncedResourcesCreateCreated{}
}

/*ProjectsSyncedResourcesCreateCreated handles this case with default header values.

SyncedResource created
*/
type ProjectsSyncedResourcesCreateCreated struct {
	Payload *models.SyncedResource
}

func (o *ProjectsSyncedResourcesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/synced-resources/][%d] projectsSyncedResourcesCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsSyncedResourcesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SyncedResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSyncedResourcesCreateBadRequest creates a ProjectsSyncedResourcesCreateBadRequest with default headers values
func NewProjectsSyncedResourcesCreateBadRequest() *ProjectsSyncedResourcesCreateBadRequest {
	return &ProjectsSyncedResourcesCreateBadRequest{}
}

/*ProjectsSyncedResourcesCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsSyncedResourcesCreateBadRequest struct {
	Payload ProjectsSyncedResourcesCreateBadRequestBody
}

func (o *ProjectsSyncedResourcesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /{namespace}/projects/{project_pk}/synced-resources/][%d] projectsSyncedResourcesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsSyncedResourcesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsSyncedResourcesCreateBadRequestBody projects synced resources create bad request body
swagger:model ProjectsSyncedResourcesCreateBadRequestBody
*/
type ProjectsSyncedResourcesCreateBadRequestBody struct {

	// folder field errors
	// Required: true
	Folder []string `json:"folder"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// provider field errors
	// Required: true
	Provider []string `json:"provider"`

	// settings field errors
	// Required: true
	Settings []string `json:"settings"`
}

// Validate validates this projects synced resources create bad request body
func (o *ProjectsSyncedResourcesCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFolder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateProvider(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSettings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsSyncedResourcesCreateBadRequestBody) validateFolder(formats strfmt.Registry) error {

	if err := validate.Required("projectsSyncedResourcesCreateBadRequest"+"."+"folder", "body", o.Folder); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsSyncedResourcesCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsSyncedResourcesCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsSyncedResourcesCreateBadRequestBody) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("projectsSyncedResourcesCreateBadRequest"+"."+"provider", "body", o.Provider); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsSyncedResourcesCreateBadRequestBody) validateSettings(formats strfmt.Registry) error {

	if err := validate.Required("projectsSyncedResourcesCreateBadRequest"+"."+"settings", "body", o.Settings); err != nil {
		return err
	}

	return nil
}

/*ProjectsSyncedResourcesCreateBody projects synced resources create body
swagger:model ProjectsSyncedResourcesCreateBody
*/
type ProjectsSyncedResourcesCreateBody struct {

	// folder
	// Required: true
	Folder *string `json:"folder"`

	// provider
	// Required: true
	Provider *string `json:"provider"`

	// settings
	Settings interface{} `json:"settings,omitempty"`
}