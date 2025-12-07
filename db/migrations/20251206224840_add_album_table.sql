-- +goose Up
-- +goose StatementBegin
CREATE TABLE albums (
  id SERIAL PRIMARY KEY,
  spotify_id TEXt NOT NULL,
  name TEXT NOT NULL,
  track_amount INTEGER NOT NULL,
  popularity INTEGER NOT NULL,

  UNIQUE(spotify_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE albums;
-- +goose StatementEnd
