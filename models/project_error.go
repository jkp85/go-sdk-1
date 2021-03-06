// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProjectError project error
// swagger:model ProjectError
type ProjectError struct {

	// Collaborators field errors.
	Collaborators []string `json:"collaborators"`

	// Description field errors.
	Description []string `json:"description"`

	// Id field errors.
	ID []string `json:"id"`

	// Last_updated field errors.
	LastUpdated []string `json:"last_updated"`

	// Name field errors.
	Name []string `json:"name"`

	// Errors not connected to any field.
	NonFieldErrors []string `json:"non_field_errors"`

	// Owner field errors.
	Owner []string `json:"owner"`

	// Private field errors.
	Private []string `json:"private"`
}

// Validate validates this project error
func (m *ProjectError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollaborators(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

	if err := m.validatePrivate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectError) validateCollaborators(formats strfmt.Registry) error {

	if swag.IsZero(m.Collaborators) { // not required
		return nil
	}

	return nil
}

func (m *ProjectError) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	return nil
}

func (m *ProjectError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *ProjectError) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	return nil
}

func (m *ProjectError) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	return nil
}

func (m *ProjectError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *ProjectError) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	return nil
}

func (m *ProjectError) validatePrivate(formats strfmt.Registry) error {

	if swag.IsZero(m.Private) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectError) UnmarshalBinary(b []byte) error {
	var res ProjectError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
