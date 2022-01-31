-- name: CreateAccount :one
INSERT INTO accounts (
  owner,balance,currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: FindAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: FindAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;