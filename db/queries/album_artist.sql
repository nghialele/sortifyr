-- name: AlbumArtistCreate :one
INSERT INTO album_artists (album_id, artist_id)
VALUES ($1, $2)
RETURNING id;

-- name: AlbumArtistDeleteByArtistAlbum :exec
DELETE FROM album_artists
WHERE artist_id = $1 AND album_id = $2;
