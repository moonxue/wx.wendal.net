package main

import (
	"github.com/wendal/goweixin"
	"log"
	"net/http"
)

type WendalWeixinHandler struct {
	*goweixin.BaseWeiXinHandler
}

func (w *WendalWeixinHandler) Text(msg goweixin.Message) (reply goweixin.Replay) {
	reply = goweixin.Replay{}
	reply.SetContent("OK")
	return reply
}

func main() {
	goweixin.DevMode = true
	goweixin.SetDebug(true)
	log.Fatal(http.ListenAndServe(":8883", &goweixin.WxHttpHandler{"", &WendalWeixinHandler{}}))
}
