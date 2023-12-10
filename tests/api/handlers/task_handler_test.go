package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/takahiromitsui/go-tasks/api/handlers"
	"github.com/takahiromitsui/go-tasks/api/models"
)

func TestGetTasks(t *testing.T) {
	router := gin.Default()
	router.GET("/api/tasks", handlers.GetTasks)
	req, err := http.NewRequest("GET", "/api/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response []models.Task
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), "")
	}

	if len(response) > 0 {
		firstTitle := response[0].Title
		if firstTitle != handlers.DummyTasks[0].Title {
			t.Errorf("Expected task title %s, but got %s", handlers.DummyTasks[0].Title, firstTitle)
		}
	} else {
		t.Error("No tasks found in the response")
	}
}
