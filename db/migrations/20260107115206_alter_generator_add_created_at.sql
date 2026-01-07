-- +goose Up
-- +goose StatementBegin
ALTER TABLE generators
ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE generators
DROP COLUMN created_at;
-- +goose StatementEnd
