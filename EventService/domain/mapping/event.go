package mapping

import (
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/dto"
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/entity"
)

func EventCreateToDto(event dto.EventCreate) dto.Event {
	return dto.Event{
		BandName:  event.BandName,
		EventCity: event.EventCity,
		Date:      event.Date,
	}
}

func EventCreateToTicketInfo(id int, event dto.EventCreate) dto.TicketInfo {
	return dto.TicketInfo{
		EventId:          id,
		EventDefault:     event.EventDefault,
		EventDefaultCost: event.EventDefaultCost,
		EventVip:         event.EventVip,
		EventVipCost:     event.EventVipCost,
		EventScene:       event.EventScene,
		EventSceneCost:   event.EventSceneCost,
	}
}

func EventToEventInfo(event entity.Event, quantity, average int) dto.EventOut {
	return dto.EventOut{
		BandName:        event.BandName,
		EventCity:       event.EventCity,
		TicketsQuantity: quantity,
		AveragePrice:    average,
		Date:            event.Date,
	}
}
