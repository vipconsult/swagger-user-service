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

// F2aAuth f2a auth
// swagger:model F2aAuth
type F2aAuth struct {

	// the  2 factor time code accuired from the google authenticator app
	// Required: true
	F2a *string `json:"f2a"`

	// the jwt token accuired form the initial login
	// Required: true
	Jwt *string `json:"jwt"`
}

// Validate validates this f2a auth
func (m *F2aAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateF2a(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateJwt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *F2aAuth) validateF2a(formats strfmt.Registry) error {

	if err := validate.Required("f2a", "body", m.F2a); err != nil {
		return err
	}

	return nil
}

func (m *F2aAuth) validateJwt(formats strfmt.Registry) error {

	if err := validate.Required("jwt", "body", m.Jwt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *F2aAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *F2aAuth) UnmarshalBinary(b []byte) error {
	var res F2aAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
