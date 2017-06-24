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

func easyjson9d322f56DecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetFleetsFleetIdWings200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFleetsFleetIdWings200OkList, 0, 1)
			} else {
				*out = GetFleetsFleetIdWings200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFleetsFleetIdWings200Ok
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
func easyjson9d322f56EncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetFleetsFleetIdWings200OkList) {
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
func (v GetFleetsFleetIdWings200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9d322f56EncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFleetsFleetIdWings200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9d322f56EncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFleetsFleetIdWings200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9d322f56DecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFleetsFleetIdWings200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9d322f56DecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjson9d322f56DecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetFleetsFleetIdWings200Ok) {
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
		case "id":
			out.Id = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "squads":
			if in.IsNull() {
				in.Skip()
				out.Squads = nil
			} else {
				in.Delim('[')
				if out.Squads == nil {
					if !in.IsDelim(']') {
						out.Squads = make([]GetFleetsFleetIdWingsSquad, 0, 2)
					} else {
						out.Squads = []GetFleetsFleetIdWingsSquad{}
					}
				} else {
					out.Squads = (out.Squads)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetFleetsFleetIdWingsSquad
					easyjson9d322f56DecodeGithubComAntihaxGoesiV12(in, &v4)
					out.Squads = append(out.Squads, v4)
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
func easyjson9d322f56EncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetFleetsFleetIdWings200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.Int64(int64(in.Id))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if len(in.Squads) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"squads\":")
		if in.Squads == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Squads {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson9d322f56EncodeGithubComAntihaxGoesiV12(out, v6)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFleetsFleetIdWings200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9d322f56EncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFleetsFleetIdWings200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9d322f56EncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFleetsFleetIdWings200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9d322f56DecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFleetsFleetIdWings200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9d322f56DecodeGithubComAntihaxGoesiV11(l, v)
}
func easyjson9d322f56DecodeGithubComAntihaxGoesiV12(in *jlexer.Lexer, out *GetFleetsFleetIdWingsSquad) {
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
		case "id":
			out.Id = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
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
func easyjson9d322f56EncodeGithubComAntihaxGoesiV12(out *jwriter.Writer, in GetFleetsFleetIdWingsSquad) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.Int64(int64(in.Id))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	out.RawByte('}')
}
