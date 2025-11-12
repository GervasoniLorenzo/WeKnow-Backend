package repository

import (
	"mime/multipart"
	"weKnow/adapter"
	"weKnow/db"
)

type ImageRepositoryInterface interface {
	// UploadImage(image string) error
	// GetImageByUuid(uuid string) (string, error)
	// GetImageByEventId(id int) (string, error)
	WriteFile(filePath string, file multipart.File) error
	GetImageBySlugAndType(entityId string, ImageType string) (string, string, error)
	GetImageUuidBySlugAndType(slug string, entity string) (string, error)
}

type ImageRepository struct {
	a  adapter.AdapterInterface
	db db.DatabaseInterface
}

func NewImageRepository(a adapter.AdapterInterface, db db.DatabaseInterface) ImageRepositoryInterface {
	return &ImageRepository{
		a:  a,
		db: db,
	}
}

func (i ImageRepository) WriteFile(filePath string, file multipart.File) error {
	return i.a.WriteFile(filePath, file)
}

func (i ImageRepository) GetImageBySlugAndType(entityId string, ImageType string) (string, string, error) {
	return i.a.GetImageBySlugDimensionAndType(entityId, ImageType)
}

func (i ImageRepository) GetImageUuidBySlugAndType(slug string, entity string) (string, error) {
	switch entity {
	case "event":
		return i.db.GetImageUuidByEventSlug(slug)
	case "release":
		return i.db.GetImageUuidByReleaseSlug(slug)
	case "artist":
		return i.db.GetImageUuidByArtistSlug(slug)
	}
	return "", nil
}
