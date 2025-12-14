package model

import (
	"time"

	"github.com/topvennie/sortifyr/pkg/sqlc"
)

type Track struct {
	ID         int
	SpotifyID  string
	Name       string
	Popularity int
	UpdatedAt  time.Time

	// Non db fields
	Artists []Artist
}

func TrackModel(t sqlc.Track) *Track {
	return &Track{
		ID:         int(t.ID),
		SpotifyID:  t.SpotifyID,
		Name:       fromString(t.Name),
		Popularity: fromInt(t.Popularity),
		UpdatedAt:  fromTime(t.UpdatedAt),
	}
}

func (t *Track) Equal(t2 Track) bool {
	return t.SpotifyID == t2.SpotifyID
}

func (t *Track) EqualEntry(t2 Track) bool {
	return t.Name == t2.Name && t.Popularity == t2.Popularity
}

type TrackArtist struct {
	ID       int
	TrackID  int
	ArtistID int
}
