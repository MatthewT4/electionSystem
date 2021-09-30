package main

import (
	"electionSystem/internal"
	"fmt"
)

func main() {
	fmt.Println("ffffp")
	internal.StartServer()
	// curl -d "{\"token\":\"ffff\",\"candidate\":\"Orudzev\"}" -X POST 127.0.0.1:8000/voit -v
}
