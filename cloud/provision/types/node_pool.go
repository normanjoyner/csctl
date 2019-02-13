// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodePool node pool
// swagger:model NodePool
type NodePool struct {

	// Autoscaling if enabled holds the configuration for an autoscaling group
	// Required: true
	Autoscaling *NodePoolAutoscaling `json:"autoscaling"`

	// Number of nodes in this node pool
	// Required: true
	Count *int32 `json:"count"`

	// Docker version running on this node pool
	// Required: true
	DockerVersion *string `json:"docker_version"`

	// Flag indicating whether etcd is running on this node pool
	// Required: true
	Etcd *bool `json:"etcd"`

	// etcd version running on this node, if applicable
	// Required: true
	EtcdVersion *string `json:"etcd_version"`

	// Node Pool ID
	// Required: true
	ID UUID `json:"id"`

	// Flag indicating whether this node pool can be scheduled on
	// Required: true
	IsSchedulable *bool `json:"is_schedulable"`

	// Kubernetes mode (master or worker) for this node pool
	// Required: true
	// Enum: [master worker]
	KubernetesMode *string `json:"kubernetes_mode"`

	// Kubernetes version running on this node pool
	// Required: true
	KubernetesVersion *string `json:"kubernetes_version"`

	// Name of this node pool
	// Required: true
	Name *string `json:"name"`

	// Operating System
	// Required: true
	Os *string `json:"os"`

	// Node Pool status
	// Required: true
	Status *NodePoolStatus `json:"status"`
}

// Validate validates this node pool
func (m *NodePool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoscaling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEtcd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEtcdVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsSchedulable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodePool) validateAutoscaling(formats strfmt.Registry) error {

	if err := validate.Required("autoscaling", "body", m.Autoscaling); err != nil {
		return err
	}

	if m.Autoscaling != nil {
		if err := m.Autoscaling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoscaling")
			}
			return err
		}
	}

	return nil
}

func (m *NodePool) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateDockerVersion(formats strfmt.Registry) error {

	if err := validate.Required("docker_version", "body", m.DockerVersion); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateEtcd(formats strfmt.Registry) error {

	if err := validate.Required("etcd", "body", m.Etcd); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateEtcdVersion(formats strfmt.Registry) error {

	if err := validate.Required("etcd_version", "body", m.EtcdVersion); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *NodePool) validateIsSchedulable(formats strfmt.Registry) error {

	if err := validate.Required("is_schedulable", "body", m.IsSchedulable); err != nil {
		return err
	}

	return nil
}

var nodePoolTypeKubernetesModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["master","worker"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodePoolTypeKubernetesModePropEnum = append(nodePoolTypeKubernetesModePropEnum, v)
	}
}

const (

	// NodePoolKubernetesModeMaster captures enum value "master"
	NodePoolKubernetesModeMaster string = "master"

	// NodePoolKubernetesModeWorker captures enum value "worker"
	NodePoolKubernetesModeWorker string = "worker"
)

// prop value enum
func (m *NodePool) validateKubernetesModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nodePoolTypeKubernetesModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NodePool) validateKubernetesMode(formats strfmt.Registry) error {

	if err := validate.Required("kubernetes_mode", "body", m.KubernetesMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateKubernetesModeEnum("kubernetes_mode", "body", *m.KubernetesMode); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateKubernetesVersion(formats strfmt.Registry) error {

	if err := validate.Required("kubernetes_version", "body", m.KubernetesVersion); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateOs(formats strfmt.Registry) error {

	if err := validate.Required("os", "body", m.Os); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodePool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodePool) UnmarshalBinary(b []byte) error {
	var res NodePool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
