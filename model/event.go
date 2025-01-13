package model

type Event struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Date     string   `json:"date"`
	Time     string   `json:"time"`
	Artists  []Artist `gorm:"many2many:event_artist;"`
}

func (e *Event) TableName() string {
	return "event"
}

type EventDto struct {
	Name      string `json:"name"`
	Location  string `json:"location"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	ArtistsId []int  `json:"artistsId"`
}
