package service

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"weKnow/model"
	"weKnow/repository"

	"github.com/google/uuid"
)

type KnownService struct {
	repo *repository.KnownRepository
}

func NewService(repo *repository.KnownRepository) *KnownService {
	return &KnownService{
		repo: repo,
	}
}

func (s *KnownService) GetJobs() []model.Job {
	return s.repo.GetJobs()
}

func (s *KnownService) SendJobEmail() {
	events, err := s.repo.GetNext3Events()
	if err == nil && len(events) > 0 {
		contacts := s.repo.GetContacts()
		eventBody := "Ecco I prossimi eventi:\n"
		for _, event := range events {
			eventBody = eventBody + fmt.Sprintf(" - %s\n", event.Name)
		}

		for _, contact := range contacts {
			email := model.Email{
				To:      contact.Email,
				Subject: "Hello",
				Body:    eventBody,
			}
			err := s.repo.SendEmail(email)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}

func (s *KnownService) SendWhatsApp() {
	contacts := s.repo.GetContacts()
	for _, c := range contacts {
		err := s.repo.SendWhatsApp(c.PhoneNumber, "YOYO TESTO PROVA")
		if err != nil {
			fmt.Print(err)
		}
	}
}

func (s *KnownService) SendEmailAndWhatsapp() {
	fmt.Println("ciao")
}

func (s *KnownService) GetArtists() []model.Artist {
	return s.repo.GetArtists()
}

func (s *KnownService) AddArtist(artist model.Artist) error {
	return s.repo.CreateArtist(artist)
}

func (s *KnownService) GetArtistImage(slug string) (string, string, error) {
	uuid := s.repo.GetArtistUuidBySlug(slug)
	return s.repo.GetArtistImage(uuid)
}

func (s *KnownService) CreateImage(handler *multipart.FileHeader, file multipart.File) (string, error) {

	mimeType := ""
	uuID := uuid.New().String()
	parts := strings.Split(handler.Filename, ".")
	if len(parts) > 1 {
		mimeType = parts[len(parts)-1]
	} else {
		return "", fmt.Errorf("Formato file non corretto")
	}

	uniqueFileName := fmt.Sprintf("%s.%s", uuID, mimeType)
	uploadDir := "./images"
	os.MkdirAll(uploadDir, os.ModePerm)

	filePath := filepath.Join(uploadDir, uniqueFileName)

	err := s.repo.WriteFile(filePath, file)

	if err != nil {
		return "", err
	}

	return uuID, nil
}

func (s *KnownService) GetEventList() ([]model.Event, error) {

	return s.repo.GetEvents()
}

func (s *KnownService) AddEvent(event model.EventDto) error {
	artists, err := s.repo.GetArtistsByIds(event.ArtistsId)
	if err != nil {
		return err
	}
	eventEntity := model.Event{
		Name:     event.Name,
		Location: event.Location,
		Date:     event.Date,
		Time:     event.Time,
	}
	return s.repo.AddEvent(eventEntity, artists)
}

func (s *KnownService) SendEventEmail(id int) error {
	event, err := s.repo.GetEventById(id)
	if err != nil {
		return err
	}
	artistBody := "Ecco gli artisti:\n"
	for _, artist := range event.Artists {
		artistBody = artistBody + fmt.Sprintf(" - %s\n", artist.Name)
	}
	body := fmt.Sprintf("Sei stato invitato al novo evento WeKnow %s\nTi aspettiamo a %s, il %s, alle ore: %s\n%s\n\n\n", event.Name, event.Location, event.Date, event.Time, artistBody)
	contacts := s.repo.GetContacts()
	for _, contact := range contacts {
		email := model.Email{
			To:      "your-email@example.com",
			Subject: "Sei stato invitato al novo evento WeKnow",
			Body:    body,
		}
		err := s.repo.SendEmail(email)
		if err != nil {
			fmt.Println("Non ho mandato la mail a " + contact.Email)
			fmt.Println(err)
		}
	}
	return nil
}

func (s *KnownService) GetArtistEvents(slug string) ([]model.Event, error) {

	return s.repo.GetArtistEvents(slug)
}
