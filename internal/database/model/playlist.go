package model

import (
	"github.com/topvennie/spotify_organizer/pkg/sqlc"
)

type Playlist struct {
	ID            int
	UserID        int
	SpotifyID     string
	OwnerUID      string
	Name          string
	Description   string
	Public        bool
	Tracks        int
	Collaborative bool

	// Non db fields
	Owner User
}

func PlaylistModel(p sqlc.Playlist) *Playlist {
	description := ""
	if p.Description.Valid {
		description = p.Description.String
	}

	return &Playlist{
		ID:            int(p.ID),
		SpotifyID:     p.SpotifyID,
		OwnerUID:      p.OwnerUid,
		Name:          p.Name,
		Description:   description,
		Public:        p.Public,
		Tracks:        int(p.Tracks),
		Collaborative: p.Collaborative,
	}
}

func PlaylistModelPopulated(p sqlc.Playlist, u sqlc.User) *Playlist {
	playlist := PlaylistModel(p)
	playlist.Owner = *UserModel(u)

	return playlist
}

func (p *Playlist) Equal(p2 Playlist) bool {
	return p.SpotifyID == p2.SpotifyID
}

func (p *Playlist) EqualEntry(p2 Playlist) bool {
	return p.OwnerUID == p2.OwnerUID && p.Name == p2.Name && p.Description == p2.Description && p.Public == p2.Public && p.Tracks == p2.Tracks && p.Collaborative == p2.Collaborative
}
