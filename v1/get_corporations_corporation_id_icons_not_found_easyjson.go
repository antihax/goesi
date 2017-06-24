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

func easyjson92dedbf5DecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetCorporationsCorporationIdIconsNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdIconsNotFoundList, 0, 4)
			} else {
				*out = GetCorporationsCorporationIdIconsNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdIconsNotFound
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
func easyjson92dedbf5EncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetCorporationsCorporationIdIconsNotFoundList) {
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
func (v GetCorporationsCorporationIdIconsNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson92dedbf5EncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdIconsNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson92dedbf5EncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdIconsNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson92dedbf5DecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdIconsNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson92dedbf5DecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjson92dedbf5DecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetCorporationsCorporationIdIconsNotFound) {
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
func easyjson92dedbf5EncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetCorporationsCorporationIdIconsNotFound) {
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
func (v GetCorporationsCorporationIdIconsNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson92dedbf5EncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdIconsNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson92dedbf5EncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdIconsNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson92dedbf5DecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdIconsNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson92dedbf5DecodeGithubComAntihaxGoesiV11(l, v)
}
