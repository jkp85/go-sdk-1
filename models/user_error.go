// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserError user error
// swagger:model UserError
type UserError struct {

	// Email field errors.
	Email []string `json:"email"`

	// First name field errors.
	FirstName []string `json:"first_name"`

	// id field errors.
	ID []string `json:"id"`

	// Last name field errors.
	LastName []string `json:"last_name"`

	// Errors not connected to any field.
	NonFieldErrors []string `json:"non_field_errors"`

	// Password field errors.
	Password []string `json:"password"`

	// profile
	Profile *UserProfileError `json:"profile,omitempty"`

	// Username field errors.
	Username []string `json:"username"`
}

// Validate validates this user error
func (m *UserError) Validate(formats strfmt.Registry) error {
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

	if err := m.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
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

func (m *UserError) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	return nil
}

func (m *UserError) validateFirstName(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstName) { // not required
		return nil
	}

	return nil
}

func (m *UserError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *UserError) validateLastName(formats strfmt.Registry) error {

	if swag.IsZero(m.LastName) { // not required
		return nil
	}

	return nil
}

func (m *UserError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *UserError) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	return nil
}

func (m *UserError) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {

		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *UserError) validateUsername(formats strfmt.Registry) error {

	if swag.IsZero(m.Username) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserError) UnmarshalBinary(b []byte) error {
	var res UserError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}