// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package goesiv2

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

func easyjsonC64e521cDecodeGithubComAntihaxGoesiV2(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdFactoryDetailsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanetsPlanetIdFactoryDetailsList, 0, 16)
			} else {
				*out = GetCharactersCharacterIdPlanetsPlanetIdFactoryDetailsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails
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
func easyjsonC64e521cEncodeGithubComAntihaxGoesiV2(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdFactoryDetailsList) {
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
func (v GetCharactersCharacterIdPlanetsPlanetIdFactoryDetailsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC64e521cEncodeGithubComAntihaxGoesiV2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdFactoryDetailsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC64e521cEncodeGithubComAntihaxGoesiV2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdFactoryDetailsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC64e521cDecodeGithubComAntihaxGoesiV2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdFactoryDetailsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC64e521cDecodeGithubComAntihaxGoesiV2(l, v)
}
func easyjsonC64e521cDecodeGithubComAntihaxGoesiV21(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) {
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
		case "schematic_id":
			out.SchematicId = int32(in.Int32())
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
func easyjsonC64e521cEncodeGithubComAntihaxGoesiV21(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SchematicId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"schematic_id\":")
		out.Int32(int32(in.SchematicId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC64e521cEncodeGithubComAntihaxGoesiV21(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC64e521cEncodeGithubComAntihaxGoesiV21(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC64e521cDecodeGithubComAntihaxGoesiV21(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC64e521cDecodeGithubComAntihaxGoesiV21(l, v)
}
