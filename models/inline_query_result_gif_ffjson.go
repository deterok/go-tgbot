// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/inline_query_result_gif.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *InlineQueryResultGif) MarshalJSON() ([]byte, error) {
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
func (mj *InlineQueryResultGif) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteByte('{')
	if len(mj.Caption) != 0 {
		buf.WriteString(`"caption":`)
		fflib.WriteJsonString(buf, string(mj.Caption))
		buf.WriteByte(',')
	}
	if mj.GifDuration != 0 {
		buf.WriteString(`"gif_duration":`)
		fflib.FormatBits2(buf, uint64(mj.GifDuration), 10, mj.GifDuration < 0)
		buf.WriteByte(',')
	}
	if mj.GifHeight != 0 {
		buf.WriteString(`"gif_height":`)
		fflib.FormatBits2(buf, uint64(mj.GifHeight), 10, mj.GifHeight < 0)
		buf.WriteByte(',')
	}
	if mj.GifURL != nil {
		buf.WriteString(`"gif_url":`)
		fflib.WriteJsonString(buf, string(*mj.GifURL))
	} else {
		buf.WriteString(`"gif_url":null`)
	}
	buf.WriteByte(',')
	if mj.GifWidth != 0 {
		buf.WriteString(`"gif_width":`)
		fflib.FormatBits2(buf, uint64(mj.GifWidth), 10, mj.GifWidth < 0)
		buf.WriteByte(',')
	}
	if mj.ID != nil {
		buf.WriteString(`"id":`)
		fflib.WriteJsonString(buf, string(*mj.ID))
	} else {
		buf.WriteString(`"id":null`)
	}
	buf.WriteByte(',')
	if mj.InputMessageContent != nil {
		buf.WriteString(`"input_message_content":`)
		/* Interface types must use runtime reflection. type=interface {} kind=interface */
		err = buf.Encode(mj.InputMessageContent)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if mj.ReplyMarkup != nil {
		if true {
			buf.WriteString(`"reply_markup":`)

			{

				err = mj.ReplyMarkup.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.ThumbURL != nil {
		buf.WriteString(`"thumb_url":`)
		fflib.WriteJsonString(buf, string(*mj.ThumbURL))
	} else {
		buf.WriteString(`"thumb_url":null`)
	}
	buf.WriteByte(',')
	if len(mj.Title) != 0 {
		buf.WriteString(`"title":`)
		fflib.WriteJsonString(buf, string(mj.Title))
		buf.WriteByte(',')
	}
	buf.WriteString(`"type":`)
	fflib.WriteJsonString(buf, string(mj.Type))
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_InlineQueryResultGifbase = iota
	ffj_t_InlineQueryResultGifno_such_key

	ffj_t_InlineQueryResultGif_Caption

	ffj_t_InlineQueryResultGif_GifDuration

	ffj_t_InlineQueryResultGif_GifHeight

	ffj_t_InlineQueryResultGif_GifURL

	ffj_t_InlineQueryResultGif_GifWidth

	ffj_t_InlineQueryResultGif_ID

	ffj_t_InlineQueryResultGif_InputMessageContent

	ffj_t_InlineQueryResultGif_ReplyMarkup

	ffj_t_InlineQueryResultGif_ThumbURL

	ffj_t_InlineQueryResultGif_Title

	ffj_t_InlineQueryResultGif_Type
)

var ffj_key_InlineQueryResultGif_Caption = []byte("caption")

var ffj_key_InlineQueryResultGif_GifDuration = []byte("gif_duration")

var ffj_key_InlineQueryResultGif_GifHeight = []byte("gif_height")

var ffj_key_InlineQueryResultGif_GifURL = []byte("gif_url")

var ffj_key_InlineQueryResultGif_GifWidth = []byte("gif_width")

var ffj_key_InlineQueryResultGif_ID = []byte("id")

var ffj_key_InlineQueryResultGif_InputMessageContent = []byte("input_message_content")

var ffj_key_InlineQueryResultGif_ReplyMarkup = []byte("reply_markup")

var ffj_key_InlineQueryResultGif_ThumbURL = []byte("thumb_url")

var ffj_key_InlineQueryResultGif_Title = []byte("title")

var ffj_key_InlineQueryResultGif_Type = []byte("type")

func (uj *InlineQueryResultGif) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *InlineQueryResultGif) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_InlineQueryResultGifbase
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
				currentKey = ffj_t_InlineQueryResultGifno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffj_key_InlineQueryResultGif_Caption, kn) {
						currentKey = ffj_t_InlineQueryResultGif_Caption
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'g':

					if bytes.Equal(ffj_key_InlineQueryResultGif_GifDuration, kn) {
						currentKey = ffj_t_InlineQueryResultGif_GifDuration
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultGif_GifHeight, kn) {
						currentKey = ffj_t_InlineQueryResultGif_GifHeight
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultGif_GifURL, kn) {
						currentKey = ffj_t_InlineQueryResultGif_GifURL
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultGif_GifWidth, kn) {
						currentKey = ffj_t_InlineQueryResultGif_GifWidth
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_InlineQueryResultGif_ID, kn) {
						currentKey = ffj_t_InlineQueryResultGif_ID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultGif_InputMessageContent, kn) {
						currentKey = ffj_t_InlineQueryResultGif_InputMessageContent
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffj_key_InlineQueryResultGif_ReplyMarkup, kn) {
						currentKey = ffj_t_InlineQueryResultGif_ReplyMarkup
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_InlineQueryResultGif_ThumbURL, kn) {
						currentKey = ffj_t_InlineQueryResultGif_ThumbURL
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultGif_Title, kn) {
						currentKey = ffj_t_InlineQueryResultGif_Title
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultGif_Type, kn) {
						currentKey = ffj_t_InlineQueryResultGif_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_InlineQueryResultGif_Type, kn) {
					currentKey = ffj_t_InlineQueryResultGif_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_InlineQueryResultGif_Title, kn) {
					currentKey = ffj_t_InlineQueryResultGif_Title
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_InlineQueryResultGif_ThumbURL, kn) {
					currentKey = ffj_t_InlineQueryResultGif_ThumbURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InlineQueryResultGif_ReplyMarkup, kn) {
					currentKey = ffj_t_InlineQueryResultGif_ReplyMarkup
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InlineQueryResultGif_InputMessageContent, kn) {
					currentKey = ffj_t_InlineQueryResultGif_InputMessageContent
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_InlineQueryResultGif_ID, kn) {
					currentKey = ffj_t_InlineQueryResultGif_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_InlineQueryResultGif_GifWidth, kn) {
					currentKey = ffj_t_InlineQueryResultGif_GifWidth
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_InlineQueryResultGif_GifURL, kn) {
					currentKey = ffj_t_InlineQueryResultGif_GifURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_InlineQueryResultGif_GifHeight, kn) {
					currentKey = ffj_t_InlineQueryResultGif_GifHeight
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_InlineQueryResultGif_GifDuration, kn) {
					currentKey = ffj_t_InlineQueryResultGif_GifDuration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_InlineQueryResultGif_Caption, kn) {
					currentKey = ffj_t_InlineQueryResultGif_Caption
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_InlineQueryResultGifno_such_key
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

				case ffj_t_InlineQueryResultGif_Caption:
					goto handle_Caption

				case ffj_t_InlineQueryResultGif_GifDuration:
					goto handle_GifDuration

				case ffj_t_InlineQueryResultGif_GifHeight:
					goto handle_GifHeight

				case ffj_t_InlineQueryResultGif_GifURL:
					goto handle_GifURL

				case ffj_t_InlineQueryResultGif_GifWidth:
					goto handle_GifWidth

				case ffj_t_InlineQueryResultGif_ID:
					goto handle_ID

				case ffj_t_InlineQueryResultGif_InputMessageContent:
					goto handle_InputMessageContent

				case ffj_t_InlineQueryResultGif_ReplyMarkup:
					goto handle_ReplyMarkup

				case ffj_t_InlineQueryResultGif_ThumbURL:
					goto handle_ThumbURL

				case ffj_t_InlineQueryResultGif_Title:
					goto handle_Title

				case ffj_t_InlineQueryResultGif_Type:
					goto handle_Type

				case ffj_t_InlineQueryResultGifno_such_key:
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

handle_Caption:

	/* handler: uj.Caption type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Caption = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_GifDuration:

	/* handler: uj.GifDuration type=int64 kind=int64 quoted=false*/

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

			uj.GifDuration = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_GifHeight:

	/* handler: uj.GifHeight type=int64 kind=int64 quoted=false*/

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

			uj.GifHeight = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_GifURL:

	/* handler: uj.GifURL type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.GifURL = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.GifURL = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_GifWidth:

	/* handler: uj.GifWidth type=int64 kind=int64 quoted=false*/

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

			uj.GifWidth = int64(tval)

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

			uj.ID = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.ID = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_InputMessageContent:

	/* handler: uj.InputMessageContent type=interface {} kind=interface quoted=false*/

	{
		/* Falling back. type=interface {} kind=interface */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.InputMessageContent)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ReplyMarkup:

	/* handler: uj.ReplyMarkup type=models.InlineKeyboardMarkup kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.ReplyMarkup = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.ReplyMarkup == nil {
			uj.ReplyMarkup = new(InlineKeyboardMarkup)
		}

		err = uj.ReplyMarkup.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ThumbURL:

	/* handler: uj.ThumbURL type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.ThumbURL = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.ThumbURL = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Title:

	/* handler: uj.Title type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Title = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Type:

	/* handler: uj.Type type=models.InlineType kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for InlineType", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Type = InlineType(string(outBuf))

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
