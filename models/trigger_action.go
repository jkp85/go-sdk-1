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

// TriggerAction trigger action
// swagger:model TriggerAction
type TriggerAction struct {

	// Name for trigger action.
	// Required: true
	ActionName *string `json:"action_name"`

	// Trigger action that represents unique identifier in UUID format.
	ID string `json:"id,omitempty"`

	// Method for trigger action, such as POST.
	// Required: true
	Method *string `json:"method"`

	// Trigger action model.
	Model string `json:"model,omitempty"`

	// Object ID Unique identifer in UUID format.
	ObjectID string `json:"object_id,omitempty"`

	// Object that represents trigger action payload.
	Payload interface{} `json:"payload,omitempty"`
}

// Validate validates this trigger action
func (m *TriggerAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TriggerAction) validateActionName(formats strfmt.Registry) error {

	if err := validate.Required("action_name", "body", m.ActionName); err != nil {
		return err
	}

	return nil
}

func (m *TriggerAction) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TriggerAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TriggerAction) UnmarshalBinary(b []byte) error {
	var res TriggerAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
