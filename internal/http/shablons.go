package http

import (
	"fmt"
	"net/http"
)

func (rout *Router) FormCandidates(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/login.html")
}

func (rout *Router) Login(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	rout.infoLog.Printf("[LOGIN]: token %v logged in", token)
	val, mes := rout.ser.Login(token)
	fmt.Println(val, mes)
	if val != true {
		rout.infoLog.Printf("[LOGIN]: token %v error: %v", token, mes)
		//TMPL := template.Must(template.ParseFiles("static/message.html"))
		var Data struct {
			Mess string
		}
		Data.Mess = mes
		rout.MessTemplate.Execute(w, Data)
		//fmt.Fprintln(w, err)
		return
	}
	rout.infoLog.Printf("[LOGIN]: the token %v has been authorized", token)
	//TMPL := template.Must(template.ParseFiles("static/shabl.html"))
	var User struct {
		Token string
		Users map[string]string
	}
	mas, err := rout.ser.GetCandidates(mes)
	if err != nil {
		rout.infoLog.Printf("[LOGIN]: ERROR GetCandidates: %v   token: %v", err.Error(), token)
		var Data struct {
			Mess string
		}
		Data.Mess = err.Error()
		rout.MessTemplate.Execute(w, Data)
		return
	}
	fmt.Println(mas)
	User.Users = mas
	User.Token = token
	rout.ShablonTemplate.Execute(w, User)
	//http.ServeFile(w, r, "static/shabl.html")
}

func (rout *Router) Admin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/lanin.html")
}
