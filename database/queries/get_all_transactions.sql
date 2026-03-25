-- name: GetAllTransactions :many
SELECT * FROM transactions;

-- name: GetAllUserTransactions :many
SELECT * FROM transactions WHERE user_id = $1;