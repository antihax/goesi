// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package goesiv3

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

func easyjson1d8817ecDecodeGithubComAntihaxGoesiV3(in *jlexer.Lexer, out *GetCharactersCharacterIdMailLabelsLabelList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdMailLabelsLabelList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdMailLabelsLabelList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdMailLabelsLabel
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
func easyjson1d8817ecEncodeGithubComAntihaxGoesiV3(out *jwriter.Writer, in GetCharactersCharacterIdMailLabelsLabelList) {
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
func (v GetCharactersCharacterIdMailLabelsLabelList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1d8817ecEncodeGithubComAntihaxGoesiV3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMailLabelsLabelList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1d8817ecEncodeGithubComAntihaxGoesiV3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLabelsLabelList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1d8817ecDecodeGithubComAntihaxGoesiV3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLabelsLabelList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1d8817ecDecodeGithubComAntihaxGoesiV3(l, v)
}
func easyjson1d8817ecDecodeGithubComAntihaxGoesiV31(in *jlexer.Lexer, out *GetCharactersCharacterIdMailLabelsLabel) {
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
		case "color":
			out.Color = string(in.String())
		case "label_id":
			out.LabelId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "unread_count":
			out.UnreadCount = int32(in.Int32())
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
func easyjson1d8817ecEncodeGithubComAntihaxGoesiV31(out *jwriter.Writer, in GetCharactersCharacterIdMailLabelsLabel) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Color != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"color\":")
		out.String(string(in.Color))
	}
	if in.LabelId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"label_id\":")
		out.Int32(int32(in.LabelId))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.UnreadCount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"unread_count\":")
		out.Int32(int32(in.UnreadCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdMailLabelsLabel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1d8817ecEncodeGithubComAntihaxGoesiV31(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMailLabelsLabel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1d8817ecEncodeGithubComAntihaxGoesiV31(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLabelsLabel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1d8817ecDecodeGithubComAntihaxGoesiV31(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLabelsLabel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1d8817ecDecodeGithubComAntihaxGoesiV31(l, v)
}
