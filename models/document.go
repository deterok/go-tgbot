// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Document document
// swagger:model Document
type Document struct {

	// file id
	FileID string `json:"file_id,omitempty"`

	// file name
	FileName string `json:"file_name,omitempty"`

	// file size
	FileSize int64 `json:"file_size,omitempty"`

	// mime type
	MimeType string `json:"mime_type,omitempty"`

	// thumb
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// Validate validates this document
func (m *Document) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThumb(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Document) validateThumb(formats strfmt.Registry) error {

	if swag.IsZero(m.Thumb) { // not required
		return nil
	}

	if m.Thumb != nil {

		if err := m.Thumb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thumb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Document) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Document) UnmarshalBinary(b []byte) error {
	var res Document
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
