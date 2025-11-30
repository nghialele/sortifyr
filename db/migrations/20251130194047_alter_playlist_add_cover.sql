-- +goose Up
-- +goose StatementBegin
ALTER TABLE playlists
ADD COLUMN cover_url TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE playlists
DROP COLUMN cover_url;
-- +goose StatementEnd
