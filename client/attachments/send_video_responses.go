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

// SendVideoReader is a Reader for the SendVideo structure.
type SendVideoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVideoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendVideoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendVideoOK creates a SendVideoOK with default headers values
func NewSendVideoOK() *SendVideoOK {
	return &SendVideoOK{}
}

/*SendVideoOK handles this case with default header values.

SendVideoOK send video o k
*/
type SendVideoOK struct {
	Payload SendVideoOKBody
}

func (o *SendVideoOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideo][%d] sendVideoOK  %+v", 200, o.Payload)
}

func (o *SendVideoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SendVideoOKBody send video o k body
swagger:model SendVideoOKBody
*/
type SendVideoOKBody struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// error code
	// Required: true
	ErrorCode *int64 `json:"error_code"`

	// ok
	// Required: true
	Ok *bool `json:"ok"`

	// result
	// Required: true
	Result *models.Message `json:"result"`
}
