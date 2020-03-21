// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PROGRAMMINGLANGUAGE p r o g RAM m i n g l a n g u a g e
// swagger:model PROGRAMMING_LANGUAGE
type PROGRAMMINGLANGUAGE string

const (

	// PROGRAMMINGLANGUAGEJAVA captures enum value "JAVA"
	PROGRAMMINGLANGUAGEJAVA PROGRAMMINGLANGUAGE = "JAVA"

	// PROGRAMMINGLANGUAGEPYTHON captures enum value "PYTHON"
	PROGRAMMINGLANGUAGEPYTHON PROGRAMMINGLANGUAGE = "PYTHON"
)

// for schema
var pROGRamMINGLANGUAGEEnum []interface{}

func init() {
	var res []PROGRAMMINGLANGUAGE
	if err := json.Unmarshal([]byte(`["JAVA","PYTHON"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pROGRamMINGLANGUAGEEnum = append(pROGRamMINGLANGUAGEEnum, v)
	}
}

func (m PROGRAMMINGLANGUAGE) validatePROGRAMMINGLANGUAGEEnum(path, location string, value PROGRAMMINGLANGUAGE) error {
	if err := validate.Enum(path, location, value, pROGRamMINGLANGUAGEEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this p r o g RAM m i n g l a n g u a g e
func (m PROGRAMMINGLANGUAGE) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePROGRAMMINGLANGUAGEEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
