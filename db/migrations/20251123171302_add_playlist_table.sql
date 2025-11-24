-- +goose Up
-- +goose StatementBegin
CREATE TABLE playlists (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users (id),
  spotify_id TEXT NOT NULL,
  owner_uid TEXT NOT NULL REFERENCES users (uid),
  name TEXT NOT NULL,
  description TEXT,
  public BOOLEAN NOT NULL,
  tracks INTEGER NOT NULL,
  collaborative BOOLEAN NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE playlists;
-- +goose StatementEnd
