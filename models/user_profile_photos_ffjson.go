// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/user_profile_photos.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *UserProfilePhotos) MarshalJSON() ([]byte, error) {
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
func (mj *UserProfilePhotos) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "photos":`)
	if mj.Photos != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Photos {
			if i != 0 {
				buf.WriteString(`,`)
			}
			if v != nil {
				buf.WriteString(`[`)
				for i, v := range v {
					if i != 0 {
						buf.WriteString(`,`)
					}

					{

						if v == nil {
							buf.WriteString("null")
							return nil
						}

						err = v.MarshalJSONBuf(buf)
						if err != nil {
							return err
						}

					}
				}
				buf.WriteString(`]`)
			} else {
				buf.WriteString(`null`)
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte(',')
	if mj.TotalCount != 0 {
		buf.WriteString(`"total_count":`)
		fflib.FormatBits2(buf, uint64(mj.TotalCount), 10, mj.TotalCount < 0)
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_UserProfilePhotosbase = iota
	ffj_t_UserProfilePhotosno_such_key

	ffj_t_UserProfilePhotos_Photos

	ffj_t_UserProfilePhotos_TotalCount
)

var ffj_key_UserProfilePhotos_Photos = []byte("photos")

var ffj_key_UserProfilePhotos_TotalCount = []byte("total_count")

func (uj *UserProfilePhotos) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *UserProfilePhotos) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_UserProfilePhotosbase
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
				currentKey = ffj_t_UserProfilePhotosno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'p':

					if bytes.Equal(ffj_key_UserProfilePhotos_Photos, kn) {
						currentKey = ffj_t_UserProfilePhotos_Photos
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_UserProfilePhotos_TotalCount, kn) {
						currentKey = ffj_t_UserProfilePhotos_TotalCount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.AsciiEqualFold(ffj_key_UserProfilePhotos_TotalCount, kn) {
					currentKey = ffj_t_UserProfilePhotos_TotalCount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_UserProfilePhotos_Photos, kn) {
					currentKey = ffj_t_UserProfilePhotos_Photos
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_UserProfilePhotosno_such_key
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

				case ffj_t_UserProfilePhotos_Photos:
					goto handle_Photos

				case ffj_t_UserProfilePhotos_TotalCount:
					goto handle_TotalCount

				case ffj_t_UserProfilePhotosno_such_key:
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

handle_Photos:

	/* handler: uj.Photos type=[][]*models.PhotoSize kind=slice quoted=false*/

	{
		/* Falling back. type=[][]*models.PhotoSize kind=slice */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Photos)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TotalCount:

	/* handler: uj.TotalCount type=int64 kind=int64 quoted=false*/

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

			uj.TotalCount = int64(tval)

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