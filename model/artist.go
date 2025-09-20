package model

type Artist struct {
	Id       int       `gorm:"id" json:"id"`
	Name     string    `gorm:"name" json:"name"`
	Slug     string    `gorm:"slug" json:"slug"`
	Uuid     string    `gorm:"uuid" json:"uuid"`
	Events   []Event   `gorm:"many2many:event_artist;" json:"events"`
	Releases []Release `gorm:"many2many:release_artist;" json:"releases"`
}

func (Artist) TableName() string {
	return "artist"
}

type ArtistBasicInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
