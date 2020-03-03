package WeHub

import "github.com/silenceper/wechat/message"

type BaseMessage struct {
	Message *message.MixMessage `json:"message"`
}

type BaseReply struct {
	Reply *message.Reply `json:"reply"`
}
