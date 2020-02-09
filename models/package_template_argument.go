// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackageTemplateArgument package template argument
// swagger:model PackageTemplateArgument
type PackageTemplateArgument struct {

	// Argument constraint
	Contrainst string `json:"contrainst,omitempty"`

	// Argument key to generate catalog entry
	Key string `json:"key,omitempty"`

	// Argument value to generate catalog entry
	Value string `json:"value,omitempty"`
}

// Validate validates this package template argument
func (m *PackageTemplateArgument) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackageTemplateArgument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageTemplateArgument) UnmarshalBinary(b []byte) error {
	var res PackageTemplateArgument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
