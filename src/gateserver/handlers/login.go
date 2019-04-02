package handlers

import (
	"gserver_gateway_demo/src/gateserver/login"
	Msg "gserver_gateway_demo/src/gateserver/protomsg"

	"github.com/gfandada/gserver/network"
	"github.com/golang/protobuf/proto"
)

// 创建游客账号
func TouristReqHandler(arg []interface{}) []interface{} {
	id, acc, pwd := login.NewTourist()
	return []interface{}{network.RawMessage{
		MsgId: uint16(501),
		MsgData: &Msg.TouristAck{
			Id:       proto.Int32(id),
			Account:  proto.String(acc),
			Password: proto.String(pwd),
		},
	}}
}

// 游客登录
func TouristLoginReqHandler(arg []interface{}) []interface{} {
	data := arg[0].(*network.RawMessage).MsgData
	userid := data.(*Msg.TouristLoginReq).GetId()
	acc := data.(*Msg.TouristLoginReq).GetAccount()
	pwd := data.(*Msg.TouristLoginReq).GetPassword()
	if ok := login.CheckTourist(acc, pwd); !ok {
		panic(1000)
	}
	return []interface{}{network.RawMessage{
		MsgId:   uint16(503),
		MsgData: &Msg.TouristLoginAck{},
	}, userid}
}
