package model

import (
	"time"

	"github.com/topvennie/sortifyr/pkg/sqlc"
)

type Show struct {
	ID            int
	SpotifyID     string
	Name          string
	EpisodeAmount int
	CoverID       string
	CoverURL      string
	UpdatedAt     time.Time
}

func ShowModel(s sqlc.Show) *Show {
	return &Show{
		ID:            int(s.ID),
		SpotifyID:     s.SpotifyID,
		Name:          fromString(s.Name),
		EpisodeAmount: fromInt(s.EpisodeAmount),
		CoverID:       fromString(s.CoverID),
		CoverURL:      fromString(s.CoverUrl),
		UpdatedAt:     fromTime(s.UpdatedAt),
	}
}

func (s *Show) Equal(s2 Show) bool {
	return s.SpotifyID == s2.SpotifyID
}

func (s *Show) EqualEntry(s2 Show) bool {
	return s.Name == s2.Name && s.EpisodeAmount == s2.EpisodeAmount && s.CoverURL == s2.CoverURL
}

type ShowUser struct {
	ID     int
	UserID int
	ShowID int
}
