package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
ConductSearchParams contains all the parameters to send to the API endpoint
for the conduct search operation typically these are written to a http.Request
*/
type ConductSearchParams struct {
	/*
	  The search query.
	*/
	Query string
}

// WriteToRequest writes these params to a swagger request
func (o *ConductSearchParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param query
	if err := r.SetQueryParam("query", o.Query); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
