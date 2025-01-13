package db

import (
	"fmt"

	"weKnow/config"
	"weKnow/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type KnownDatabase struct {
	*gorm.DB
}

func NewDataBase(config *config.KnownConfig) *KnownDatabase {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	return &KnownDatabase{
		DB: gormDB,
	}
}

func (db *KnownDatabase) GetJobs() []model.Job {
	var jobs []model.Job
	db.Find(&jobs)
	return jobs
}

func (db *KnownDatabase) GetContacts() []model.Contact {
	var contacts []model.Contact
	db.Find(&contacts)
	return contacts
}

func (db *KnownDatabase) GetArtists() []model.Artist {
	var artists []model.Artist
	db.Find(&artists)
	return artists
}

func (db *KnownDatabase) GetArtistUuidBySlug(slug string) string {
	var artist model.Artist
	db.Where("slug = ?", slug).First(&artist)
	return artist.Uuid
}

func (db *KnownDatabase) AddArtist(artist model.Artist) error {
	return db.Create(&artist).Error
}

func (db *KnownDatabase) GetEvents() ([]model.Event, error) {
	var events []model.Event
	return events, db.Preload("Artists").Find(&events).Error
}

func (db *KnownDatabase) AddEvent(event model.Event, artists []model.Artist) error {
	err := db.Create(&event).Error
	if err != nil {
		return err
	}
	if err := db.Model(&event).Association("Artists").Append(artists); err != nil {
		return err
	}
	return nil
}

func (db *KnownDatabase) GetArtistsByIds(artistIds []int) ([]model.Artist, error) {
	var artists []model.Artist
	err := db.Where("id IN ?", artistIds).Find(&artists).Error
	return artists, err
}

func (db *KnownDatabase) GetEventById(id int) (model.Event, error) {
	var event model.Event
	err := db.Preload("Artists").First(&event, id).Error
	return event, err
}

func (db *KnownDatabase) GetNext3Events() ([]model.Event, error) {
	var events []model.Event
	return events, db.Preload("Artists").Order("date ASC").Where("date >= now()").Limit(3).Find(&events).Error
}

func (db *KnownDatabase) GetArtistEvents(slug string) ([]model.Event, error) {
	var events []model.Event
	return events, db.Joins("JOIN event_artist ON event_artist.event_id = event.id").
		Joins("JOIN artist ON artist.id = event_artist.artist_id").
		Where("artist.slug = ?", slug).
		Order("date ASC").
		Find(&events).Error
}
