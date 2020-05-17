// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// GetNextStepsReader is a Reader for the GetNextSteps structure.
type GetNextStepsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNextStepsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNextStepsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNextStepsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNextStepsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNextStepsOK creates a GetNextStepsOK with default headers values
func NewGetNextStepsOK() *GetNextStepsOK {
	return &GetNextStepsOK{}
}

/*GetNextStepsOK handles this case with default header values.

Success.
*/
type GetNextStepsOK struct {
	Payload models.Steps
}

func (o *GetNextStepsOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/instructions][%d] getNextStepsOK  %+v", 200, o.Payload)
}

func (o *GetNextStepsOK) GetPayload() models.Steps {
	return o.Payload
}

func (o *GetNextStepsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNextStepsNotFound creates a GetNextStepsNotFound with default headers values
func NewGetNextStepsNotFound() *GetNextStepsNotFound {
	return &GetNextStepsNotFound{}
}

/*GetNextStepsNotFound handles this case with default header values.

Error.
*/
type GetNextStepsNotFound struct {
	Payload *models.Error
}

func (o *GetNextStepsNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/instructions][%d] getNextStepsNotFound  %+v", 404, o.Payload)
}

func (o *GetNextStepsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNextStepsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNextStepsInternalServerError creates a GetNextStepsInternalServerError with default headers values
func NewGetNextStepsInternalServerError() *GetNextStepsInternalServerError {
	return &GetNextStepsInternalServerError{}
}

/*GetNextStepsInternalServerError handles this case with default header values.

Error.
*/
type GetNextStepsInternalServerError struct {
	Payload *models.Error
}

func (o *GetNextStepsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/instructions][%d] getNextStepsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNextStepsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNextStepsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}