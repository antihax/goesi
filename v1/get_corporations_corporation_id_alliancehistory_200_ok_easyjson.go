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

func easyjsonE012c571DecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetCorporationsCorporationIdAlliancehistory200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdAlliancehistory200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdAlliancehistory200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdAlliancehistory200Ok
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
func easyjsonE012c571EncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetCorporationsCorporationIdAlliancehistory200OkList) {
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
func (v GetCorporationsCorporationIdAlliancehistory200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE012c571EncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdAlliancehistory200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE012c571EncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdAlliancehistory200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE012c571DecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdAlliancehistory200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE012c571DecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjsonE012c571DecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetCorporationsCorporationIdAlliancehistory200Ok) {
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
		case "alliance":
			easyjsonE012c571DecodeGithubComAntihaxGoesiV12(in, &out.Alliance)
		case "record_id":
			out.RecordId = int32(in.Int32())
		case "start_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartDate).UnmarshalJSON(data))
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
func easyjsonE012c571EncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetCorporationsCorporationIdAlliancehistory200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance\":")
		easyjsonE012c571EncodeGithubComAntihaxGoesiV12(out, in.Alliance)
	}
	if in.RecordId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"record_id\":")
		out.Int32(int32(in.RecordId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"start_date\":")
		out.Raw((in.StartDate).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdAlliancehistory200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE012c571EncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdAlliancehistory200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE012c571EncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdAlliancehistory200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE012c571DecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdAlliancehistory200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE012c571DecodeGithubComAntihaxGoesiV11(l, v)
}
func easyjsonE012c571DecodeGithubComAntihaxGoesiV12(in *jlexer.Lexer, out *GetCorporationsCorporationIdAlliancehistoryAlliance) {
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
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "is_deleted":
			out.IsDeleted = bool(in.Bool())
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
func easyjsonE012c571EncodeGithubComAntihaxGoesiV12(out *jwriter.Writer, in GetCorporationsCorporationIdAlliancehistoryAlliance) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllianceId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance_id\":")
		out.Int32(int32(in.AllianceId))
	}
	if in.IsDeleted {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_deleted\":")
		out.Bool(bool(in.IsDeleted))
	}
	out.RawByte('}')
}
