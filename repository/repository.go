package repository

import (
	"weKnow/adapter"
	"weKnow/db"
)

type KnownRepository struct {
	ArtistRepositoryInterface
	EventRepositoryInterface
	UtilityRepositoryInterface
	ReleaseRepositoryInterface
	ImageRepositoryInterface
}

type KnownRepositoryInterface interface {
	ArtistRepositoryInterface
	EventRepositoryInterface
	UtilityRepositoryInterface
	ReleaseRepositoryInterface
	ImageRepositoryInterface
}

func NewRepository(db db.DatabaseInterface, a adapter.AdapterInterface) KnownRepositoryInterface {
	return KnownRepository{
		EventRepositoryInterface:   NewEventRepository(db, nil),
		ArtistRepositoryInterface:  NewArtistRepository(db, a),
		ReleaseRepositoryInterface: NewReleaseRepository(db),
		UtilityRepositoryInterface: NewUtilityRepository(db),
		ImageRepositoryInterface:   NewImageRepository(a, db),
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
