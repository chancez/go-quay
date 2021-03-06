package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type ChangeTagImageReader struct {
	formats strfmt.Registry
}

func (o *ChangeTagImageReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result ChangeTagImageOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result ChangeTagImageBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("changeTagImageBadRequest", &result, response.Code())

	case 401:
		var result ChangeTagImageUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("changeTagImageUnauthorized", &result, response.Code())

	case 403:
		var result ChangeTagImageForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("changeTagImageForbidden", &result, response.Code())

	case 404:
		var result ChangeTagImageNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("changeTagImageNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*
Successful invocation
*/
type ChangeTagImageOK struct {
}

func (o *ChangeTagImageOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type ChangeTagImageBadRequest struct {
	Payload *models.GeneralError
}

func (o *ChangeTagImageBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

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
type ChangeTagImageUnauthorized struct {
}

func (o *ChangeTagImageUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type ChangeTagImageForbidden struct {
}

func (o *ChangeTagImageForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type ChangeTagImageNotFound struct {
}

func (o *ChangeTagImageNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
