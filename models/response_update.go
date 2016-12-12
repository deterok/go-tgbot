package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// ResponseUpdate response update
// swagger:model ResponseUpdate
type ResponseUpdate struct {

	// description
	Description string `json:"description,omitempty"`

	// error code
	ErrorCode int64 `json:"error_code,omitempty"`

	// ok
	Ok bool `json:"ok,omitempty"`

	// result
	Result []*Update `json:"result"`
}

// Validate validates this response update
func (m *ResponseUpdate) Validate(formats strfmt.Registry) error {
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

func (m *ResponseUpdate) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	for i := 0; i < len(m.Result); i++ {

		if swag.IsZero(m.Result[i]) { // not required
			continue
		}

		if m.Result[i] != nil {

			if err := m.Result[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
