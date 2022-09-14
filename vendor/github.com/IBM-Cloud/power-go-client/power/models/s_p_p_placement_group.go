// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SPPPlacementGroup s p p placement group
//
// swagger:model SPPPlacementGroup
type SPPPlacementGroup struct {

	// The id of the Shared Processor Pool Placement Group
	// Required: true
	ID *string `json:"id"`

	// The list of Shared Processor Pool names that are a member of the Shared Processor Pool Placement Group
	MemberSharedProcessorPools []string `json:"memberSharedProcessorPools"`

	// The name of the Shared Processor Pool Placement Group
	// Required: true
	Name *string `json:"name"`

	// The Shared Processor Pool Placement Group policy
	// Required: true
	Policy *string `json:"policy"`
}

// Validate validates this s p p placement group
func (m *SPPPlacementGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SPPPlacementGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SPPPlacementGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SPPPlacementGroup) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", m.Policy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this s p p placement group based on context it is used
func (m *SPPPlacementGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SPPPlacementGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SPPPlacementGroup) UnmarshalBinary(b []byte) error {
	var res SPPPlacementGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}