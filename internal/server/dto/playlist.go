package dto

import (
	"github.com/topvennie/spotify_organizer/internal/database/model"
)

type Playlist struct {
	ID            int    `json:"id" validate:"required"`
	SpotifyID     string `json:"spotify_id"`
	Owner         User   `json:"owner,omitzero"`
	Name          string `json:"name"`
	Description   string `json:"description,omitzero"`
	Public        bool   `json:"public"`
	Tracks        int    `json:"tracks"`
	Collaborative bool   `json:"collaborative"`
}

func PlaylistDTO(playlist *model.Playlist, user *model.User) Playlist {
	return Playlist{
		ID:            playlist.ID,
		SpotifyID:     playlist.SpotifyID,
		Owner:         UserDTO(user),
		Name:          playlist.Name,
		Description:   playlist.Description,
		Public:        playlist.Public,
		Tracks:        playlist.Tracks,
		Collaborative: playlist.Collaborative,
	}
}

func (p Playlist) ToModel(userID int) *model.Playlist {
	return &model.Playlist{
		ID:            p.ID,
		UserID:        userID,
		SpotifyID:     p.SpotifyID,
		OwnerUID:      p.Owner.UID,
		Name:          p.Name,
		Description:   p.Description,
		Public:        p.Public,
		Tracks:        p.Tracks,
		Collaborative: p.Collaborative,
		Owner:         *p.Owner.ToModel(),
	}
}
