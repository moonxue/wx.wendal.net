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
	return goweixin.ReplyText("Got it")
}

func (w *WendalWeixinHandler) Location(msg goweixin.Message) (reply goweixin.Replay) {
	return goweixin.ReplyTextf("Mark: %s %s %s", msg.Location_X(), msg.Location_Y(), msg.Label())
}

func main() {
	log.Fatal(http.ListenAndServe(":8883", &goweixin.WxHttpHandler{"", &WendalWeixinHandler{}}))
}

func init() {
	goweixin.DevMode = true
	goweixin.SetDebug(true)
}
