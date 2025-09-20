package repository

import (
	"mime/multipart"
	"weKnow/db"
	"weKnow/model"
)

type KnownRepository struct {
	*EventRepository
	*ArtistRepository
	*UtilityRepository
	*ReleaseRepository
}

type KnownRepositoryInterface interface {
	GetJobs() []model.Job
	GetContacts() []model.Contact
	SendEmail(email model.Email) error
	SendWhatsApp(recipent string, message string) error
	GetArtists() []model.Artist
	GetArtistImage(uuid string) (string, string, error)
	WriteFile(filePath string, file multipart.File) error
}

func NewRepository(db *db.KnownDatabase) *KnownRepository {
	return &KnownRepository{
		EventRepository: &EventRepository{
			dataBase: db,
		},
		ArtistRepository: &ArtistRepository{
			dataBase: db,
		},
		ReleaseRepository: &ReleaseRepository{
			dataBase: db,
		},
		UtilityRepository: &UtilityRepository{
			dataBase: db,
		},
	}
}

// func (r *KnownRepository) GetJobs() []model.Job {
// 	return r.dataBase.GetJobs()
// }

// func (r *KnownRepository) SendEmail(email model.Email) error {
// 	return r.adapter.SendEmail(email)
// }

// func (r *KnownRepository) SendWhatsApp(recipent string, message string) error {
// 	return r.adapter.SendWhatsApp(recipent, message)
// }

// func (r *KnownRepository) WriteFile(filePath string, file multipart.File) error {
// 	return r.adapter.WriteFile(filePath, file)
// }
