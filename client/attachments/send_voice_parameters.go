// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSendVoiceParams creates a new SendVoiceParams object
// with the default values initialized.
func NewSendVoiceParams() *SendVoiceParams {
	var ()
	return &SendVoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendVoiceParamsWithTimeout creates a new SendVoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendVoiceParamsWithTimeout(timeout time.Duration) *SendVoiceParams {
	var ()
	return &SendVoiceParams{

		timeout: timeout,
	}
}

// NewSendVoiceParamsWithContext creates a new SendVoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendVoiceParamsWithContext(ctx context.Context) *SendVoiceParams {
	var ()
	return &SendVoiceParams{

		Context: ctx,
	}
}

// NewSendVoiceParamsWithHTTPClient creates a new SendVoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendVoiceParamsWithHTTPClient(client *http.Client) *SendVoiceParams {
	var ()
	return &SendVoiceParams{
		HTTPClient: client,
	}
}

/*SendVoiceParams contains all the parameters to send to the API endpoint
for the send voice operation typically these are written to a http.Request
*/
type SendVoiceParams struct {

	/*Caption*/
	Caption *string
	/*ChatID*/
	ChatID string
	/*DisableNotification*/
	DisableNotification *bool
	/*Duration*/
	Duration *int64
	/*ParseMode*/
	ParseMode *string
	/*ReplyMarkup
	  json string of reply_markup object

	*/
	ReplyMarkup *string
	/*ReplyToMessageID*/
	ReplyToMessageID *int64
	/*Token
	  bot's token to authorize the request

	*/
	Token *string
	/*Voice*/
	Voice runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send voice params
func (o *SendVoiceParams) WithTimeout(timeout time.Duration) *SendVoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send voice params
func (o *SendVoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send voice params
func (o *SendVoiceParams) WithContext(ctx context.Context) *SendVoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send voice params
func (o *SendVoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send voice params
func (o *SendVoiceParams) WithHTTPClient(client *http.Client) *SendVoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send voice params
func (o *SendVoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaption adds the caption to the send voice params
func (o *SendVoiceParams) WithCaption(caption *string) *SendVoiceParams {
	o.SetCaption(caption)
	return o
}

// SetCaption adds the caption to the send voice params
func (o *SendVoiceParams) SetCaption(caption *string) {
	o.Caption = caption
}

// WithChatID adds the chatID to the send voice params
func (o *SendVoiceParams) WithChatID(chatID string) *SendVoiceParams {
	o.SetChatID(chatID)
	return o
}

// SetChatID adds the chatId to the send voice params
func (o *SendVoiceParams) SetChatID(chatID string) {
	o.ChatID = chatID
}

// WithDisableNotification adds the disableNotification to the send voice params
func (o *SendVoiceParams) WithDisableNotification(disableNotification *bool) *SendVoiceParams {
	o.SetDisableNotification(disableNotification)
	return o
}

// SetDisableNotification adds the disableNotification to the send voice params
func (o *SendVoiceParams) SetDisableNotification(disableNotification *bool) {
	o.DisableNotification = disableNotification
}

// WithDuration adds the duration to the send voice params
func (o *SendVoiceParams) WithDuration(duration *int64) *SendVoiceParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the send voice params
func (o *SendVoiceParams) SetDuration(duration *int64) {
	o.Duration = duration
}

// WithParseMode adds the parseMode to the send voice params
func (o *SendVoiceParams) WithParseMode(parseMode *string) *SendVoiceParams {
	o.SetParseMode(parseMode)
	return o
}

// SetParseMode adds the parseMode to the send voice params
func (o *SendVoiceParams) SetParseMode(parseMode *string) {
	o.ParseMode = parseMode
}

// WithReplyMarkup adds the replyMarkup to the send voice params
func (o *SendVoiceParams) WithReplyMarkup(replyMarkup *string) *SendVoiceParams {
	o.SetReplyMarkup(replyMarkup)
	return o
}

// SetReplyMarkup adds the replyMarkup to the send voice params
func (o *SendVoiceParams) SetReplyMarkup(replyMarkup *string) {
	o.ReplyMarkup = replyMarkup
}

// WithReplyToMessageID adds the replyToMessageID to the send voice params
func (o *SendVoiceParams) WithReplyToMessageID(replyToMessageID *int64) *SendVoiceParams {
	o.SetReplyToMessageID(replyToMessageID)
	return o
}

// SetReplyToMessageID adds the replyToMessageId to the send voice params
func (o *SendVoiceParams) SetReplyToMessageID(replyToMessageID *int64) {
	o.ReplyToMessageID = replyToMessageID
}

// WithToken adds the token to the send voice params
func (o *SendVoiceParams) WithToken(token *string) *SendVoiceParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send voice params
func (o *SendVoiceParams) SetToken(token *string) {
	o.Token = token
}

// WithVoice adds the voice to the send voice params
func (o *SendVoiceParams) WithVoice(voice runtime.NamedReadCloser) *SendVoiceParams {
	o.SetVoice(voice)
	return o
}

// SetVoice adds the voice to the send voice params
func (o *SendVoiceParams) SetVoice(voice runtime.NamedReadCloser) {
	o.Voice = voice
}

// WriteToRequest writes these params to a swagger request
func (o *SendVoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Caption != nil {

		// form param caption
		var frCaption string
		if o.Caption != nil {
			frCaption = *o.Caption
		}
		fCaption := frCaption
		if fCaption != "" {
			if err := r.SetFormParam("caption", fCaption); err != nil {
				return err
			}
		}

	}

	// form param chat_id
	frChatID := o.ChatID
	fChatID := frChatID
	if fChatID != "" {
		if err := r.SetFormParam("chat_id", fChatID); err != nil {
			return err
		}
	}

	if o.DisableNotification != nil {

		// form param disable_notification
		var frDisableNotification bool
		if o.DisableNotification != nil {
			frDisableNotification = *o.DisableNotification
		}
		fDisableNotification := swag.FormatBool(frDisableNotification)
		if fDisableNotification != "" {
			if err := r.SetFormParam("disable_notification", fDisableNotification); err != nil {
				return err
			}
		}

	}

	if o.Duration != nil {

		// form param duration
		var frDuration int64
		if o.Duration != nil {
			frDuration = *o.Duration
		}
		fDuration := swag.FormatInt64(frDuration)
		if fDuration != "" {
			if err := r.SetFormParam("duration", fDuration); err != nil {
				return err
			}
		}

	}

	if o.ParseMode != nil {

		// form param parse_mode
		var frParseMode string
		if o.ParseMode != nil {
			frParseMode = *o.ParseMode
		}
		fParseMode := frParseMode
		if fParseMode != "" {
			if err := r.SetFormParam("parse_mode", fParseMode); err != nil {
				return err
			}
		}

	}

	if o.ReplyMarkup != nil {

		// form param reply_markup
		var frReplyMarkup string
		if o.ReplyMarkup != nil {
			frReplyMarkup = *o.ReplyMarkup
		}
		fReplyMarkup := frReplyMarkup
		if fReplyMarkup != "" {
			if err := r.SetFormParam("reply_markup", fReplyMarkup); err != nil {
				return err
			}
		}

	}

	if o.ReplyToMessageID != nil {

		// form param reply_to_message_id
		var frReplyToMessageID int64
		if o.ReplyToMessageID != nil {
			frReplyToMessageID = *o.ReplyToMessageID
		}
		fReplyToMessageID := swag.FormatInt64(frReplyToMessageID)
		if fReplyToMessageID != "" {
			if err := r.SetFormParam("reply_to_message_id", fReplyToMessageID); err != nil {
				return err
			}
		}

	}

	if o.Token != nil {

		// path param token
		if err := r.SetPathParam("token", *o.Token); err != nil {
			return err
		}

	}

	// form file param voice
	if err := r.SetFileParam("voice", o.Voice); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
