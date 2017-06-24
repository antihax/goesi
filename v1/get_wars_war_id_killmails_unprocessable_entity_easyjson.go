// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package goesiv1

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

func easyjsonD22e9c8fDecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetWarsWarIdKillmailsUnprocessableEntityList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetWarsWarIdKillmailsUnprocessableEntityList, 0, 4)
			} else {
				*out = GetWarsWarIdKillmailsUnprocessableEntityList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetWarsWarIdKillmailsUnprocessableEntity
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
func easyjsonD22e9c8fEncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetWarsWarIdKillmailsUnprocessableEntityList) {
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
func (v GetWarsWarIdKillmailsUnprocessableEntityList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD22e9c8fEncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWarsWarIdKillmailsUnprocessableEntityList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD22e9c8fEncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWarsWarIdKillmailsUnprocessableEntityList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD22e9c8fDecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWarsWarIdKillmailsUnprocessableEntityList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD22e9c8fDecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjsonD22e9c8fDecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetWarsWarIdKillmailsUnprocessableEntity) {
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
func easyjsonD22e9c8fEncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetWarsWarIdKillmailsUnprocessableEntity) {
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
func (v GetWarsWarIdKillmailsUnprocessableEntity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD22e9c8fEncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWarsWarIdKillmailsUnprocessableEntity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD22e9c8fEncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWarsWarIdKillmailsUnprocessableEntity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD22e9c8fDecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWarsWarIdKillmailsUnprocessableEntity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD22e9c8fDecodeGithubComAntihaxGoesiV11(l, v)
}
