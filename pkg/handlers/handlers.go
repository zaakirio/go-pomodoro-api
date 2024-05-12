package handlers

import (
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/zaakirio/go-pomodoro-api/pkg/models"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var FirestoreClient *firestore.Client

func GetPomodoros(c echo.Context) error {
	iter := FirestoreClient.Collection("pomodoros").Documents(c.Request().Context())
	var pomodoros []models.Pomodoro
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		var pomodoro models.Pomodoro
		doc.DataTo(&pomodoro)
		pomodoro.ID = doc.Ref.ID
		pomodoros = append(pomodoros, pomodoro)
	}
	return c.JSON(http.StatusOK, pomodoros)
}

func CreatePomodoro(c echo.Context) error {
	pomodoro := &models.Pomodoro{}
	if err := c.Bind(pomodoro); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pomodoro.StartTime = time.Now()
	pomodoro.Completed = false

	_, _, err := FirestoreClient.Collection("pomodoros").Add(c.Request().Context(), pomodoro)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, pomodoro)
}

func UpdatePomodoro(c echo.Context) error {
	id := c.Param("id")
	pomodoro := &models.Pomodoro{}
	if err := c.Bind(pomodoro); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	docRef := FirestoreClient.Collection("pomodoros").Doc(id)
	_, err := docRef.Set(c.Request().Context(), pomodoro)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return c.JSON(http.StatusNotFound, "Pomodoro session not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pomodoro)
}

func DeletePomodoro(c echo.Context) error {
	id := c.Param("id")
	_, err := FirestoreClient.Collection("pomodoros").Doc(id).Delete(c.Request().Context())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return c.JSON(http.StatusNotFound, "Pomodoro session not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
