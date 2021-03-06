package internal

import (
	"electionSystem/internal/blogic"
	"electionSystem/internal/http"
	mongodb "electionSystem/pkg/mongoDB"
	"fmt"
	"log"
)

func StartServer() {
	client, err := mongodb.NewClient("mongodb+srv://cluster0.lbets.mongodb.net/myFirstDatabase", "Mathew", "8220")
	if err != nil {
		log.Fatal(err)
	}
	name := "ElectionsDB"
	/*for i := 0; i < 20; i++ {
		fmt.Println(blogic.GenerateToken(10))
	}*/
	db := client.Database(name)
	router := http.NewRouter(db)
	router.Start() // Start http server

	logic := blogic.CreateBVoting(db)
	res, strr := logic.Vote("ffff", "Orudzev")
	fmt.Println(string(res) + " " + strr)
	//logic.BAddVoted()
}
