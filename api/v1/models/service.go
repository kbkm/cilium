package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Service Collection of endpoints to be served
// swagger:model Service
type Service struct {

	// List of backend addresses
	BackendAddresses []*BackendAddress `json:"backend-addresses"`

	// flags
	Flags *ServiceFlags `json:"flags,omitempty"`

	// Frontend address
	// Required: true
	FrontendAddress *FrontendAddress `json:"frontend-address"`

	// Unique identification
	ID int64 `json:"id,omitempty"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackendAddresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFlags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFrontendAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) validateBackendAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.BackendAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.BackendAddresses); i++ {

		if swag.IsZero(m.BackendAddresses[i]) { // not required
			continue
		}

		if m.BackendAddresses[i] != nil {

			if err := m.BackendAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backend-addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateFlags(formats strfmt.Registry) error {

	if swag.IsZero(m.Flags) { // not required
		return nil
	}

	if m.Flags != nil {

		if err := m.Flags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags")
			}
			return err
		}
	}

	return nil
}

func (m *Service) validateFrontendAddress(formats strfmt.Registry) error {

	if err := validate.Required("frontend-address", "body", m.FrontendAddress); err != nil {
		return err
	}

	if m.FrontendAddress != nil {

		if err := m.FrontendAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frontend-address")
			}
			return err
		}
	}

	return nil
}

// ServiceFlags Optional service configuration flags
// swagger:model ServiceFlags
type ServiceFlags struct {

	// Frontend to backend translation activated
	ActiveFrontend bool `json:"active-frontend,omitempty"`

	// Perform direct server return
	DirectServerReturn bool `json:"direct-server-return,omitempty"`
}

// Validate validates this service flags
func (m *ServiceFlags) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
