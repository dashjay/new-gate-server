// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package WeHub

import (
	json "encoding/json"
	xml "encoding/xml"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	message "github.com/silenceper/wechat/message"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson666b529aDecodeNewGateServerWeHub(in *jlexer.Lexer, out *BaseReply) {
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
		case "reply":
			if in.IsNull() {
				in.Skip()
				out.Reply = nil
			} else {
				if out.Reply == nil {
					out.Reply = new(message.Reply)
				}
				easyjson666b529aDecodeGithubComSilenceperWechatMessage(in, out.Reply)
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
func easyjson666b529aEncodeNewGateServerWeHub(out *jwriter.Writer, in BaseReply) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"reply\":"
		out.RawString(prefix[1:])
		if in.Reply == nil {
			out.RawString("null")
		} else {
			easyjson666b529aEncodeGithubComSilenceperWechatMessage(out, *in.Reply)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BaseReply) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson666b529aEncodeNewGateServerWeHub(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BaseReply) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson666b529aEncodeNewGateServerWeHub(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BaseReply) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson666b529aDecodeNewGateServerWeHub(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BaseReply) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson666b529aDecodeNewGateServerWeHub(l, v)
}
func easyjson666b529aDecodeGithubComSilenceperWechatMessage(in *jlexer.Lexer, out *message.Reply) {
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
		case "MsgType":
			out.MsgType = message.MsgType(in.String())
		case "MsgData":
			if m, ok := out.MsgData.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.MsgData.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.MsgData = in.Interface()
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
func easyjson666b529aEncodeGithubComSilenceperWechatMessage(out *jwriter.Writer, in message.Reply) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"MsgType\":"
		out.RawString(prefix[1:])
		out.String(string(in.MsgType))
	}
	{
		const prefix string = ",\"MsgData\":"
		out.RawString(prefix)
		if m, ok := in.MsgData.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.MsgData.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.MsgData))
		}
	}
	out.RawByte('}')
}
func easyjson666b529aDecodeNewGateServerWeHub1(in *jlexer.Lexer, out *BaseMessage) {
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
			if in.IsNull() {
				in.Skip()
				out.Message = nil
			} else {
				if out.Message == nil {
					out.Message = new(message.MixMessage)
				}
				easyjson666b529aDecodeGithubComSilenceperWechatMessage1(in, out.Message)
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
func easyjson666b529aEncodeNewGateServerWeHub1(out *jwriter.Writer, in BaseMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix[1:])
		if in.Message == nil {
			out.RawString("null")
		} else {
			easyjson666b529aEncodeGithubComSilenceperWechatMessage1(out, *in.Message)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BaseMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson666b529aEncodeNewGateServerWeHub1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BaseMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson666b529aEncodeNewGateServerWeHub1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BaseMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson666b529aDecodeNewGateServerWeHub1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BaseMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson666b529aDecodeNewGateServerWeHub1(l, v)
}
func easyjson666b529aDecodeGithubComSilenceperWechatMessage1(in *jlexer.Lexer, out *message.MixMessage) {
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
		case "MsgID":
			out.MsgID = int64(in.Int64())
		case "Content":
			out.Content = string(in.String())
		case "Recognition":
			out.Recognition = string(in.String())
		case "PicURL":
			out.PicURL = string(in.String())
		case "MediaID":
			out.MediaID = string(in.String())
		case "Format":
			out.Format = string(in.String())
		case "ThumbMediaID":
			out.ThumbMediaID = string(in.String())
		case "LocationX":
			out.LocationX = float64(in.Float64())
		case "LocationY":
			out.LocationY = float64(in.Float64())
		case "Scale":
			out.Scale = float64(in.Float64())
		case "Label":
			out.Label = string(in.String())
		case "Title":
			out.Title = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "URL":
			out.URL = string(in.String())
		case "Event":
			out.Event = message.EventType(in.String())
		case "EventKey":
			out.EventKey = string(in.String())
		case "Ticket":
			out.Ticket = string(in.String())
		case "Latitude":
			out.Latitude = string(in.String())
		case "Longitude":
			out.Longitude = string(in.String())
		case "Precision":
			out.Precision = string(in.String())
		case "MenuID":
			out.MenuID = string(in.String())
		case "Status":
			out.Status = string(in.String())
		case "SessionFrom":
			out.SessionFrom = string(in.String())
		case "ScanCodeInfo":
			easyjson666b529aDecode(in, &out.ScanCodeInfo)
		case "SendPicsInfo":
			easyjson666b529aDecode1(in, &out.SendPicsInfo)
		case "SendLocationInfo":
			easyjson666b529aDecode2(in, &out.SendLocationInfo)
		case "InfoType":
			out.InfoType = message.InfoType(in.String())
		case "AppID":
			out.AppID = string(in.String())
		case "ComponentVerifyTicket":
			out.ComponentVerifyTicket = string(in.String())
		case "AuthorizerAppid":
			out.AuthorizerAppid = string(in.String())
		case "AuthorizationCode":
			out.AuthorizationCode = string(in.String())
		case "AuthorizationCodeExpiredTime":
			out.AuthorizationCodeExpiredTime = int64(in.Int64())
		case "PreAuthCode":
			out.PreAuthCode = string(in.String())
		case "CardID":
			out.CardID = string(in.String())
		case "RefuseReason":
			out.RefuseReason = string(in.String())
		case "IsGiveByFriend":
			out.IsGiveByFriend = int32(in.Int32())
		case "FriendUserName":
			out.FriendUserName = string(in.String())
		case "UserCardCode":
			out.UserCardCode = string(in.String())
		case "OldUserCardCode":
			out.OldUserCardCode = string(in.String())
		case "OuterStr":
			out.OuterStr = string(in.String())
		case "IsRestoreMemberCard":
			out.IsRestoreMemberCard = int32(in.Int32())
		case "UnionID":
			out.UnionID = string(in.String())
		case "IsRisky":
			out.IsRisky = bool(in.Bool())
		case "ExtraInfoJSON":
			out.ExtraInfoJSON = string(in.String())
		case "TraceID":
			out.TraceID = string(in.String())
		case "StatusCode":
			out.StatusCode = int(in.Int())
		case "DeviceType":
			out.DeviceType = string(in.String())
		case "DeviceID":
			out.DeviceID = string(in.String())
		case "SessionID":
			out.SessionID = string(in.String())
		case "OpenID":
			out.OpenID = string(in.String())
		case "XMLName":
			easyjson666b529aDecodeEncodingXml(in, &out.XMLName)
		case "ToUserName":
			out.ToUserName = message.CDATA(in.String())
		case "FromUserName":
			out.FromUserName = message.CDATA(in.String())
		case "CreateTime":
			out.CreateTime = int64(in.Int64())
		case "MsgType":
			out.MsgType = message.MsgType(in.String())
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
func easyjson666b529aEncodeGithubComSilenceperWechatMessage1(out *jwriter.Writer, in message.MixMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"MsgID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.MsgID))
	}
	{
		const prefix string = ",\"Content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"Recognition\":"
		out.RawString(prefix)
		out.String(string(in.Recognition))
	}
	{
		const prefix string = ",\"PicURL\":"
		out.RawString(prefix)
		out.String(string(in.PicURL))
	}
	{
		const prefix string = ",\"MediaID\":"
		out.RawString(prefix)
		out.String(string(in.MediaID))
	}
	{
		const prefix string = ",\"Format\":"
		out.RawString(prefix)
		out.String(string(in.Format))
	}
	{
		const prefix string = ",\"ThumbMediaID\":"
		out.RawString(prefix)
		out.String(string(in.ThumbMediaID))
	}
	{
		const prefix string = ",\"LocationX\":"
		out.RawString(prefix)
		out.Float64(float64(in.LocationX))
	}
	{
		const prefix string = ",\"LocationY\":"
		out.RawString(prefix)
		out.Float64(float64(in.LocationY))
	}
	{
		const prefix string = ",\"Scale\":"
		out.RawString(prefix)
		out.Float64(float64(in.Scale))
	}
	{
		const prefix string = ",\"Label\":"
		out.RawString(prefix)
		out.String(string(in.Label))
	}
	{
		const prefix string = ",\"Title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"URL\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"Event\":"
		out.RawString(prefix)
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"EventKey\":"
		out.RawString(prefix)
		out.String(string(in.EventKey))
	}
	{
		const prefix string = ",\"Ticket\":"
		out.RawString(prefix)
		out.String(string(in.Ticket))
	}
	{
		const prefix string = ",\"Latitude\":"
		out.RawString(prefix)
		out.String(string(in.Latitude))
	}
	{
		const prefix string = ",\"Longitude\":"
		out.RawString(prefix)
		out.String(string(in.Longitude))
	}
	{
		const prefix string = ",\"Precision\":"
		out.RawString(prefix)
		out.String(string(in.Precision))
	}
	{
		const prefix string = ",\"MenuID\":"
		out.RawString(prefix)
		out.String(string(in.MenuID))
	}
	{
		const prefix string = ",\"Status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"SessionFrom\":"
		out.RawString(prefix)
		out.String(string(in.SessionFrom))
	}
	{
		const prefix string = ",\"ScanCodeInfo\":"
		out.RawString(prefix)
		easyjson666b529aEncode(out, in.ScanCodeInfo)
	}
	{
		const prefix string = ",\"SendPicsInfo\":"
		out.RawString(prefix)
		easyjson666b529aEncode1(out, in.SendPicsInfo)
	}
	{
		const prefix string = ",\"SendLocationInfo\":"
		out.RawString(prefix)
		easyjson666b529aEncode2(out, in.SendLocationInfo)
	}
	{
		const prefix string = ",\"InfoType\":"
		out.RawString(prefix)
		out.String(string(in.InfoType))
	}
	{
		const prefix string = ",\"AppID\":"
		out.RawString(prefix)
		out.String(string(in.AppID))
	}
	{
		const prefix string = ",\"ComponentVerifyTicket\":"
		out.RawString(prefix)
		out.String(string(in.ComponentVerifyTicket))
	}
	{
		const prefix string = ",\"AuthorizerAppid\":"
		out.RawString(prefix)
		out.String(string(in.AuthorizerAppid))
	}
	{
		const prefix string = ",\"AuthorizationCode\":"
		out.RawString(prefix)
		out.String(string(in.AuthorizationCode))
	}
	{
		const prefix string = ",\"AuthorizationCodeExpiredTime\":"
		out.RawString(prefix)
		out.Int64(int64(in.AuthorizationCodeExpiredTime))
	}
	{
		const prefix string = ",\"PreAuthCode\":"
		out.RawString(prefix)
		out.String(string(in.PreAuthCode))
	}
	{
		const prefix string = ",\"CardID\":"
		out.RawString(prefix)
		out.String(string(in.CardID))
	}
	{
		const prefix string = ",\"RefuseReason\":"
		out.RawString(prefix)
		out.String(string(in.RefuseReason))
	}
	{
		const prefix string = ",\"IsGiveByFriend\":"
		out.RawString(prefix)
		out.Int32(int32(in.IsGiveByFriend))
	}
	{
		const prefix string = ",\"FriendUserName\":"
		out.RawString(prefix)
		out.String(string(in.FriendUserName))
	}
	{
		const prefix string = ",\"UserCardCode\":"
		out.RawString(prefix)
		out.String(string(in.UserCardCode))
	}
	{
		const prefix string = ",\"OldUserCardCode\":"
		out.RawString(prefix)
		out.String(string(in.OldUserCardCode))
	}
	{
		const prefix string = ",\"OuterStr\":"
		out.RawString(prefix)
		out.String(string(in.OuterStr))
	}
	{
		const prefix string = ",\"IsRestoreMemberCard\":"
		out.RawString(prefix)
		out.Int32(int32(in.IsRestoreMemberCard))
	}
	{
		const prefix string = ",\"UnionID\":"
		out.RawString(prefix)
		out.String(string(in.UnionID))
	}
	{
		const prefix string = ",\"IsRisky\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsRisky))
	}
	{
		const prefix string = ",\"ExtraInfoJSON\":"
		out.RawString(prefix)
		out.String(string(in.ExtraInfoJSON))
	}
	{
		const prefix string = ",\"TraceID\":"
		out.RawString(prefix)
		out.String(string(in.TraceID))
	}
	{
		const prefix string = ",\"StatusCode\":"
		out.RawString(prefix)
		out.Int(int(in.StatusCode))
	}
	{
		const prefix string = ",\"DeviceType\":"
		out.RawString(prefix)
		out.String(string(in.DeviceType))
	}
	{
		const prefix string = ",\"DeviceID\":"
		out.RawString(prefix)
		out.String(string(in.DeviceID))
	}
	{
		const prefix string = ",\"SessionID\":"
		out.RawString(prefix)
		out.String(string(in.SessionID))
	}
	{
		const prefix string = ",\"OpenID\":"
		out.RawString(prefix)
		out.String(string(in.OpenID))
	}
	{
		const prefix string = ",\"XMLName\":"
		out.RawString(prefix)
		easyjson666b529aEncodeEncodingXml(out, in.XMLName)
	}
	{
		const prefix string = ",\"ToUserName\":"
		out.RawString(prefix)
		out.String(string(in.ToUserName))
	}
	{
		const prefix string = ",\"FromUserName\":"
		out.RawString(prefix)
		out.String(string(in.FromUserName))
	}
	{
		const prefix string = ",\"CreateTime\":"
		out.RawString(prefix)
		out.Int64(int64(in.CreateTime))
	}
	{
		const prefix string = ",\"MsgType\":"
		out.RawString(prefix)
		out.String(string(in.MsgType))
	}
	out.RawByte('}')
}
func easyjson666b529aDecodeEncodingXml(in *jlexer.Lexer, out *xml.Name) {
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
		case "Space":
			out.Space = string(in.String())
		case "Local":
			out.Local = string(in.String())
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
func easyjson666b529aEncodeEncodingXml(out *jwriter.Writer, in xml.Name) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Space\":"
		out.RawString(prefix[1:])
		out.String(string(in.Space))
	}
	{
		const prefix string = ",\"Local\":"
		out.RawString(prefix)
		out.String(string(in.Local))
	}
	out.RawByte('}')
}
func easyjson666b529aDecode2(in *jlexer.Lexer, out *struct {
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     float64 `xml:"Scale"`
	Label     string  `xml:"Label"`
	Poiname   string  `xml:"Poiname"`
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "LocationX":
			out.LocationX = float64(in.Float64())
		case "LocationY":
			out.LocationY = float64(in.Float64())
		case "Scale":
			out.Scale = float64(in.Float64())
		case "Label":
			out.Label = string(in.String())
		case "Poiname":
			out.Poiname = string(in.String())
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
func easyjson666b529aEncode2(out *jwriter.Writer, in struct {
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     float64 `xml:"Scale"`
	Label     string  `xml:"Label"`
	Poiname   string  `xml:"Poiname"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"LocationX\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.LocationX))
	}
	{
		const prefix string = ",\"LocationY\":"
		out.RawString(prefix)
		out.Float64(float64(in.LocationY))
	}
	{
		const prefix string = ",\"Scale\":"
		out.RawString(prefix)
		out.Float64(float64(in.Scale))
	}
	{
		const prefix string = ",\"Label\":"
		out.RawString(prefix)
		out.String(string(in.Label))
	}
	{
		const prefix string = ",\"Poiname\":"
		out.RawString(prefix)
		out.String(string(in.Poiname))
	}
	out.RawByte('}')
}
func easyjson666b529aDecode1(in *jlexer.Lexer, out *struct {
	Count   int32              `xml:"Count"`
	PicList []message.EventPic `xml:"PicList>item"`
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Count":
			out.Count = int32(in.Int32())
		case "PicList":
			if in.IsNull() {
				in.Skip()
				out.PicList = nil
			} else {
				in.Delim('[')
				if out.PicList == nil {
					if !in.IsDelim(']') {
						out.PicList = make([]message.EventPic, 0, 4)
					} else {
						out.PicList = []message.EventPic{}
					}
				} else {
					out.PicList = (out.PicList)[:0]
				}
				for !in.IsDelim(']') {
					var v1 message.EventPic
					easyjson666b529aDecodeGithubComSilenceperWechatMessage2(in, &v1)
					out.PicList = append(out.PicList, v1)
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
func easyjson666b529aEncode1(out *jwriter.Writer, in struct {
	Count   int32              `xml:"Count"`
	PicList []message.EventPic `xml:"PicList>item"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Count\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.Count))
	}
	{
		const prefix string = ",\"PicList\":"
		out.RawString(prefix)
		if in.PicList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.PicList {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson666b529aEncodeGithubComSilenceperWechatMessage2(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson666b529aDecodeGithubComSilenceperWechatMessage2(in *jlexer.Lexer, out *message.EventPic) {
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
		case "PicMd5Sum":
			out.PicMd5Sum = string(in.String())
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
func easyjson666b529aEncodeGithubComSilenceperWechatMessage2(out *jwriter.Writer, in message.EventPic) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"PicMd5Sum\":"
		out.RawString(prefix[1:])
		out.String(string(in.PicMd5Sum))
	}
	out.RawByte('}')
}
func easyjson666b529aDecode(in *jlexer.Lexer, out *struct {
	ScanType   string `xml:"ScanType"`
	ScanResult string `xml:"ScanResult"`
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ScanType":
			out.ScanType = string(in.String())
		case "ScanResult":
			out.ScanResult = string(in.String())
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
func easyjson666b529aEncode(out *jwriter.Writer, in struct {
	ScanType   string `xml:"ScanType"`
	ScanResult string `xml:"ScanResult"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ScanType\":"
		out.RawString(prefix[1:])
		out.String(string(in.ScanType))
	}
	{
		const prefix string = ",\"ScanResult\":"
		out.RawString(prefix)
		out.String(string(in.ScanResult))
	}
	out.RawByte('}')
}
