-- +goose Up
-- +goose StatementBegin
CREATE TABLE shows (
  id SERIAL PRIMARY KEY,
  spotify_id TEXT NOT NULL,
  episode_amount INTEGER NOT NULL,
  name TEXT NOT NULL,

  UNIQUE(spotify_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE shows;
-- +goose StatementEnd
