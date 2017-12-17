// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson285795f9DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationCorporationIdMiningObservers200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationCorporationIdMiningObservers200OkList, 0, 1)
			} else {
				*out = GetCorporationCorporationIdMiningObservers200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationCorporationIdMiningObservers200Ok
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
func easyjson285795f9EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationCorporationIdMiningObservers200OkList) {
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
func (v GetCorporationCorporationIdMiningObservers200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson285795f9EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationCorporationIdMiningObservers200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson285795f9EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningObservers200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson285795f9DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningObservers200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson285795f9DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson285795f9DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationCorporationIdMiningObservers200Ok) {
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
		case "last_updated":
			out.LastUpdated = string(in.String())
		case "observer_id":
			out.ObserverId = int64(in.Int64())
		case "observer_type":
			out.ObserverType = string(in.String())
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
func easyjson285795f9EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationCorporationIdMiningObservers200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastUpdated != "" {
		const prefix string = ",\"last_updated\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LastUpdated))
	}
	if in.ObserverId != 0 {
		const prefix string = ",\"observer_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ObserverId))
	}
	if in.ObserverType != "" {
		const prefix string = ",\"observer_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ObserverType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationCorporationIdMiningObservers200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson285795f9EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationCorporationIdMiningObservers200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson285795f9EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningObservers200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson285795f9DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningObservers200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson285795f9DecodeGithubComAntihaxGoesiEsi1(l, v)
}