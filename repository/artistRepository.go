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
	CheckArtistSlugExists(slug string) (bool, error)
	UpdateArtist(artist model.Artist) error
	DeleteArtist(id int) error
	GetArtistDetailsById(id int) (model.Artist, error)
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
func (r *ArtistRepository) CheckArtistSlugExists(slug string) (bool, error) {
	return r.dataBase.SlugAlreadyExist(slug, "artist")
}

func (r *ArtistRepository) UpdateArtist(artist model.Artist) error {
	return r.dataBase.UpdateArtist(artist)
}

func (r *ArtistRepository) DeleteArtist(id int) error {
	return r.dataBase.DeleteArtist(id)
}

func (r *ArtistRepository) GetArtistDetailsById(id int) (model.Artist, error) {
	return r.dataBase.GetArtistDetailsById(id)
}
