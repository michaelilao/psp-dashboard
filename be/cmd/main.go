package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"psp-dashboard-be/cmd/api"
	"time"

	_ "psp-dashboard-be/docs"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Expense Management Dashboard API
// @version 1.0
// @description Expense Tracking Dashboard for PSP Take Home Assignment
// @host localhost:8080
// @BasePath /

func main() {

	if os.Getenv("DOCKER") == "" {
    err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}



	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, dbPort)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("failed to ping mongodb:", err)
	}
	log.Println("successfully connected to mongo db")


	bePort := os.Getenv("BE_PORT")
	server := api.NewAPIServer(":"+bePort, client)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
