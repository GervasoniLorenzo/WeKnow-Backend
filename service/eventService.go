package service

import (
	"fmt"
	"weKnow/config"
	"weKnow/model"
	"weKnow/repository"
	"weKnow/utils"
)

type EventServiceInterface interface {
	GetNextEvent(view string) (model.EventResponseDto, error)
	GetPastEvents() ([]model.EventResponseDto, error)
	GetUpcomingEvents() ([]model.EventResponseDto, error)
	SendEventEmail(id int) error
	GetArtistEvents(slug string) ([]model.Event, error)
	AdminGetEventList() ([]model.Event, error)
	AdminDeleteEvent(id int) error
	AdminUpdateEvent(event model.UpdateEventDto) error
	AdminCreateEvent(event model.EventDto) error
}

type EventService struct {
	eventRepo   repository.EventRepositoryInterface
	artistRepo  repository.ArtistRepositoryInterface
	utilityRepo repository.UtilityRepositoryInterface
	imageRepo   repository.ImageRepositoryInterface
	cfg         config.KnownConfig
	u           utils.UtilsInterface
}

func NewEventService(er repository.EventRepositoryInterface, ar repository.ArtistRepositoryInterface, ur repository.UtilityRepositoryInterface, ir repository.ImageRepositoryInterface, cfg config.KnownConfig) EventServiceInterface {
	return &EventService{
		eventRepo:   er,
		artistRepo:  ar,
		utilityRepo: ur,
		imageRepo:   ir,
		cfg:         cfg,
		u:           utils.NewUtils(),
	}
}

func (s *EventService) GetNextEvent(view string) (model.EventResponseDto, error) {

	ev, err := s.eventRepo.GetNextEvent()
	if err != nil {
		return model.EventResponseDto{}, err
	}
	return s.formatEvent(ev, ""), nil
}

func (s *EventService) GetPastEvents() ([]model.EventResponseDto, error) {

	event, err := s.eventRepo.GetPastEvents()
	if err != nil {
		return nil, err
	}
	return s.formatEvents(event), nil
}
func (s *EventService) GetUpcomingEvents() ([]model.EventResponseDto, error) {

	event, err := s.eventRepo.GetUpComingEvents()
	if err != nil {
		return nil, err
	}
	return s.formatEvents(event), nil
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

func (s *EventService) AdminGetEventList() ([]model.Event, error) {
	return s.eventRepo.AdminGetEventList()
}

func (s *EventService) formatEvents(events []model.Event) []model.EventResponseDto {
	var formattedEvents []model.EventResponseDto
	for _, event := range events {
		formattedEvents = append(formattedEvents, s.formatEvent(event, ""))
	}
	return formattedEvents
}

func (s *EventService) formatEvent(e model.Event, view string) model.EventResponseDto {
	day := e.Date.Day()
	month := model.MONTHS[int(e.Date.Month())-1]
	year := e.Date.Year()
	size := "flyer"
	if view == "mobile" {
		size = "small"
	}
	image := fmt.Sprintf("%s/event/image/%s?size=%s", s.cfg.App.Host, e.Slug, size)
	return model.EventResponseDto{
		Id:       e.Id,
		Name:     e.Name,
		Slug:     e.Slug,
		Location: e.Location,
		Date:     fmt.Sprintf("%d %s %d", day, month, year),
		Artists:  e.Artists,
		Image:    image,
	}
}

func (s *EventService) AdminDeleteEvent(id int) error {
	return s.eventRepo.AdminDeleteEvent(id)
}

func (s *EventService) AdminUpdateEvent(event model.UpdateEventDto) error {
	artists, err := s.artistRepo.GetArtistsByIds(event.ArtistsId)
	if err != nil {
		return err
	}

	eventEntity := model.Event{
		Id:       event.Id,
		Name:     event.Name,
		Location: event.Location,
		Date:     event.Date,
		Artists:  artists,
	}

	return s.eventRepo.AdminUpdateEvent(eventEntity)
}

func (s *EventService) AdminCreateEvent(event model.EventDto) error {
	artists, err := s.artistRepo.GetArtistsByIds(event.ArtistsId)
	if err != nil {
		return err
	}

	slug := s.u.GenerateSlug(event.Name)
	// Assegna il valore di controllo a una variabile, poi ciclaci sopra
	exists, err := s.eventRepo.CheckEventSlugExists(slug)
	if err != nil {
		return err
	}
	count := 0
	for exists {
		count++
		exists, err = s.eventRepo.CheckEventSlugExists(fmt.Sprintf("%s-%v", slug, count))
		if err != nil {
			return err
		}
		if exists {
			slug = fmt.Sprintf("%s-%d", slug, count)
		}
	}

	eventEntity := model.Event{
		Name:      event.Name,
		Location:  event.Location,
		Date:      event.Date,
		Artists:   artists,
		Slug:      slug,
		ImageUuid: event.ImageUuid,
	}

	return s.eventRepo.AdminAddEvent(eventEntity)
}
