// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetGameHighScoresOKBody get game high scores o k body
// swagger:model getGameHighScoresOKBody
type GetGameHighScoresOKBody struct {

	// description
	Description string `json:"description,omitempty"`

	// error code
	ErrorCode int64 `json:"error_code,omitempty"`

	// ok
	Ok bool `json:"ok,omitempty"`

	// result
	Result *GameHighScore `json:"result,omitempty"`
}

// Validate validates this get game high scores o k body
func (m *GetGameHighScoresOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetGameHighScoresOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {

		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetGameHighScoresOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetGameHighScoresOKBody) UnmarshalBinary(b []byte) error {
	var res GetGameHighScoresOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
