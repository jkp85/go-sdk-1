package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkp85/go-sdk/models"
)

// UsersPartialUpdateReader is a Reader for the UsersPartialUpdate structure.
type UsersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUsersPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersPartialUpdateOK creates a UsersPartialUpdateOK with default headers values
func NewUsersPartialUpdateOK() *UsersPartialUpdateOK {
	return &UsersPartialUpdateOK{}
}

/*UsersPartialUpdateOK handles this case with default header values.

User updated
*/
type UsersPartialUpdateOK struct {
	Payload *models.User
}

func (o *UsersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/users/{id}/][%d] usersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersPartialUpdateBadRequest creates a UsersPartialUpdateBadRequest with default headers values
func NewUsersPartialUpdateBadRequest() *UsersPartialUpdateBadRequest {
	return &UsersPartialUpdateBadRequest{}
}

/*UsersPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type UsersPartialUpdateBadRequest struct {
	Payload UsersPartialUpdateBadRequestBody
}

func (o *UsersPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/users/{id}/][%d] usersPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersPartialUpdateNotFound creates a UsersPartialUpdateNotFound with default headers values
func NewUsersPartialUpdateNotFound() *UsersPartialUpdateNotFound {
	return &UsersPartialUpdateNotFound{}
}

/*UsersPartialUpdateNotFound handles this case with default header values.

User not found
*/
type UsersPartialUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *UsersPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/users/{id}/][%d] usersPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UsersPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UsersPartialUpdateBadRequestBody users partial update bad request body
swagger:model UsersPartialUpdateBadRequestBody
*/
type UsersPartialUpdateBadRequestBody struct {

	// email firld errors
	// Required: true
	Email []string `json:"email"`

	// first_name firld errors
	// Required: true
	FirstName []string `json:"first_name"`

	// id firld errors
	// Required: true
	ID []string `json:"id"`

	// last_name firld errors
	// Required: true
	LastName []string `json:"last_name"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// password firld errors
	// Required: true
	Password []string `json:"password"`

	// profile
	// Required: true
	Profile *UsersPartialUpdateBadRequestBodyProfile `json:"profile"`

	// username firld errors
	// Required: true
	Username []string `json:"username"`
}

// Validate validates this users partial update bad request body
func (o *UsersPartialUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateFirstName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateProfile(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersPartialUpdateBadRequestBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("usersPartialUpdateBadRequest"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBody) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("usersPartialUpdateBadRequest"+"."+"first_name", "body", o.FirstName); err != nil {
		return err
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("usersPartialUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBody) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("usersPartialUpdateBadRequest"+"."+"last_name", "body", o.LastName); err != nil {
		return err
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("usersPartialUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("usersPartialUpdateBadRequest"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBody) validateProfile(formats strfmt.Registry) error {

	if err := validate.Required("usersPartialUpdateBadRequest"+"."+"profile", "body", o.Profile); err != nil {
		return err
	}

	if o.Profile != nil {

		if err := o.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usersPartialUpdateBadRequest" + "." + "profile")
			}
			return err
		}
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("usersPartialUpdateBadRequest"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

/*UsersPartialUpdateBadRequestBodyProfile users partial update bad request body profile
swagger:model UsersPartialUpdateBadRequestBodyProfile
*/
type UsersPartialUpdateBadRequestBodyProfile struct {

	// avatar_url firld errors
	AvatarURL []string `json:"avatar_url"`

	// bio firld errors
	Bio []string `json:"bio"`

	// company firld errors
	Company []string `json:"company"`

	// location firld errors
	Location []string `json:"location"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// timezone firld errors
	Timezone []string `json:"timezone"`

	// url firld errors
	URL []string `json:"url"`
}

// Validate validates this users partial update bad request body profile
func (o *UsersPartialUpdateBadRequestBodyProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvatarURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateBio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCompany(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateLocation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTimezone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersPartialUpdateBadRequestBodyProfile) validateAvatarURL(formats strfmt.Registry) error {

	if swag.IsZero(o.AvatarURL) { // not required
		return nil
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBodyProfile) validateBio(formats strfmt.Registry) error {

	if swag.IsZero(o.Bio) { // not required
		return nil
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBodyProfile) validateCompany(formats strfmt.Registry) error {

	if swag.IsZero(o.Company) { // not required
		return nil
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBodyProfile) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(o.Location) { // not required
		return nil
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBodyProfile) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBodyProfile) validateTimezone(formats strfmt.Registry) error {

	if swag.IsZero(o.Timezone) { // not required
		return nil
	}

	return nil
}

func (o *UsersPartialUpdateBadRequestBodyProfile) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(o.URL) { // not required
		return nil
	}

	return nil
}

/*UsersPartialUpdateBody users partial update body
swagger:model UsersPartialUpdateBody
*/
type UsersPartialUpdateBody struct {

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// profile
	Profile *models.UserProfile `json:"profile,omitempty"`

	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username,omitempty"`
}
