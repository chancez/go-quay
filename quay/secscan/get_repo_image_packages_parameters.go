package secscan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
GetRepoImagePackagesParams contains all the parameters to send to the API endpoint
for the get repo image packages operation typically these are written to a http.Request
*/
type GetRepoImagePackagesParams struct {
	/*
	  The image ID
	*/
	Imageid string
	/*
	  The full path of the repository. e.g. namespace/name
	*/
	Repository string
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoImagePackagesParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param imageid
	if err := r.SetPathParam("imageid", o.Imageid); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
