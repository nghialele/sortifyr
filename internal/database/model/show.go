package model

import "github.com/topvennie/sortifyr/pkg/sqlc"

type Show struct {
	ID            int
	SpotifyID     string
	Name          string
	EpisodeAmount int
}

func ShowModel(s sqlc.Show) *Show {
	return &Show{
		ID:            int(s.ID),
		SpotifyID:     s.SpotifyID,
		Name:          s.Name,
		EpisodeAmount: int(s.EpisodeAmount),
	}
}

func (s *Show) EqualEntry(s2 Show) bool {
	return s.Name == s2.Name && s.EpisodeAmount == s2.EpisodeAmount
}
