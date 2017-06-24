// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package goesiv3

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF540c98aDecodeGithubComAntihaxGoesiV3(in *jlexer.Lexer, out *GetUniverseTypesTypeIdNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseTypesTypeIdNotFoundList, 0, 4)
			} else {
				*out = GetUniverseTypesTypeIdNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseTypesTypeIdNotFound
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF540c98aEncodeGithubComAntihaxGoesiV3(out *jwriter.Writer, in GetUniverseTypesTypeIdNotFoundList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseTypesTypeIdNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF540c98aEncodeGithubComAntihaxGoesiV3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF540c98aEncodeGithubComAntihaxGoesiV3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF540c98aDecodeGithubComAntihaxGoesiV3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF540c98aDecodeGithubComAntihaxGoesiV3(l, v)
}
func easyjsonF540c98aDecodeGithubComAntihaxGoesiV31(in *jlexer.Lexer, out *GetUniverseTypesTypeIdNotFound) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "error":
			out.Error_ = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF540c98aEncodeGithubComAntihaxGoesiV31(out *jwriter.Writer, in GetUniverseTypesTypeIdNotFound) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"error\":")
		out.String(string(in.Error_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseTypesTypeIdNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF540c98aEncodeGithubComAntihaxGoesiV31(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF540c98aEncodeGithubComAntihaxGoesiV31(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF540c98aDecodeGithubComAntihaxGoesiV31(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF540c98aDecodeGithubComAntihaxGoesiV31(l, v)
}
