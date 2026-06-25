package services

import (
	"context"
	"log"

	eventDto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/events"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListEvents(ctx context.Context) ([]eventDto.EventResponse, error) {
	events, err := repositories.GetEvents(ctx)
	if err != nil {
		log.Println("Error while fetching events from the database: ", err)
		return nil, ErrEventListFailed
	}

	log.Println("Success fetching events from the database")
	return eventDto.NewEventResponseList(events), nil
}

func FindEvent(ctx context.Context, id int64) (*eventDto.EventResponse, error) {
	event, err := repositories.GetEventById(ctx, id)
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

func CreateEvent(ctx context.Context, event eventDto.CreateEventRequest) error {
	affected, err := repositories.CreateEvent(ctx, event.ToModel())
	if err != nil {
		log.Println("Error while creating event in the database: ", err)
		return ErrCreateEventFailed
	}
	if affected == 0 {
		log.Println("Failed to register event")
		return ErrCreateEventFailed
	}

	log.Println("Success creating event in the database")
	return nil
}

func DeleteEvent(ctx context.Context, id int64) error {
	affected, err := repositories.DeleteEvent(ctx, id)
	if err != nil {
		log.Println("Error while deleting event in the database: ", err)
		return ErrDeleteEventFailed
	}
	if affected == 0 {
		log.Println("Failed to delete event in the database")
		return ErrDeleteEventFailed
	}
	
	log.Println("Event deleted successfully")
	return nil
}