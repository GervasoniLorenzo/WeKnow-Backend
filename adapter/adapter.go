package adapter

import (
	"fmt"
	"mime/multipart"

	"weKnow/config"
)

type KnownAdapter struct {
	config *config.KnownConfig
}

type AdapterInterface interface {
	ServeImage(path string) (string, string, error)
	CreateImage(image string) (string, string, error)
	WriteFile(filePath string, file multipart.File) error
	GetImageBySlugDimensionAndType(entityId string, ImageType string) (string, string, error)
}

func NewAdapter() AdapterInterface {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
	}
	return &KnownAdapter{
		config: c,
	}
}
