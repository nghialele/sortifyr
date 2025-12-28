-- +goose Up
-- +goose StatementBegin
ALTER TABLE album_users
DROP CONSTRAINT album_users_user_id_album_id_key;

ALTER TABLE playlist_users
DROP CONSTRAINT playlist_users_user_id_playlist_id_key;

ALTER TABLE show_users
DROP CONSTRAINT show_users_user_id_show_id_key;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE show_users
ADD CONSTRAINT show_users_user_id_show_id_key UNIQUE (user_id, show_id);

ALTER TABLE playlist_users
ADD CONSTRAINT playlist_users_user_id_playlist_id_key UNIQUE (user_id, playlist_id);

ALTER TABLE album_users
ADD CONSTRAINT album_users_user_id_album_id_key UNIQUE (user_id, album_id);
-- +goose StatementEnd
