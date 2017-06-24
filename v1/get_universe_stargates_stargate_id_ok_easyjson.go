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

func easyjson3bdd99b0DecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetUniverseStargatesStargateIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseStargatesStargateIdOkList, 0, 1)
			} else {
				*out = GetUniverseStargatesStargateIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseStargatesStargateIdOk
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
func easyjson3bdd99b0EncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetUniverseStargatesStargateIdOkList) {
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
func (v GetUniverseStargatesStargateIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3bdd99b0EncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStargatesStargateIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3bdd99b0EncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStargatesStargateIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3bdd99b0DecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStargatesStargateIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3bdd99b0DecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjson3bdd99b0DecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetUniverseStargatesStargateIdOk) {
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
		case "destination":
			(out.Destination).UnmarshalEasyJSON(in)
		case "name":
			out.Name = string(in.String())
		case "position":
			easyjson3bdd99b0DecodeGithubComAntihaxGoesiV12(in, &out.Position)
		case "stargate_id":
			out.StargateId = int32(in.Int32())
		case "system_id":
			out.SystemId = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
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
func easyjson3bdd99b0EncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetUniverseStargatesStargateIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"destination\":")
		(in.Destination).MarshalEasyJSON(out)
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
		easyjson3bdd99b0EncodeGithubComAntihaxGoesiV12(out, in.Position)
	}
	if in.StargateId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"stargate_id\":")
		out.Int32(int32(in.StargateId))
	}
	if in.SystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"system_id\":")
		out.Int32(int32(in.SystemId))
	}
	if in.TypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type_id\":")
		out.Int32(int32(in.TypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseStargatesStargateIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3bdd99b0EncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStargatesStargateIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3bdd99b0EncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStargatesStargateIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3bdd99b0DecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStargatesStargateIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3bdd99b0DecodeGithubComAntihaxGoesiV11(l, v)
}
func easyjson3bdd99b0DecodeGithubComAntihaxGoesiV12(in *jlexer.Lexer, out *GetUniverseStargatesStargateIdPosition) {
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
func easyjson3bdd99b0EncodeGithubComAntihaxGoesiV12(out *jwriter.Writer, in GetUniverseStargatesStargateIdPosition) {
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
