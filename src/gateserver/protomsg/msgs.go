package protomsg

import (
	"github.com/gfandada/gserver/network"
)

func NewMsgCoder() *network.MsgManager {
	coder := network.NewMsgManager()
	coder.Register(&network.RawMessage{
		MsgId:   uint16(0),
		MsgData: &HeartbeatReq{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(1),
		MsgData: &HeartbeatAck{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2),
		MsgData: &ErrorAck{},
	})
	// for login
	coder.Register(&network.RawMessage{
		MsgId:   uint16(500),
		MsgData: &TouristReq{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(501),
		MsgData: &TouristAck{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(502),
		MsgData: &TouristLoginReq{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(503),
		MsgData: &TouristLoginAck{},
	})
	return coder
}
