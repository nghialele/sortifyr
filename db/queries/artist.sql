-- name: ArtistGetAll :many
SELECT *
FROM artists;

-- name: ArtistGetBySpotify :one
SELECT *
FROM artists
WHERE spotify_id = $1;

-- name: ArtistGetByAlbum :many
SELECT a.*
FROM artists a
LEFT JOIN album_artists a_a ON a_a.artist_id = a.id
WHERE a_a.album_id = $1;

-- name: ArtistGetByTrack :many
SELECT a.*
FROM artists a
LEFT JOIN track_artists t_a ON t_a.artist_id = a.id
WHERE t_a.track_id = $1;

-- name: ArtistCreate :one
INSERT INTO artists (spotify_id, name, followers, popularity, cover_id, cover_url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: ArtistUpdate :exec
UPDATE artists
SET
  name = coalesce(sqlc.narg('name'), name),
  followers = coalesce(sqlc.narg('followers'), followers),
  popularity = coalesce(sqlc.narg('popularity'), popularity),
  cover_id = coalesce(sqlc.narg('cover_id'), cover_id),
  cover_url = coalesce(sqlc.narg('cover_url'), cover_url),
  updated_at = NOW()
WHERE id = $1;
