-- name: TrackArtistCreate :one
INSERT INTO track_artists (track_id, artist_id)
VALUES ($1, $2)
RETURNING id;

-- name: TrackArtistDeleteByArtistTrack :exec
DELETE FROM track_artists
WHERE artist_id = $1 AND track_id = $2;

