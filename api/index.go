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

func handleOnetext(w http.ResponseWriter, r *http.Request) {
	var jsonData []OnetextData
	jsonBuffer, _ := ioutil.ReadFile("../OneText-Library.json")
	_ = json.Unmarshal(jsonBuffer, &jsonData)
	n := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes, _ := json.MarshalIndent(jsonData[n.Intn(len(jsonData))], "", "    ")
	_, _ = fmt.Fprint(w, string(bytes))
}
