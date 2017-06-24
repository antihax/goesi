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

func easyjson5182cc53DecodeGithubComAntihaxGoesiV1(in *jlexer.Lexer, out *GetCharactersCharacterIdIndustryJobs200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdIndustryJobs200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdIndustryJobs200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdIndustryJobs200Ok
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
func easyjson5182cc53EncodeGithubComAntihaxGoesiV1(out *jwriter.Writer, in GetCharactersCharacterIdIndustryJobs200OkList) {
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
func (v GetCharactersCharacterIdIndustryJobs200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5182cc53EncodeGithubComAntihaxGoesiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdIndustryJobs200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5182cc53EncodeGithubComAntihaxGoesiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdIndustryJobs200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5182cc53DecodeGithubComAntihaxGoesiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdIndustryJobs200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5182cc53DecodeGithubComAntihaxGoesiV1(l, v)
}
func easyjson5182cc53DecodeGithubComAntihaxGoesiV11(in *jlexer.Lexer, out *GetCharactersCharacterIdIndustryJobs200Ok) {
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
		case "activity_id":
			out.ActivityId = int32(in.Int32())
		case "blueprint_id":
			out.BlueprintId = int64(in.Int64())
		case "blueprint_location_id":
			out.BlueprintLocationId = int64(in.Int64())
		case "blueprint_type_id":
			out.BlueprintTypeId = int32(in.Int32())
		case "completed_character_id":
			out.CompletedCharacterId = int32(in.Int32())
		case "completed_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CompletedDate).UnmarshalJSON(data))
			}
		case "cost":
			out.Cost = float32(in.Float32())
		case "duration":
			out.Duration = int32(in.Int32())
		case "end_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EndDate).UnmarshalJSON(data))
			}
		case "facility_id":
			out.FacilityId = int64(in.Int64())
		case "installer_id":
			out.InstallerId = int32(in.Int32())
		case "job_id":
			out.JobId = int32(in.Int32())
		case "licensed_runs":
			out.LicensedRuns = int32(in.Int32())
		case "output_location_id":
			out.OutputLocationId = int64(in.Int64())
		case "pause_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.PauseDate).UnmarshalJSON(data))
			}
		case "probability":
			out.Probability = float32(in.Float32())
		case "product_type_id":
			out.ProductTypeId = int32(in.Int32())
		case "runs":
			out.Runs = int32(in.Int32())
		case "start_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartDate).UnmarshalJSON(data))
			}
		case "station_id":
			out.StationId = int64(in.Int64())
		case "status":
			out.Status = string(in.String())
		case "successful_runs":
			out.SuccessfulRuns = int32(in.Int32())
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
func easyjson5182cc53EncodeGithubComAntihaxGoesiV11(out *jwriter.Writer, in GetCharactersCharacterIdIndustryJobs200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ActivityId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"activity_id\":")
		out.Int32(int32(in.ActivityId))
	}
	if in.BlueprintId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"blueprint_id\":")
		out.Int64(int64(in.BlueprintId))
	}
	if in.BlueprintLocationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"blueprint_location_id\":")
		out.Int64(int64(in.BlueprintLocationId))
	}
	if in.BlueprintTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"blueprint_type_id\":")
		out.Int32(int32(in.BlueprintTypeId))
	}
	if in.CompletedCharacterId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"completed_character_id\":")
		out.Int32(int32(in.CompletedCharacterId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"completed_date\":")
		out.Raw((in.CompletedDate).MarshalJSON())
	}
	if in.Cost != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cost\":")
		out.Float32(float32(in.Cost))
	}
	if in.Duration != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"duration\":")
		out.Int32(int32(in.Duration))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"end_date\":")
		out.Raw((in.EndDate).MarshalJSON())
	}
	if in.FacilityId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"facility_id\":")
		out.Int64(int64(in.FacilityId))
	}
	if in.InstallerId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"installer_id\":")
		out.Int32(int32(in.InstallerId))
	}
	if in.JobId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"job_id\":")
		out.Int32(int32(in.JobId))
	}
	if in.LicensedRuns != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"licensed_runs\":")
		out.Int32(int32(in.LicensedRuns))
	}
	if in.OutputLocationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"output_location_id\":")
		out.Int64(int64(in.OutputLocationId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"pause_date\":")
		out.Raw((in.PauseDate).MarshalJSON())
	}
	if in.Probability != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"probability\":")
		out.Float32(float32(in.Probability))
	}
	if in.ProductTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"product_type_id\":")
		out.Int32(int32(in.ProductTypeId))
	}
	if in.Runs != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"runs\":")
		out.Int32(int32(in.Runs))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"start_date\":")
		out.Raw((in.StartDate).MarshalJSON())
	}
	if in.StationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"station_id\":")
		out.Int64(int64(in.StationId))
	}
	if in.Status != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"status\":")
		out.String(string(in.Status))
	}
	if in.SuccessfulRuns != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"successful_runs\":")
		out.Int32(int32(in.SuccessfulRuns))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdIndustryJobs200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5182cc53EncodeGithubComAntihaxGoesiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdIndustryJobs200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5182cc53EncodeGithubComAntihaxGoesiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdIndustryJobs200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5182cc53DecodeGithubComAntihaxGoesiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdIndustryJobs200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5182cc53DecodeGithubComAntihaxGoesiV11(l, v)
}
