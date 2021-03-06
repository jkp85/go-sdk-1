// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PlanError plan error
// swagger:model PlanError
type PlanError struct {

	// amount field errors
	Amount []string `json:"amount"`

	// created field errors
	Created []string `json:"created"`

	// currency field errors
	Currency []string `json:"currency"`

	// id field errors
	ID []string `json:"id"`

	// interval field errors
	Interval []string `json:"interval"`

	// interval_count field errors
	IntervalCount []string `json:"interval_count"`

	// livemode field errors
	Livemode []string `json:"livemode"`

	// metadata field errors
	Metadata []string `json:"metadata"`

	// name field errors
	Name []string `json:"name"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// statement_descriptor field errors
	StatementDescriptor []string `json:"statement_descriptor"`

	// stripe_id field errors
	StripeID []string `json:"stripe_id"`

	// trial period days field errors
	TrialPeriodDays []string `json:"trial_period_days"`
}

// Validate validates this plan error
func (m *PlanError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
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

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIntervalCount(formats); err != nil {
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

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatementDescriptor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStripeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrialPeriodDays(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanError) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.Interval) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateIntervalCount(formats strfmt.Registry) error {

	if swag.IsZero(m.IntervalCount) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateLivemode(formats strfmt.Registry) error {

	if swag.IsZero(m.Livemode) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateStatementDescriptor(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementDescriptor) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateStripeID(formats strfmt.Registry) error {

	if swag.IsZero(m.StripeID) { // not required
		return nil
	}

	return nil
}

func (m *PlanError) validateTrialPeriodDays(formats strfmt.Registry) error {

	if swag.IsZero(m.TrialPeriodDays) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanError) UnmarshalBinary(b []byte) error {
	var res PlanError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
