-- name: UpdateUserDetails :one
UPDATE users SET
    first_name = $1,
    last_name = $2,
    phone_number = $3,
    address = $4,
    city = $5,
    state = $6,
    zip_code = $7,
    country = $8,
    company_name = $9,
    updated_at = $10
WHERE id = $11
RETURNING id, first_name, email, last_name, phone_number, address, city, state, zip_code, country, company_name, updated_at;
