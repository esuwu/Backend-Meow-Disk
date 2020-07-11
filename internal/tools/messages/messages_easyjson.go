// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package messages

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

func easyjson66c1e240DecodeMainInternalToolsMessages(in *jlexer.Lexer, out *MessageString) {
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
		case "message":
			out.Message = string(in.String())
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
func easyjson66c1e240EncodeMainInternalToolsMessages(out *jwriter.Writer, in MessageString) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessageString) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeMainInternalToolsMessages(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessageString) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeMainInternalToolsMessages(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessageString) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeMainInternalToolsMessages(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessageString) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeMainInternalToolsMessages(l, v)
}