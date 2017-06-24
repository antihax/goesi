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

func easyjson2e2eb27aDecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetUniverseRegionsRegionIdNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseRegionsRegionIdNotFoundList, 0, 4)
			} else {
				*out = GetUniverseRegionsRegionIdNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseRegionsRegionIdNotFound
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
func easyjson2e2eb27aEncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetUniverseRegionsRegionIdNotFoundList) {
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
func (v GetUniverseRegionsRegionIdNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2e2eb27aEncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseRegionsRegionIdNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2e2eb27aEncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseRegionsRegionIdNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2e2eb27aDecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseRegionsRegionIdNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2e2eb27aDecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjson2e2eb27aDecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetUniverseRegionsRegionIdNotFound) {
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
func easyjson2e2eb27aEncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetUniverseRegionsRegionIdNotFound) {
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
func (v GetUniverseRegionsRegionIdNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2e2eb27aEncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseRegionsRegionIdNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2e2eb27aEncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseRegionsRegionIdNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2e2eb27aDecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseRegionsRegionIdNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2e2eb27aDecodeGithubComAntihaxGoesiV11(l, v)
}
