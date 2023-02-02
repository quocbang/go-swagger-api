// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountData account data
//
// swagger:model AccountData
type AccountData []*AccountDataItems0

// Validate validates this account data
func (m AccountData) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this account data based on the context it is used
func (m AccountData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {
			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// AccountDataItems0 account data items0
//
// swagger:model AccountDataItems0
type AccountDataItems0 struct {

	// 人員工號
	// Example: 0
	// Required: true
	EmployeeID *string `json:"employeeID"`

	// roles
	// Required: true
	Roles Roles `json:"roles"`
}

// Validate validates this account data items0
func (m *AccountDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmployeeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountDataItems0) validateEmployeeID(formats strfmt.Registry) error {

	if err := validate.Required("employeeID", "body", m.EmployeeID); err != nil {
		return err
	}

	return nil
}

func (m *AccountDataItems0) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	if err := m.Roles.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("roles")
		}
		return err
	}

	return nil
}

// ContextValidate validate this account data items0 based on the context it is used
func (m *AccountDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountDataItems0) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Roles.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("roles")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountDataItems0) UnmarshalBinary(b []byte) error {
	var res AccountDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}