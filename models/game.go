package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Game game
// swagger:model Game
type Game struct {

	// animation
	Animation *Animation `json:"animation,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// photo
	// Required: true
	Photo []*PhotoSize `json:"photo"`

	// text
	Text string `json:"text,omitempty"`

	// text entities
	TextEntities []*MessageEntity `json:"text_entities"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this game
func (m *Game) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnimation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhoto(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTextEntities(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Game) validateAnimation(formats strfmt.Registry) error {

	if swag.IsZero(m.Animation) { // not required
		return nil
	}

	if m.Animation != nil {

		if err := m.Animation.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Game) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Game) validatePhoto(formats strfmt.Registry) error {

	if err := validate.Required("photo", "body", m.Photo); err != nil {
		return err
	}

	for i := 0; i < len(m.Photo); i++ {

		if swag.IsZero(m.Photo[i]) { // not required
			continue
		}

		if m.Photo[i] != nil {

			if err := m.Photo[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Game) validateTextEntities(formats strfmt.Registry) error {

	if swag.IsZero(m.TextEntities) { // not required
		return nil
	}

	for i := 0; i < len(m.TextEntities); i++ {

		if swag.IsZero(m.TextEntities[i]) { // not required
			continue
		}

		if m.TextEntities[i] != nil {

			if err := m.TextEntities[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Game) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}
