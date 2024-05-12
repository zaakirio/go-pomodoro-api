package main

import (
	"context"
	"fmt"
	"log"

	firebase "cloud.google.com/go/firestore"

	"github.com/labstack/echo/v4"
	"github.com/zaakirio/go-pomodoro-api/pkg/routes"
	"google.golang.org/api/option"
)

func main() {
	e := echo.New()

	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")

	firestoreClient, err := firebase.NewClient(ctx, "go-pomodoro-api", opt)
	if err != nil {
		log.Fatalf("error initializing firebase: %v\n", err)
	}

	routes.RegisterRoutes(e, firestoreClient)
	printRoutes(e)

	log.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Println("\033[1;36mRegistered routes:\033[0m")
	for _, route := range e.Routes() {
		method := route.Method
		path := route.Path
		fmt.Printf("\033[1;36m%-6s ==> %-25s\n", method, path)
	}
}
