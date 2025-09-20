package repository

import (
	"weKnow/adapter"
	"weKnow/db"
	"weKnow/model"
)

type EventRepositoryInterface interface {
	GetNextEvent() (model.Event, error)
	GetPastEvents() ([]model.Event, error)
	GetUpComingEvents() ([]model.Event, error)
	GetEventById(id int) (model.Event, error)
	GetNext3Events() ([]model.Event, error)
	GetArtistEvents(slug string) ([]model.Event, error)
	AddEvent(event model.Event, artists []model.Artist) error
}

type EventRepository struct {
	dataBase *db.KnownDatabase
	adapter  *adapter.KnownAdapter
}

func (r *EventRepository) GetNextEvent() (model.Event, error) {
	return r.dataBase.GetNextEvent()
}
func (r *EventRepository) GetPastEvents() ([]model.Event, error) {
	return r.dataBase.GetPastEvents()
}
func (r *EventRepository) GetUpComingEvents() ([]model.Event, error) {
	return r.dataBase.GetUpComingEvents()
}
func (r *EventRepository) GetEventById(id int) (model.Event, error) {
	return r.dataBase.GetEventById(id)
}

func (r *EventRepository) GetArtistEvents(slug string) ([]model.Event, error) {
	return r.dataBase.GetArtistEvents(slug)
}
func (r *EventRepository) AddEvent(event model.Event, artists []model.Artist) error {
	return r.dataBase.AddEvent(event, artists)
}
