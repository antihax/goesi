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

func easyjsonD380af9cDecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *PostUniverseNamesIdsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostUniverseNamesIdsList, 0, 2)
			} else {
				*out = PostUniverseNamesIdsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostUniverseNamesIds
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
func easyjsonD380af9cEncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in PostUniverseNamesIdsList) {
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
func (v PostUniverseNamesIdsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD380af9cEncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUniverseNamesIdsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD380af9cEncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUniverseNamesIdsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD380af9cDecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUniverseNamesIdsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD380af9cDecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjsonD380af9cDecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *PostUniverseNamesIds) {
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
		case "ids":
			if in.IsNull() {
				in.Skip()
				out.Ids = nil
			} else {
				in.Delim('[')
				if out.Ids == nil {
					if !in.IsDelim(']') {
						out.Ids = make([]int32, 0, 16)
					} else {
						out.Ids = []int32{}
					}
				} else {
					out.Ids = (out.Ids)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Ids = append(out.Ids, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonD380af9cEncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in PostUniverseNamesIds) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Ids) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ids\":")
		if in.Ids == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Ids {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUniverseNamesIds) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD380af9cEncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUniverseNamesIds) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD380af9cEncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUniverseNamesIds) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD380af9cDecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUniverseNamesIds) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD380af9cDecodeGithubComAntihaxGoesiV11(l, v)
}
