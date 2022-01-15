package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type OnetextData struct {
	Text string   `json:"text"`
	By   string   `json:"by,omitempty"`
	From string   `json:"from,omitempty"`
	Time []string `json:"time"`
	Uri  string   `json:"uri,omitempty"`
}

var JsonData []OnetextData

func init() {
	resp, err := http.Get("https://raw.githubusercontent.com/lz233/OneText-Library/master/OneText-Library.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	jsonByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(jsonByte, &JsonData)
}

func HandleOnetext(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	n := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes, _ := json.MarshalIndent(JsonData[n.Intn(len(JsonData))], "", "    ")
	_, _ = fmt.Fprint(w, string(bytes))
}
