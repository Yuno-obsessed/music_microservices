package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"

	"projects/music_microservices/StorageService/domain/dto"
	mocks "projects/music_microservices/StorageService/infra/mocks/ticket"
	"projects/music_microservices/StorageService/infra/server/handlers"
)

func TestTicket_GetEntity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicket := mocks.NewMockTicketInterface(ctrl)
	ginCtx, _ := gin.CreateTestContext(httptest.NewRecorder())

	expected := dto.TicketOut{
		DefaultQuantity: 300,
		VipQuantity:     100,
		SceneQuantity:   150,
		DefaultCost:     40,
		VipCost:         60,
		SceneCost:       90,
	}

	mockTicket.EXPECT().GetEntity(gomock.Any()).Return(expected)

	ticket := &handlers.Ticket{
		*mockTicket,
		logger.NewLogger(),
	}
	ticket.GetEntity(ginCtx)

	assert.Equal(t, http.StatusOK, ginCtx.Writer.Status())
	var responseDto dto.TicketOut
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&responseDto)
	if err != nil {
		t.Errorf("expected: %v\ngot:%v", expected, responseDto)
	}
	assert.Equal(t, expected, responseDto)
}
