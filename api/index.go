package api

import (
	"github.com/XiaoMengXinX/OneTextAPI-Go"
	"net/http"
)

var o onetext.OneText

func init() {
	err := o.GetUrl("https://raw.githubusercontent.com/lz233/OneText-Library/master/OneText-Library.json")
	if err != nil {
		panic(err)
	}
}

func HandleOnetext(w http.ResponseWriter, r *http.Request) {
	o.Handler(w, r)
}
