-- name: GetCurrencies :many
SELECT *
FROM currencies
WHERE active = TRUE
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CreateCurrency :one
INSERT INTO "currencies" ("name", "icon", "active")
VALUES(
        sqlc.arg('name'),
        sqlc.arg('icon'),
        sqlc.arg('active')
    )
RETURNING *;

-- name: UpdateCurrency :one
UPDATE currencies
SET "name" = coallesce(sqlc.narg('name'), name),
    "icon" = coallesce(sqlc.narg('icon'), icon),
    "active" = coallesce(sqlc.narg('active'), active)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteCurrency :exec
DELETE FROM currencies
WHERE id = sqlc.arg('id');

-- name: GetCurrencyByName :one
SELECT *
FROM currencies
WHERE name = sqlc.arg('name');