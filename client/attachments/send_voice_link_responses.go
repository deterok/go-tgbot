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

// SendVoiceLinkReader is a Reader for the SendVoiceLink structure.
type SendVoiceLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVoiceLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendVoiceLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendVoiceLinkOK creates a SendVoiceLinkOK with default headers values
func NewSendVoiceLinkOK() *SendVoiceLinkOK {
	return &SendVoiceLinkOK{}
}

/*SendVoiceLinkOK handles this case with default header values.

SendVoiceLinkOK send voice link o k
*/
type SendVoiceLinkOK struct {
	Payload SendVoiceLinkOKBody
}

func (o *SendVoiceLinkOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVoice#link][%d] sendVoiceLinkOK  %+v", 200, o.Payload)
}

func (o *SendVoiceLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SendVoiceLinkBody send voice link body
swagger:model SendVoiceLinkBody
*/
type SendVoiceLinkBody struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// chat id
	// Required: true
	ChatID interface{} `json:"chat_id"`

	// disable notification
	DisableNotification bool `json:"disable_notification,omitempty"`

	// duration
	Duration int64 `json:"duration,omitempty"`

	// reply markup
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`

	// reply to message id
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

	// voice
	// Required: true
	Voice *string `json:"voice"`
}

/*SendVoiceLinkOKBody send voice link o k body
swagger:model SendVoiceLinkOKBody
*/
type SendVoiceLinkOKBody struct {

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
