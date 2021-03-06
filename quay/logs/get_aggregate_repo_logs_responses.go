package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type GetAggregateRepoLogsReader struct {
	formats strfmt.Registry
}

func (o *GetAggregateRepoLogsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetAggregateRepoLogsOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result GetAggregateRepoLogsBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getAggregateRepoLogsBadRequest", &result, response.Code())

	case 401:
		var result GetAggregateRepoLogsUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getAggregateRepoLogsUnauthorized", &result, response.Code())

	case 403:
		var result GetAggregateRepoLogsForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getAggregateRepoLogsForbidden", &result, response.Code())

	case 404:
		var result GetAggregateRepoLogsNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getAggregateRepoLogsNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*
Successful invocation
*/
type GetAggregateRepoLogsOK struct {
}

func (o *GetAggregateRepoLogsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type GetAggregateRepoLogsBadRequest struct {
	Payload *models.GeneralError
}

func (o *GetAggregateRepoLogsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

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
type GetAggregateRepoLogsUnauthorized struct {
}

func (o *GetAggregateRepoLogsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type GetAggregateRepoLogsForbidden struct {
}

func (o *GetAggregateRepoLogsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type GetAggregateRepoLogsNotFound struct {
}

func (o *GetAggregateRepoLogsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
