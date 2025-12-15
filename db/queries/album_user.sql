-- name: AlbumUserCreate :one
INSERT INTO album_users (user_id, album_id)
VALUES ($1, $2)
RETURNING id;

-- name: AlbumUserDeleteByUserAlbum :exec
UPDATE album_users
SET deleted_at = NOW()
WHERE user_id = $1 AND album_id = $2;
