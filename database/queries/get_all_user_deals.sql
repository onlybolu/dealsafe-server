-- name: GetAllUserDeals :many
SELECT * FROM deals WHERE user_id = $1;


-- name: GetUserDealsByID :one
SELECT * FROM deals WHERE id = $1;