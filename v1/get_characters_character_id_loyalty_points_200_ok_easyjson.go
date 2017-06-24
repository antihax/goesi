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

func easyjson35c79e8eDecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetCharactersCharacterIdLoyaltyPoints200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdLoyaltyPoints200OkList, 0, 8)
			} else {
				*out = GetCharactersCharacterIdLoyaltyPoints200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdLoyaltyPoints200Ok
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
func easyjson35c79e8eEncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetCharactersCharacterIdLoyaltyPoints200OkList) {
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
func (v GetCharactersCharacterIdLoyaltyPoints200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson35c79e8eEncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdLoyaltyPoints200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson35c79e8eEncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdLoyaltyPoints200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson35c79e8eDecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdLoyaltyPoints200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson35c79e8eDecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjson35c79e8eDecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetCharactersCharacterIdLoyaltyPoints200Ok) {
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
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "loyalty_points":
			out.LoyaltyPoints = int32(in.Int32())
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
func easyjson35c79e8eEncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetCharactersCharacterIdLoyaltyPoints200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CorporationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"corporation_id\":")
		out.Int32(int32(in.CorporationId))
	}
	if in.LoyaltyPoints != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"loyalty_points\":")
		out.Int32(int32(in.LoyaltyPoints))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdLoyaltyPoints200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson35c79e8eEncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdLoyaltyPoints200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson35c79e8eEncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdLoyaltyPoints200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson35c79e8eDecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdLoyaltyPoints200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson35c79e8eDecodeGithubComAntihaxGoesiV11(l, v)
}
