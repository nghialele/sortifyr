-- name: PlaylistGet :one
SELECT *
FROM playlists
WHERE id = $1;

-- name: PlaylistGetBySpotify :one
SELECT *
FROM playlists
WHERE spotify_id = $1;

-- name: PlaylistGetByUserWithOwner :many
SELECT sqlc.embed(p), sqlc.embed(u)
FROM playlists p
LEFT JOIN playlist_users pu ON pu.playlist_id = p.id
LEFT JOIN users u ON u.id = p.owner_id
WHERE pu.user_id = $1 AND pu.deleted_at IS NULL
ORDER BY p.name;

-- name: PlaylistCreate :one
INSERT INTO playlists (spotify_id, owner_id, name, description, public, track_amount, collaborative, cover_id, cover_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;

-- name: PlaylistUpdateBySpotify :exec
UPDATE playlists
SET 
  owner_id = coalesce(sqlc.narg('owner_id'), owner_id),
  name = coalesce(sqlc.narg('name'), name),
  description = coalesce(sqlc.narg('description'), description),
  public = coalesce(sqlc.narg('public'), public),
  track_amount = coalesce(sqlc.narg('track_amount'), track_amount),
  collaborative = coalesce(sqlc.narg('collaborative'), collaborative),
  cover_id = coalesce(sqlc.narg('cover_id'), cover_id),
  cover_url = coalesce(sqlc.narg('cover_url'), cover_url),
  updated_at = NOW()
WHERE spotify_id = $1;

