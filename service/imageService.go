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
	CreateImage(handler *multipart.FileHeader, file multipart.File) (string, error)
	GetEventImage(slug string, Imagetype string) (string, string, error)
	GetArtistsImage(id string, Imagetype string) (string, string, error)
	GetReleaseImage(id string, Imagetype string) (string, string, error)
}

func NewImageService(img repository.ImageRepositoryInterface) ImageServiceInterface {
	return &ImageService{ir: img}
}

func (s *ImageService) CreateImage(handler *multipart.FileHeader, file multipart.File) (string, error) {

	mimeType := ""
	uuID := uuid.New().String()
	parts := strings.Split(handler.Filename, ".")
	if len(parts) > 1 {
		mimeType = parts[len(parts)-1]
	} else {
		return "", fmt.Errorf("Formato file non corretto")
	}

	uniqueFileName := fmt.Sprintf("%s.%s", uuID, mimeType)
	uploadDir := "./images"
	os.MkdirAll(uploadDir, os.ModePerm)

	filePath := filepath.Join(uploadDir, uniqueFileName)

	err := s.ir.WriteFile(filePath, file)
	fmt.Println(filePath)
	if err != nil {
		return "", err
	}

	return uuID, nil
}

func (s *ImageService) GetEventImage(slug string, Imagetype string) (string, string, error) {

	return s.ir.GetImageBySlugDimensionAndType(slug, "event", Imagetype)
}
func (s *ImageService) GetArtistsImage(id string, Imagetype string) (string, string, error) {

	return s.ir.GetImageBySlugDimensionAndType(id, "artist", Imagetype)
}
func (s *ImageService) GetReleaseImage(id string, Imagetype string) (string, string, error) {

	return s.ir.GetImageBySlugDimensionAndType(id, "release", Imagetype)
}
