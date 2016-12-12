package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// InlineQueryResultVenue inline query result venue
// swagger:model InlineQueryResultVenue
type InlineQueryResultVenue struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// foursquare id
	FoursquareID string `json:"foursquare_id,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	InputMessageContent interface{} `json:"input_message_content,omitempty"`

	// latitude
	// Required: true
	Latitude *float64 `json:"latitude"`

	// longitude
	// Required: true
	Longitude *float64 `json:"longitude"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// thumb height
	ThumbHeight int64 `json:"thumb_height,omitempty"`

	// thumb url
	ThumbURL string `json:"thumb_url,omitempty"`

	// thumb width
	ThumbWidth int64 `json:"thumb_width,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	// Required: true
	Type InlineType `json:"type"`
}

// Validate validates this inline query result venue
func (m *InlineQueryResultVenue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLatitude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineQueryResultVenue) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultVenue) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultVenue) validateLatitude(formats strfmt.Registry) error {

	if err := validate.Required("latitude", "body", m.Latitude); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultVenue) validateLongitude(formats strfmt.Registry) error {

	if err := validate.Required("longitude", "body", m.Longitude); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultVenue) validateReplyMarkup(formats strfmt.Registry) error {

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

func (m *InlineQueryResultVenue) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultVenue) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		return err
	}

	return nil
}
