// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CollaboratorError collaborator error
// swagger:model CollaboratorError
type CollaboratorError struct {

	// Email field errors.
	Email []string `json:"email"`

	// First name field errors.
	FirstName []string `json:"first_name"`

	// Id field errors this
	ID []string `json:"id"`

	// Joined field errors.
	Joined []string `json:"joined"`

	// Last name field errors.
	LastName []string `json:"last_name"`

	// Member field errors.
	// Required: true
	Member []string `json:"member"`

	// Errors not connected to any field.
	NonFieldErrors []string `json:"non_field_errors"`

	// Owner field errors.
	Owner []string `json:"owner"`

	// Permissions field errors.
	// Required: true
	Permissions []string `json:"permissions"`

	// Username field errors.
	Username []string `json:"username"`
}

// Validate validates this collaborator error
func (m *CollaboratorError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateJoined(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMember(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollaboratorError) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	return nil
}

func (m *CollaboratorError) validateFirstName(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstName) { // not required
		return nil
	}

	return nil
}

func (m *CollaboratorError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *CollaboratorError) validateJoined(formats strfmt.Registry) error {

	if swag.IsZero(m.Joined) { // not required
		return nil
	}

	return nil
}

func (m *CollaboratorError) validateLastName(formats strfmt.Registry) error {

	if swag.IsZero(m.LastName) { // not required
		return nil
	}

	return nil
}

func (m *CollaboratorError) validateMember(formats strfmt.Registry) error {

	if err := validate.Required("member", "body", m.Member); err != nil {
		return err
	}

	return nil
}

func (m *CollaboratorError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *CollaboratorError) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	return nil
}

func (m *CollaboratorError) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}

func (m *CollaboratorError) validateUsername(formats strfmt.Registry) error {

	if swag.IsZero(m.Username) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CollaboratorError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollaboratorError) UnmarshalBinary(b []byte) error {
	var res CollaboratorError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
