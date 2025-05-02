-- name: ListInstalledPaks :many
SELECT *
FROM installed_paks
ORDER BY name;

-- name: Install :exec
INSERT INTO installed_paks (display_name, name, version, type, can_uninstall)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateVersion :exec
UPDATE installed_paks
SET version = ?
WHERE name = ?;

-- name: Uninstall :exec
DELETE
FROM installed_paks
WHERE name = ?;