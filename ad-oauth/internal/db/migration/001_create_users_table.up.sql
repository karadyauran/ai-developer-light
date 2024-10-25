SET TIMEZONE = 'UTC';

-- creating the users table
CREATE TABLE IF NOT EXISTS users (
                                     id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
                                     github_id BIGINT UNIQUE NOT NULL,
                                     username VARCHAR(255) NOT NULL,
                                     avatar_url VARCHAR(255),
                                     email VARCHAR(255),
                                     token VARCHAR(255) NOT NULL,
                                     created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
                                     updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

-- Creating a unique index on 'github_id'
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_github_id ON users (github_id);

-- Creating a trigger function for automatic updating 'updated_at'
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to call a function before updating a record
CREATE TRIGGER trigger_update_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();