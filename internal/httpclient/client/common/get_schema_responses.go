// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// GetSchemaReader is a Reader for the GetSchema structure.
type GetSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSchemaOK creates a GetSchemaOK with default headers values
func NewGetSchemaOK() *GetSchemaOK {
	return &GetSchemaOK{}
}

/*GetSchemaOK handles this case with default header values.

The raw identity traits schema
*/
type GetSchemaOK struct {
	Payload interface{}
}

func (o *GetSchemaOK) Error() string {
	return fmt.Sprintf("[GET /schemas/{id}][%d] getSchemaOK  %+v", 200, o.Payload)
}

func (o *GetSchemaOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemaNotFound creates a GetSchemaNotFound with default headers values
func NewGetSchemaNotFound() *GetSchemaNotFound {
	return &GetSchemaNotFound{}
}

/*GetSchemaNotFound handles this case with default header values.

genericError
*/
type GetSchemaNotFound struct {
	Payload *models.GenericError
}

func (o *GetSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /schemas/{id}][%d] getSchemaNotFound  %+v", 404, o.Payload)
}

func (o *GetSchemaNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemaInternalServerError creates a GetSchemaInternalServerError with default headers values
func NewGetSchemaInternalServerError() *GetSchemaInternalServerError {
	return &GetSchemaInternalServerError{}
}

/*GetSchemaInternalServerError handles this case with default header values.

genericError
*/
type GetSchemaInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schemas/{id}][%d] getSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSchemaInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
