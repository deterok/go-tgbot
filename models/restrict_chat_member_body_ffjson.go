// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/restrict_chat_member_body.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *RestrictChatMemberBody) MarshalJSON() ([]byte, error) {
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
func (mj *RestrictChatMemberBody) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteByte('{')
	if mj.CanAddWebPagePreviews != false {
		if mj.CanAddWebPagePreviews {
			buf.WriteString(`"can_add_web_page_previews":true`)
		} else {
			buf.WriteString(`"can_add_web_page_previews":false`)
		}
		buf.WriteByte(',')
	}
	if mj.CanSendMediaMessages != false {
		if mj.CanSendMediaMessages {
			buf.WriteString(`"can_send_media_messages":true`)
		} else {
			buf.WriteString(`"can_send_media_messages":false`)
		}
		buf.WriteByte(',')
	}
	if mj.CanSendMessages != false {
		if mj.CanSendMessages {
			buf.WriteString(`"can_send_messages":true`)
		} else {
			buf.WriteString(`"can_send_messages":false`)
		}
		buf.WriteByte(',')
	}
	if mj.CanSendOtherMessages != false {
		if mj.CanSendOtherMessages {
			buf.WriteString(`"can_send_other_messages":true`)
		} else {
			buf.WriteString(`"can_send_other_messages":false`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"chat_id":`)
	/* Interface types must use runtime reflection. type=interface {} kind=interface */
	err = buf.Encode(mj.ChatID)
	if err != nil {
		return err
	}
	buf.WriteByte(',')
	if mj.UntilDate != 0 {
		buf.WriteString(`"until_date":`)
		fflib.FormatBits2(buf, uint64(mj.UntilDate), 10, mj.UntilDate < 0)
		buf.WriteByte(',')
	}
	if mj.UserID != nil {
		buf.WriteString(`"user_id":`)
		fflib.FormatBits2(buf, uint64(*mj.UserID), 10, *mj.UserID < 0)
	} else {
		buf.WriteString(`"user_id":null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_RestrictChatMemberBodybase = iota
	ffj_t_RestrictChatMemberBodyno_such_key

	ffj_t_RestrictChatMemberBody_CanAddWebPagePreviews

	ffj_t_RestrictChatMemberBody_CanSendMediaMessages

	ffj_t_RestrictChatMemberBody_CanSendMessages

	ffj_t_RestrictChatMemberBody_CanSendOtherMessages

	ffj_t_RestrictChatMemberBody_ChatID

	ffj_t_RestrictChatMemberBody_UntilDate

	ffj_t_RestrictChatMemberBody_UserID
)

var ffj_key_RestrictChatMemberBody_CanAddWebPagePreviews = []byte("can_add_web_page_previews")

var ffj_key_RestrictChatMemberBody_CanSendMediaMessages = []byte("can_send_media_messages")

var ffj_key_RestrictChatMemberBody_CanSendMessages = []byte("can_send_messages")

var ffj_key_RestrictChatMemberBody_CanSendOtherMessages = []byte("can_send_other_messages")

var ffj_key_RestrictChatMemberBody_ChatID = []byte("chat_id")

var ffj_key_RestrictChatMemberBody_UntilDate = []byte("until_date")

var ffj_key_RestrictChatMemberBody_UserID = []byte("user_id")

func (uj *RestrictChatMemberBody) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *RestrictChatMemberBody) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_RestrictChatMemberBodybase
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
				currentKey = ffj_t_RestrictChatMemberBodyno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffj_key_RestrictChatMemberBody_CanAddWebPagePreviews, kn) {
						currentKey = ffj_t_RestrictChatMemberBody_CanAddWebPagePreviews
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_RestrictChatMemberBody_CanSendMediaMessages, kn) {
						currentKey = ffj_t_RestrictChatMemberBody_CanSendMediaMessages
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_RestrictChatMemberBody_CanSendMessages, kn) {
						currentKey = ffj_t_RestrictChatMemberBody_CanSendMessages
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_RestrictChatMemberBody_CanSendOtherMessages, kn) {
						currentKey = ffj_t_RestrictChatMemberBody_CanSendOtherMessages
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_RestrictChatMemberBody_ChatID, kn) {
						currentKey = ffj_t_RestrictChatMemberBody_ChatID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffj_key_RestrictChatMemberBody_UntilDate, kn) {
						currentKey = ffj_t_RestrictChatMemberBody_UntilDate
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_RestrictChatMemberBody_UserID, kn) {
						currentKey = ffj_t_RestrictChatMemberBody_UserID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_RestrictChatMemberBody_UserID, kn) {
					currentKey = ffj_t_RestrictChatMemberBody_UserID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_RestrictChatMemberBody_UntilDate, kn) {
					currentKey = ffj_t_RestrictChatMemberBody_UntilDate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_RestrictChatMemberBody_ChatID, kn) {
					currentKey = ffj_t_RestrictChatMemberBody_ChatID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_RestrictChatMemberBody_CanSendOtherMessages, kn) {
					currentKey = ffj_t_RestrictChatMemberBody_CanSendOtherMessages
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_RestrictChatMemberBody_CanSendMessages, kn) {
					currentKey = ffj_t_RestrictChatMemberBody_CanSendMessages
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_RestrictChatMemberBody_CanSendMediaMessages, kn) {
					currentKey = ffj_t_RestrictChatMemberBody_CanSendMediaMessages
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_RestrictChatMemberBody_CanAddWebPagePreviews, kn) {
					currentKey = ffj_t_RestrictChatMemberBody_CanAddWebPagePreviews
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_RestrictChatMemberBodyno_such_key
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

				case ffj_t_RestrictChatMemberBody_CanAddWebPagePreviews:
					goto handle_CanAddWebPagePreviews

				case ffj_t_RestrictChatMemberBody_CanSendMediaMessages:
					goto handle_CanSendMediaMessages

				case ffj_t_RestrictChatMemberBody_CanSendMessages:
					goto handle_CanSendMessages

				case ffj_t_RestrictChatMemberBody_CanSendOtherMessages:
					goto handle_CanSendOtherMessages

				case ffj_t_RestrictChatMemberBody_ChatID:
					goto handle_ChatID

				case ffj_t_RestrictChatMemberBody_UntilDate:
					goto handle_UntilDate

				case ffj_t_RestrictChatMemberBody_UserID:
					goto handle_UserID

				case ffj_t_RestrictChatMemberBodyno_such_key:
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

handle_CanAddWebPagePreviews:

	/* handler: uj.CanAddWebPagePreviews type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.CanAddWebPagePreviews = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.CanAddWebPagePreviews = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CanSendMediaMessages:

	/* handler: uj.CanSendMediaMessages type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.CanSendMediaMessages = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.CanSendMediaMessages = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CanSendMessages:

	/* handler: uj.CanSendMessages type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.CanSendMessages = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.CanSendMessages = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CanSendOtherMessages:

	/* handler: uj.CanSendOtherMessages type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.CanSendOtherMessages = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.CanSendOtherMessages = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ChatID:

	/* handler: uj.ChatID type=interface {} kind=interface quoted=false*/

	{
		/* Falling back. type=interface {} kind=interface */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.ChatID)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_UntilDate:

	/* handler: uj.UntilDate type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.UntilDate = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_UserID:

	/* handler: uj.UserID type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

			uj.UserID = nil

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			ttypval := int64(tval)
			uj.UserID = &ttypval

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
