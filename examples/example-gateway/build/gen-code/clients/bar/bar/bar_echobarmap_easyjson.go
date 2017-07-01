// Code generated by zanzibar
// @generated
// Checksum : QMUNhiOKElofMVqRwQyE6A==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
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

func easyjsonE298db6DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap(in *jlexer.Lexer, out *Bar_EchoBarMap_Result) {
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
		case "success":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Success = make(map[UUID]*BarResponse)
				} else {
					out.Success = nil
				}
				for !in.IsDelim('}') {
					key := UUID(in.String())
					in.WantColon()
					var v1 *BarResponse
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(BarResponse)
						}
						easyjsonE298db6DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in, &*v1)
					}
					(out.Success)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
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
func easyjsonE298db6EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap(out *jwriter.Writer, in Bar_EchoBarMap_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"success\":")
	if in.Success == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in.Success {
			if !v2First {
				out.RawByte(',')
			}
			v2First = false
			out.String(string(v2Name))
			out.RawByte(':')
			if v2Value == nil {
				out.RawString("null")
			} else {
				easyjsonE298db6EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out, *v2Value)
			}
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_EchoBarMap_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE298db6EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_EchoBarMap_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE298db6EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_EchoBarMap_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE298db6DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_EchoBarMap_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE298db6DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap(l, v)
}
func easyjsonE298db6DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in *jlexer.Lexer, out *BarResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var IntWithRangeSet bool
	var IntWithoutRangeSet bool
	var MapIntWithRangeSet bool
	var MapIntWithoutRangeSet bool
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
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "intWithRange":
			out.IntWithRange = int32(in.Int32())
			IntWithRangeSet = true
		case "intWithoutRange":
			out.IntWithoutRange = int32(in.Int32())
			IntWithoutRangeSet = true
		case "mapIntWithRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithRange = make(map[string]int32)
				} else {
					out.MapIntWithRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 int32
					v3 = int32(in.Int32())
					(out.MapIntWithRange)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithRangeSet = true
		case "mapIntWithoutRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithoutRange = make(map[string]int32)
				} else {
					out.MapIntWithoutRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 int32
					v4 = int32(in.Int32())
					(out.MapIntWithoutRange)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithoutRangeSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !IntWithRangeSet {
		in.AddError(fmt.Errorf("key 'intWithRange' is required"))
	}
	if !IntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'intWithoutRange' is required"))
	}
	if !MapIntWithRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithRange' is required"))
	}
	if !MapIntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithoutRange' is required"))
	}
}
func easyjsonE298db6EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out *jwriter.Writer, in BarResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithRange\":")
	out.Int32(int32(in.IntWithRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithoutRange\":")
	out.Int32(int32(in.IntWithoutRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithRange\":")
	if in.MapIntWithRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v5First := true
		for v5Name, v5Value := range in.MapIntWithRange {
			if !v5First {
				out.RawByte(',')
			}
			v5First = false
			out.String(string(v5Name))
			out.RawByte(':')
			out.Int32(int32(v5Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithoutRange\":")
	if in.MapIntWithoutRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v6First := true
		for v6Name, v6Value := range in.MapIntWithoutRange {
			if !v6First {
				out.RawByte(',')
			}
			v6First = false
			out.String(string(v6Name))
			out.RawByte(':')
			out.Int32(int32(v6Value))
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}
func easyjsonE298db6DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap1(in *jlexer.Lexer, out *Bar_EchoBarMap_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var ArgSet bool
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
		case "arg":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Arg = make(map[string]UUID)
				} else {
					out.Arg = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 UUID
					v7 = UUID(in.String())
					(out.Arg)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
			ArgSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !ArgSet {
		in.AddError(fmt.Errorf("key 'arg' is required"))
	}
}
func easyjsonE298db6EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap1(out *jwriter.Writer, in Bar_EchoBarMap_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"arg\":")
	if in.Arg == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v8First := true
		for v8Name, v8Value := range in.Arg {
			if !v8First {
				out.RawByte(',')
			}
			v8First = false
			out.String(string(v8Name))
			out.RawByte(':')
			out.String(string(v8Value))
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_EchoBarMap_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE298db6EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_EchoBarMap_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE298db6EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_EchoBarMap_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE298db6DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_EchoBarMap_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE298db6DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarEchoBarMap1(l, v)
}
