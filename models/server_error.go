// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerError server error
// swagger:model ServerError
type ServerError struct {

	// config field errors.
	Config []string `json:"config"`

	// connected field errors.
	Connected []string `json:"connected"`

	// created_at field errors.
	CreatedAt []string `json:"created_at"`

	// endpoint field errors.
	Endpoint []string `json:"endpoint"`

	// host field errors.
	Host []string `json:"host"`

	// id field errors.
	ID []string `json:"id"`

	// image_name field errors.
	ImageName []string `json:"image_name"`

	// logs_url field errors.
	LogsURL []string `json:"logs_url"`

	// name field errors.
	Name []string `json:"name"`

	// Errors not connected to any field.
	NonFieldErrors []string `json:"non_field_errors"`

	// server_size field errors.
	ServerSize []string `json:"server_size"`

	// startup_script field errors.
	StartupScript []string `json:"startup_script"`

	// status field errors.
	Status []string `json:"status"`

	// status_url field errors.
	StatusURL []string `json:"status_url"`
}

// Validate validates this server error
func (m *ServerError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnected(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLogsURL(formats); err != nil {
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

	if err := m.validateServerSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartupScript(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatusURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerError) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateConnected(formats strfmt.Registry) error {

	if swag.IsZero(m.Connected) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateEndpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.Endpoint) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateHost(formats strfmt.Registry) error {

	if swag.IsZero(m.Host) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateImageName(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageName) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateLogsURL(formats strfmt.Registry) error {

	if swag.IsZero(m.LogsURL) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateServerSize(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerSize) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateStartupScript(formats strfmt.Registry) error {

	if swag.IsZero(m.StartupScript) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	return nil
}

func (m *ServerError) validateStatusURL(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusURL) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerError) UnmarshalBinary(b []byte) error {
	var res ServerError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}