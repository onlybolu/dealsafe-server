-- name: CreateUser :one
INSERT INTO users (email, password, test_pub_key, test_priv_key, terms_accepted)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, created_at, test_pub_key, test_priv_key, updated_at;