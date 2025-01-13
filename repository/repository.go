package repository

import (
	"mime/multipart"
	"weKnow/adapter"
	"weKnow/db"
	"weKnow/model"
)

type KnownRepository struct {
	dataBase *db.KnownDatabase
	adapter  *adapter.KnownAdapter
}

func NewRepository(db *db.KnownDatabase) *KnownRepository {
	return &KnownRepository{
		dataBase: db,
		adapter:  adapter.NewAdapter(),
	}
}

func (r *KnownRepository) GetJobs() []model.Job {
	return r.dataBase.GetJobs()
}

func (r *KnownRepository) GetContacts() []model.Contact {
	return r.dataBase.GetContacts()
}

func (r *KnownRepository) SendEmail(email model.Email) error {
	return r.adapter.SendEmail(email)
}

func (r *KnownRepository) SendWhatsApp(recipent string, message string) error {
	return r.adapter.SendWhatsApp(recipent, message)
}

func (r *KnownRepository) GetArtists() []model.Artist {
	return r.dataBase.GetArtists()
}

func (r *KnownRepository) GetArtistImage(uuid string) (string, string, error) {
	return r.adapter.ServeImage("images/" + uuid + ".png")
}

func (r *KnownRepository) WriteFile(filePath string, file multipart.File) error {
	return r.adapter.WriteFile(filePath, file)
}

func (r *KnownRepository) GetArtistUuidBySlug(slug string) string {
	return r.dataBase.GetArtistUuidBySlug(slug)
}

func (r *KnownRepository) CreateArtist(artist model.Artist) error {
	return r.dataBase.AddArtist(artist)
}

func (r *KnownRepository) GetEvents() ([]model.Event, error) {
	return r.dataBase.GetEvents()
}

func (r *KnownRepository) AddEvent(event model.Event, artists []model.Artist) error {
	return r.dataBase.AddEvent(event, artists)
}

func (r *KnownRepository) GetArtistsByIds(artistIds []int) ([]model.Artist, error) {
	return r.dataBase.GetArtistsByIds(artistIds)
}

func (r *KnownRepository) GetEventById(id int) (model.Event, error) {
	return r.dataBase.GetEventById(id)
}

func (r *KnownRepository) GetNext3Events() ([]model.Event, error) {
	return r.dataBase.GetNext3Events()
}

func (r *KnownRepository) GetArtistEvents(slug string) ([]model.Event, error) {
	return r.dataBase.GetArtistEvents(slug)
}