// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package signalr

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

func easyjson766e99b1DecodeGithubComPhilippseithSignalr(in *jlexer.Lexer, out *jsonInvocationMessage) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = int(in.Int())
		case "target":
			out.Target = string(in.String())
		case "invocationId":
			out.InvocationID = string(in.String())
		case "arguments":
			if in.IsNull() {
				in.Skip()
				out.Arguments = nil
			} else {
				in.Delim('[')
				if out.Arguments == nil {
					if !in.IsDelim(']') {
						out.Arguments = make([]json.RawMessage, 0, 2)
					} else {
						out.Arguments = []json.RawMessage{}
					}
				} else {
					out.Arguments = (out.Arguments)[:0]
				}
				for !in.IsDelim(']') {
					var v1 json.RawMessage
					if data := in.Raw(); in.Ok() {
						in.AddError((v1).UnmarshalJSON(data))
					}
					out.Arguments = append(out.Arguments, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "streamIds":
			if in.IsNull() {
				in.Skip()
				out.StreamIds = nil
			} else {
				in.Delim('[')
				if out.StreamIds == nil {
					if !in.IsDelim(']') {
						out.StreamIds = make([]string, 0, 4)
					} else {
						out.StreamIds = []string{}
					}
				} else {
					out.StreamIds = (out.StreamIds)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.StreamIds = append(out.StreamIds, v2)
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
func easyjson766e99b1EncodeGithubComPhilippseithSignalr(out *jwriter.Writer, in jsonInvocationMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Type))
	}
	{
		const prefix string = ",\"target\":"
		out.RawString(prefix)
		out.String(string(in.Target))
	}
	{
		const prefix string = ",\"invocationId\":"
		out.RawString(prefix)
		out.String(string(in.InvocationID))
	}
	{
		const prefix string = ",\"arguments\":"
		out.RawString(prefix)
		if in.Arguments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Arguments {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.Raw((v4).MarshalJSON())
			}
			out.RawByte(']')
		}
	}
	if len(in.StreamIds) != 0 {
		const prefix string = ",\"streamIds\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.StreamIds {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v jsonInvocationMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson766e99b1EncodeGithubComPhilippseithSignalr(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v jsonInvocationMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson766e99b1EncodeGithubComPhilippseithSignalr(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *jsonInvocationMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson766e99b1DecodeGithubComPhilippseithSignalr(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *jsonInvocationMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson766e99b1DecodeGithubComPhilippseithSignalr(l, v)
}
func easyjson766e99b1DecodeGithubComPhilippseithSignalr1(in *jlexer.Lexer, out *jsonError) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjson766e99b1EncodeGithubComPhilippseithSignalr1(out *jwriter.Writer, in jsonError) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v jsonError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson766e99b1EncodeGithubComPhilippseithSignalr1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v jsonError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson766e99b1EncodeGithubComPhilippseithSignalr1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *jsonError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson766e99b1DecodeGithubComPhilippseithSignalr1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *jsonError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson766e99b1DecodeGithubComPhilippseithSignalr1(l, v)
}
func easyjson766e99b1DecodeGithubComPhilippseithSignalr2(in *jlexer.Lexer, out *JSONHubProtocol) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjson766e99b1EncodeGithubComPhilippseithSignalr2(out *jwriter.Writer, in JSONHubProtocol) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JSONHubProtocol) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson766e99b1EncodeGithubComPhilippseithSignalr2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JSONHubProtocol) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson766e99b1EncodeGithubComPhilippseithSignalr2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JSONHubProtocol) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson766e99b1DecodeGithubComPhilippseithSignalr2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JSONHubProtocol) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson766e99b1DecodeGithubComPhilippseithSignalr2(l, v)
}
