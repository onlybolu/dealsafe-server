-- name: GetAllTransactions :many
SELECT * FROM transactions;

-- name: GetAllUserTransactions :many
SELECT * FROM transactions WHERE user_id = $1;

-- name: GetTransactionByID :one
SELECT * FROM transactions WHERE id = $1;