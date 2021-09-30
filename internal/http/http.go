package http

import (
	"electionSystem/internal/blogic"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

type Router struct {
	ser blogic.IBVoting
}

func NewRouter(db *mongo.Database) *Router {
	return &Router{blogic.CreateBVoting(db)}
}

func (rout *Router) Start() {
	fmt.Println("Server Started")
	rou := mux.NewRouter()
	//r.HandleFunc("/getdata", rout.GetData)
	rou.HandleFunc("/voit", rout.Voit)
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
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
