package http

import (
	"fmt"
	"html/template"
	"net/http"
)

func (rout *Router) FormCandidates(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/login.html")
}

func (rout *Router) Login(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	fmt.Println(token)
	val, err := rout.ser.Login(token)
	fmt.Println(val, err)
	if val != true {
		fmt.Fprintln(w, err)
		return
	}
	TMPL := template.Must(template.ParseFiles("static/shabl.html"))
	var User struct {
		Token string
	}
	User.Token = token
	TMPL.Execute(w, User)
	//http.ServeFile(w, r, "static/shabl.html")
}
