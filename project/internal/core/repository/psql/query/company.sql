-- name: CreateCompany :one
INSERT INTO company (company_name)
VALUES ($1) RETURNING *;

-- name: GetCompany :one
SELECT *
FROM company
WHERE id = sqlc.arg('id');

-- name: GetAllCompanies :many
SELECT *
FROM company
WHERE TRUE
    AND CASE
        WHEN @company_name::VARCHAR = '' THEN TRUE
        ELSE company_name = @company_name
    END
ORDER BY created_at DESC OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetAllCompaniesCount :one
SELECT COUNT(*)
FROM company
WHERE TRUE
    AND CASE
        WHEN @company_name::VARCHAR = '' THEN TRUE
        ELSE company_name = @company_name
    END;

-- name: UpdateCompany :exec
UPDATE company
SET company_name = coallesce(sqlc.narg('company_name'), company_name),
    updated_at = now()
WHERE id = sqlc.arg('id');

-- name: DeleteCompany :exec
DELETE FROM company
WHERE id = sqlc.arg('id');

-- name: SoftDeleteCompany :exec
UPDATE company
SET deleted_at = now()
WHERE id = sqlc.arg('id');