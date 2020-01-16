// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RuleHandler rule handler
// swagger:model ruleHandler
type RuleHandler struct {

	// Config contains the configuration for the handler. Please read the user
	// guide for a complete list of each handler's available settings.
	Config interface{} `json:"config,omitempty"`

	// Handler identifies the implementation which will be used to handle this specific request. Please read the user
	// guide for a complete list of available handlers.
	Handler string `json:"handler,omitempty"`
}

// Validate validates this rule handler
func (m *RuleHandler) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuleHandler) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleHandler) UnmarshalBinary(b []byte) error {
	var res RuleHandler
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}