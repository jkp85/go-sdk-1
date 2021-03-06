// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CardError card error
// swagger:model CardError
type CardError struct {

	// address_city field errors
	AddressCity []string `json:"address_city"`

	// address_country field errors
	AddressCountry []string `json:"address_country"`

	// address_line1 field errors
	AddressLine1 []string `json:"address_line1"`

	// address_line1_check field errors
	AddressLine1Check []string `json:"address_line1_check"`

	// address_line2 field errors
	AddressLine2 []string `json:"address_line2"`

	// address_state field errors
	AddressState []string `json:"address_state"`

	// address_zip field errors
	AddressZip []string `json:"address_zip"`

	// address_zip_check field errors
	AddressZipCheck []string `json:"address_zip_check"`

	// brand field errors
	Brand []string `json:"brand"`

	// created field errors
	Created []string `json:"created"`

	// customer field errors
	Customer []string `json:"customer"`

	// cvc_check field errors
	CvcCheck []string `json:"cvc_check"`

	// exp_month field errors
	ExpMonth []string `json:"exp_month"`

	// exp_year field errors
	ExpYear []string `json:"exp_year"`

	// fingerprint field errors
	Fingerprint []string `json:"fingerprint"`

	// funding field errors
	Funding []string `json:"funding"`

	// id field errors
	ID []string `json:"id"`

	// last4 field errors
	Last4 []string `json:"last4"`

	// name field errors
	Name []string `json:"name"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// stripe_id field errors
	StripeID []string `json:"stripe_id"`

	// token field errors
	Token []string `json:"token"`
}

// Validate validates this card error
func (m *CardError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressCity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressLine1(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressLine1Check(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressLine2(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressZip(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddressZipCheck(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBrand(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCvcCheck(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExpMonth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExpYear(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFingerprint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFunding(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLast4(formats); err != nil {
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

	if err := m.validateStripeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardError) validateAddressCity(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressCity) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateAddressCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressCountry) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateAddressLine1(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine1) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateAddressLine1Check(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine1Check) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateAddressLine2(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine2) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateAddressState(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressState) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateAddressZip(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressZip) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateAddressZipCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressZipCheck) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateBrand(formats strfmt.Registry) error {

	if swag.IsZero(m.Brand) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateCustomer(formats strfmt.Registry) error {

	if swag.IsZero(m.Customer) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateCvcCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.CvcCheck) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateExpMonth(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpMonth) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateExpYear(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpYear) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateFingerprint(formats strfmt.Registry) error {

	if swag.IsZero(m.Fingerprint) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateFunding(formats strfmt.Registry) error {

	if swag.IsZero(m.Funding) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateLast4(formats strfmt.Registry) error {

	if swag.IsZero(m.Last4) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateStripeID(formats strfmt.Registry) error {

	if swag.IsZero(m.StripeID) { // not required
		return nil
	}

	return nil
}

func (m *CardError) validateToken(formats strfmt.Registry) error {

	if swag.IsZero(m.Token) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CardError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardError) UnmarshalBinary(b []byte) error {
	var res CardError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
