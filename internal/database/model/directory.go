package model

import (
	"slices"

	"github.com/topvennie/spotify_organizer/pkg/sqlc"
	"github.com/topvennie/spotify_organizer/pkg/utils"
)

type Directory struct {
	ID       int
	UserID   int
	Name     string
	ParentID int

	// Non db fields
	Playlists []Playlist
}

func DirectoryModel(d sqlc.Directory) *Directory {
	parentID := 0
	if d.ParentID.Valid {
		parentID = int(d.ParentID.Int32)
	}

	return &Directory{
		ID:       int(d.ID),
		UserID:   int(d.UserID),
		Name:     d.Name,
		ParentID: parentID,
	}
}

func (d *Directory) Equal(d2 Directory) bool {
	values := d.UserID == d2.UserID && d.Name == d2.Name && d.ParentID == d2.ParentID
	if !values {
		return false
	}

	p := utils.SliceMap(d.Playlists, func(p Playlist) int { return p.ID })
	p2 := utils.SliceMap(d2.Playlists, func(p Playlist) int { return p.ID })

	slices.Sort(p)
	slices.Sort(p2)

	return slices.Equal(p, p2)
}

type DirectoryPlaylist struct {
	ID          int
	DirectoryID int
	PlaylistID  int
}

func DirectoryPlaylistModel(d sqlc.DirectoryPlaylist) *DirectoryPlaylist {
	return &DirectoryPlaylist{
		ID:          int(d.ID),
		DirectoryID: int(d.DirectoryID),
		PlaylistID:  int(d.PlaylistID),
	}
}
