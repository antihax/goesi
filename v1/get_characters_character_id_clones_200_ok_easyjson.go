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

func easyjsonF7130648DecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetCharactersCharacterIdClones200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdClones200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdClones200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdClones200Ok
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
func easyjsonF7130648EncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetCharactersCharacterIdClones200OkList) {
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
func (v GetCharactersCharacterIdClones200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF7130648EncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdClones200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF7130648EncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdClones200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF7130648DecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdClones200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF7130648DecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjsonF7130648DecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetCharactersCharacterIdClones200Ok) {
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
		case "implants":
			if in.IsNull() {
				in.Skip()
				out.Implants = nil
			} else {
				in.Delim('[')
				if out.Implants == nil {
					if !in.IsDelim(']') {
						out.Implants = make([]int32, 0, 16)
					} else {
						out.Implants = []int32{}
					}
				} else {
					out.Implants = (out.Implants)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Implants = append(out.Implants, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "location_type":
			out.LocationType = string(in.String())
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
func easyjsonF7130648EncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetCharactersCharacterIdClones200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Implants) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"implants\":")
		if in.Implants == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Implants {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	if in.LocationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"location_id\":")
		out.Int64(int64(in.LocationId))
	}
	if in.LocationType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"location_type\":")
		out.String(string(in.LocationType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdClones200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF7130648EncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdClones200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF7130648EncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdClones200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF7130648DecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdClones200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF7130648DecodeGithubComAntihaxGoesiV11(l, v)
}
