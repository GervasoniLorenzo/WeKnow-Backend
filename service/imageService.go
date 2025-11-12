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
	GetArtistsImage(slug string) (string, string, error)
	GetReleaseImage(slug string) (string, string, error)
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
	imageUuid, err := s.ir.GetImageUuidBySlugAndType(slug, "event")
	if err != nil {
		return "", "", err
	}
	return s.ir.GetImageBySlugAndType(imageUuid, "event")
}
func (s *ImageService) GetArtistsImage(slug string) (string, string, error) {
	imageUuid, err := s.ir.GetImageUuidBySlugAndType(slug, "artist")
	if err != nil {
		return "", "", err
	}
	return s.ir.GetImageBySlugAndType(imageUuid, "artist")
}
func (s *ImageService) GetReleaseImage(slug string) (string, string, error) {
	imageUuid, err := s.ir.GetImageUuidBySlugAndType(slug, "release")
	if err != nil {
		return "", "", err
	}
	return s.ir.GetImageBySlugAndType(imageUuid, "release")
}
