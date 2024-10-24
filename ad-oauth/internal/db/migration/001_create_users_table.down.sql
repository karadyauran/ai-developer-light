-- Removing a trigger
DROP TRIGGER IF EXISTS trigger_update_updated_at ON users;

-- Removing a trigger function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Removing an index
DROP INDEX IF EXISTS idx_users_github_id;

-- Dropping table 'users'
DROP TABLE IF EXISTS users;