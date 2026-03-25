-- name: GetAllUserDeals :many
SELECT * FROM deals WHERE user_id = $1;
