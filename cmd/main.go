package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"github.com/zaakirio/go-pomodoro-api/pkg/routes"
	"google.golang.org/api/option"
)

func main() {
	e := echo.New()

	opt := option.WithCredentialsFile("firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase: %v\n", err)
	}

	// Get Firestore client
	firestoreClient, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}

	// Register routes and pass the Firestore client
	routes.RegisterRoutes(e, firestoreClient)

	// Start the server
	log.Fatal(e.Start(":8080"))
}
