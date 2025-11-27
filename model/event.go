package model

import (
	"time"
)

var MONTHS = []string{
	"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno",
	"luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre",
}

type Event struct {
	Id        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Slug      string     `gorm:"column:slug" json:"slug"`
	Location  string     `gorm:"column:location" json:"location"`
	Date      *time.Time `gorm:"column:date" json:"date"`
	Artists   []Artist   `gorm:"many2many:event_artist" json:"artists"`
	ImageUuid *string    `gorm:"column:image_uuid" json:"imageUuid"`
}

func (e *Event) TableName() string {
	return "event"
}

type EventResponseDto struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Slug     string   `json:"slug"`
	Location string   `json:"location"`
	Date     string   `json:"date"`
	Artists  []Artist `json:"artists"`
	Image    string   `json:"image"`
}

type EventDto struct {
	Name      string     `json:"name"`
	Location  string     `json:"location"`
	Date      *time.Time `json:"date"`
	ArtistsId []int      `json:"artistsIds"`
	ImageUuid *string    `json:"imageUuid,omitempty"`
}

type EventBasicDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
	Slug string `json:"slug"`
}

type UpdateEventDto struct {
	Id int `json:"id"`
	EventDto
}

type EventArtist struct {
	EventID  int `gorm:"primaryKey;autoIncrement:false"`
	ArtistID int `gorm:"primaryKey;autoIncrement:false"`
	// niente DeletedAt qui, altrimenti devi gestire unique con deleted_at
}

func (EventArtist) TableName() string { return "event_artists" }
