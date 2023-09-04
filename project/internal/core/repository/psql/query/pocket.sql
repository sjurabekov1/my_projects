-- name: CreatePocket :copyfrom 
INSERT INTO pocket (
        "name",
        "title",
        "user_id",
        "company_id",
        "type",
        "icon",
        "color",
        "salesforce_id"
    )
VALUES (
        sqlc.arg('name'),
        sqlc.arg('title'),
        sqlc.arg('user_id'),
        sqlc.arg('company_id'),
        sqlc.arg('type'),
        sqlc.arg('icon'),
        sqlc.arg('color'),
        sqlc.arg('salesforce_id')
    );

-- name: GetPocketsByUserID :many
SELECT p.*,
    pt.name AS "pocket_type_name",
    c.company_name AS "company_name"
FROM pocket p
    JOIN pocket_type pt ON p.type = pt.id
    JOIN company c ON p.company_id = c.id
WHERE user_id = sqlc.arg('user_id');