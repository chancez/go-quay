package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
Description of an updated application.

swagger:model UpdateApp
*/
type UpdateApp struct {

	/* The URI for the application's homepage

	Required: true
	*/
	ApplicationURI string `json:"application_uri,omitempty"`

	/* The e-mail address of the avatar to use for the application
	 */
	AvatarEmail string `json:"avatar_email,omitempty"`

	/* The human-readable description for the application
	 */
	Description string `json:"description,omitempty"`

	/* The name of the application

	Required: true
	*/
	Name string `json:"name,omitempty"`

	/* The URI for the application's OAuth redirect

	Required: true
	*/
	RedirectURI string `json:"redirect_uri,omitempty"`
}

// Validate validates this update app
func (m *UpdateApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirectURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateApp) validateApplicationURI(formats strfmt.Registry) error {

	if err := validate.Required("application_uri", "body", string(m.ApplicationURI)); err != nil {
		return err
	}

	return nil
}

func (m *UpdateApp) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *UpdateApp) validateRedirectURI(formats strfmt.Registry) error {

	if err := validate.Required("redirect_uri", "body", string(m.RedirectURI)); err != nil {
		return err
	}

	return nil
}
