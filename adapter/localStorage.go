package adapter

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func (a *KnownAdapter) ServeImage(path string) (string, string, error) {

	// found := false
	// for _, image := range imagePath {
	// 	i, err := os.Stat(image)
	// 	fmt.Println(i)
	// 	if os.IsExist(err) {
	// 		found = true
	// 		path = image
	// 		break
	// 	}
	// }

	// if !found {
	// 	fmt.Println("File immagine non trovato")
	// 	return "", "", fmt.Errorf("file immagine non trovato")
	// }
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Errore nell'apertura del file immagine")
		return "", "", err
	}
	defer file.Close()

	// Determina il tipo MIME in base all'estensione del file
	var mimeType string
	switch filepath.Ext(path) {
	case ".jpg", ".jpeg":
		mimeType = "image/jpeg"
	case ".png":
		mimeType = "image/png"
	default:
		fmt.Println("Formato immagine non supportato")
		return "", "", err
	}

	return path, mimeType, nil
}

func (a *KnownAdapter) CreateImage(image string) (string, string, error) {
	return "", "", nil
}

func (a *KnownAdapter) WriteFile(filePath string, file multipart.File) error {
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	return nil
}

func (a *KnownAdapter) GetImageBySlugDimensionAndType(entityId string, ImageType string) (string, string, error) {
	extensions := []string{".png", ".jpg", ".jpeg"}
	var path string
	var file *os.File
	var err error

	for _, ext := range extensions {
		path = fmt.Sprintf("./images/%s/%s%s", ImageType, entityId, ext)
		file, err = os.Open(path)
		if err == nil {
			break
		}
	}

	if err != nil {
		fmt.Println("Errore nell'apertura del file immagine:", err)
		return "", "", fmt.Errorf("image not found for entityId %s with type %s", entityId, ImageType)
	}
	defer file.Close()

	var mimeType string
	switch filepath.Ext(path) {
	case ".jpg", ".jpeg":
		mimeType = "image/jpeg"
	case ".png":
		mimeType = "image/png"
	default:
		fmt.Println("Formato immagine non supportato:", filepath.Ext(path))
		return "", "", fmt.Errorf("unsupported image format: %s", filepath.Ext(path))
	}

	return path, mimeType, nil
}


