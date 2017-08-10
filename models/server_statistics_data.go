// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerStatisticsData server statistics data
// swagger:model ServerStatisticsData
type ServerStatisticsData struct {

	// id field errors.
	ID []string `json:"id"`

	// Errors not connected to any field.
	NonFieldErrors []string `json:"non_field_errors"`

	// size field errors.
	Size []string `json:"size"`

	// start field errors.
	Start []string `json:"start"`

	// stop field errors.
	Stop []string `json:"stop"`
}

// Validate validates this server statistics data
func (m *ServerStatisticsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStop(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerStatisticsData) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *ServerStatisticsData) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *ServerStatisticsData) validateSize(formats strfmt.Registry) error {

	if swag.IsZero(m.Size) { // not required
		return nil
	}

	return nil
}

func (m *ServerStatisticsData) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	return nil
}

func (m *ServerStatisticsData) validateStop(formats strfmt.Registry) error {

	if swag.IsZero(m.Stop) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerStatisticsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerStatisticsData) UnmarshalBinary(b []byte) error {
	var res ServerStatisticsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}