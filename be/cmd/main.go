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

	// Load .env from file when in dev mode (Non-Docker)
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("could not connect" , err)
	}

	// ping db
	pingCtx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)	
	defer cancel()

	retryInterval := 2 * time.Second
	for {
		attemptCtx, attemptCancel := context.WithTimeout(pingCtx, retryInterval)
		err := client.Ping(attemptCtx, nil)
		attemptCancel() // clean up context

		if err == nil {
			log.Println("successfully connected to mongo db")
			break
		}

		select {
			case <-pingCtx.Done():
				log.Fatal("timed out waiting for MongoDB to respond to ping:", err)
			case <-time.After(retryInterval):
				log.Println("MongoDB ping failed, retrying...", err)
			}
	}
	log.Println("successfully connected to mongo db")
	bePort := os.Getenv("BE_PORT")
	server := api.NewAPIServer(":"+bePort, client)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
