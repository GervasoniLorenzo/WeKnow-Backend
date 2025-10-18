package repository

import (
	"mime/multipart"
	"weKnow/adapter"
)

type ImageRepositoryInterface interface {
	// UploadImage(image string) error
	// GetImageByUuid(uuid string) (string, error)
	// GetImageByEventId(id int) (string, error)
	WriteFile(filePath string, file multipart.File) error
	GetImageBySlugDimensionAndType(entityId string, ImageType string, dimension string) (string, string, error)
}

type ImageRepository struct {
	a adapter.AdapterInterface
}

func NewImageRepository(a adapter.AdapterInterface) ImageRepositoryInterface {
	return &ImageRepository{
		a: a,
	}
}

func (i ImageRepository) WriteFile(filePath string, file multipart.File) error {
	return i.a.WriteFile(filePath, file)
}

func (i ImageRepository) GetImageBySlugDimensionAndType(entityId string, ImageType string, dimension string) (string, string, error) {
	return i.a.GetImageBySlugDimensionAndType(entityId, ImageType, dimension)
}
