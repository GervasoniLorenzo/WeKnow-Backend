package model

type Release struct {
	ID          string        `json:"id" gorm:"primaryKey"`
	Slug        string        `json:"slug" gorm:"column:slug;not null;unique"`
	Title       string        `json:"title" gorm:"column:name;not null"`
	ReleaseDate string        `json:"release_date" gorm:"column:release_date;not null"`
	Links       []ReleaseLink `json:"links" gorm:"foreignKey:ReleaseID;references:ID"`
	Artist      []Artist      `json:"artist" gorm:"many2many:release_artist;"`
}

type ReleaseDto struct {
	Title       string        `json:"title"`
	ReleaseDate string        `json:"release_date"`
	Links       []ReleaseLink `json:"links"`
	ArtistIds   []int         `json:"artist_ids"`
}

type ReleaseLink struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Platform  string `json:"platform" gorm:"column:platform;not null"`
	URL       string `json:"url" gorm:"column:url;not null"`
	ReleaseID string `json:"release_id" gorm:"column:release_id;not null"`
}

func (Release) TableName() string {
	return "release"
}

func (ReleaseLink) TableName() string {
	return "release_link"
}

type ReleaseResponseDto struct {
	ID          string        `json:"id"`
	Slug        string        `json:"slug"`
	Title       string        `json:"title"`
	ReleaseDate string        `json:"release_date"`
	Links       []ReleaseLink `json:"links"`
	Artists     string        `json:"artists"`
}

func FormatRelease(release Release) ReleaseResponseDto {
	artists := ""
	for _, artist := range release.Artist {
		artists += artist.Name + ", "
	}
	if len(artists) > 0 {
		artists = artists[:len(artists)-2] // Remove trailing comma and space
	}

	return ReleaseResponseDto{
		ID:          release.ID,
		Slug:        release.Slug,
		Title:       release.Title,
		ReleaseDate: release.ReleaseDate,
		Links:       release.Links,
		Artists:     artists,
	}
}
