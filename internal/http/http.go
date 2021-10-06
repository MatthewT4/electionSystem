package http

import (
	"electionSystem/internal/blogic"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Router struct {
	ser             blogic.IBVoting
	infoLog         *log.Logger
	MessTemplate    *template.Template
	ShablonTemplate *template.Template
}

func NewRouter(db *mongo.Database) *Router {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return &Router{blogic.CreateBVoting(db),
		log.New(f, "INFO\t", log.Ldate|log.Ltime),
		template.Must(template.ParseFiles("static/message.html")),
		template.Must(template.ParseFiles("static/shabl.html"))}
}

func (rout *Router) Start() {
	rout.infoLog.Println("Server Started")
	fmt.Println("Server Started")
	//fs := http.FileServer(http.Dir("static"))
	staticHandler := http.StripPrefix("/data/", http.FileServer(http.Dir("static")))
	rou := mux.NewRouter()
	rou.HandleFunc("/get_data", rout.GetData)
	//rou.Handle("/data/", staticHandler)
	rou.HandleFunc("/vote", rout.Vote)
	rou.HandleFunc("/", rout.FormCandidates)
	rou.HandleFunc("/login", rout.Login)
	rou.Handle("/data/", staticHandler)
	//r.HandleFunc("/screen_register", rout.NewScreen)
	//rou.Handle("/", r)
	/*
		rScreen := rou.PathPrefix("/api/scr").Subrouter()
		rScreen.HandleFunc("/getdata", rout.GetImageScreen)
		rScreen.Use(rout.ScreenAuthentication)

		rAdm := rou.PathPrefix("/api/admin").Subrouter()
		//r.HandleFunc("/", HomeHandler)
		//r.HandleFunc("/login", LoginUser)
		rAdm.HandleFunc("/addcatalog", rout.AddCatalog).Methods("POST")
		rAdm.HandleFunc("/insertdata", rout.InsertDataInCatalog).Methods("POST")
		rAdm.HandleFunc("/getcatalogs", rout.GetCatalogs).Methods("GET")
		rAdm.HandleFunc("/getdata", rout.GetDataInCatalog).Methods("GET")
		//r.HandleFunc("/articles", ArticlesHandler)
		rAdm.Use(rout.AdminAuthentication)*/

	srv := &http.Server{
		Handler: rou,
		Addr:    ":80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
