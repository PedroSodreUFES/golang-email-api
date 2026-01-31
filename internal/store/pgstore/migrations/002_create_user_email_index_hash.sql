-- Write your migrate up statements here
CREATE UNIQUE INDEX idx_users_email_hash
    ON users(email);
---- create above / drop below ----
DROP INDEX idx_users_email_hash;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.