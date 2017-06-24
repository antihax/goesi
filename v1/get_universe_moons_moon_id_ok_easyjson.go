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

func easyjsonE06267a0DecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetUniverseMoonsMoonIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseMoonsMoonIdOkList, 0, 1)
			} else {
				*out = GetUniverseMoonsMoonIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseMoonsMoonIdOk
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
func easyjsonE06267a0EncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetUniverseMoonsMoonIdOkList) {
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
func (v GetUniverseMoonsMoonIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE06267a0EncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseMoonsMoonIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE06267a0EncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseMoonsMoonIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE06267a0DecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseMoonsMoonIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE06267a0DecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjsonE06267a0DecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetUniverseMoonsMoonIdOk) {
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
		case "moon_id":
			out.MoonId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "position":
			easyjsonE06267a0DecodeGithubComAntihaxGoesiV12(in, &out.Position)
		case "system_id":
			out.SystemId = int32(in.Int32())
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
func easyjsonE06267a0EncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetUniverseMoonsMoonIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MoonId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"moon_id\":")
		out.Int32(int32(in.MoonId))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"position\":")
		easyjsonE06267a0EncodeGithubComAntihaxGoesiV12(out, in.Position)
	}
	if in.SystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"system_id\":")
		out.Int32(int32(in.SystemId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseMoonsMoonIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE06267a0EncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseMoonsMoonIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE06267a0EncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseMoonsMoonIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE06267a0DecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseMoonsMoonIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE06267a0DecodeGithubComAntihaxGoesiV11(l, v)
}
func easyjsonE06267a0DecodeGithubComAntihaxGoesiV12(in *jlexer.Lexer, out *GetUniverseMoonsMoonIdPosition) {
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
		case "x":
			out.X = float32(in.Float32())
		case "y":
			out.Y = float32(in.Float32())
		case "z":
			out.Z = float32(in.Float32())
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
func easyjsonE06267a0EncodeGithubComAntihaxGoesiV12(out *jwriter.Writer, in GetUniverseMoonsMoonIdPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"x\":")
		out.Float32(float32(in.X))
	}
	if in.Y != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"y\":")
		out.Float32(float32(in.Y))
	}
	if in.Z != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"z\":")
		out.Float32(float32(in.Z))
	}
	out.RawByte('}')
}
