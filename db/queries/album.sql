-- name: AlbumGetBySpotify :one
SELECT *
FROM albums
WHERE spotify_id = $1;

-- name: AlbumCreate :one
INSERT INTO albums (spotify_id, name, track_amount, popularity)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: AlbumUpdate :exec
UPDATE albums
SET name = $2, track_amount = $3, popularity = $4
WHERE id = $1;
