package db

import (
	"errors"
	"fmt"

	"weKnow/config"
	"weKnow/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type KnownDatabase struct {
	*gorm.DB
}

type DatabaseInterface interface {
	GetJobs() []model.Job
	GetContacts() []model.Contact

	//Artists
	GetArtists() []model.Artist
	GetArtistUuidBySlug(slug string) string
	AddArtist(artist model.Artist) error
	GetArtistsByIds(artistIds []int) ([]model.Artist, error)
	GetArtistEvents(slug string) ([]model.Event, error)
	GetArtistDetailsBySlug(slug string) (model.Artist, error)

	//Events
	GetNextEvent() (model.Event, error)
	AddEvent(event model.Event) error
	GetEventById(id int) (model.Event, error)
	GetNext3Events() ([]model.Event, error)
	GetPastEvents() ([]model.Event, error)
	GetUpComingEvents() ([]model.Event, error)
	AdminGetEventList() ([]model.Event, error)
	DeleteEvent(id int) error
	UpdateEvent(event model.Event) error
	EventSlugAlreadyExist(slug string) (bool, error)

	GetReleases() ([]model.Release, error)
}

func NewDataBase(config *config.KnownConfig) DatabaseInterface {

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

	if err := gormDB.AutoMigrate(&model.Event{}, &model.Artist{}, &model.EventArtist{}); err != nil {
		panic(fmt.Errorf("failed to migrate database: %w", err))
	}
	if err = gormDB.SetupJoinTable(&model.Event{}, "Artists", &model.EventArtist{}); err != nil {
		panic(fmt.Errorf("failed to setup join table: %w", err))
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

func (db *KnownDatabase) GetNextEvent() (model.Event, error) {
	var event model.Event
	return event, db.
		Preload("Artists").
		Where("date >= now()").Order("date ASC").First(&event).Error
}

func (db *KnownDatabase) AddEvent(event model.Event) error {
	err := db.Create(&event).Error
	if err != nil {
		return err
	}
	// if err := db.Model(&event).Association("Artists").Append(artists); err != nil {
	// 	return err
	// }
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

func (db *KnownDatabase) GetPastEvents() ([]model.Event, error) {
	var events []model.Event
	return events, db.
		Preload("Artists").
		Order("date ASC").
		Where("date < now()").
		Find(&events).
		Error
}

func (db *KnownDatabase) GetUpComingEvents() ([]model.Event, error) {
	var events []model.Event
	return events, db.
		Preload("Artists").
		Order("date ASC").
		Where("date >= now()").
		Offset(1).
		Find(&events).
		Error
}

func (db *KnownDatabase) GetReleases() ([]model.Release, error) {
	var releases []model.Release
	err := db.
		Preload("Artist").
		Preload("Links").
		Where("release_date <= now()").
		Order("release_date DESC").
		Limit(9).
		Find(&releases).
		Error
	if err != nil {
		return nil, err
	}
	return releases, nil
}

func (db *KnownDatabase) GetArtistDetailsBySlug(slug string) (model.Artist, error) {
	var artist model.Artist
	err := db.
		Preload("Events", func(db *gorm.DB) *gorm.DB {
			return db.Order("date ASC")
		}).
		Preload("Releases", func(db *gorm.DB) *gorm.DB {
			return db.Order("release_date DESC")
		}).
		Where("slug = ?", slug).
		First(&artist).Error
	return artist, err
}

func (db *KnownDatabase) AdminGetEventList() ([]model.Event, error) {
	var events []model.Event
	return events, db.
		Preload("Artists").
		Where("date >= now()").
		Order("date DESC").
		Find(&events).
		Error
}

func (db *KnownDatabase) DeleteEvent(id int) error {
	return db.Delete(&model.Event{}, id).Error
}

func (db *KnownDatabase) UpdateEvent(event model.Event) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	var persisted model.Event
	if err := tx.First(&persisted, "id = ?", event.Id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&persisted).
		Omit("Artists", "Artists.*").
		Updates(map[string]any{
			"name":     event.Name,
			"location": event.Location,
			"date":     event.Date, // *time.Time ok
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&persisted).Association("Artists").Replace(event.Artists); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (db *KnownDatabase) EventSlugAlreadyExist(slug string) (bool, error) {
	var exist int
	err := db.
		Model(&model.Event{}).
		Select("count(id)").
		Where("slug = ?", slug).
		Scan(&exist).
		Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return exist > 0, nil
}
