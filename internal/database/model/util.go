package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func fromString(s pgtype.Text) string {
	result := ""
	if s.Valid {
		result = s.String
	}

	return result
}

func fromInt(i pgtype.Int4) int {
	result := 0
	if i.Valid {
		result = int(i.Int32)
	}

	return result
}

func fromBool(b pgtype.Bool) bool {
	result := false
	if b.Valid {
		result = b.Bool
	}

	return result
}

func fromTime(t pgtype.Timestamptz) time.Time {
	result := time.Time{}
	if t.Valid {
		result = t.Time
	}

	return result
}
