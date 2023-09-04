-- name: CreateBankAccount :one
INSERT INTO bank_account (
        bank_account_name,
        bic_code,
        iban_code,
        swift_code,
        region,
        registration_number,
        address,
        bank_address,
        phone_number,
        payment_reference,
        description,
        salesforce_id,
        owner_name,
        transactions_reference,
        receiving_party_bank_details,
        currency
    )
VALUES (
        sqlc.arg('bank_account_name'),
        sqlc.arg('bic_code'),
        sqlc.arg('iban_code'),
        sqlc.arg('swift_code'),
        sqlc.arg('region'),
        sqlc.arg('registration_number'),
        sqlc.arg('address'),
        sqlc.arg('bank_address'),
        sqlc.arg('phone_number'),
        sqlc.arg('payment_reference'),
        sqlc.arg('description'),
        sqlc.arg('salesforce_id'),
        sqlc.arg('owner_name'),
        sqlc.arg('transactions_reference'),
        sqlc.arg('receiving_party_bank_details'),
        sqlc.arg('currency')
    )
RETURNING *;

-- name: GetBankAccount :one
SELECT *
FROM bank_account
WHERE id = sqlc.arg('id');

-- name: GetAllBankAccounts :many
SELECT ba.*,
    c.name AS currency_name
FROM bank_account AS ba
    LEFT JOIN currencies AS c ON ba.currency = c.id
WHERE ba.bank_account_name ilike '%' || sqlc.arg('bank_account_name') || '%'
    AND CASE
        WHEN @region::VARCHAR = '' THEN TRUE
        ELSE ba.region = @region
    END
    AND CASE
        WHEN @registration_number::VARCHAR = '' THEN TRUE
        ELSE ba.registration_number = @registration_number
    END
    AND CASE
        WHEN @salesforce_id::VARCHAR = '' THEN TRUE
        ELSE ba.salesforce_id = @salesforce_id
    END
ORDER BY ba.created_at DESC OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetBankAccountsCount :one
SELECT COUNT(*) AS COUNT
FROM bank_account
WHERE bank_account_name ilike '%' || sqlc.arg('bank_account_name') || '%'
    AND CASE
        WHEN @region::VARCHAR = '' THEN TRUE
        ELSE region = @region
    END
    AND CASE
        WHEN @registration_number::VARCHAR = '' THEN TRUE
        ELSE registration_number = @registration_number
    END
    AND CASE
        WHEN @salesforce_id::VARCHAR = '' THEN TRUE
        ELSE salesforce_id = @salesforce_id
    END;

-- name: GetBankAccountById :one
SELECT ba.*,
    c.name AS currency_name
FROM bank_account AS ba
    LEFT JOIN currencies AS c ON ba.currency = c.id
WHERE ba.id = sqlc.arg('id');

-- name: UpdateBankAccount :exec
UPDATE bank_account
SET bic_code = COALESCE(sqlc.narg('bic_code'), bic_code),
    iban_code = COALESCE(sqlc.narg('iban_code'), iban_code),
    swift_code = COALESCE(sqlc.narg('swift_code'), swift_code),
    region = COALESCE(sqlc.narg('region'), region),
    address = COALESCE(sqlc.narg('address'), address),
    bank_address = COALESCE(sqlc.narg('bank_address'), bank_address),
    phone_number = COALESCE(sqlc.narg('phone_number'), phone_number),
    description = COALESCE(sqlc.narg('description'), description),
    salesforce_id = COALESCE(sqlc.narg('salesforce_id'), salesforce_id),
    owner_name = COALESCE(sqlc.narg('owner_name'), owner_name),
    currency = COALESCE(sqlc.narg('currency'), currency),
    bank_account_name = COALESCE(
        sqlc.narg('bank_account_name'),
        bank_account_name
    ),
    registration_number = COALESCE(
        sqlc.narg('registration_number'),
        registration_number
    ),
    payment_reference = COALESCE(
        sqlc.narg('payment_reference'),
        payment_reference
    ),
    transactions_reference = COALESCE(
        sqlc.narg('transactions_reference'),
        transactions_reference
    ),
    receiving_party_bank_details = COALESCE(
        sqlc.narg('receiving_party_bank_details'),
        receiving_party_bank_details
    ),
    updated_at = NOW()
WHERE id = sqlc.arg('id');

-- name: SetSalesforceId :exec
UPDATE bank_account
SET salesforce_id = sqlc.arg('salesforce_id')
WHERE id = sqlc.arg('id');

-- name: DeleteBankAccount :exec
DELETE FROM bank_account
WHERE id = sqlc.arg('id');