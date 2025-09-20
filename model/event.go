package model

import "time"

var MONTHS = []string{
	"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno",
	"luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre",
}

type Event struct {
	Id         int          `gorm:"column:id"`
	Name       string       `gorm:"column:name" json:"name"`
	Location   string       `gorm:"column:location" json:"location"`
	Date       *time.Time   `gorm:"column:date" json:"date"`
	Artists    []Artist     `gorm:"many2many:event_artist" json:"artists"`
	EventImage []EventImage `gorm:"foreignKey:EventId" json:"images"`
}

func (e *Event) TableName() string {
	return "event"
}

type EventResponseDto struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Date     string   `json:"date"`
	Artists  []Artist `json:"artists"`
	Image    string   `json:"image"`
}

type EventDto struct {
	Name      string       `json:"name"`
	Location  string       `json:"location"`
	Date      *time.Time   `json:"date"`
	Images    []EventImage `json:"imageUrl"`
	ArtistsId []int        `json:"artistsId"`
}

type EventImage struct {
	Id      int    `gorm:"column:id" json:"id"`
	EventId int    `gorm:"column:event_id"`
	Url     string `gorm:"column:url" json:"url"`
	Type    string `gorm:"column:type" json:"type"`
}

func (e *EventImage) TableName() string {
	return "event_image"
}
