package routes

import (
	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/zaakirio/go-pomodoro-api/pkg/handlers"
)

func RegisterRoutes(e *echo.Echo, firestoreClient *firestore.Client) {
	handlers.FirestoreClient = firestoreClient
	e.GET("/pomodoros", handlers.GetPomodoros)
	e.POST("/pomodoros", handlers.CreatePomodoro)
	e.PUT("/pomodoros/:id", handlers.UpdatePomodoro)
	e.DELETE("/pomodoros/:id", handlers.DeletePomodoro)
}
