-- name: CreateDeal :one
INSERT INTO deals (
    user_id,
    title,
    description,
    amount,
    category,
    seller_email,
    seller_id,
    buyer_email,
    buyer_id,
    role,
    currency,
    stage,
    fee_payer,
    inspection_period,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
) RETURNING *;