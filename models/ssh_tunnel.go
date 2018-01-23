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

// SSHTunnel Ssh tunnel
// swagger:model SshTunnel
type SSHTunnel struct {

	// SSH tunnel destination endpoint.
	// Required: true
	Endpoint *string `json:"endpoint"`

	// Host, usually IPv4, for SSH tunnel.
	// Required: true
	Host *string `json:"host"`

	// SSH tunnel unique identifier in UUID format.
	ID string `json:"id,omitempty"`

	// Local source port for SSH tunnel.
	// Required: true
	LocalPort *int64 `json:"local_port"`

	// SSH tunnel name.
	// Required: true
	Name *string `json:"name"`

	// Remote port to establish SSH tunnel.
	// Required: true
	RemotePort *int64 `json:"remote_port"`

	// Server name.
	Server string `json:"server,omitempty"`

	// SSH tunnel user name.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this Ssh tunnel
func (m *SSHTunnel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocalPort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRemotePort(formats); err != nil {
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

func (m *SSHTunnel) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

func (m *SSHTunnel) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *SSHTunnel) validateLocalPort(formats strfmt.Registry) error {

	if err := validate.Required("local_port", "body", m.LocalPort); err != nil {
		return err
	}

	return nil
}

func (m *SSHTunnel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SSHTunnel) validateRemotePort(formats strfmt.Registry) error {

	if err := validate.Required("remote_port", "body", m.RemotePort); err != nil {
		return err
	}

	return nil
}

func (m *SSHTunnel) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSHTunnel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHTunnel) UnmarshalBinary(b []byte) error {
	var res SSHTunnel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
