package model

import (
	"time"

	"github.com/topvennie/sortifyr/pkg/sqlc"
)

type Artist struct {
	ID         int
	SpotifyID  string
	Name       string
	Followers  int
	Popularity int
	CoverID    string
	CoverURL   string
	UpdatedAt  time.Time
}

func ArtistModel(a sqlc.Artist) *Artist {
	return &Artist{
		ID:         int(a.ID),
		SpotifyID:  a.SpotifyID,
		Name:       fromString(a.Name),
		Followers:  fromInt(a.Followers),
		Popularity: fromInt(a.Popularity),
		CoverID:    fromString(a.CoverID),
		CoverURL:   fromString(a.CoverUrl),
		UpdatedAt:  fromTime(a.UpdatedAt),
	}
}

func (a *Artist) Equal(a2 Artist) bool {
	return a.SpotifyID == a2.SpotifyID
}

func (a *Artist) EqualEntry(a2 Artist) bool {
	return a.Name == a2.Name && a.Followers == a2.Followers && a.Popularity == a2.Popularity && a.CoverURL == a2.CoverURL
}
