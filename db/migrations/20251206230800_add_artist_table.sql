-- +goose Up
-- +goose StatementBegin
CREATE TABLE artists (
  id SERIAL PRIMARY KEY,
  spotify_id TEXT NOT NULL,
  name TEXT NOT NULL,
  followers INTEGER NOT NULL,
  popularity INTEGER NOT NULL,

  UNIQUE (spotify_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE artists;
-- +goose StatementEnd
