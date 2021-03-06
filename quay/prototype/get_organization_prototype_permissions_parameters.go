package prototype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
GetOrganizationPrototypePermissionsParams contains all the parameters to send to the API endpoint
for the get organization prototype permissions operation typically these are written to a http.Request
*/
type GetOrganizationPrototypePermissionsParams struct {
	/*
	  The name of the organization
	*/
	Orgname string
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationPrototypePermissionsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
