// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostInstallProgressParams host install progress params
//
// swagger:model host-install-progress-params
type HostInstallProgressParams struct {

	// current stage
	// Required: true
	CurrentStage HostStage `json:"current_stage"`

	// progress info
	ProgressInfo string `json:"progress_info,omitempty" gorm:"type:varchar(2048)"`
}

// Validate validates this host install progress params
func (m *HostInstallProgressParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentStage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostInstallProgressParams) validateCurrentStage(formats strfmt.Registry) error {

	if err := m.CurrentStage.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("current_stage")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostInstallProgressParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostInstallProgressParams) UnmarshalBinary(b []byte) error {
	var res HostInstallProgressParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}