-- +goose Up
-- +goose StatementBegin
ALTER TABLE playlist_tracks
ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE playlist_tracks
DROP COLUMN created_at;
-- +goose StatementEnd
