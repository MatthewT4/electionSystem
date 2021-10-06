package http

import (
	"fmt"
	"net/http"
)

func (rout *Router) GetData(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	nameElection := r.FormValue("election")
	if login != "root" || password != "ffdd4518" {
		http.Error(w, "err", 404)
	}
	code, data := rout.ser.GetVoteInElection(nameElection)
	if code == 200 {
		fmt.Fprintln(w, data)
		return
	}
	http.Error(w, data, 500)
}

func (rout *Router) Vote(w http.ResponseWriter, r *http.Request) {

	var OnlyVoitVar struct {
		Candidate string `json:"candidate"`
		Token     string `json:"token"`
	}
	OnlyVoitVar.Token = r.FormValue("token")
	OnlyVoitVar.Candidate = r.FormValue("candidate")
	rout.infoLog.Printf("[VOTE]: token: %v candidate: %v", OnlyVoitVar.Token, OnlyVoitVar.Candidate)
	fmt.Printf("[VOTE]: token: %v candidate: %v", OnlyVoitVar.Token, OnlyVoitVar.Candidate)
	if OnlyVoitVar.Candidate == "" {
		rout.infoLog.Printf("[VOTE]: candidate is null, token: %v", OnlyVoitVar.Token)
		var Data struct {
			Mess string
		}
		Data.Mess = "Похоже вы не выбрали ни одного кандидата"
		rout.MessTemplate.Execute(w, Data)
		//http.Error(w, "candidate is null :(", 404)
		return
	}
	if OnlyVoitVar.Token == "" {
		rout.infoLog.Printf("[VOTE]:token is null, candidate: %v", OnlyVoitVar.Candidate)
		http.Error(w, "token is null :(", 404)
		return
	}
	//fmt.Println(OnlyVoitVar.Token + "   " + OnlyVoitVar.Candidate)
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
	fmt.Println(string(code) + mess)
	if code == 200 {
		http.ServeFile(w, r, "static/done.html")
		rout.infoLog.Printf("[VOTE]: token %v OK", OnlyVoitVar.Token)
		return
	}
	rout.infoLog.Printf("[VOTE]: ERROR request voted: token: %v code: %v  error: %v", OnlyVoitVar.Token, code, mess)
	var Data struct {
		Mess string
	}
	Data.Mess = mess
	rout.MessTemplate.Execute(w, Data)
}
