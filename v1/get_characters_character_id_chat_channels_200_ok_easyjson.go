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

func easyjson15e3e24fDecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannels200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdChatChannels200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdChatChannels200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdChatChannels200Ok
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
func easyjson15e3e24fEncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetCharactersCharacterIdChatChannels200OkList) {
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
func (v GetCharactersCharacterIdChatChannels200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15e3e24fEncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdChatChannels200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15e3e24fEncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannels200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15e3e24fDecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannels200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15e3e24fDecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjson15e3e24fDecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannels200Ok) {
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
		case "allowed":
			if in.IsNull() {
				in.Skip()
				out.Allowed = nil
			} else {
				in.Delim('[')
				if out.Allowed == nil {
					if !in.IsDelim(']') {
						out.Allowed = make([]GetCharactersCharacterIdChatChannelsAllowed, 0, 2)
					} else {
						out.Allowed = []GetCharactersCharacterIdChatChannelsAllowed{}
					}
				} else {
					out.Allowed = (out.Allowed)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdChatChannelsAllowed
					easyjson15e3e24fDecodeGithubComAntihaxGoesiV12(in, &v4)
					out.Allowed = append(out.Allowed, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "blocked":
			if in.IsNull() {
				in.Skip()
				out.Blocked = nil
			} else {
				in.Delim('[')
				if out.Blocked == nil {
					if !in.IsDelim(']') {
						out.Blocked = make([]GetCharactersCharacterIdChatChannelsBlocked, 0, 1)
					} else {
						out.Blocked = []GetCharactersCharacterIdChatChannelsBlocked{}
					}
				} else {
					out.Blocked = (out.Blocked)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetCharactersCharacterIdChatChannelsBlocked
					easyjson15e3e24fDecodeGithubComAntihaxGoesiV13(in, &v5)
					out.Blocked = append(out.Blocked, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "channel_id":
			out.ChannelId = int32(in.Int32())
		case "comparison_key":
			out.ComparisonKey = string(in.String())
		case "has_password":
			out.HasPassword = bool(in.Bool())
		case "motd":
			out.Motd = string(in.String())
		case "muted":
			if in.IsNull() {
				in.Skip()
				out.Muted = nil
			} else {
				in.Delim('[')
				if out.Muted == nil {
					if !in.IsDelim(']') {
						out.Muted = make([]GetCharactersCharacterIdChatChannelsMuted, 0, 1)
					} else {
						out.Muted = []GetCharactersCharacterIdChatChannelsMuted{}
					}
				} else {
					out.Muted = (out.Muted)[:0]
				}
				for !in.IsDelim(']') {
					var v6 GetCharactersCharacterIdChatChannelsMuted
					easyjson15e3e24fDecodeGithubComAntihaxGoesiV14(in, &v6)
					out.Muted = append(out.Muted, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "operators":
			if in.IsNull() {
				in.Skip()
				out.Operators = nil
			} else {
				in.Delim('[')
				if out.Operators == nil {
					if !in.IsDelim(']') {
						out.Operators = make([]GetCharactersCharacterIdChatChannelsOperator, 0, 2)
					} else {
						out.Operators = []GetCharactersCharacterIdChatChannelsOperator{}
					}
				} else {
					out.Operators = (out.Operators)[:0]
				}
				for !in.IsDelim(']') {
					var v7 GetCharactersCharacterIdChatChannelsOperator
					easyjson15e3e24fDecodeGithubComAntihaxGoesiV15(in, &v7)
					out.Operators = append(out.Operators, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "owner_id":
			out.OwnerId = int32(in.Int32())
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
func easyjson15e3e24fEncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetCharactersCharacterIdChatChannels200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Allowed) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"allowed\":")
		if in.Allowed == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Allowed {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjson15e3e24fEncodeGithubComAntihaxGoesiV12(out, v9)
			}
			out.RawByte(']')
		}
	}
	if len(in.Blocked) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"blocked\":")
		if in.Blocked == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Blocked {
				if v10 > 0 {
					out.RawByte(',')
				}
				easyjson15e3e24fEncodeGithubComAntihaxGoesiV13(out, v11)
			}
			out.RawByte(']')
		}
	}
	if in.ChannelId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"channel_id\":")
		out.Int32(int32(in.ChannelId))
	}
	if in.ComparisonKey != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"comparison_key\":")
		out.String(string(in.ComparisonKey))
	}
	if in.HasPassword {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"has_password\":")
		out.Bool(bool(in.HasPassword))
	}
	if in.Motd != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"motd\":")
		out.String(string(in.Motd))
	}
	if len(in.Muted) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"muted\":")
		if in.Muted == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Muted {
				if v12 > 0 {
					out.RawByte(',')
				}
				easyjson15e3e24fEncodeGithubComAntihaxGoesiV14(out, v13)
			}
			out.RawByte(']')
		}
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if len(in.Operators) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"operators\":")
		if in.Operators == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Operators {
				if v14 > 0 {
					out.RawByte(',')
				}
				easyjson15e3e24fEncodeGithubComAntihaxGoesiV15(out, v15)
			}
			out.RawByte(']')
		}
	}
	if in.OwnerId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"owner_id\":")
		out.Int32(int32(in.OwnerId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdChatChannels200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15e3e24fEncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdChatChannels200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15e3e24fEncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannels200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15e3e24fDecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannels200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15e3e24fDecodeGithubComAntihaxGoesiV11(l, v)
}
func easyjson15e3e24fDecodeGithubComAntihaxGoesiV15(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannelsOperator) {
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
		case "accessor_id":
			out.AccessorId = int32(in.Int32())
		case "accessor_type":
			out.AccessorType = string(in.String())
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
func easyjson15e3e24fEncodeGithubComAntihaxGoesiV15(out *jwriter.Writer, in GetCharactersCharacterIdChatChannelsOperator) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AccessorId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_id\":")
		out.Int32(int32(in.AccessorId))
	}
	if in.AccessorType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_type\":")
		out.String(string(in.AccessorType))
	}
	out.RawByte('}')
}
func easyjson15e3e24fDecodeGithubComAntihaxGoesiV14(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannelsMuted) {
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
		case "accessor_id":
			out.AccessorId = int32(in.Int32())
		case "accessor_type":
			out.AccessorType = string(in.String())
		case "end_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EndAt).UnmarshalJSON(data))
			}
		case "reason":
			out.Reason = string(in.String())
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
func easyjson15e3e24fEncodeGithubComAntihaxGoesiV14(out *jwriter.Writer, in GetCharactersCharacterIdChatChannelsMuted) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AccessorId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_id\":")
		out.Int32(int32(in.AccessorId))
	}
	if in.AccessorType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_type\":")
		out.String(string(in.AccessorType))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"end_at\":")
		out.Raw((in.EndAt).MarshalJSON())
	}
	if in.Reason != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"reason\":")
		out.String(string(in.Reason))
	}
	out.RawByte('}')
}
func easyjson15e3e24fDecodeGithubComAntihaxGoesiV13(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannelsBlocked) {
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
		case "accessor_id":
			out.AccessorId = int32(in.Int32())
		case "accessor_type":
			out.AccessorType = string(in.String())
		case "end_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EndAt).UnmarshalJSON(data))
			}
		case "reason":
			out.Reason = string(in.String())
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
func easyjson15e3e24fEncodeGithubComAntihaxGoesiV13(out *jwriter.Writer, in GetCharactersCharacterIdChatChannelsBlocked) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AccessorId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_id\":")
		out.Int32(int32(in.AccessorId))
	}
	if in.AccessorType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_type\":")
		out.String(string(in.AccessorType))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"end_at\":")
		out.Raw((in.EndAt).MarshalJSON())
	}
	if in.Reason != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"reason\":")
		out.String(string(in.Reason))
	}
	out.RawByte('}')
}
func easyjson15e3e24fDecodeGithubComAntihaxGoesiV12(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannelsAllowed) {
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
		case "accessor_id":
			out.AccessorId = int32(in.Int32())
		case "accessor_type":
			out.AccessorType = string(in.String())
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
func easyjson15e3e24fEncodeGithubComAntihaxGoesiV12(out *jwriter.Writer, in GetCharactersCharacterIdChatChannelsAllowed) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AccessorId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_id\":")
		out.Int32(int32(in.AccessorId))
	}
	if in.AccessorType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_type\":")
		out.String(string(in.AccessorType))
	}
	out.RawByte('}')
}
