package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/takahiromitsui/go-tasks/api/models"
)

var dummyTasks = []models.Task {
	{
		ID:          uuid.New(),
		UserID:      uuid.New(),
		Title:       "Complete Project Proposal",
		Description: "Write a detailed proposal for the upcoming project.",
		Category:    "Work",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, 
	{
		ID:          uuid.New(),
		UserID:      uuid.New(),
		Title:       "Go for a Run",
		Description: "Run for at least 5 kilometers in the park.",
		Category:    "Personal",
		Completed:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},{
		ID:          uuid.New(),
		UserID:      uuid.New(),
		Title:       "Read Go Programming Book",
		Description: "Read the 'Programming in Go' book for at least 1 hour.",
		Category:    "Personal",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}


func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, dummyTasks)
}