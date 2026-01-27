-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(80) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    profile_picture VARCHAR(300) DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS emails(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    wasSeen BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT now(),
    id_receiver INT NOT NULL,
    id_sender INT NOT NULL,
    CONSTRAINT fk_id_receiver
        FOREIGN KEY(id_receiver) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_id_sender
        FOREIGN KEY(id_sender) REFERENCES users(id) ON DELETE SET NULL
);
---- create above / drop below ----
DROP TABLE IF EXISTS users
DROP TABLE IF EXISTS emails
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
