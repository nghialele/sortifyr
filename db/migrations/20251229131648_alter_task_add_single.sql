-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks
ADD COLUMN recurring BOOLEAN NOT NULL DEFAULT true;

ALTER TABLE tasks
ALTER COLUMN recurring DROP DEFAULT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tasks
DROP COLUMN recurring;
-- +goose StatementEnd
