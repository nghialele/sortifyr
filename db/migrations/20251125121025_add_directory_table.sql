-- +goose Up
-- +goose StatementBegin
CREATE TABLE directories (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users (id),
  name TEXT NOT NULL,
  parent_id INTEGER REFERENCES directories (id) ON DELETE CASCADE
);

CREATE TABLE directory_playlists (
  id SERIAL PRIMARY KEY,
  directory_id INTEGER NOT NULL REFERENCES directories (id) ON DELETE CASCADE,
  playlist_id INTEGER NOT NULL REFERENCES playlists (id) ON DELETE CASCADE,

  UNIQUE (directory_id, playlist_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE directory_playlists;

DROP TABLE directories;
-- +goose StatementEnd
