-- +goose Up
-- +goose StatementBegin
ALTER TABLE artists
ADD COLUMN cover_url TEXT;

ALTER TABLE artists
ADD COLUMN cover_id TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE artists
DROP COLUMN cover_id;

ALTER TABLE artists
DROP COLUMN cover_url;
-- +goose StatementEnd
