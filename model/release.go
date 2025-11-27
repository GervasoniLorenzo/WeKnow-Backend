package model

import "time"

type Release struct {
	ID        int           `json:"id" gorm:"primaryKey"`
	Slug      string        `json:"slug" gorm:"column:slug;not null;unique"`
	Title     string        `json:"title" gorm:"column:title;not null"`
	Label     string        `json:"label" gorm:"column:label;"`
	Date      *time.Time    `json:"date" gorm:"column:date;not null"`
	Links     []ReleaseLink `json:"links" gorm:"foreignKey:ReleaseID;references:ID"`
	Artists   []Artist      `json:"artists" gorm:"many2many:release_artist;"`
	ImageUuid *string       `json:"imageUrl" gorm:"column:image_uuid;"`
}

type ReleaseDto struct {
	Title     string        `json:"title"`
	Date      *time.Time    `json:"date"`
	Label     string        `json:"label"`
	Links     []ReleaseLink `json:"links"`
	ArtistIds []int         `json:"artistIds"`
	ImageUuid string        `json:"imageUuid"`
}

type ReleaseLink struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Platform  string `json:"platform" gorm:"column:platform;not null"`
	URL       string `json:"url" gorm:"column:url;not null"`
	ReleaseID int    `json:"release_id" gorm:"column:release_id;not null"`
}

func (Release) TableName() string {
	return "release"
}

func (ReleaseLink) TableName() string {
	return "release_link"
}

type ReleaseResponseDto struct {
	ID       int           `json:"id"`
	Slug     string        `json:"slug"`
	Title    string        `json:"title"`
	Date     *time.Time    `json:"date"`
	Links    []ReleaseLink `json:"links"`
	Artists  []string      `json:"artists"`
	Label    string        `json:"label"`
	ImageUrl string        `json:"imageUrl"`
}

func FormatRelease(release Release) ReleaseResponseDto {
	artists := []string{}
	for _, artist := range release.Artists {
		artists = append(artists, artist.Name)
	}

	return ReleaseResponseDto{
		ID:      release.ID,
		Slug:    release.Slug,
		Title:   release.Title,
		Date:    release.Date,
		Links:   release.Links,
		Artists: artists,
	}
}
