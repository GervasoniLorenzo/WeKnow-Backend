package db

import (
	"errors"
	"fmt"
	"strings"

	"weKnow/config"
	"weKnow/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type KnownDatabase struct {
	*gorm.DB
}

type DatabaseInterface interface {
	AddArtist(artist model.Artist) error
	AddEvent(event model.Event) error
	AdminGetEventList() ([]model.Event, error)
	CreateRelease(release model.Release) error
	DeleteArtist(id int) error
	DeleteEvent(id int) error
	DeleteRelease(id int) error
	GetArtistDetailsById(id int) (model.Artist, error)
	GetArtistDetailsBySlug(slug string) (model.Artist, error)
	GetArtistEvents(slug string) ([]model.Event, error)
	GetArtistUuidBySlug(slug string) string
	GetArtists() []model.Artist
	GetArtistsByIds(artistIds []int) ([]model.Artist, error)
	GetContacts() []model.Contact
	GetEventById(id int) (model.Event, error)
	GetImageUuidByEventSlug(slug string) (string, error)
	GetImageUuidByReleaseSlug(slug string) (string, error)
	GetImageUuidByArtistSlug(slug string) (string, error)
	GetJobs() []model.Job
	GetNext3Events() ([]model.Event, error)
	GetNextEvent() (model.Event, error)
	GetPastEvents() ([]model.Event, error)
	GetReleases() ([]model.Release, error)
	GetUpComingEvents() ([]model.Event, error)
	SlugAlreadyExist(slug string, slugEntity string) (bool, error)
	UpdateArtist(artist model.Artist) error
	UpdateEvent(event model.Event) error
	UpdateRelease(release model.Release) error
}

func AddArtistPlaceholder() {} // placeholder to avoid unused-export issue if needed (no-op)

// NewDataBase constructs the KnownDatabase
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

func (db *KnownDatabase) AddArtist(artist model.Artist) error {
	return db.Create(&artist).Error
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

func (db *KnownDatabase) AdminGetEventList() ([]model.Event, error) {
	var events []model.Event
	return events, db.
		Preload("Artists").
		Where("date >= now()").
		Order("date DESC").
		Find(&events).
		Error
}

func (db *KnownDatabase) CreateRelease(release model.Release) error {
	return db.Create(&release).Error
}

func (db *KnownDatabase) DeleteArtist(id int) error {
	return db.Delete(&model.Artist{}, id).Error
}

func (db *KnownDatabase) DeleteEvent(id int) error {
	return db.Delete(&model.Event{}, id).Error
}

func (db *KnownDatabase) DeleteRelease(id int) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&model.Release{ID: id}).Association("Artists").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("release_id = ?", id).Delete(&model.ReleaseLink{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&model.Release{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (db *KnownDatabase) GetArtistDetailsById(id int) (model.Artist, error) {
	var artist model.Artist
	err := db.
		Where("id = ?", id).
		First(&artist).Error
	return artist, err
}

func (db *KnownDatabase) GetArtistDetailsBySlug(slug string) (model.Artist, error) {
	var artist model.Artist
	err := db.
		Preload("Events", func(db *gorm.DB) *gorm.DB {
			return db.Order("date ASC")
		}).
		Preload("Releases", func(db *gorm.DB) *gorm.DB {
			return db.Order("date DESC")
		}).
		Where("slug = ?", slug).
		First(&artist).Error
	return artist, err
}

func (db *KnownDatabase) GetArtistEvents(slug string) ([]model.Event, error) {
	var events []model.Event
	return events, db.Joins("JOIN event_artist ON event_artist.event_id = event.id").
		Joins("JOIN artist ON artist.id = event_artist.artist_id").
		Where("artist.slug = ?", slug).
		Order("date ASC").
		Find(&events).Error
}

func (db *KnownDatabase) GetArtistUuidBySlug(slug string) string {
	var artist model.Artist
	db.Where("slug = ?", slug).First(&artist)
	return *artist.ImageUuid
}

func (db *KnownDatabase) GetArtists() []model.Artist {
	var artists []model.Artist
	db.Find(&artists).Order("name ASC")
	return artists
}

func (db *KnownDatabase) GetArtistsByIds(artistIds []int) ([]model.Artist, error) {
	var artists []model.Artist
	err := db.Where("id IN ?", artistIds).Find(&artists).Error
	return artists, err
}

func (db *KnownDatabase) GetContacts() []model.Contact {
	var contacts []model.Contact
	db.Find(&contacts)
	return contacts
}

func (db *KnownDatabase) GetEventById(id int) (model.Event, error) {
	var event model.Event
	err := db.Preload("Artists").First(&event, id).Error
	return event, err
}

func (db *KnownDatabase) GetImageUuidByEventSlug(slug string) (string, error) {
	var event model.Event
	err := db.
		Select("image_uuid").
		Where("slug = ?", slug).
		First(&event).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	if event.ImageUuid == nil {
		return "", errors.New("image uuid is nil")
	}
	return *event.ImageUuid, nil
}

func (db *KnownDatabase) GetImageUuidByReleaseSlug(slug string) (string, error) {
	var release model.Release
	err := db.
		Select("image_uuid").
		Where("slug = ?", slug).
		First(&release).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	if release.ImageUuid == nil {
		return "", errors.New("image uuid is nil")
	}
	return *release.ImageUuid, nil
}

func (db *KnownDatabase) GetImageUuidByArtistSlug(slug string) (string, error) {
	var artist model.Artist
	err := db.
		Select("image_uuid").
		Where("slug = ?", slug).
		First(&artist).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	if artist.ImageUuid == nil {
		return "", errors.New("image uuid is nil")
	}
	return *artist.ImageUuid, nil
}

func (db *KnownDatabase) GetJobs() []model.Job {
	var jobs []model.Job
	db.Find(&jobs)
	return jobs
}

func (db *KnownDatabase) GetNext3Events() ([]model.Event, error) {
	var events []model.Event
	return events, db.Preload("Artists").Order("date ASC").Where("date >= now()").Limit(3).Find(&events).Error
}

func (db *KnownDatabase) GetNextEvent() (model.Event, error) {
	var event model.Event
	return event, db.
		Preload("Artists").
		Where("date >= now()").Order("date ASC").First(&event).Error
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

func (db *KnownDatabase) GetReleases() ([]model.Release, error) {
	var releases []model.Release
	err := db.
		Preload("Artists").
		Preload("Links").
		Where("date::timestamptz <= now()").
		Order("date DESC").
		Limit(9).
		Find(&releases).
		Error
	if err != nil {
		return nil, err
	}
	return releases, nil
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

func (db *KnownDatabase) SlugAlreadyExist(slug string, slugEntity string) (bool, error) {

	var exist int
	var entity interface{}
	switch slugEntity {
	case "event":
		entity = &model.Event{}
	default:
		entity = &model.Artist{}
	}

	err := db.
		Model(entity).
		Select("count(id)").
		Where("slug = ?", slug).
		Scan(&exist).
		Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return exist > 0, nil
}

func (db *KnownDatabase) UpdateArtist(artist model.Artist) error {
	return db.Model(&artist).
		Omit("Events", "Events.*", "Releases", "Releases.*").
		Updates(map[string]any{
			"name":       artist.Name,
			"bio":        artist.Bio,
			"image_uuid": artist.ImageUuid,
		}).Error
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
			"name":       event.Name,
			"location":   event.Location,
			"date":       event.Date,
			"image_uuid": event.ImageUuid,
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

func (db *KnownDatabase) UpdateRelease(release model.Release) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Release{}).
			Where("id = ?", release.ID).
			Updates(map[string]any{
				"title":      release.Title,
				"date":       release.Date,
				"label":      release.Label,
				"image_uuid": release.ImageUuid,
			}).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Release{ID: release.ID}).
			Association("Artists").Replace(release.Artists); err != nil {
			return err
		}

		var existing []model.ReleaseLink
		if err := tx.Where("release_id = ?", release.ID).
			Find(&existing).Error; err != nil {
			return err
		}

		exByKey := make(map[string]model.ReleaseLink, len(existing))
		key := func(p, u string) string {
			return strings.ToLower(strings.TrimSpace(p)) + "|" + strings.TrimSpace(u)
		}
		for _, e := range existing {
			exByKey[key(e.Platform, e.URL)] = e
		}

		for i := range release.Links {
			// assicura FK
			release.Links[i].ReleaseID = release.ID

			// se ID mancante e già esiste un link con stessa chiave → usa l'ID esistente (update, non insert)
			if release.Links[i].ID == 0 {
				if ex, ok := exByKey[key(release.Links[i].Platform, release.Links[i].URL)]; ok {
					release.Links[i].ID = ex.ID
				}
			}
		}

		// 4) replace finale (aggiunge/aggiorna/rimuove)
		if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).
			Model(&model.Release{ID: release.ID}).
			Association("Links").
			Unscoped().
			Replace(release.Links); err != nil {
			return err
		}

		return nil
	})
}
