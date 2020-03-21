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

// LIBRARYTYPE l i b r a r y t y p e
// swagger:model LIBRARY_TYPE
type LIBRARYTYPE string

const (

	// LIBRARYTYPEMODEL captures enum value "MODEL"
	LIBRARYTYPEMODEL LIBRARYTYPE = "MODEL"

	// LIBRARYTYPEINTERFACE captures enum value "INTERFACE"
	LIBRARYTYPEINTERFACE LIBRARYTYPE = "INTERFACE"

	// LIBRARYTYPECLIENT captures enum value "CLIENT"
	LIBRARYTYPECLIENT LIBRARYTYPE = "CLIENT"
)

// for schema
var lIBRARYTYPEEnum []interface{}

func init() {
	var res []LIBRARYTYPE
	if err := json.Unmarshal([]byte(`["MODEL","INTERFACE","CLIENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lIBRARYTYPEEnum = append(lIBRARYTYPEEnum, v)
	}
}

func (m LIBRARYTYPE) validateLIBRARYTYPEEnum(path, location string, value LIBRARYTYPE) error {
	if err := validate.Enum(path, location, value, lIBRARYTYPEEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this l i b r a r y t y p e
func (m LIBRARYTYPE) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLIBRARYTYPEEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
