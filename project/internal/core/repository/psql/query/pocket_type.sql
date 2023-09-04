-- name: GetAllPocketTypes :many
SELECT *
FROM pocket_type
WHERE TRUE
    AND CASE
        WHEN @required = FALSE THEN TRUE
        ELSE required = @required
    END OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');