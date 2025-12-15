package model

import (
	"time"

	"github.com/topvennie/sortifyr/pkg/sqlc"
)

type Album struct {
	ID          int
	SpotifyID   string
	Name        string
	TrackAmount int
	Popularity  int
	CoverID     string
	CoverURL    string
	UpdatedAt   time.Time

	// Non db fields
	Artists []Artist
}

func AlbumModel(a sqlc.Album) *Album {
	return &Album{
		ID:          int(a.ID),
		SpotifyID:   a.SpotifyID,
		Name:        fromString(a.Name),
		TrackAmount: fromInt(a.TrackAmount),
		Popularity:  fromInt(a.Popularity),
		CoverID:     fromString(a.CoverID),
		CoverURL:    fromString(a.CoverUrl),
		UpdatedAt:   fromTime(a.UpdatedAt),
	}
}

func (a *Album) Equal(a2 Album) bool {
	return a.SpotifyID == a2.SpotifyID
}

func (a *Album) EqualEntry(a2 Album) bool {
	return a.Name == a2.Name && a.TrackAmount == a2.TrackAmount && a.Popularity == a2.Popularity && a.CoverURL != a2.CoverURL
}

type AlbumUser struct {
	ID        int
	UserID    int
	AlbumID   int
	DeletedAt time.Time
}

type AlbumArtist struct {
	ID       int
	AlbumID  int
	ArtistID int
}
