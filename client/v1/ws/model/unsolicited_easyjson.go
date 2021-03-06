// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel(in *jlexer.Lexer, out *SystemConnection) {
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
		case "topic":
			out.Topic = TopicType(in.String())
		case "success":
			out.Success = string(in.String())
		case "hb":
			out.HB = int64(in.Int64())
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
func easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel(out *jwriter.Writer, in SystemConnection) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"topic\":"
		out.RawString(prefix[1:])
		out.String(string(in.Topic))
	}
	{
		const prefix string = ",\"success\":"
		out.RawString(prefix)
		out.String(string(in.Success))
	}
	{
		const prefix string = ",\"hb\":"
		out.RawString(prefix)
		out.Int64(int64(in.HB))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SystemConnection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SystemConnection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SystemConnection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SystemConnection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel(l, v)
}
func easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel1(in *jlexer.Lexer, out *Notifications) {
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
		case "topic":
			out.Topic = TopicType(in.String())
		case "args":
			easyjson4b879c23Decode(in, &out.Args)
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
func easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel1(out *jwriter.Writer, in Notifications) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"topic\":"
		out.RawString(prefix[1:])
		out.String(string(in.Topic))
	}
	{
		const prefix string = ",\"args\":"
		out.RawString(prefix)
		easyjson4b879c23Encode(out, in.Args)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Notifications) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Notifications) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Notifications) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Notifications) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel1(l, v)
}
func easyjson4b879c23Decode(in *jlexer.Lexer, out *struct {
	Id    string `json:"id"`
	Text  string `json:"text"`
	Title string `json:"title"`
	Url   string `json:"url"`
}) {
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
		case "id":
			out.Id = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "url":
			out.Url = string(in.String())
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
func easyjson4b879c23Encode(out *jwriter.Writer, in struct {
	Id    string `json:"id"`
	Text  string `json:"text"`
	Title string `json:"title"`
	Url   string `json:"url"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.Url))
	}
	out.RawByte('}')
}
func easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel2(in *jlexer.Lexer, out *Bulletins) {
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
		case "topic":
			out.Topic = TopicType(in.String())
		case "args":
			easyjson4b879c23Decode1(in, &out.Args)
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
func easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel2(out *jwriter.Writer, in Bulletins) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"topic\":"
		out.RawString(prefix[1:])
		out.String(string(in.Topic))
	}
	{
		const prefix string = ",\"args\":"
		out.RawString(prefix)
		easyjson4b879c23Encode1(out, in.Args)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bulletins) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bulletins) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bulletins) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bulletins) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel2(l, v)
}
func easyjson4b879c23Decode1(in *jlexer.Lexer, out *struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}) {
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
		case "id":
			out.Id = string(in.String())
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
func easyjson4b879c23Encode1(out *jwriter.Writer, in struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	out.RawByte('}')
}
func easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel3(in *jlexer.Lexer, out *AuthStatus) {
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
		case "topic":
			out.Topic = TopicType(in.String())
		case "args":
			easyjson4b879c23Decode2(in, &out.Args)
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
func easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel3(out *jwriter.Writer, in AuthStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"topic\":"
		out.RawString(prefix[1:])
		out.String(string(in.Topic))
	}
	{
		const prefix string = ",\"args\":"
		out.RawString(prefix)
		easyjson4b879c23Encode2(out, in.Args)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4b879c23EncodeGithubComKamaiuIbCpGoClientV1WsModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4b879c23DecodeGithubComKamaiuIbCpGoClientV1WsModel3(l, v)
}
func easyjson4b879c23Decode2(in *jlexer.Lexer, out *struct {
	Authenticated bool `json:"authenticated"`
	Competing     bool `json:"competing"`
}) {
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
		case "authenticated":
			out.Authenticated = bool(in.Bool())
		case "competing":
			out.Competing = bool(in.Bool())
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
func easyjson4b879c23Encode2(out *jwriter.Writer, in struct {
	Authenticated bool `json:"authenticated"`
	Competing     bool `json:"competing"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"authenticated\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Authenticated))
	}
	{
		const prefix string = ",\"competing\":"
		out.RawString(prefix)
		out.Bool(bool(in.Competing))
	}
	out.RawByte('}')
}
