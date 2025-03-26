package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConectToDb() *mongo.Client {
    	err := godotenv.Load(".env") // Get the environment set up (currently just using localhost and the DB I set up)
	if err != nil { // Return an error if the environment isn't set up
		log.Fatal("Error loading environment: check \".env\" file", err)
	}


	URI := os.Getenv("URI") // Gets the database's URI from the ".env"
	dbOpts := options.Client().ApplyURI(URI) // Parses the URI 
	db, err := mongo.Connect(context.Background(), dbOpts) // Gets our database connection

	if err != nil {
		log.Fatal("Error connecting to database", err)
        print("error!")
	}


	err = db.Ping(context.Background(), nil) // Checks for a ping

	if err != nil {
		log.Fatal(err)
        print("error!")
	}

	fmt.Println("Connected to database") // Prints that the connection was made

    return db // returns pointer to database
}
