package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	URL      string `gorm:"column:url" json:"url"`
	MimeType string `gorm:"column:mime_type" json:"mimeType"`
	Size     int64  `gorm:"column:size" json:"size"`
}
