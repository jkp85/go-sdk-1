package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Server server
// swagger:model Server
type Server struct {

	// config
	Config interface{} `json:"config,omitempty"`

	// connected
	// Required: true
	Connected []string `json:"connected"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// environment resources
	// Required: true
	EnvironmentResources *string `json:"environment_resources"`

	// environment type
	// Required: true
	EnvironmentType *string `json:"environment_type"`

	// id
	ID string `json:"id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// startup script
	StartupScript string `json:"startup_script,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this server
func (m *Server) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnected(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnvironmentResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnvironmentType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Server) validateConnected(formats strfmt.Registry) error {

	if err := validate.Required("connected", "body", m.Connected); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateEnvironmentResources(formats strfmt.Registry) error {

	if err := validate.Required("environment_resources", "body", m.EnvironmentResources); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateEnvironmentType(formats strfmt.Registry) error {

	if err := validate.Required("environment_type", "body", m.EnvironmentType); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
