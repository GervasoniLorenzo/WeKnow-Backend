package model

type Artist struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"name" json:"name"`
	Slug      string    `gorm:"slug" json:"slug"`
	ImageUuid *string   `gorm:"image_uuid" json:"uuid"`
	Bio       string    `gorm:"bio" json:"bio"`
	Events    []Event   `gorm:"many2many:event_artist;" json:"events"`
	Releases  []Release `gorm:"many2many:release_artist;" json:"releases"`
}

type ArtistDto struct {
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	ImageUuid string `json:"image_uuid"`
}

func (Artist) TableName() string {
	return "artist"
}

type ArtistBasicInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
