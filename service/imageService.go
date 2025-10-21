package service

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"weKnow/repository"

	"github.com/google/uuid"
)

type ImageService struct {
	ir repository.ImageRepositoryInterface
}

type ImageServiceInterface interface {
	CreateImage(handler *multipart.FileHeader, file multipart.File, entity string) (string, error)
	GetEventImage(slug string) (string, string, error)
	GetArtistsImage(id string) (string, string, error)
	GetReleaseImage(id string) (string, string, error)
}

func NewImageService(img repository.ImageRepositoryInterface) ImageServiceInterface {
	return &ImageService{ir: img}
}

func (s *ImageService) CreateImage(handler *multipart.FileHeader, file multipart.File, entity string) (string, error) {

	mimeType := ""
	uuID := uuid.New().String()
	parts := strings.Split(handler.Filename, ".")
	if len(parts) <= 1 {
		return "", fmt.Errorf("formato file non corretto")
	}
	mimeType = parts[len(parts)-1]

	uniqueFileName := fmt.Sprintf("%s.%s", uuID, mimeType)
	uploadDir := fmt.Sprintf("./images/%s/", entity)
	os.MkdirAll(uploadDir, os.ModePerm)
	filePath := filepath.Join(uploadDir, uniqueFileName)

	err := s.ir.WriteFile(filePath, file)
	fmt.Println(filePath)
	if err != nil {
		return "", err
	}

	return uuID, nil
}

func (s *ImageService) GetEventImage(slug string) (string, string, error) {
	imageUuid, err := s.ir.GetImageUuidByEventSlug(slug)
	if err != nil {
		return "", "", err
	}
	return s.ir.GetImageBySlugDimensionAndType(imageUuid, "event")
}
func (s *ImageService) GetArtistsImage(id string) (string, string, error) {

	return s.ir.GetImageBySlugDimensionAndType(id, "artist")
}
func (s *ImageService) GetReleaseImage(id string) (string, string, error) {

	return s.ir.GetImageBySlugDimensionAndType(id, "release")
}
