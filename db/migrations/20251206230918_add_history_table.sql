-- +goose Up
-- +goose StatementBegin
CREATE TABLE history (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  track_id INTEGER NOT NULL REFERENCES tracks (id) ON DELETE CASCADE,
  played_at TIMESTAMPTZ NOT NULL,
  album_id INTEGER REFERENCES albums (id),
  artist_id INTEGER REFERENCES artists (id),
  playlist_id INTEGER REFERENCES playlists (id),
  show_id INTEGER REFERENCES shows (id),

  CONSTRAINT history_exactly_one_source CHECK (
    (
      (album_id IS NOT NULL)::int +
      (artist_id IS NOT NULL)::int +
      (playlist_id IS NOT NULL)::int +
      (show_id IS NOT NULL)::int
    ) <= 1
  )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE history;
-- +goose StatementEnd
