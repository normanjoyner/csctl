// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PluginCatalog A Containership plugin definition
// swagger:model PluginCatalog
type PluginCatalog struct {

	// Container Network Interface (CNI) Plugins
	CNI []*PluginDefinition `json:"CNI,omitempty"`

	// Container Storage Interface (CSI) Plugins
	CSI []*PluginDefinition `json:"CSI,omitempty"`

	// Cloud Controller Manager Plugins
	CloudControllerManager []*PluginDefinition `json:"cloud_controller_manager,omitempty"`

	// Cluster Management Plugins
	ClusterManagement []*PluginDefinition `json:"cluster_management,omitempty"`

	// Log Plugins
	Logs []*PluginDefinition `json:"logs,omitempty"`

	// Metric Plugins
	Metrics []*PluginDefinition `json:"metrics,omitempty"`
}

// Validate validates this plugin catalog
func (m *PluginCatalog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCNI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCSI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudControllerManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterManagement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginCatalog) validateCNI(formats strfmt.Registry) error {

	if swag.IsZero(m.CNI) { // not required
		return nil
	}

	for i := 0; i < len(m.CNI); i++ {
		if swag.IsZero(m.CNI[i]) { // not required
			continue
		}

		if m.CNI[i] != nil {
			if err := m.CNI[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CNI" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginCatalog) validateCSI(formats strfmt.Registry) error {

	if swag.IsZero(m.CSI) { // not required
		return nil
	}

	for i := 0; i < len(m.CSI); i++ {
		if swag.IsZero(m.CSI[i]) { // not required
			continue
		}

		if m.CSI[i] != nil {
			if err := m.CSI[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CSI" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginCatalog) validateCloudControllerManager(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudControllerManager) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudControllerManager); i++ {
		if swag.IsZero(m.CloudControllerManager[i]) { // not required
			continue
		}

		if m.CloudControllerManager[i] != nil {
			if err := m.CloudControllerManager[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloud_controller_manager" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginCatalog) validateClusterManagement(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterManagement) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterManagement); i++ {
		if swag.IsZero(m.ClusterManagement[i]) { // not required
			continue
		}

		if m.ClusterManagement[i] != nil {
			if err := m.ClusterManagement[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_management" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginCatalog) validateLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginCatalog) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(m.Metrics); i++ {
		if swag.IsZero(m.Metrics[i]) { // not required
			continue
		}

		if m.Metrics[i] != nil {
			if err := m.Metrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginCatalog) UnmarshalBinary(b []byte) error {
	var res PluginCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
