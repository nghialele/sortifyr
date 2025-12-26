-- +goose Up
-- +goose StatementBegin
ALTER TABLE playlists
ADD COLUMN snapshot_id TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE playlists
DROP COLUMN snapshot_id;
-- +goose StatementEnd
