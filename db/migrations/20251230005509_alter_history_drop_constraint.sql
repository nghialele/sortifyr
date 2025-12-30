-- +goose Up
-- +goose StatementBegin
ALTER TABLE history
DROP CONSTRAINT history_exactly_one_source;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE history
ADD CONSTRAINT history_exactly_one_source CHECK (
    (
      (album_id IS NOT NULL)::int +
      (artist_id IS NOT NULL)::int +
      (playlist_id IS NOT NULL)::int +
      (show_id IS NOT NULL)::int
    ) <= 1
);
-- +goose StatementEnd
