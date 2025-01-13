package model

type Artist struct {
	Id     int     `gorm:"id" json:"id"`
	Name   string  `gorm:"name" json:"name"`
	Slug   string  `gorm:"slug" json:"slug"`
	Uuid   string  `gorm:"uuid" json:"uuid"`
	Events []Event `gorm:"many2many:event_artist;"`
}

func (Artist) TableName() string {
	return "artist"
}
