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

func easyjsonE838f8a7DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetContractsPublicRegionId200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetContractsPublicRegionId200OkList, 0, 1)
			} else {
				*out = GetContractsPublicRegionId200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetContractsPublicRegionId200Ok
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
func easyjsonE838f8a7EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetContractsPublicRegionId200OkList) {
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
func (v GetContractsPublicRegionId200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE838f8a7EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetContractsPublicRegionId200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE838f8a7EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetContractsPublicRegionId200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE838f8a7DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetContractsPublicRegionId200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE838f8a7DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonE838f8a7DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetContractsPublicRegionId200Ok) {
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
		case "buyout":
			out.Buyout = float64(in.Float64())
		case "collateral":
			out.Collateral = float64(in.Float64())
		case "contract_id":
			out.ContractId = int32(in.Int32())
		case "date_expired":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateExpired).UnmarshalJSON(data))
			}
		case "date_issued":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateIssued).UnmarshalJSON(data))
			}
		case "days_to_complete":
			out.DaysToComplete = int32(in.Int32())
		case "end_location_id":
			out.EndLocationId = int64(in.Int64())
		case "for_corporation":
			out.ForCorporation = bool(in.Bool())
		case "issuer_corporation_id":
			out.IssuerCorporationId = int32(in.Int32())
		case "issuer_id":
			out.IssuerId = int32(in.Int32())
		case "price":
			out.Price = float64(in.Float64())
		case "reward":
			out.Reward = float64(in.Float64())
		case "start_location_id":
			out.StartLocationId = int64(in.Int64())
		case "title":
			out.Title = string(in.String())
		case "type":
			out.Type_ = string(in.String())
		case "volume":
			out.Volume = float64(in.Float64())
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
func easyjsonE838f8a7EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetContractsPublicRegionId200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Buyout != 0 {
		const prefix string = ",\"buyout\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Buyout))
	}
	if in.Collateral != 0 {
		const prefix string = ",\"collateral\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Collateral))
	}
	if in.ContractId != 0 {
		const prefix string = ",\"contract_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ContractId))
	}
	if true {
		const prefix string = ",\"date_expired\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.DateExpired).MarshalJSON())
	}
	if true {
		const prefix string = ",\"date_issued\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.DateIssued).MarshalJSON())
	}
	if in.DaysToComplete != 0 {
		const prefix string = ",\"days_to_complete\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.DaysToComplete))
	}
	if in.EndLocationId != 0 {
		const prefix string = ",\"end_location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.EndLocationId))
	}
	if in.ForCorporation {
		const prefix string = ",\"for_corporation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ForCorporation))
	}
	if in.IssuerCorporationId != 0 {
		const prefix string = ",\"issuer_corporation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.IssuerCorporationId))
	}
	if in.IssuerId != 0 {
		const prefix string = ",\"issuer_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.IssuerId))
	}
	if in.Price != 0 {
		const prefix string = ",\"price\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Price))
	}
	if in.Reward != 0 {
		const prefix string = ",\"reward\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Reward))
	}
	if in.StartLocationId != 0 {
		const prefix string = ",\"start_location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.StartLocationId))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Type_ != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type_))
	}
	if in.Volume != 0 {
		const prefix string = ",\"volume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Volume))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetContractsPublicRegionId200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE838f8a7EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetContractsPublicRegionId200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE838f8a7EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetContractsPublicRegionId200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE838f8a7DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetContractsPublicRegionId200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE838f8a7DecodeGithubComAntihaxGoesiEsi1(l, v)
}