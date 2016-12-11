package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// InlineKeyboardButton inline keyboard button
// swagger:model InlineKeyboardButton
type InlineKeyboardButton struct {

	// callback data
	CallbackData string `json:"callback_data,omitempty"`

	// callback game
	CallbackGame CallbackGame `json:"callback_game,omitempty"`

	// switch inline query
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// switch inline query current chat
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this inline keyboard button
func (m *InlineKeyboardButton) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
