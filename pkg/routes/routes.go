package routes

import (
	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/zaakirio/go-pomodoro-api/pkg/handlers"
)

func RegisterRoutes(e *echo.Echo, firestoreClient *firestore.Client) {
	handlers.FirestoreClient = firestoreClient

	e.GET("/pomodoros/fetch", handlers.GetPomodoros)
	e.GET("/pomodoros/fetch/:id", handlers.GetPomodoroById)
	e.POST("/pomodoros/create", handlers.CreatePomodoro)
	e.PUT("/pomodoros/update/:id", handlers.UpdatePomodoro)
	e.DELETE("/pomodoros/delete/:id", handlers.DeletePomodoro)
}
