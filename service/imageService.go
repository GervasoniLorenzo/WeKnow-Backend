package service

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"weKnow/adapter"

	"github.com/google/uuid"
)

type ImageService struct {
	adapter *adapter.KnownAdapter
}

func NewImageService() *ImageService {
	return &ImageService{}
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

	err := s.adapter.WriteFile(filePath, file)
	fmt.Println(filePath)
	if err != nil {
		return "", err
	}

	return uuID, nil
}

func (s *ImageService) GetEventImage(id string, Imagetype string) (string, string, error) {

	return s.adapter.GetImageByIdDimensionAndType(id, "event", Imagetype)
}
func (s *ImageService) GetArtistsImage(id string, Imagetype string) (string, string, error) {

	return s.adapter.GetImageByIdDimensionAndType(id, "artist", Imagetype)
}
func (s *ImageService) GetReleaseImage(id string, Imagetype string) (string, string, error) {

	return s.adapter.GetImageByIdDimensionAndType(id, "release", Imagetype)
}
