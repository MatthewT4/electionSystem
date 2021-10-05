package http

import (
	"fmt"
	"net/http"
)

func (rout *Router) Voit(w http.ResponseWriter, r *http.Request) {

	var OnlyVoitVar struct {
		Candidate string `json:"candidate"`
		Token     string `json:"token"`
	}
	OnlyVoitVar.Token = r.FormValue("token")
	OnlyVoitVar.Candidate = r.FormValue("candidate")
	if OnlyVoitVar.Candidate == "" {
		http.Error(w, "candidate not found :(", 404)
		return
	}
	if OnlyVoitVar.Token == "" {
		http.Error(w, "token not found :(", 404)
		return
	}
	fmt.Println(OnlyVoitVar.Token + "   " + OnlyVoitVar.Candidate)
	/*
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Server Error", 500)
			return
		}
		err = json.Unmarshal(body, &OnlyVoitVar)
		if err != nil {
			http.Error(w, "Parse JSON error", 400)
			return
		}*/
	code, mess := rout.ser.Vote(OnlyVoitVar.Token, OnlyVoitVar.Candidate)
	if code == 200 {
		http.ServeFile(w, r, "static/done.html")
		return
	}
	w.WriteHeader(code)
	fmt.Fprintln(w, mess)
}
