// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PackageTemplate package template
// swagger:model PackageTemplate
type PackageTemplate struct {

	// Catalog entry arguments
	Arguments []*PackageTemplateArgument `json:"arguments"`

	// Name
	Name string `json:"name,omitempty"`
}

// Validate validates this package template
func (m *PackageTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageTemplate) validateArguments(formats strfmt.Registry) error {

	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	for i := 0; i < len(m.Arguments); i++ {
		if swag.IsZero(m.Arguments[i]) { // not required
			continue
		}

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackageTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageTemplate) UnmarshalBinary(b []byte) error {
	var res PackageTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
