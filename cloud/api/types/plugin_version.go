// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginVersion A Containership plugin version definition
// swagger:model PluginVersion
type PluginVersion struct {

	// Kubernetes and upgrade compatibility
	// Required: true
	Compatibility *PluginCompatibility `json:"compatibility"`

	// needed keys
	// Required: true
	ConfigurationSchema []string `json:"configuration_schema"`

	// If the plugin is disabled
	Disabled bool `json:"disabled,omitempty"`

	// The plugin version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this plugin version
func (m *PluginVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompatibility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurationSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginVersion) validateCompatibility(formats strfmt.Registry) error {

	if err := validate.Required("compatibility", "body", m.Compatibility); err != nil {
		return err
	}

	if m.Compatibility != nil {
		if err := m.Compatibility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compatibility")
			}
			return err
		}
	}

	return nil
}

func (m *PluginVersion) validateConfigurationSchema(formats strfmt.Registry) error {

	if err := validate.Required("configuration_schema", "body", m.ConfigurationSchema); err != nil {
		return err
	}

	return nil
}

func (m *PluginVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginVersion) UnmarshalBinary(b []byte) error {
	var res PluginVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
