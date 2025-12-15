-- +goose Up
-- +goose StatementBegin
ALTER TABLE album_users
ADD COLUMN deleted_at TIMESTAMPTZ;

ALTER TABLE playlist_tracks
ADD COLUMN deleted_at TIMESTAMPTZ;

ALTER TABLE playlist_users
ADD COLUMN deleted_at TIMESTAMPTZ;

ALTER TABLE show_users
ADD COLUMN deleted_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE show_users
DROP COLUMN deleted_at;

ALTER TABLE playlist_users
DROP COLUMN deleted_at;

ALTER TABLE playlist_tracks
DROP COLUMN deleted_at;

ALTER TABLE album_users
DROP COLUMN deleted_at;
-- +goose StatementEnd
