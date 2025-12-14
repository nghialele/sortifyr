package repository

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func toString(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: s != ""}
}

func toInt(i int) pgtype.Int4 {
	return pgtype.Int4{Int32: int32(i), Valid: i != 0}
}

func toBool(b bool) pgtype.Bool {
	return pgtype.Bool{Bool: b, Valid: true}
}

func toTime(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t, Valid: !t.IsZero()}
}
