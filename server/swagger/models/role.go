// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Role 授權角色
// 定義來源參考: https://gitlab.kenda.com.tw/kenda/mcom/-/blob/${xxx}/utils/roles/roles.proto `Role` type.
// in which `${xxx}` inside the URL reference is the specified branch name of the corresponding features.
//
//
// swagger:model Role
type Role int64

// Validate validates this role
func (m Role) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this role based on context it is used
func (m Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}