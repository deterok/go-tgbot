// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InlineQueryResultVoice inline query result voice
// swagger:model InlineQueryResultVoice
type InlineQueryResultVoice struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	InputMessageContent interface{} `json:"input_message_content,omitempty"`

	// parse mode
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	// Required: true
	Type InlineType `json:"type"`

	// voice duration
	VoiceDuration int64 `json:"voice_duration,omitempty"`

	// voice url
	// Required: true
	VoiceURL *string `json:"voice_url"`
}

// Validate validates this inline query result voice
func (m *InlineQueryResultVoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParseMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyMarkup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVoiceURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineQueryResultVoice) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultVoice) validateParseMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ParseMode) { // not required
		return nil
	}

	if err := m.ParseMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("parse_mode")
		}
		return err
	}

	return nil
}

func (m *InlineQueryResultVoice) validateReplyMarkup(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyMarkup) { // not required
		return nil
	}

	if m.ReplyMarkup != nil {

		if err := m.ReplyMarkup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reply_markup")
			}
			return err
		}
	}

	return nil
}

func (m *InlineQueryResultVoice) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultVoice) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *InlineQueryResultVoice) validateVoiceURL(formats strfmt.Registry) error {

	if err := validate.Required("voice_url", "body", m.VoiceURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineQueryResultVoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineQueryResultVoice) UnmarshalBinary(b []byte) error {
	var res InlineQueryResultVoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
