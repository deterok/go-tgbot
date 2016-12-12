package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// InlineQueryResultCachedVideo inline query result cached video
// swagger:model InlineQueryResultCachedVideo
type InlineQueryResultCachedVideo struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	InputMessageContent interface{} `json:"input_message_content,omitempty"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	// Required: true
	Type InlineType `json:"type"`

	// video file id
	// Required: true
	VideoFileID *string `json:"video_file_id"`
}

// Validate validates this inline query result cached video
func (m *InlineQueryResultCachedVideo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyMarkup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVideoFileID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineQueryResultCachedVideo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultCachedVideo) validateReplyMarkup(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyMarkup) { // not required
		return nil
	}

	if m.ReplyMarkup != nil {

		if err := m.ReplyMarkup.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *InlineQueryResultCachedVideo) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultCachedVideo) validateVideoFileID(formats strfmt.Registry) error {

	if err := validate.Required("video_file_id", "body", m.VideoFileID); err != nil {
		return err
	}

	return nil
}
