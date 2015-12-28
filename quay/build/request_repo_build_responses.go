package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type RequestRepoBuildReader struct {
	formats strfmt.Registry
}

func (o *RequestRepoBuildReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result RequestRepoBuildOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result RequestRepoBuildBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("requestRepoBuildBadRequest", &result, response.Code())

	case 401:
		var result RequestRepoBuildUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("requestRepoBuildUnauthorized", &result, response.Code())

	case 403:
		var result RequestRepoBuildForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("requestRepoBuildForbidden", &result, response.Code())

	case 404:
		var result RequestRepoBuildNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("requestRepoBuildNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*
Successful invocation
*/
type RequestRepoBuildOK struct {
}

func (o *RequestRepoBuildOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type RequestRepoBuildBadRequest struct {
	Payload *models.GeneralError
}

func (o *RequestRepoBuildBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

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
type RequestRepoBuildUnauthorized struct {
}

func (o *RequestRepoBuildUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type RequestRepoBuildForbidden struct {
}

func (o *RequestRepoBuildForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type RequestRepoBuildNotFound struct {
}

func (o *RequestRepoBuildNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}