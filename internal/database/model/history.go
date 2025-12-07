package model

import (
	"time"

	"github.com/topvennie/sortifyr/pkg/sqlc"
)

type History struct {
	ID         int
	UserID     int
	TrackID    int
	PlayedAt   time.Time
	AlbumID    int
	ArtistID   int
	PlaylistID int
	ShowID     int
}

func HistoryModel(h sqlc.History) *History {
	albumID := 0
	if h.AlbumID.Valid {
		albumID = int(h.AlbumID.Int32)
	}
	artistID := 0
	if h.ArtistID.Valid {
		artistID = int(h.ArtistID.Int32)
	}
	playlistID := 0
	if h.PlaylistID.Valid {
		playlistID = int(h.PlaylistID.Int32)
	}
	showID := 0
	if h.ShowID.Valid {
		showID = int(h.ShowID.Int32)
	}

	return &History{
		UserID:     int(h.UserID),
		ID:         int(h.ID),
		TrackID:    int(h.TrackID),
		PlayedAt:   h.PlayedAt.Time,
		AlbumID:    albumID,
		ArtistID:   artistID,
		PlaylistID: playlistID,
		ShowID:     showID,
	}
}
