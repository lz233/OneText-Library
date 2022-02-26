package api

import (
	onetext "github.com/XiaoMengXinX/OneTextAPI-Go"
	"net/http"
	"onetextJson"
)

var o = onetext.New()

func init() {
	o.ReadBytes(text.Onetext)
}

func HandleOnetext(w http.ResponseWriter, r *http.Request) {
	o.Handler(w, r)
}
