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

	"github.com/3Blades/go-sdk/models"
)

// UsersCreateReader is a Reader for the UsersCreate structure.
type UsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUsersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersCreateCreated creates a UsersCreateCreated with default headers values
func NewUsersCreateCreated() *UsersCreateCreated {
	return &UsersCreateCreated{}
}

/*UsersCreateCreated handles this case with default header values.

User created
*/
type UsersCreateCreated struct {
	Payload *models.User
}

func (o *UsersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /{namespace}/users/][%d] usersCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersCreateBadRequest creates a UsersCreateBadRequest with default headers values
func NewUsersCreateBadRequest() *UsersCreateBadRequest {
	return &UsersCreateBadRequest{}
}

/*UsersCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type UsersCreateBadRequest struct {
	Payload UsersCreateBadRequestBody
}

func (o *UsersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /{namespace}/users/][%d] usersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UsersCreateBadRequestBody users create bad request body
swagger:model UsersCreateBadRequestBody
*/
type UsersCreateBadRequestBody struct {

	// email field errors
	// Required: true
	Email []string `json:"email"`

	// first_name field errors
	// Required: true
	FirstName []string `json:"first_name"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// last_name field errors
	// Required: true
	LastName []string `json:"last_name"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// password field errors
	// Required: true
	Password []string `json:"password"`

	// profile
	// Required: true
	Profile *UsersCreateBadRequestBodyProfile `json:"profile"`

	// username field errors
	// Required: true
	Username []string `json:"username"`
}

// Validate validates this users create bad request body
func (o *UsersCreateBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *UsersCreateBadRequestBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("usersCreateBadRequest"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *UsersCreateBadRequestBody) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("usersCreateBadRequest"+"."+"first_name", "body", o.FirstName); err != nil {
		return err
	}

	return nil
}

func (o *UsersCreateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("usersCreateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *UsersCreateBadRequestBody) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("usersCreateBadRequest"+"."+"last_name", "body", o.LastName); err != nil {
		return err
	}

	return nil
}

func (o *UsersCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("usersCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *UsersCreateBadRequestBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("usersCreateBadRequest"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

func (o *UsersCreateBadRequestBody) validateProfile(formats strfmt.Registry) error {

	if err := validate.Required("usersCreateBadRequest"+"."+"profile", "body", o.Profile); err != nil {
		return err
	}

	if o.Profile != nil {

		if err := o.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usersCreateBadRequest" + "." + "profile")
			}
			return err
		}
	}

	return nil
}

func (o *UsersCreateBadRequestBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("usersCreateBadRequest"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

/*UsersCreateBadRequestBodyProfile users create bad request body profile
swagger:model UsersCreateBadRequestBodyProfile
*/
type UsersCreateBadRequestBodyProfile struct {

	// avatar_url field errors
	AvatarURL []string `json:"avatar_url"`

	// bio field errors
	Bio []string `json:"bio"`

	// company field errors
	Company []string `json:"company"`

	// location field errors
	Location []string `json:"location"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// timezone field errors
	Timezone []string `json:"timezone"`

	// url field errors
	URL []string `json:"url"`
}

// Validate validates this users create bad request body profile
func (o *UsersCreateBadRequestBodyProfile) Validate(formats strfmt.Registry) error {
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

func (o *UsersCreateBadRequestBodyProfile) validateAvatarURL(formats strfmt.Registry) error {

	if swag.IsZero(o.AvatarURL) { // not required
		return nil
	}

	return nil
}

func (o *UsersCreateBadRequestBodyProfile) validateBio(formats strfmt.Registry) error {

	if swag.IsZero(o.Bio) { // not required
		return nil
	}

	return nil
}

func (o *UsersCreateBadRequestBodyProfile) validateCompany(formats strfmt.Registry) error {

	if swag.IsZero(o.Company) { // not required
		return nil
	}

	return nil
}

func (o *UsersCreateBadRequestBodyProfile) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(o.Location) { // not required
		return nil
	}

	return nil
}

func (o *UsersCreateBadRequestBodyProfile) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (o *UsersCreateBadRequestBodyProfile) validateTimezone(formats strfmt.Registry) error {

	if swag.IsZero(o.Timezone) { // not required
		return nil
	}

	return nil
}

func (o *UsersCreateBadRequestBodyProfile) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(o.URL) { // not required
		return nil
	}

	return nil
}

/*UsersCreateBody users create body
swagger:model UsersCreateBody
*/
type UsersCreateBody struct {

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// password
	// Required: true
	Password *string `json:"password"`

	// profile
	// Required: true
	Profile *models.UserProfile `json:"profile"`

	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	// Required: true
	Username *string `json:"username"`
}