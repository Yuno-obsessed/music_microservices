package event_test

import (
	"bytes"
	"encoding/json"
	"event-service/domain/entity"
	"event-service/infra/server/handlers/event"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWatchEvent(t *testing.T) {
	router := gin.Default()
	eventInfo := entity.Event{
		EventId:  uuid.New().String(),
		BandName: "Eluveitie",
		EventCity: entity.City{
			Id: uuid.New().String(),
			CountryId: entity.Country{
				Id:          uuid.New().String(),
				CountryName: "Italy",
			},
			CityName: "Rome",
		},
		Date: time.Date(2023, 11, 23, 21, 00, 00, 0, time.UTC),
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
		if events[i].EventCity.CityName == "Rome" {
			matches++
		}
	}
	if matches == 0 {
		t.Errorf("No requested city in response")
	}
}
