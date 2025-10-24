package repository

import (
	"weKnow/adapter"
	"weKnow/db"
	"weKnow/model"
)

type EventRepository struct {
	dataBase db.DatabaseInterface
	adapter  adapter.AdapterInterface
}

type EventRepositoryInterface interface {
	GetNextEvent() (model.Event, error)
	GetPastEvents() ([]model.Event, error)
	GetUpComingEvents() ([]model.Event, error)
	GetEventById(id int) (model.Event, error)
	GetArtistEvents(slug string) ([]model.Event, error)
	// ADMIN
	AdminAddEvent(event model.Event) error
	AdminGetEventList() ([]model.Event, error)
	AdminDeleteEvent(id int) error
	AdminUpdateEvent(event model.Event) error
	CheckEventSlugExists(slug string) (bool, error)
}

func NewEventRepository(db db.DatabaseInterface, adapter adapter.AdapterInterface) EventRepositoryInterface {
	return &EventRepository{
		dataBase: db,
		adapter:  adapter,
	}
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
func (r *EventRepository) AddEvent(event model.Event) error {
	return r.dataBase.AddEvent(event)
}
func (r *EventRepository) AdminGetEventList() ([]model.Event, error) {
	return r.dataBase.AdminGetEventList()
}

func (r *EventRepository) AdminAddEvent(event model.Event) error {
	return r.dataBase.AddEvent(event)
}

func (r *EventRepository) AdminDeleteEvent(id int) error {
	return r.dataBase.DeleteEvent(id)
}

func (r *EventRepository) AdminUpdateEvent(event model.Event) error {
	return r.dataBase.UpdateEvent(event)
}

func (r *EventRepository) CheckEventSlugExists(slug string) (bool, error) {
	return r.dataBase.SlugAlreadyExist(slug, "event")
}
