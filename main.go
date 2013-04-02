package main

import (
	"github.com/wendal/goweixin"
	"net/http"
)

type WendalWeixinHandler struct {
	*goweixin.BaseWeiXinHandler
}

func (w *WendalWeixinHandler) Text(msg goweixin.Message) (reply goweixin.Replay) {
	reply = goweixin.Replay{}
	reply.SetContent("OK")
	return
}

func main() {
	http.ListenAndServe(":8884", &goweixin.WxHttpHandler{"", &WendalWeixinHandler{}})
}
