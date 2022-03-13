// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateTodo TODO更新
//
// swagger:model updateTodo
type UpdateTodo struct {

	// 完了
	// Example: false
	Done bool `json:"done,omitempty"`

	// TODO名
	// Example: TODO名
	Name string `json:"name,omitempty"`
}

// Validate validates this update todo
func (m *UpdateTodo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update todo based on context it is used
func (m *UpdateTodo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTodo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTodo) UnmarshalBinary(b []byte) error {
	var res UpdateTodo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
