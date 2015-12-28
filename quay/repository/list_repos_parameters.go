package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*
ListReposParams contains all the parameters to send to the API endpoint
for the list repos operation typically these are written to a http.Request
*/
type ListReposParams struct {
	/*
	  Whether to include when the repository was last modified.
	*/
	LastModified bool
	/*
	  Limit on the number of results (int)
	*/
	Limit int64
	/*
	  Filters the repositories returned to this namespace
	*/
	Namespace string
	/*
	  Offset page number. (int)
	*/
	Page int64
	/*
	  Whether to include the repository's popularity metric.
	*/
	Popularity bool
	/*
	  Adds any repositories visible to the user by virtue of being public
	*/
	Public bool
	/*
	  Filters the repositories returned to those starred by the user
	*/
	Starred bool
}

// WriteToRequest writes these params to a swagger request
func (o *ListReposParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param last_modified
	if err := r.SetQueryParam("last_modified", swag.FormatBool(o.LastModified)); err != nil {
		return err
	}

	// query param limit
	if err := r.SetQueryParam("limit", swag.FormatInt64(o.Limit)); err != nil {
		return err
	}

	// query param namespace
	if err := r.SetQueryParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param page
	if err := r.SetQueryParam("page", swag.FormatInt64(o.Page)); err != nil {
		return err
	}

	// query param popularity
	if err := r.SetQueryParam("popularity", swag.FormatBool(o.Popularity)); err != nil {
		return err
	}

	// query param public
	if err := r.SetQueryParam("public", swag.FormatBool(o.Public)); err != nil {
		return err
	}

	// query param starred
	if err := r.SetQueryParam("starred", swag.FormatBool(o.Starred)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}