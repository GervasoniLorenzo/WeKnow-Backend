package service

import (
	"fmt"
	"weKnow/model"
	"weKnow/repository"
)

type EventServiceInterface interface {
	GetNextEvent() ([]model.Event, error)
	GetPastEvents() ([]model.Event, error)
	GetUpcomingEvents() ([]model.Event, error)
	SendEventEmail(id int) error
	GetArtistEvents(slug string) ([]model.Event, error)
}

type EventService struct {
	eventRepo   *repository.EventRepository
	artistRepo  *repository.ArtistRepository
	utilityRepo *repository.UtilityRepository
}

func NewEventService(eventRepo *repository.EventRepository, artistRepo *repository.ArtistRepository, utilityRepo *repository.UtilityRepository) *EventService {
	return &EventService{
		eventRepo:   eventRepo,
		artistRepo:  artistRepo,
		utilityRepo: utilityRepo,
	}
}

func (s *EventService) GetNextEvent() (model.EventResponseDto, error) {

	ev, err := s.eventRepo.GetNextEvent()
	if err != nil {
		return model.EventResponseDto{}, err
	}
	return formatEvent(ev), nil
}

func (s *EventService) GetPastEvents() ([]model.EventResponseDto, error) {

	event, err := s.eventRepo.GetPastEvents()
	if err != nil {
		return nil, err
	}
	return formatEvents(event), nil
}
func (s *EventService) GetUpcomingEvents() ([]model.EventResponseDto, error) {

	event, err := s.eventRepo.GetUpComingEvents()
	if err != nil {
		return nil, err
	}
	return formatEvents(event), nil
}
func (s *EventService) SendEventEmail(id int) error {
	event, err := s.eventRepo.GetEventById(id)
	if err != nil {
		return err
	}
	artistBody := "Ecco gli artisti:\n"
	for _, artist := range event.Artists {
		artistBody = artistBody + fmt.Sprintf(" - %s\n", artist.Name)
	}
	// body := fmt.Sprintf("Sei stato invitato al novo evento WeKnow %s\nTi aspettiamo a %s, il %s\n%s\n\n\n", event.Name, event.Location, event.Date, artistBody)
	contacts := s.utilityRepo.GetContacts()
	for _, contact := range contacts {
		// email := model.Email{
		// 	To:      "your-email@example.com",
		// 	Subject: "Sei stato invitato al novo evento WeKnow",
		// 	Body:    body,
		// }
		// err := s.repo.SendEmail(email)
		// if err != nil {
		// 	fmt.Println("Non ho mandato la mail a " + contact.Email)
		// 	fmt.Println(err)
		// }
		fmt.Printf("Non ho mandato la mail a %s\n", contact.Email)
	}
	return nil
}

func (s *EventService) GetArtistEvents(slug string) ([]model.Event, error) {

	return s.eventRepo.GetArtistEvents(slug)
}

func (s *EventService) AddEvent(event model.EventDto) error {
	artists, err := s.artistRepo.GetArtistsByIds(event.ArtistsId)
	if err != nil {
		return err
	}
	eventEntity := model.Event{
		Name:     event.Name,
		Location: event.Location,
		Date:     event.Date,
	}
	return s.eventRepo.AddEvent(eventEntity, artists)
}

func formatEvents(events []model.Event) []model.EventResponseDto {
	var formattedEvents []model.EventResponseDto
	for _, event := range events {
		formattedEvents = append(formattedEvents, formatEvent(event))
	}
	return formattedEvents
}

func formatEvent(e model.Event) model.EventResponseDto {
	day := e.Date.Day()
	month := model.MONTHS[int(e.Date.Month())-1]
	year := e.Date.Year()
	image := ""
	if len(e.EventImage) > 0 {
		image = e.EventImage[0].Url
	}
	return model.EventResponseDto{
		Id:       e.Id,
		Name:     e.Name,
		Location: e.Location,
		Date:     fmt.Sprintf("%d %s %d", day, month, year),
		Artists:  e.Artists,
		Image:    image,
	}
}
