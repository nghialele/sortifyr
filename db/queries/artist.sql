-- name: ArtistGetBySpotify :one
SELECT *
FROM artists
WHERE spotify_id = $1;

-- name: ArtistCreate :one
INSERT INTO artists (spotify_id, name, followers, popularity)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: ArtistUpdate :exec
UPDATE artists
SET name = $2, followers = $3, popularity = $4
WHERE id = $1;
