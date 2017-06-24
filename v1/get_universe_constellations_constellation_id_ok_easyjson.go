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

func easyjson535e99f6DecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetUniverseConstellationsConstellationIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseConstellationsConstellationIdOkList, 0, 1)
			} else {
				*out = GetUniverseConstellationsConstellationIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseConstellationsConstellationIdOk
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
func easyjson535e99f6EncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetUniverseConstellationsConstellationIdOkList) {
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
func (v GetUniverseConstellationsConstellationIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson535e99f6EncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseConstellationsConstellationIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson535e99f6EncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseConstellationsConstellationIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson535e99f6DecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseConstellationsConstellationIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson535e99f6DecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjson535e99f6DecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetUniverseConstellationsConstellationIdOk) {
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
		case "constellation_id":
			out.ConstellationId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "position":
			easyjson535e99f6DecodeGithubComAntihaxGoesiV12(in, &out.Position)
		case "region_id":
			out.RegionId = int32(in.Int32())
		case "systems":
			if in.IsNull() {
				in.Skip()
				out.Systems = nil
			} else {
				in.Delim('[')
				if out.Systems == nil {
					if !in.IsDelim(']') {
						out.Systems = make([]int32, 0, 16)
					} else {
						out.Systems = []int32{}
					}
				} else {
					out.Systems = (out.Systems)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Systems = append(out.Systems, v4)
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
func easyjson535e99f6EncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetUniverseConstellationsConstellationIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConstellationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"constellation_id\":")
		out.Int32(int32(in.ConstellationId))
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
		easyjson535e99f6EncodeGithubComAntihaxGoesiV12(out, in.Position)
	}
	if in.RegionId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"region_id\":")
		out.Int32(int32(in.RegionId))
	}
	if len(in.Systems) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"systems\":")
		if in.Systems == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Systems {
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
func (v GetUniverseConstellationsConstellationIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson535e99f6EncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseConstellationsConstellationIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson535e99f6EncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseConstellationsConstellationIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson535e99f6DecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseConstellationsConstellationIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson535e99f6DecodeGithubComAntihaxGoesiV11(l, v)
}
func easyjson535e99f6DecodeGithubComAntihaxGoesiV12(in *jlexer.Lexer, out *GetUniverseConstellationsConstellationIdPosition) {
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
func easyjson535e99f6EncodeGithubComAntihaxGoesiV12(out *jwriter.Writer, in GetUniverseConstellationsConstellationIdPosition) {
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
