package services

import (
	"log"

	eventDto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/events"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListEvents() ([]eventDto.EventResponse, error) {
	events, err := repositories.GetEvents()
	if err != nil {
		log.Println("Error while fetching events from the database: ", err)
		return nil, ErrEventListFailed
	}

	log.Println("Success fetching events from the database")
	return eventDto.NewEventResponseList(events), nil
}

func FindEvent(id int64) (*eventDto.EventResponse, error) {
	event, err := repositories.GetEventById(id)
	if err != nil {
		log.Println("Error while fetching event from the database: ", err)
		return nil, err
	}
	if event == nil {
		log.Println("No event with the id ", id, " was found in the database")
		return nil, ErrEventNotFound
	}

	log.Println("Success while fetching event with the provided id")
	
	result := eventDto.NewEventResponse(*event)
	return &result, nil
}
