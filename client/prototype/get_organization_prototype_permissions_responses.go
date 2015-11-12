package prototype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type GetOrganizationPrototypePermissionsReader struct {
	formats strfmt.Registry
}

func (o *GetOrganizationPrototypePermissionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetOrganizationPrototypePermissionsOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result GetOrganizationPrototypePermissionsBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationPrototypePermissionsBadRequest", &result, response.Code())

	case 401:
		var result GetOrganizationPrototypePermissionsUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationPrototypePermissionsUnauthorized", &result, response.Code())

	case 403:
		var result GetOrganizationPrototypePermissionsForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationPrototypePermissionsForbidden", &result, response.Code())

	case 404:
		var result GetOrganizationPrototypePermissionsNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationPrototypePermissionsNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type GetOrganizationPrototypePermissionsOK struct {
}

func (o *GetOrganizationPrototypePermissionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type GetOrganizationPrototypePermissionsBadRequest struct {
}

func (o *GetOrganizationPrototypePermissionsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type GetOrganizationPrototypePermissionsUnauthorized struct {
}

func (o *GetOrganizationPrototypePermissionsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type GetOrganizationPrototypePermissionsForbidden struct {
}

func (o *GetOrganizationPrototypePermissionsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type GetOrganizationPrototypePermissionsNotFound struct {
}

func (o *GetOrganizationPrototypePermissionsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}