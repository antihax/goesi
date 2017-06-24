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

func easyjson67db0fd7DecodeGithubComAntihaxGoesiV2(in *jlexer.Lexer, out *GetCharactersCharacterIdSearchOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdSearchOkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdSearchOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdSearchOk
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
func easyjson67db0fd7EncodeGithubComAntihaxGoesiV2(out *jwriter.Writer, in GetCharactersCharacterIdSearchOkList) {
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
func (v GetCharactersCharacterIdSearchOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson67db0fd7EncodeGithubComAntihaxGoesiV2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSearchOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson67db0fd7EncodeGithubComAntihaxGoesiV2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSearchOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson67db0fd7DecodeGithubComAntihaxGoesiV2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSearchOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson67db0fd7DecodeGithubComAntihaxGoesiV2(l, v)
}
func easyjson67db0fd7DecodeGithubComAntihaxGoesiV21(in *jlexer.Lexer, out *GetCharactersCharacterIdSearchOk) {
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
		case "agent":
			if in.IsNull() {
				in.Skip()
				out.Agent = nil
			} else {
				in.Delim('[')
				if out.Agent == nil {
					if !in.IsDelim(']') {
						out.Agent = make([]int32, 0, 16)
					} else {
						out.Agent = []int32{}
					}
				} else {
					out.Agent = (out.Agent)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Agent = append(out.Agent, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "alliance":
			if in.IsNull() {
				in.Skip()
				out.Alliance = nil
			} else {
				in.Delim('[')
				if out.Alliance == nil {
					if !in.IsDelim(']') {
						out.Alliance = make([]int32, 0, 16)
					} else {
						out.Alliance = []int32{}
					}
				} else {
					out.Alliance = (out.Alliance)[:0]
				}
				for !in.IsDelim(']') {
					var v5 int32
					v5 = int32(in.Int32())
					out.Alliance = append(out.Alliance, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "character":
			if in.IsNull() {
				in.Skip()
				out.Character = nil
			} else {
				in.Delim('[')
				if out.Character == nil {
					if !in.IsDelim(']') {
						out.Character = make([]int32, 0, 16)
					} else {
						out.Character = []int32{}
					}
				} else {
					out.Character = (out.Character)[:0]
				}
				for !in.IsDelim(']') {
					var v6 int32
					v6 = int32(in.Int32())
					out.Character = append(out.Character, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "constellation":
			if in.IsNull() {
				in.Skip()
				out.Constellation = nil
			} else {
				in.Delim('[')
				if out.Constellation == nil {
					if !in.IsDelim(']') {
						out.Constellation = make([]int32, 0, 16)
					} else {
						out.Constellation = []int32{}
					}
				} else {
					out.Constellation = (out.Constellation)[:0]
				}
				for !in.IsDelim(']') {
					var v7 int32
					v7 = int32(in.Int32())
					out.Constellation = append(out.Constellation, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "corporation":
			if in.IsNull() {
				in.Skip()
				out.Corporation = nil
			} else {
				in.Delim('[')
				if out.Corporation == nil {
					if !in.IsDelim(']') {
						out.Corporation = make([]int32, 0, 16)
					} else {
						out.Corporation = []int32{}
					}
				} else {
					out.Corporation = (out.Corporation)[:0]
				}
				for !in.IsDelim(']') {
					var v8 int32
					v8 = int32(in.Int32())
					out.Corporation = append(out.Corporation, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "faction":
			if in.IsNull() {
				in.Skip()
				out.Faction = nil
			} else {
				in.Delim('[')
				if out.Faction == nil {
					if !in.IsDelim(']') {
						out.Faction = make([]int32, 0, 16)
					} else {
						out.Faction = []int32{}
					}
				} else {
					out.Faction = (out.Faction)[:0]
				}
				for !in.IsDelim(']') {
					var v9 int32
					v9 = int32(in.Int32())
					out.Faction = append(out.Faction, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "inventorytype":
			if in.IsNull() {
				in.Skip()
				out.Inventorytype = nil
			} else {
				in.Delim('[')
				if out.Inventorytype == nil {
					if !in.IsDelim(']') {
						out.Inventorytype = make([]int32, 0, 16)
					} else {
						out.Inventorytype = []int32{}
					}
				} else {
					out.Inventorytype = (out.Inventorytype)[:0]
				}
				for !in.IsDelim(']') {
					var v10 int32
					v10 = int32(in.Int32())
					out.Inventorytype = append(out.Inventorytype, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "region":
			if in.IsNull() {
				in.Skip()
				out.Region = nil
			} else {
				in.Delim('[')
				if out.Region == nil {
					if !in.IsDelim(']') {
						out.Region = make([]int32, 0, 16)
					} else {
						out.Region = []int32{}
					}
				} else {
					out.Region = (out.Region)[:0]
				}
				for !in.IsDelim(']') {
					var v11 int32
					v11 = int32(in.Int32())
					out.Region = append(out.Region, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "solarsystem":
			if in.IsNull() {
				in.Skip()
				out.Solarsystem = nil
			} else {
				in.Delim('[')
				if out.Solarsystem == nil {
					if !in.IsDelim(']') {
						out.Solarsystem = make([]int32, 0, 16)
					} else {
						out.Solarsystem = []int32{}
					}
				} else {
					out.Solarsystem = (out.Solarsystem)[:0]
				}
				for !in.IsDelim(']') {
					var v12 int32
					v12 = int32(in.Int32())
					out.Solarsystem = append(out.Solarsystem, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "station":
			if in.IsNull() {
				in.Skip()
				out.Station = nil
			} else {
				in.Delim('[')
				if out.Station == nil {
					if !in.IsDelim(']') {
						out.Station = make([]int32, 0, 16)
					} else {
						out.Station = []int32{}
					}
				} else {
					out.Station = (out.Station)[:0]
				}
				for !in.IsDelim(']') {
					var v13 int32
					v13 = int32(in.Int32())
					out.Station = append(out.Station, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "structure":
			if in.IsNull() {
				in.Skip()
				out.Structure = nil
			} else {
				in.Delim('[')
				if out.Structure == nil {
					if !in.IsDelim(']') {
						out.Structure = make([]int64, 0, 8)
					} else {
						out.Structure = []int64{}
					}
				} else {
					out.Structure = (out.Structure)[:0]
				}
				for !in.IsDelim(']') {
					var v14 int64
					v14 = int64(in.Int64())
					out.Structure = append(out.Structure, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "wormhole":
			if in.IsNull() {
				in.Skip()
				out.Wormhole = nil
			} else {
				in.Delim('[')
				if out.Wormhole == nil {
					if !in.IsDelim(']') {
						out.Wormhole = make([]int32, 0, 16)
					} else {
						out.Wormhole = []int32{}
					}
				} else {
					out.Wormhole = (out.Wormhole)[:0]
				}
				for !in.IsDelim(']') {
					var v15 int32
					v15 = int32(in.Int32())
					out.Wormhole = append(out.Wormhole, v15)
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
func easyjson67db0fd7EncodeGithubComAntihaxGoesiV21(out *jwriter.Writer, in GetCharactersCharacterIdSearchOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Agent) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"agent\":")
		if in.Agent == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.Agent {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v17))
			}
			out.RawByte(']')
		}
	}
	if len(in.Alliance) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance\":")
		if in.Alliance == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.Alliance {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v19))
			}
			out.RawByte(']')
		}
	}
	if len(in.Character) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"character\":")
		if in.Character == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.Character {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v21))
			}
			out.RawByte(']')
		}
	}
	if len(in.Constellation) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"constellation\":")
		if in.Constellation == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v22, v23 := range in.Constellation {
				if v22 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v23))
			}
			out.RawByte(']')
		}
	}
	if len(in.Corporation) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"corporation\":")
		if in.Corporation == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v24, v25 := range in.Corporation {
				if v24 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v25))
			}
			out.RawByte(']')
		}
	}
	if len(in.Faction) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"faction\":")
		if in.Faction == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.Faction {
				if v26 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v27))
			}
			out.RawByte(']')
		}
	}
	if len(in.Inventorytype) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"inventorytype\":")
		if in.Inventorytype == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v28, v29 := range in.Inventorytype {
				if v28 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v29))
			}
			out.RawByte(']')
		}
	}
	if len(in.Region) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"region\":")
		if in.Region == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v30, v31 := range in.Region {
				if v30 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v31))
			}
			out.RawByte(']')
		}
	}
	if len(in.Solarsystem) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"solarsystem\":")
		if in.Solarsystem == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v32, v33 := range in.Solarsystem {
				if v32 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v33))
			}
			out.RawByte(']')
		}
	}
	if len(in.Station) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"station\":")
		if in.Station == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v34, v35 := range in.Station {
				if v34 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v35))
			}
			out.RawByte(']')
		}
	}
	if len(in.Structure) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"structure\":")
		if in.Structure == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v36, v37 := range in.Structure {
				if v36 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v37))
			}
			out.RawByte(']')
		}
	}
	if len(in.Wormhole) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"wormhole\":")
		if in.Wormhole == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v38, v39 := range in.Wormhole {
				if v38 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v39))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdSearchOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson67db0fd7EncodeGithubComAntihaxGoesiV21(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSearchOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson67db0fd7EncodeGithubComAntihaxGoesiV21(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSearchOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson67db0fd7DecodeGithubComAntihaxGoesiV21(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSearchOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson67db0fd7DecodeGithubComAntihaxGoesiV21(l, v)
}
