package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type DeleteOrganizationTeamMemberReader struct {
	formats strfmt.Registry
}

func (o *DeleteOrganizationTeamMemberReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		var result DeleteOrganizationTeamMemberNoContent
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result DeleteOrganizationTeamMemberBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteOrganizationTeamMemberBadRequest", &result, response.Code())

	case 401:
		var result DeleteOrganizationTeamMemberUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteOrganizationTeamMemberUnauthorized", &result, response.Code())

	case 403:
		var result DeleteOrganizationTeamMemberForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteOrganizationTeamMemberForbidden", &result, response.Code())

	case 404:
		var result DeleteOrganizationTeamMemberNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteOrganizationTeamMemberNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*
Deleted
*/
type DeleteOrganizationTeamMemberNoContent struct {
}

func (o *DeleteOrganizationTeamMemberNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type DeleteOrganizationTeamMemberBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteOrganizationTeamMemberBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*
Session required
*/
type DeleteOrganizationTeamMemberUnauthorized struct {
}

func (o *DeleteOrganizationTeamMemberUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type DeleteOrganizationTeamMemberForbidden struct {
}

func (o *DeleteOrganizationTeamMemberForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type DeleteOrganizationTeamMemberNotFound struct {
}

func (o *DeleteOrganizationTeamMemberNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}