package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*Member Member member

swagger:model Member
*/
type Member struct {

	/* Avatar avatar
	 */
	Avatar *Avatar `json:"avatar,omitempty"`

	/* Name name
	 */
	Name *string `json:"name,omitempty"`

	/* Repositories repositories
	 */
	Repositories []string `json:"repositories,omitempty"`

	/* Teams teams
	 */
	Teams []*Team `json:"teams,omitempty"`
}

// Validate validates this member
func (m *Member) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepositories(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Member) validateRepositories(formats strfmt.Registry) error {

	if swag.IsZero(m.Repositories) { // not required
		return nil
	}

	for i := 0; i < len(m.Repositories); i++ {

		if err := validate.RequiredString("repositories"+"."+strconv.Itoa(i), "body", string(m.Repositories[i])); err != nil {
			return err
		}

	}

	return nil
}

func (m *Member) validateTeams(formats strfmt.Registry) error {

	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {

		if m.Teams[i] != nil {

			if err := m.Teams[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}