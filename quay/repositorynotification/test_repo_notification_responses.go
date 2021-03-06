package repositorynotification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type TestRepoNotificationReader struct {
	formats strfmt.Registry
}

func (o *TestRepoNotificationReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result TestRepoNotificationOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result TestRepoNotificationBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("testRepoNotificationBadRequest", &result, response.Code())

	case 401:
		var result TestRepoNotificationUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("testRepoNotificationUnauthorized", &result, response.Code())

	case 403:
		var result TestRepoNotificationForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("testRepoNotificationForbidden", &result, response.Code())

	case 404:
		var result TestRepoNotificationNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("testRepoNotificationNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*
Successful invocation
*/
type TestRepoNotificationOK struct {
}

func (o *TestRepoNotificationOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type TestRepoNotificationBadRequest struct {
	Payload *models.GeneralError
}

func (o *TestRepoNotificationBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

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
type TestRepoNotificationUnauthorized struct {
}

func (o *TestRepoNotificationUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type TestRepoNotificationForbidden struct {
}

func (o *TestRepoNotificationForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type TestRepoNotificationNotFound struct {
}

func (o *TestRepoNotificationNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
