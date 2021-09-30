package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (rout *Router) Voit(w http.ResponseWriter, r *http.Request) {
	var OnlyVoitVar struct {
		Candidate string `json:"candidate"`
		Token     string `json:"token"`
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Server Error", 500)
		return
	}
	err = json.Unmarshal(body, &OnlyVoitVar)
	if err != nil {
		http.Error(w, "Parse JSON error", 400)
		return
	}
	code, mess := rout.ser.Vote(OnlyVoitVar.Token, OnlyVoitVar.Candidate)
	w.WriteHeader(code)
	fmt.Fprintln(w, mess)
}
