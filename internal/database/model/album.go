package model

import "github.com/topvennie/sortifyr/pkg/sqlc"

type Album struct {
	ID          int
	SpotifyID   string
	Name        string
	TrackAmount int
	Popularity  int
}

func AlbumModel(a sqlc.Album) *Album {
	return &Album{
		ID:          int(a.ID),
		SpotifyID:   a.SpotifyID,
		Name:        a.Name,
		TrackAmount: int(a.TrackAmount),
		Popularity:  int(a.Popularity),
	}
}

func (a *Album) EqualEntry(a2 Album) bool {
	return a.Name == a2.Name && a.TrackAmount == a2.TrackAmount && a.Popularity == a2.Popularity
}
