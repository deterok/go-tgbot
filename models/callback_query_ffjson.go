// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/callback_query.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *CallbackQuery) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *CallbackQuery) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.ChatInstance) != 0 {
		buf.WriteString(`"chat_instance":`)
		fflib.WriteJsonString(buf, string(mj.ChatInstance))
		buf.WriteByte(',')
	}
	if len(mj.Data) != 0 {
		buf.WriteString(`"data":`)
		fflib.WriteJsonString(buf, string(mj.Data))
		buf.WriteByte(',')
	}
	if mj.From != nil {
		if true {
			/* Struct fall back. type=models.User kind=struct */
			buf.WriteString(`"from":`)
			err = buf.Encode(mj.From)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if len(mj.GameShortName) != 0 {
		buf.WriteString(`"game_short_name":`)
		fflib.WriteJsonString(buf, string(mj.GameShortName))
		buf.WriteByte(',')
	}
	if len(mj.ID) != 0 {
		buf.WriteString(`"id":`)
		fflib.WriteJsonString(buf, string(mj.ID))
		buf.WriteByte(',')
	}
	if len(mj.InlineMessageID) != 0 {
		buf.WriteString(`"inline_message_id":`)
		fflib.WriteJsonString(buf, string(mj.InlineMessageID))
		buf.WriteByte(',')
	}
	if mj.Message != nil {
		if true {
			/* Struct fall back. type=models.Message kind=struct */
			buf.WriteString(`"message":`)
			err = buf.Encode(mj.Message)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_CallbackQuerybase = iota
	ffj_t_CallbackQueryno_such_key

	ffj_t_CallbackQuery_ChatInstance

	ffj_t_CallbackQuery_Data

	ffj_t_CallbackQuery_From

	ffj_t_CallbackQuery_GameShortName

	ffj_t_CallbackQuery_ID

	ffj_t_CallbackQuery_InlineMessageID

	ffj_t_CallbackQuery_Message
)

var ffj_key_CallbackQuery_ChatInstance = []byte("chat_instance")

var ffj_key_CallbackQuery_Data = []byte("data")

var ffj_key_CallbackQuery_From = []byte("from")

var ffj_key_CallbackQuery_GameShortName = []byte("game_short_name")

var ffj_key_CallbackQuery_ID = []byte("id")

var ffj_key_CallbackQuery_InlineMessageID = []byte("inline_message_id")

var ffj_key_CallbackQuery_Message = []byte("message")

func (uj *CallbackQuery) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *CallbackQuery) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_CallbackQuerybase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_CallbackQueryno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffj_key_CallbackQuery_ChatInstance, kn) {
						currentKey = ffj_t_CallbackQuery_ChatInstance
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_CallbackQuery_Data, kn) {
						currentKey = ffj_t_CallbackQuery_Data
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffj_key_CallbackQuery_From, kn) {
						currentKey = ffj_t_CallbackQuery_From
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'g':

					if bytes.Equal(ffj_key_CallbackQuery_GameShortName, kn) {
						currentKey = ffj_t_CallbackQuery_GameShortName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_CallbackQuery_ID, kn) {
						currentKey = ffj_t_CallbackQuery_ID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CallbackQuery_InlineMessageID, kn) {
						currentKey = ffj_t_CallbackQuery_InlineMessageID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_CallbackQuery_Message, kn) {
						currentKey = ffj_t_CallbackQuery_Message
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_CallbackQuery_Message, kn) {
					currentKey = ffj_t_CallbackQuery_Message
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CallbackQuery_InlineMessageID, kn) {
					currentKey = ffj_t_CallbackQuery_InlineMessageID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_CallbackQuery_ID, kn) {
					currentKey = ffj_t_CallbackQuery_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CallbackQuery_GameShortName, kn) {
					currentKey = ffj_t_CallbackQuery_GameShortName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_CallbackQuery_From, kn) {
					currentKey = ffj_t_CallbackQuery_From
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_CallbackQuery_Data, kn) {
					currentKey = ffj_t_CallbackQuery_Data
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CallbackQuery_ChatInstance, kn) {
					currentKey = ffj_t_CallbackQuery_ChatInstance
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_CallbackQueryno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_CallbackQuery_ChatInstance:
					goto handle_ChatInstance

				case ffj_t_CallbackQuery_Data:
					goto handle_Data

				case ffj_t_CallbackQuery_From:
					goto handle_From

				case ffj_t_CallbackQuery_GameShortName:
					goto handle_GameShortName

				case ffj_t_CallbackQuery_ID:
					goto handle_ID

				case ffj_t_CallbackQuery_InlineMessageID:
					goto handle_InlineMessageID

				case ffj_t_CallbackQuery_Message:
					goto handle_Message

				case ffj_t_CallbackQueryno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_ChatInstance:

	/* handler: uj.ChatInstance type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ChatInstance = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Data:

	/* handler: uj.Data type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Data = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_From:

	/* handler: uj.From type=models.User kind=struct quoted=false*/

	{
		/* Falling back. type=models.User kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.From)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_GameShortName:

	/* handler: uj.GameShortName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.GameShortName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ID:

	/* handler: uj.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_InlineMessageID:

	/* handler: uj.InlineMessageID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.InlineMessageID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Message:

	/* handler: uj.Message type=models.Message kind=struct quoted=false*/

	{
		/* Falling back. type=models.Message kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Message)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
