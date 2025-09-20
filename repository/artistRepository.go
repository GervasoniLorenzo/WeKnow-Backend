package repository

import (
	"weKnow/adapter"
	"weKnow/db"
	"weKnow/model"
)

type ArtistRepository struct {
	dataBase *db.KnownDatabase
	adapter  *adapter.KnownAdapter
}

func NewArtistRepository() *ArtistRepository {
	return &ArtistRepository{}
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