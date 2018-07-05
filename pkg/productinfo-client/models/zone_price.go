// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ZonePrice ZonePrice struct for displaying proce information per zone
// swagger:model ZonePrice
type ZonePrice struct {

	// price
	Price float64 `json:"price,omitempty"`

	// zone
	Zone string `json:"zone,omitempty"`
}

// Validate validates this zone price
func (m *ZonePrice) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ZonePrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZonePrice) UnmarshalBinary(b []byte) error {
	var res ZonePrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
