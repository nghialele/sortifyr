-- name: AlbumGetAll :many
SELECT *
FROM albums;

-- name: AlbumGetBySpotify :one
SELECT *
FROM albums
WHERE spotify_id = $1;

-- name: AlbumGetByUser :many
SELECT a.*
FROM albums a
LEFT JOIN album_users au on au.album_id = a.id
WHERE au.user_id = $1;

-- name: AlbumCreate :one
INSERT INTO albums (spotify_id, name, track_amount, popularity, cover_id, cover_url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: AlbumUpdate :exec
UPDATE albums
SET 
  name = coalesce(sqlc.narg('name'), name),
  track_amount = coalesce(sqlc.narg('track_amount'), track_amount),
  popularity = coalesce(sqlc.narg('popularity'), popularity),
  cover_id = coalesce(sqlc.narg('cover_id'), cover_id),
  cover_url = coalesce(sqlc.narg('cover_url'), cover_url),
  updated_at = NOW()
WHERE id = $1;
