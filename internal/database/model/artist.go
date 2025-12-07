package model

import "github.com/topvennie/sortifyr/pkg/sqlc"

type Artist struct {
	ID         int
	SpotifyID  string
	Name       string
	Followers  int
	Popularity int
}

func ArtistModel(a sqlc.Artist) *Artist {
	return &Artist{
		ID:         int(a.ID),
		SpotifyID:  a.SpotifyID,
		Name:       a.Name,
		Followers:  int(a.Followers),
		Popularity: int(a.Popularity),
	}
}

func (a *Artist) EqualEntry(a2 Artist) bool {
	return a.Name == a2.Name && a.Followers == a2.Followers && a.Popularity == a2.Popularity
}
