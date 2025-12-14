-- +goose Up
-- +goose StatementBegin
CREATE TABLE track_artists (
  id SERIAL PRIMARY KEY,
  artist_id INTEGER NOT NULL REFERENCES artists (id) ON DELETE CASCADE,
  track_id INTEGER NOT NULL REFERENCES tracks (id) ON DELETE CASCADE
);

CREATE TABLE album_artists (
  id SERIAL PRIMARY KEY,
  artist_id INTEGER NOT NULL REFERENCES artists (id) ON DELETE CASCADE,
  album_id INTEGER NOT NULL REFERENCES albums (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE album_artists;

DROP TABLe track_artists;
-- +goose StatementEnd
