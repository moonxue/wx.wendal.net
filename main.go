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
	return
}

func (w *WendalWeixinHandler) Location(msg goweixin.Message) (reply goweixin.Replay) {
	reply = goweixin.Replay{}
	log.Println("Mark : " + msg.Label())
	reply.SetContent("Mark : " + msg.Label())
	return
}

func main() {
	log.Fatal(http.ListenAndServe(":8883", &goweixin.WxHttpHandler{"", &WendalWeixinHandler{}}))
}

func init() {
	goweixin.DevMode = true
	goweixin.SetDebug(true)
}
