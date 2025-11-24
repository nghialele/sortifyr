package model

import (
	"time"

	"github.com/topvennie/spotify_organizer/pkg/sqlc"
)

type Setting struct {
	ID         int
	UserID     int
	LastUpdate time.Time
}

func SettingModel(s sqlc.Setting) *Setting {
	lastUpdate := time.Time{}
	if s.LastUpdated.Valid {
		lastUpdate = s.LastUpdated.Time
	}

	return &Setting{
		ID:         int(s.ID),
		UserID:     int(s.UserID),
		LastUpdate: lastUpdate,
	}
}
