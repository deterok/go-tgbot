// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/add_sticker_to_set_link_body.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *AddStickerToSetLinkBody) MarshalJSON() ([]byte, error) {
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
func (mj *AddStickerToSetLinkBody) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	if mj.Emojis != nil {
		buf.WriteString(`{"emojis":`)
		fflib.WriteJsonString(buf, string(*mj.Emojis))
	} else {
		buf.WriteString(`{"emojis":null`)
	}
	buf.WriteByte(',')
	if mj.MaskPosition != nil {
		if true {
			/* Struct fall back. type=models.MaskPosition kind=struct */
			buf.WriteString(`"mask_position":`)
			err = buf.Encode(mj.MaskPosition)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if mj.Name != nil {
		buf.WriteString(`"name":`)
		fflib.WriteJsonString(buf, string(*mj.Name))
	} else {
		buf.WriteString(`"name":null`)
	}
	if mj.PngSticker != nil {
		buf.WriteString(`,"png_sticker":`)
		fflib.WriteJsonString(buf, string(*mj.PngSticker))
	} else {
		buf.WriteString(`,"png_sticker":null`)
	}
	if mj.UserID != nil {
		buf.WriteString(`,"user_id":`)
		fflib.FormatBits2(buf, uint64(*mj.UserID), 10, *mj.UserID < 0)
	} else {
		buf.WriteString(`,"user_id":null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_AddStickerToSetLinkBodybase = iota
	ffj_t_AddStickerToSetLinkBodyno_such_key

	ffj_t_AddStickerToSetLinkBody_Emojis

	ffj_t_AddStickerToSetLinkBody_MaskPosition

	ffj_t_AddStickerToSetLinkBody_Name

	ffj_t_AddStickerToSetLinkBody_PngSticker

	ffj_t_AddStickerToSetLinkBody_UserID
)

var ffj_key_AddStickerToSetLinkBody_Emojis = []byte("emojis")

var ffj_key_AddStickerToSetLinkBody_MaskPosition = []byte("mask_position")

var ffj_key_AddStickerToSetLinkBody_Name = []byte("name")

var ffj_key_AddStickerToSetLinkBody_PngSticker = []byte("png_sticker")

var ffj_key_AddStickerToSetLinkBody_UserID = []byte("user_id")

func (uj *AddStickerToSetLinkBody) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *AddStickerToSetLinkBody) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_AddStickerToSetLinkBodybase
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
				currentKey = ffj_t_AddStickerToSetLinkBodyno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffj_key_AddStickerToSetLinkBody_Emojis, kn) {
						currentKey = ffj_t_AddStickerToSetLinkBody_Emojis
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_AddStickerToSetLinkBody_MaskPosition, kn) {
						currentKey = ffj_t_AddStickerToSetLinkBody_MaskPosition
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffj_key_AddStickerToSetLinkBody_Name, kn) {
						currentKey = ffj_t_AddStickerToSetLinkBody_Name
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_AddStickerToSetLinkBody_PngSticker, kn) {
						currentKey = ffj_t_AddStickerToSetLinkBody_PngSticker
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffj_key_AddStickerToSetLinkBody_UserID, kn) {
						currentKey = ffj_t_AddStickerToSetLinkBody_UserID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_AddStickerToSetLinkBody_UserID, kn) {
					currentKey = ffj_t_AddStickerToSetLinkBody_UserID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AddStickerToSetLinkBody_PngSticker, kn) {
					currentKey = ffj_t_AddStickerToSetLinkBody_PngSticker
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_AddStickerToSetLinkBody_Name, kn) {
					currentKey = ffj_t_AddStickerToSetLinkBody_Name
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AddStickerToSetLinkBody_MaskPosition, kn) {
					currentKey = ffj_t_AddStickerToSetLinkBody_MaskPosition
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AddStickerToSetLinkBody_Emojis, kn) {
					currentKey = ffj_t_AddStickerToSetLinkBody_Emojis
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_AddStickerToSetLinkBodyno_such_key
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

				case ffj_t_AddStickerToSetLinkBody_Emojis:
					goto handle_Emojis

				case ffj_t_AddStickerToSetLinkBody_MaskPosition:
					goto handle_MaskPosition

				case ffj_t_AddStickerToSetLinkBody_Name:
					goto handle_Name

				case ffj_t_AddStickerToSetLinkBody_PngSticker:
					goto handle_PngSticker

				case ffj_t_AddStickerToSetLinkBody_UserID:
					goto handle_UserID

				case ffj_t_AddStickerToSetLinkBodyno_such_key:
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

handle_Emojis:

	/* handler: uj.Emojis type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.Emojis = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.Emojis = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MaskPosition:

	/* handler: uj.MaskPosition type=models.MaskPosition kind=struct quoted=false*/

	{
		/* Falling back. type=models.MaskPosition kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.MaskPosition)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: uj.Name type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.Name = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.Name = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PngSticker:

	/* handler: uj.PngSticker type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.PngSticker = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.PngSticker = &tval

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
