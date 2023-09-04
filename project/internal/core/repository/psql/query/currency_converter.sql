-- name: GetCurrencyConvertor :many
SELECT *
FROM currency_convertor
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetCurrencyConvertorByCurrency :one
SELECT *
FROM currency_convertor
WHERE primary_currency = sqlc.arg('primary_currency')
    AND secondary_currency = sqlc.arg('secondary_currency');

-- name: UpdateCurrencyConvertor :one
INSERT INTO currency_convertor (
        "id",
        "primary_currency",
        "secondary_currency",
        "amount"
    )
VALUES (
        sqlc.arg('id'),
        sqlc.arg('primary_currency'),
        sqlc.arg('secondary_currency'),
        sqlc.arg('amount')
    ) ON CONFLICT ON CONSTRAINT uq_primary_currency_secondary_currency DO
UPDATE
SET amount = sqlc.arg('amount')
RETURNING "id";