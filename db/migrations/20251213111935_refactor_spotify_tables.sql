-- +goose Up
-- +goose StatementBegin

-- album

ALTER TABLE albums
ALTER COLUMN name DROP NOT NULL;

ALTER TABLE albums
ALTER COLUMN track_amount DROP NOT NULL;

ALTER TABLE albums
ALTER COLUMN popularity DROP NOT NULL;

ALTER TABLE albums
ADD COLUMN updated_at TIMESTAMPTZ;

-- artist

ALTER TABLE artists
ALTER COLUMN name DROP NOT NULL;

ALTER TABLE artists
ALTER COLUMN followers DROP NOT NULL;

ALTER TABLE artists
ALTER COLUMN popularity DROP NOT NULL;

ALTER TABLE artists
ADD COLUMN updated_at TIMESTAMPTZ;

-- playlist

ALTER TABLE playlists
ADD COLUMN owner_id INTEGER REFERENCES users (id);

ALTER TABLE playlists
DROP COLUMN owner_uid;

ALTER TABLE playlists
ALTER COLUMN name DROP NOT NULL;

ALTER TABLE playlists
ALTER COLUMN public DROP NOT NULL;

ALTER TABLE playlists
ALTER COLUMN track_amount DROP NOT NULL;

ALTER TABLE playlists
ALTER COLUMN collaborative DROP NOT NULL;

ALTER TABLE playlists
ADD COLUMN updated_at TIMESTAMPTZ;

-- show

ALTER TABLE shows
ALTER COLUMN name DROP NOT NULL;

ALTER TABLE shows
ALTER COLUMN episode_amount DROP NOT NULL;

ALTER TABLE shows
ADD COLUMN updated_at TIMESTAMPTZ;

-- track

ALTER TABLE tracks
ALTER COLUMN name DROP NOT NULL;

ALTER TABLE tracks
ALTER COLUMN popularity DROP NOT NULL;

ALTER TABLE tracks
ADD COLUMN updated_at TIMESTAMPTZ;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- shows

ALTER TABLE tracks
DROP COLUMN updated_at TIMESTAMPTZ;

ALTER TABLE tracks
ALTER COLUMN popularity SET NOT NULL;

ALTER TABLE tracks
ALTER COLUMN name SET NOT NULL;

-- shows

ALTER TABLE shows
DROP COLUMN updated_at TIMESTAMPTZ;

ALTER TABLE shows
ALTER COLUMN episode_amount SET NOT NULL;

ALTER TABLE shows
ALTER COLUMN name SET NOT NULL;

-- playlists

ALTER TABLE playlists
DROP COLUMN updated_at;

ALTER TABLE playlists
ALTER COLUMN collaborative SET NOT NULL;

ALTER TABLE playlists
ALTER COLUMN track_amount SET NOT NULL;

ALTER TABLE playlists
ALTER COLUMN public SET NOT NULL;

ALTER TABLE playlists
ALTER COLUMN name SET NOT NULL;

ALTER TABLE playlists
ADD COLUMN owner_uid TEXT REFERENCES users (uid);

ALTER TABLE playlists
DROP COLUMN owner_id;

-- artists

ALTER TABLE artists
DROP COLUMN updated_at TIMESTAMPTZ;

ALTER TABLE artists
ALTER COLUMN popularity SET NOT NULL;

ALTER TABLE artists
ALTER COLUMN followers SET NOT NULL;

ALTER TABLE artists
ALTER COLUMN name SET NOT NULL;

-- albums

ALTER TABLE albums
DROP COLUMN updated_at TIMESTAMPTZ;

ALTER TABLE albums
ALTER COLUMN popularity SET NOT NULL;

ALTER TABLE albums
ALTER COLUMN track_amount SET NOT NULL;

ALTER TABLE albums
ALTER COLUMN name SET NOT NULL;

-- +goose StatementEnd
