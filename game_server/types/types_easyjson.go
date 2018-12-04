// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

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

func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes(in *jlexer.Lexer, out *ServerResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StatusSet bool
	var MessageSet bool
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
		case "status":
			out.Status = string(in.String())
			StatusSet = true
		case "message":
			out.Message = string(in.String())
			MessageSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StatusSet {
		in.AddError(fmt.Errorf("key 'status' is required"))
	}
	if !MessageSet {
		in.AddError(fmt.Errorf("key 'message' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes(out *jwriter.Writer, in ServerResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServerResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServerResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServerResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServerResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes1(in *jlexer.Lexer, out *GameOver) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var WinnerSet bool
	var FromSet bool
	var ToSet bool
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
		case "winner":
			out.Winner = bool(in.Bool())
			WinnerSet = true
		case "from":
			out.From = int(in.Int())
			FromSet = true
		case "to":
			out.To = int(in.Int())
			ToSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !WinnerSet {
		in.AddError(fmt.Errorf("key 'winner' is required"))
	}
	if !FromSet {
		in.AddError(fmt.Errorf("key 'from' is required"))
	}
	if !ToSet {
		in.AddError(fmt.Errorf("key 'to' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes1(out *jwriter.Writer, in GameOver) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"winner\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Winner))
	}
	{
		const prefix string = ",\"from\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.From))
	}
	{
		const prefix string = ",\"to\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.To))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GameOver) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GameOver) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GameOver) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GameOver) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes1(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes2(in *jlexer.Lexer, out *WeaponChangeRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var CharacterPositionSet bool
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
		case "character_position":
			out.CharacterPosition = int(in.Int())
			CharacterPositionSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !CharacterPositionSet {
		in.AddError(fmt.Errorf("key 'character_position' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes2(out *jwriter.Writer, in WeaponChangeRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"character_position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.CharacterPosition))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WeaponChangeRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WeaponChangeRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WeaponChangeRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WeaponChangeRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes2(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes3(in *jlexer.Lexer, out *AddWeapon) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var CoordinatesSet bool
	var WeaponSet bool
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
		case "coordinates":
			out.Coordinates = int(in.Int())
			CoordinatesSet = true
		case "weapon":
			out.Weapon = string(in.String())
			WeaponSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !CoordinatesSet {
		in.AddError(fmt.Errorf("key 'coordinates' is required"))
	}
	if !WeaponSet {
		in.AddError(fmt.Errorf("key 'weapon' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes3(out *jwriter.Writer, in AddWeapon) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"coordinates\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Coordinates))
	}
	{
		const prefix string = ",\"weapon\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Weapon))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddWeapon) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddWeapon) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddWeapon) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddWeapon) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes3(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes4(in *jlexer.Lexer, out *Attack) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var WinnerSet bool
	var LoserSet bool
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
		case "winner":
			(out.Winner).UnmarshalEasyJSON(in)
			WinnerSet = true
		case "loser":
			(out.Loser).UnmarshalEasyJSON(in)
			LoserSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !WinnerSet {
		in.AddError(fmt.Errorf("key 'winner' is required"))
	}
	if !LoserSet {
		in.AddError(fmt.Errorf("key 'loser' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes4(out *jwriter.Writer, in Attack) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"winner\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Winner).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"loser\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Loser).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Attack) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Attack) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Attack) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Attack) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes4(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes5(in *jlexer.Lexer, out *AttackingСharacter) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var CoordinatesSet bool
	var WeaponSet bool
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
		case "coordinates":
			out.Coordinates = int(in.Int())
			CoordinatesSet = true
		case "weapon":
			out.Weapon = string(in.String())
			WeaponSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !CoordinatesSet {
		in.AddError(fmt.Errorf("key 'coordinates' is required"))
	}
	if !WeaponSet {
		in.AddError(fmt.Errorf("key 'weapon' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes5(out *jwriter.Writer, in AttackingСharacter) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"coordinates\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Coordinates))
	}
	{
		const prefix string = ",\"weapon\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Weapon))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AttackingСharacter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AttackingСharacter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AttackingСharacter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AttackingСharacter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes5(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes6(in *jlexer.Lexer, out *MoveCharacter) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var FromSet bool
	var ToSet bool
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
		case "from":
			out.From = int(in.Int())
			FromSet = true
		case "to":
			out.To = int(in.Int())
			ToSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !FromSet {
		in.AddError(fmt.Errorf("key 'from' is required"))
	}
	if !ToSet {
		in.AddError(fmt.Errorf("key 'to' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes6(out *jwriter.Writer, in MoveCharacter) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"from\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.From))
	}
	{
		const prefix string = ",\"to\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.To))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MoveCharacter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MoveCharacter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MoveCharacter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MoveCharacter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes6(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes7(in *jlexer.Lexer, out *DownloadMap) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('[')
		v1 := 0
		for !in.IsDelim(']') {
			if v1 < 42 {
				if in.IsNull() {
					in.Skip()
					(*out)[v1] = nil
				} else {
					if (*out)[v1] == nil {
						(*out)[v1] = new(MapCell)
					}
					(*(*out)[v1]).UnmarshalEasyJSON(in)
				}
				v1++
			} else {
				in.SkipRecursive()
			}
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes7(out *jwriter.Writer, in DownloadMap) {
	out.RawByte('[')
	for v2 := range in {
		if v2 > 0 {
			out.RawByte(',')
		}
		if (in)[v2] == nil {
			out.RawString("null")
		} else {
			(*(in)[v2]).MarshalEasyJSON(out)
		}
	}
	out.RawByte(']')
}

// MarshalJSON supports json.Marshaler interface
func (v DownloadMap) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DownloadMap) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DownloadMap) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DownloadMap) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes7(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes8(in *jlexer.Lexer, out *MapCell) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var UserSet bool
	var WeaponSet bool
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
		case "user":
			out.User = bool(in.Bool())
			UserSet = true
		case "weapon":
			if in.IsNull() {
				in.Skip()
				out.Weapon = nil
			} else {
				if out.Weapon == nil {
					out.Weapon = new(string)
				}
				*out.Weapon = string(in.String())
			}
			WeaponSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !UserSet {
		in.AddError(fmt.Errorf("key 'user' is required"))
	}
	if !WeaponSet {
		in.AddError(fmt.Errorf("key 'weapon' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes8(out *jwriter.Writer, in MapCell) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.User))
	}
	{
		const prefix string = ",\"weapon\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Weapon == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Weapon))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MapCell) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MapCell) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MapCell) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MapCell) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes8(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes9(in *jlexer.Lexer, out *ReassignWeapons) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var NewWeaponSet bool
	var CharacterPositionSet bool
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
		case "new_weapon":
			out.NewWeapon = string(in.String())
			NewWeaponSet = true
		case "character_position":
			out.CharacterPosition = int(in.Int())
			CharacterPositionSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !NewWeaponSet {
		in.AddError(fmt.Errorf("key 'new_weapon' is required"))
	}
	if !CharacterPositionSet {
		in.AddError(fmt.Errorf("key 'character_position' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes9(out *jwriter.Writer, in ReassignWeapons) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"new_weapon\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.NewWeapon))
	}
	{
		const prefix string = ",\"character_position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.CharacterPosition))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReassignWeapons) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReassignWeapons) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReassignWeapons) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReassignWeapons) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes9(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes10(in *jlexer.Lexer, out *AttemptGoToCell) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var FromSet bool
	var ToSet bool
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
		case "from":
			out.From = int(in.Int())
			FromSet = true
		case "to":
			out.To = int(in.Int())
			ToSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !FromSet {
		in.AddError(fmt.Errorf("key 'from' is required"))
	}
	if !ToSet {
		in.AddError(fmt.Errorf("key 'to' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes10(out *jwriter.Writer, in AttemptGoToCell) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"from\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.From))
	}
	{
		const prefix string = ",\"to\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.To))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AttemptGoToCell) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AttemptGoToCell) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AttemptGoToCell) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AttemptGoToCell) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes10(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes11(in *jlexer.Lexer, out *UploadMap) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var WeaponsSet bool
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
		case "weapons":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v3 := 0
				for !in.IsDelim(']') {
					if v3 < 14 {
						(out.Weapons)[v3] = string(in.String())
						v3++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
			WeaponsSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !WeaponsSet {
		in.AddError(fmt.Errorf("key 'weapons' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes11(out *jwriter.Writer, in UploadMap) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"weapons\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.RawByte('[')
		for v4 := range in.Weapons {
			if v4 > 0 {
				out.RawByte(',')
			}
			out.String(string((in.Weapons)[v4]))
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UploadMap) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UploadMap) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UploadMap) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UploadMap) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes11(l, v)
}
func easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes12(in *jlexer.Lexer, out *Event) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var MethodSet bool
	var ParameterSet bool
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
		case "method":
			out.Method = string(in.String())
			MethodSet = true
		case "parameter":
			(out.Parameter).UnmarshalEasyJSON(in)
			ParameterSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !MethodSet {
		in.AddError(fmt.Errorf("key 'method' is required"))
	}
	if !ParameterSet {
		in.AddError(fmt.Errorf("key 'parameter' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes12(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"parameter\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Parameter).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComGoParkMailRu2018242GameServerTypes12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComGoParkMailRu2018242GameServerTypes12(l, v)
}