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

func easyjson9709e0a2DecodeGithubComAntihaxGoesiV2(in *jlexer.Lexer, out *GetUniverseStationsStationIdNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseStationsStationIdNotFoundList, 0, 4)
			} else {
				*out = GetUniverseStationsStationIdNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseStationsStationIdNotFound
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
func easyjson9709e0a2EncodeGithubComAntihaxGoesiV2(out *jwriter.Writer, in GetUniverseStationsStationIdNotFoundList) {
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
func (v GetUniverseStationsStationIdNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9709e0a2EncodeGithubComAntihaxGoesiV2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStationsStationIdNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9709e0a2EncodeGithubComAntihaxGoesiV2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStationsStationIdNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9709e0a2DecodeGithubComAntihaxGoesiV2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStationsStationIdNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9709e0a2DecodeGithubComAntihaxGoesiV2(l, v)
}
func easyjson9709e0a2DecodeGithubComAntihaxGoesiV21(in *jlexer.Lexer, out *GetUniverseStationsStationIdNotFound) {
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
func easyjson9709e0a2EncodeGithubComAntihaxGoesiV21(out *jwriter.Writer, in GetUniverseStationsStationIdNotFound) {
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
func (v GetUniverseStationsStationIdNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9709e0a2EncodeGithubComAntihaxGoesiV21(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStationsStationIdNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9709e0a2EncodeGithubComAntihaxGoesiV21(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStationsStationIdNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9709e0a2DecodeGithubComAntihaxGoesiV21(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStationsStationIdNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9709e0a2DecodeGithubComAntihaxGoesiV21(l, v)
}
