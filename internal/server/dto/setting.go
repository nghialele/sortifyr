package dto

import (
	"time"

	"github.com/topvennie/spotify_organizer/internal/database/model"
)

type Setting struct {
	ID         int       `json:"id"`
	LastUpdate time.Time `json:"last_update,omitzero"`
}

func SettingDTO(s *model.Setting) Setting {
	return Setting{
		ID:         s.ID,
		LastUpdate: s.LastUpdate,
	}
}
