-- name: SettingGetByUser :one
SELECT *
FROM settings
WHERE user_id = $1;

-- name: SettingCreate :one
INSERT INTO settings (user_id)
VALUES ($1)
RETURNING id;

-- name: SettingUpdateByUser :exec
UPDATE settings
SET last_updated = $2
WHERE user_id = $1;
