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

func easyjson7d8677f5DecodeGithubComAntihaxGoesiV3(in *jlexer.Lexer, out *ForbiddenList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ForbiddenList, 0, 4)
			} else {
				*out = ForbiddenList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Forbidden
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
func easyjson7d8677f5EncodeGithubComAntihaxGoesiV3(out *jwriter.Writer, in ForbiddenList) {
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
func (v ForbiddenList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d8677f5EncodeGithubComAntihaxGoesiV3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ForbiddenList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d8677f5EncodeGithubComAntihaxGoesiV3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ForbiddenList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d8677f5DecodeGithubComAntihaxGoesiV3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ForbiddenList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d8677f5DecodeGithubComAntihaxGoesiV3(l, v)
}
func easyjson7d8677f5DecodeGithubComAntihaxGoesiV31(in *jlexer.Lexer, out *Forbidden) {
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
func easyjson7d8677f5EncodeGithubComAntihaxGoesiV31(out *jwriter.Writer, in Forbidden) {
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
func (v Forbidden) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d8677f5EncodeGithubComAntihaxGoesiV31(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Forbidden) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d8677f5EncodeGithubComAntihaxGoesiV31(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Forbidden) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d8677f5DecodeGithubComAntihaxGoesiV31(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Forbidden) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d8677f5DecodeGithubComAntihaxGoesiV31(l, v)
}
