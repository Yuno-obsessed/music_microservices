package event_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Yuno-obsessed/music_microservices/EventService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/EventService/infra/server/handlers/event"
	"github.com/gin-gonic/gin"
)

func TestWatchEvent(t *testing.T) {
	router := gin.Default()
	eventInfo := entity.Event{
		EventId:   1,
		BandName:  "Eluveitie",
		EventCity: "2",
		Date:      time.Date(2023, 11, 23, 21, 0o0, 0o0, 0, time.UTC),
	}
	mock, err := json.Marshal(eventInfo)
	if err != nil {
		t.Errorf("error marshalling mocks, %v", err)
	}

	// Add your handler to the router
	router.GET("/api/v1/event/watch/:city", event.EventInfo)

	// Create a request to send to the handler
	req, _ := http.NewRequest("GET", "http://localhost:8084/api/v1/event/watch/Rome", bytes.NewBuffer(mock))
	w := httptest.NewRecorder()

	// Call the handler through the Gin engine
	router.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code: %d", w.Code)
	}

	// Check the response body
	var events []entity.Event
	err = json.NewDecoder(w.Body).Decode(&events)
	if err != nil {
		t.Errorf("Error parsing payload, %v", err)
	}
	matches := 0
	for i := 0; i < len(events); i++ {
		if events[i].EventCity == "2" {
			matches++
		}
	}
	if matches == 0 {
		t.Errorf("No requested city in response")
	}
}
