// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// SendVideoLinkReader is a Reader for the SendVideoLink structure.
type SendVideoLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVideoLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendVideoLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendVideoLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendVideoLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendVideoLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendVideoLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSendVideoLinkEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendVideoLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendVideoLinkOK creates a SendVideoLinkOK with default headers values
func NewSendVideoLinkOK() *SendVideoLinkOK {
	return &SendVideoLinkOK{}
}

/*SendVideoLinkOK handles this case with default header values.

SendVideoLinkOK send video link o k
*/
type SendVideoLinkOK struct {
	Payload *models.ResponseMessage
}

func (o *SendVideoLinkOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo#link][%d] sendVideoLinkOK  %+v", 200, o.Payload)
}

func (o *SendVideoLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoLinkBadRequest creates a SendVideoLinkBadRequest with default headers values
func NewSendVideoLinkBadRequest() *SendVideoLinkBadRequest {
	return &SendVideoLinkBadRequest{}
}

/*SendVideoLinkBadRequest handles this case with default header values.

Bad Request
*/
type SendVideoLinkBadRequest struct {
	Payload *models.Error
}

func (o *SendVideoLinkBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo#link][%d] sendVideoLinkBadRequest  %+v", 400, o.Payload)
}

func (o *SendVideoLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoLinkUnauthorized creates a SendVideoLinkUnauthorized with default headers values
func NewSendVideoLinkUnauthorized() *SendVideoLinkUnauthorized {
	return &SendVideoLinkUnauthorized{}
}

/*SendVideoLinkUnauthorized handles this case with default header values.

Unauthorized
*/
type SendVideoLinkUnauthorized struct {
	Payload *models.Error
}

func (o *SendVideoLinkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo#link][%d] sendVideoLinkUnauthorized  %+v", 401, o.Payload)
}

func (o *SendVideoLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoLinkForbidden creates a SendVideoLinkForbidden with default headers values
func NewSendVideoLinkForbidden() *SendVideoLinkForbidden {
	return &SendVideoLinkForbidden{}
}

/*SendVideoLinkForbidden handles this case with default header values.

Forbidden
*/
type SendVideoLinkForbidden struct {
	Payload *models.Error
}

func (o *SendVideoLinkForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo#link][%d] sendVideoLinkForbidden  %+v", 403, o.Payload)
}

func (o *SendVideoLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoLinkNotFound creates a SendVideoLinkNotFound with default headers values
func NewSendVideoLinkNotFound() *SendVideoLinkNotFound {
	return &SendVideoLinkNotFound{}
}

/*SendVideoLinkNotFound handles this case with default header values.

Not Found
*/
type SendVideoLinkNotFound struct {
	Payload *models.Error
}

func (o *SendVideoLinkNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo#link][%d] sendVideoLinkNotFound  %+v", 404, o.Payload)
}

func (o *SendVideoLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoLinkEnhanceYourCalm creates a SendVideoLinkEnhanceYourCalm with default headers values
func NewSendVideoLinkEnhanceYourCalm() *SendVideoLinkEnhanceYourCalm {
	return &SendVideoLinkEnhanceYourCalm{}
}

/*SendVideoLinkEnhanceYourCalm handles this case with default header values.

Flood
*/
type SendVideoLinkEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SendVideoLinkEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo#link][%d] sendVideoLinkEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SendVideoLinkEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoLinkInternalServerError creates a SendVideoLinkInternalServerError with default headers values
func NewSendVideoLinkInternalServerError() *SendVideoLinkInternalServerError {
	return &SendVideoLinkInternalServerError{}
}

/*SendVideoLinkInternalServerError handles this case with default header values.

Internal
*/
type SendVideoLinkInternalServerError struct {
	Payload *models.Error
}

func (o *SendVideoLinkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo#link][%d] sendVideoLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *SendVideoLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
