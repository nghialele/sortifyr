-- +goose Up
-- +goose StatementBegin
ALTER TABLE tracks
DROP CONSTRAINT tracks_spotify_id_key;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tracks
ADD CONSTRAINT tracks_spotify_id_key UNIQUE (spotify_id);
-- +goose StatementEnd
