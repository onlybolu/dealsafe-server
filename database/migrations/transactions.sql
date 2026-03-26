CREATE TYPE tx_status AS ENUM ('pending', 'funded', 'released', 'cancelled');
CREATE TYPE tx_role AS ENUM ('buyer', 'seller');
CREATE TYPE tx_stage AS ENUM ('terms_accepted', 'buyer_funded', 'seller_delivers', 'inspection', 'inspection_approved', 'inspection_rejected', 'dispute', 'dispute_resolved', 'completed');

CREATE TABLE IF NOT EXISTS milestones (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255),
    percentage DECIMAL(10, 2),
    due_date TIMESTAMP,
    amount DECIMAL(10, 2),
    status tx_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    transaction_id VARCHAR(255),
    role tx_role,
    item_category VARCHAR(255),
    buyer_email VARCHAR(255),
    seller_email VARCHAR(255),
    amount DECIMAL(10, 2),
    name VARCHAR(255),
    currency VARCHAR(255),
    stage tx_stage,
    milestone_id INTEGER REFERENCES milestones(id) ON DELETE CASCADE,
    description TEXT,
    inspection_period INTEGER,
    status tx_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);