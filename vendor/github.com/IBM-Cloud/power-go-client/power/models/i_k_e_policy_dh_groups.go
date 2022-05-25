// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IKEPolicyDhGroups i k e policy dh groups
// Example: [1,2,5,14,19,20,24]
//
// swagger:model IKEPolicyDhGroups
type IKEPolicyDhGroups []float64

// Validate validates this i k e policy dh groups
func (m IKEPolicyDhGroups) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this i k e policy dh groups based on context it is used
func (m IKEPolicyDhGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}