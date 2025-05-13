package main

import (
	"log"
	"psp-dashboard-be/cmd/api"
)

func main() {

	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	log.Println("server running on port 8080")
}
