package repository

type ImageRepositoryInterface interface {
	UploadImage(image string) error
	GetImageByUuid(uuid string) (string, error)
	GetImageByEventId(id int) (string, error)
}

type ImageRepository struct {
}

func NewImageRepository() *ImageRepository {
	return &ImageRepository{}
}
