// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustomerError customer error
// swagger:model CustomerError
type CustomerError struct {

	// account_balance field errors
	AccountBalance []string `json:"account_balance"`

	// created field errors
	Created []string `json:"created"`

	// currency field errors
	Currency []string `json:"currency"`

	// default_source field errors
	DefaultSource []string `json:"default_source"`

	// id field errors
	ID []string `json:"id"`

	// last_invoice_sync field errors
	LastInvoiceSync []string `json:"last_invoice_sync"`

	// livemode field errors
	Livemode []string `json:"livemode"`

	// metadata field errors
	Metadata []string `json:"metadata"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// stripe_id field errors
	StripeID []string `json:"stripe_id"`

	// user field errors
	User []string `json:"user"`
}

// Validate validates this customer error
func (m *CustomerError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountBalance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDefaultSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastInvoiceSync(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLivemode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStripeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerError) validateAccountBalance(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountBalance) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateDefaultSource(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultSource) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateLastInvoiceSync(formats strfmt.Registry) error {

	if swag.IsZero(m.LastInvoiceSync) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateLivemode(formats strfmt.Registry) error {

	if swag.IsZero(m.Livemode) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateStripeID(formats strfmt.Registry) error {

	if swag.IsZero(m.StripeID) { // not required
		return nil
	}

	return nil
}

func (m *CustomerError) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerError) UnmarshalBinary(b []byte) error {
	var res CustomerError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}