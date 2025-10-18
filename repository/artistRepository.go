package repository

import (
	"weKnow/adapter"
	"weKnow/db"
	"weKnow/model"
)

type ArtistRepository struct {
	dataBase db.DatabaseInterface
	adapter  adapter.AdapterInterface
}

type ArtistRepositoryInterface interface {
	CreateArtist(artist model.Artist) error
	GetArtistsByIds(artistIds []int) ([]model.Artist, error)
	GetArtistUuidBySlug(slug string) string
	GetArtists() []model.Artist
	GetArtistImage(uuid string) (string, string, error)
	GetArtistDetailsBySlug(artistSlug string) (model.Artist, error)
}

func NewArtistRepository(db db.DatabaseInterface, a adapter.AdapterInterface) ArtistRepositoryInterface {
	return &ArtistRepository{
		dataBase: db,
		adapter:  a,
	}
}

func (r *ArtistRepository) CreateArtist(artist model.Artist) error {
	return r.dataBase.AddArtist(artist)
}

func (r *ArtistRepository) GetArtistsByIds(artistIds []int) ([]model.Artist, error) {
	return r.dataBase.GetArtistsByIds(artistIds)
}
func (r *ArtistRepository) GetArtistUuidBySlug(slug string) string {
	return r.dataBase.GetArtistUuidBySlug(slug)
}
func (r *ArtistRepository) GetArtists() []model.Artist {
	return r.dataBase.GetArtists()
}

func (r *ArtistRepository) GetArtistImage(uuid string) (string, string, error) {
	return r.adapter.ServeImage("images/" + uuid + ".png")
}

func (r *ArtistRepository) GetArtistDetailsBySlug(artistSlug string) (model.Artist, error) {
	return r.dataBase.GetArtistDetailsBySlug(artistSlug)
}
