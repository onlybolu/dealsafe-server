CREATE TYPE kyc_status AS ENUM ('pending', 'approved', 'rejected');

CREATE TABLE IF NOT EXISTS users(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255),
    address VARCHAR(255),
    city VARCHAR(255),
    company_name VARCHAR(255),
    state VARCHAR(255),
    zip_code VARCHAR(255),
    terms_accepted BOOLEAN DEFAULT FALSE,
    country VARCHAR(255),
    test_pub_key VARCHAR(255),
    test_priv_key VARCHAR(255),
    live_pub_key VARCHAR(255),
    live_priv_key VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_kyc (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    id_type VARCHAR(50),
    id_number VARCHAR(255),
    id_url VARCHAR(500),
    status kyc_status DEFAULT 'pending',
    rejection_reason TEXT,
    cac_number VARCHAR(255),
    cac_url VARCHAR(500),
    verified_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_bank_details (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_number VARCHAR(255),
    account_name VARCHAR(255),
    bank_name VARCHAR(255),
    bank_code VARCHAR(255),
    is_default BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

