-- +goose Up
-- +goose StatementBegin
ALTER TABLE playlists
ADD COLUMN cover_id TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE playlists
DROP COLUMN cover_id;
-- +goose StatementEnd
