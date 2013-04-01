package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/call", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			w.Write([]byte(req.Form.Get("echostr")))
			return
		}
	})
	http.ListenAndServe(":8884", nil)
}
